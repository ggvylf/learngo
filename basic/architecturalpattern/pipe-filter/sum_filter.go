package main

import (
	"errors"
)

var SumFilterWrongFormatError = errors.New("input data must be []int")

type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Porcess(data Request) (Response, error) {
	parts, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}

	sum := 0
	for _, part := range parts {
		sum += part
	}

	return sum, nil
}
