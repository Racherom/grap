package main_test

import (
	"fmt"
	"testing"
)

type color [4]byte

func TestBla(t *testing.T) {
	a := color{
		1,
		2,
		3,
		4,
	}

	fmt.Println(a)

	for i := range a {
		fmt.Println(a[i])
	}
}
