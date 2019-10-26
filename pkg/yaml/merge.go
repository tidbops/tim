package yaml

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/imdario/mergo"
	"github.com/ngaut/log"
	yaml "gopkg.in/mikefarah/yaml.v2"
)

func Merge(overwrite bool, append bool, input string, filesToMerge ...string) (string, error) {
	docIndexIntn := 0

	if input == "" {
		return "", errors.New("must provide filename")
	}

	var stream io.Reader
	if input == "-" {
		stream = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(input) // nolint gosec
		if err != nil {
			return "", err
		}
		defer safelyCloseFile(file)
		stream = file
	}

	var updateData = func(dataBucket interface{}, currentIndex int) (interface{}, error) {
		log.Debugf("Merging doc %v", currentIndex)
		if currentIndex == docIndexIntn {
			var mergedData map[interface{}]interface{}
			// merge only works for maps, so put everything in a temporary
			// map
			var mapDataBucket = make(map[interface{}]interface{})
			mapDataBucket["root"] = dataBucket
			if err := merge(&mergedData, mapDataBucket, overwrite, append); err != nil {
				return nil, err
			}

			for _, f := range filesToMerge {
				var fileToMerge interface{}
				if err := readData(f, 0, &fileToMerge); err != nil {
					if err == io.EOF {
						continue
					}
					return nil, err
				}
				mapDataBucket["root"] = fileToMerge
				if err := merge(&mergedData, mapDataBucket, overwrite, append); err != nil {
					return nil, err
				}
			}
			return mergedData["root"], nil
		}
		return dataBucket, nil
	}
	yaml.DefaultMapType = reflect.TypeOf(map[interface{}]interface{}{})

	defer func() {
		yaml.DefaultMapType = reflect.TypeOf(yaml.MapSlice{})
	}()

	return readAndUpdate(stream, updateData)
}

type updateDataFn func(dataBucket interface{}, currentIndex int) (interface{}, error)

func readAndUpdate(reader io.Reader, updateData updateDataFn) (string, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(buf)
	encoder := yaml.NewEncoder(writer)

	yamDecoder := mapYamlDecoder(updateData, encoder)
	if err := yamDecoder(yaml.NewDecoder(reader)); err != nil {
		return "", err
	}

	if err := writer.Flush(); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func mapYamlDecoder(updateData updateDataFn, encoder *yaml.Encoder) yamlDecoderFn {
	return func(decoder *yaml.Decoder) error {
		var (
			dataBucket   interface{}
			currentIndex = 0
			err          error
		)

		for {
			log.Debugf("Read doc %v", currentIndex)
			err = decoder.Decode(&dataBucket)

			if err == io.EOF {
				return nil
			} else if err != nil {
				return fmt.Errorf("faied to read document at index %v, %v", currentIndex, err)
			}

			dataBucket, err = updateData(dataBucket, currentIndex)
			if err != nil {
				return fmt.Errorf("failed to update document at index %v, %v", currentIndex, err)
			}

			err = encoder.Encode(dataBucket)

			if err != nil {
				return fmt.Errorf("failed to write document at index %v, %v", currentIndex, err)
			}
			currentIndex = currentIndex + 1
		}
	}
}

type yamlDecoderFn func(*yaml.Decoder) error

func safelyCloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Error(err.Error())
	}
}

func merge(dst interface{}, src interface{}, overwrite bool, append bool) error {
	if overwrite {
		return mergo.Merge(dst, src, mergo.WithOverride)
	} else if append {
		return mergo.Merge(dst, src, mergo.WithAppendSlice)
	}
	return mergo.Merge(dst, src)
}

func readData(filename string, indexToRead int, parsedData interface{}) error {
	return readStream(filename, func(decoder *yaml.Decoder) error {
		for currentIndex := 0; currentIndex < indexToRead; currentIndex++ {
			if err := decoder.Decode(parsedData); err != nil {
				return fmt.Errorf("failed to process document at index %v, %v", currentIndex, err)
			}
		}
		return decoder.Decode(parsedData)
	})
}

func readStream(filename string, yamlDecoder yamlDecoderFn) error {
	if filename == "" {
		return errors.New("must provide filename")
	}

	var stream io.Reader
	if filename == "-" {
		stream = bufio.NewReader(os.Stdin)
	} else {
		file, err := os.Open(filename) // nolint gosec
		if err != nil {
			return err
		}
		defer safelyCloseFile(file)
		stream = file
	}
	return yamlDecoder(yaml.NewDecoder(stream))
}
