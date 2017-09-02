package mole

import (
	"reflect"
	"testing"
)

type NestedEmbeded struct {
	NestEmbStr  string `mole:"name"`
	NestEmbInt  int    `mole:"name"`
	NestEmbBool bool   `mole:"name"`
}

type Embeded struct {
	EmbStr        string `mole:"name"`
	EmbInt        int    `mole:"name"`
	EmbBool       bool   `mole:"name"`
	NestedEmbeded `mole:"name"`
}

type S struct {
	Str        string
	Int        int
	privateInt int
	Bool       bool
	Slicy      []float64
	Embeded
	AnonStruct struct {
		Str  string
		Int  int
		Bool bool
	}
	privateAnonStruct struct {
		Str  string
		Int  int
		Bool bool
	}
	AnonStructSlice []struct {
		Str  string
		Int  int
		Bool bool
	}
}

func TestFlattern(t *testing.T) {

	s := S{}
	styp := reflect.TypeOf(s)

	flattern(styp)

}
