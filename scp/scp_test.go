package scp

import (
	"testing"
	"github.com/bypdhu/ssh-executor/log"
	"fmt"
	"os"
	"io"
)

func TestWalkDir(t *testing.T) {
	r,_ := getFullPath("../../")
	files, err := walkDir(r)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(2)
	}
	for i, f := range files {
		fmt.Printf("%d %s\n", i, f)
	}
}

func TestGetFullPath(t *testing.T) {
	path := "."
	fmt.Println(getFullPath(path))
}

func TestPushFileData(t *testing.T) {
	var w io.WriteCloser
	path := "scp.go"
	toName := "scp.go.toName"
	perm := true
	pushFileData(w, path, toName, perm)
	fmt.Println(w)

}