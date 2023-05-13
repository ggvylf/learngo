package main

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input must a string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Porcess(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
