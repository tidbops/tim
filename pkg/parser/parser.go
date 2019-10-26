package parser

import (
	"fmt"
	"github.com/tidbops/tim/pkg/utils"
	"io/ioutil"
	"strings"
)

const (
	NewConfigStart    = "@new"
	DeleteConfigStart = "@delete"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParserFile(
	srcPath string,
	path string,
	prefix string,
) (string, string, error) {
	data, err := ioutil.ReadFile(srcPath)
	if err != nil {
		return "", "", err
	}

	lines := strings.Split(string(data), "\n")

	var (
		newConfigLines    []string
		deleteConfigLines []string
		newStart          bool
		newEnd            bool
		deleteStart       bool
		deleteEnd         bool
	)

	for _, line := range lines {
		switch {
		case strings.Contains(line, NewConfigStart):
			newStart = true
			if deleteEnd {
				deleteEnd = false
			}
		case strings.Contains(line, DeleteConfigStart):
			deleteStart = true
			if newStart {
				newStart = false
			}
		default:
		}

		if newStart && !newEnd {
			newConfigLines = append(newConfigLines, line)
		}

		if deleteStart && !deleteEnd {
			deleteConfigLines = append(deleteConfigLines, line)
		}
	}

	newRuleFile := fmt.Sprintf("%s/%s-newrule.yml", path, prefix)
	if err := utils.WriteLines(newConfigLines, newRuleFile); err != nil {
		return "", "", err
	}

	deleteRuleFile := fmt.Sprintf("%s/%s-deleterule.yml", path, prefix)
	if err := utils.WriteLines(deleteConfigLines, deleteRuleFile); err != nil {
		return "", "", err
	}

	return newRuleFile, deleteRuleFile, nil
}
