package utils

import (
	"fmt"
	"testing"
)

func TestGetFullPath(t *testing.T) {
	path := "."
	fmt.Println(GetFullPath(path))
}

func TestReadLine(t *testing.T) {
	cp, _ := GetFullPath(".")
	ls, _ := ReadLine(cp + "/../testdata/ips2.txt")
	for i, l := range ls {
		fmt.Printf("%d %s", i + 1, l)
	}
}

func TestReadLineNotEmpty(t *testing.T) {
	cp, _ := GetFullPath(".")
	ls, _ := ReadLineNotEmpty(cp + "/../testdata/ips2.txt")
	for i, l := range ls {
		fmt.Printf("%d %s\n", i + 1, l)
	}
}
