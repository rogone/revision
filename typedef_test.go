package tags

import (
	"fmt"
	"testing"
)

func TestRedef(t *testing.T) {
	type RedefT int
	const (
		RedefOne RedefT = iota
		RedefTwo
		RedefThree
	)

	type T struct {
		I RedefT `default:"2"`
	}

	var v T

	err := Revise(&v)
	fmt.Println(v, err)
}

func TestAlias(t *testing.T) {
	type RedefT = int
	const (
		RedefOne RedefT = iota
		RedefTwo
		RedefThree
	)

	type T struct {
		I RedefT `default:"2"`
	}

	var v T

	err := Revise(&v)
	fmt.Println(v, err)
}
