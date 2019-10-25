package diff

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kylelemons/godebug/pretty"
	"github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
)

func DiffYaml(file1, file2 string, color bool) (string, error) {
	formatter := newFormatter(color)

	if err := stat(file1, file2); err != nil {
		return "", err
	}

	yaml1, err := unmarshal(file1)
	if err != nil {
		return "", err
	}
	yaml2, err := unmarshal(file2)
	if err != nil {
		return "", err
	}

	diff := computeDiff(formatter, yaml1, yaml2)

	return diff, nil
}

func stat(filenames ...string) error {
	for _, filename := range filenames {
		if filename == "-" {
			continue
		}
		_, err := os.Stat(filename)
		if err != nil {
			return fmt.Errorf("cannot find file: %v. Does it exist?", filename)
		}
	}
	return nil
}

func unmarshal(filename string) (interface{}, error) {
	var contents []byte
	var err error
	if filename == "-" {
		contents, err = ioutil.ReadAll(os.Stdin)
	} else {
		contents, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		return nil, err
	}
	var ret interface{}
	err = yaml.Unmarshal(contents, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func computeDiff(formatter aurora.Aurora, a interface{}, b interface{}) string {
	diffs := make([]string, 0)
	for _, s := range strings.Split(pretty.Compare(a, b), "\n") {
		switch {
		case strings.HasPrefix(s, "+"):
			diffs = append(diffs, formatter.Bold(formatter.Green(s)).String())
		case strings.HasPrefix(s, "-"):
			diffs = append(diffs, formatter.Bold(formatter.Red(s)).String())
		}
	}
	return strings.Join(diffs, "\n")
}

func newFormatter(color bool) aurora.Aurora {
	if color {
		return aurora.NewAurora(true)
	}

	return aurora.NewAurora(false)
}
