package main

import (
	"fmt"
	"testing"
)

func TestProcessUserData(T *testing.T) {

	err := processUserData("test")
	fmt.Println(err)
}

func TestProcessMetaData(T *testing.T) {

	_ = processMetaData("test")
}
