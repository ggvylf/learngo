package main

import (
	"errors"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data must be []string")

type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (ti *ToIntFilter) Porcess(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}

	return ret, nil
}
