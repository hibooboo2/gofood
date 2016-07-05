package main

import (
	"errors"
	"strconv"
	"strings"
)

// For reading in arrays from postgres
type FloatSlice []float64
type IntSlice []int

type Recipe struct {
	Id            int
	IngredientIds IntSlice `db:"ingredient_ids"`
	Name          string
	Steps         string
	Tags          IntSlice
}

func (s *IntSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	asString := string(asBytes)
	(*s) = strToIntSlice(asString)
	return nil
}

func (s *FloatSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return error(errors.New("Scan source was not []byte"))
	}
	asString := string(asBytes)
	(*s) = strToFloatSlice(asString, *s)
	return nil
}

func strToFloatSlice(s string, a []float64) []float64 {
	r := strings.Trim(s, "{}")
	if a == nil {
		a = make([]float64, 0, 10)
	}
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.ParseFloat(t, 64)
		a = append(a, i)
	}
	return a
}

func strToIntSlice(s string) []int {
	r := strings.Trim(s, "{}")
	a := make([]int, 0, 10)
	for _, t := range strings.Split(r, ",") {
		i, _ := strconv.Atoi(t)
		a = append(a, i)
	}
	return a
}
