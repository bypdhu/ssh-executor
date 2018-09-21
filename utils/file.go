package utils

import (
	"os"
	"bufio"
	//"strings"
	"io"
	"os/user"
	"strings"
	"path/filepath"
)

func ReadLine(fileName string) (res []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	for {
		// ignore case: line is larger than buffer size.
		line, err1 := buf.ReadString('\n')
		//line = strings.TrimSpace(line)
		if err1 != nil {
			if err1 == io.EOF {
				res = append(res, line)
				return res, err1
			}
			return
		}
		if line != "" {
			res = append(res, line)
		}
	}
	return
}

func ReadLineNotEmpty(fileName string) (res []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	for {
		// ignore case: line is larger than buffer size.
		line, err1 := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err1 != nil {
			if err1 == io.EOF {
				if line != "" {
					res = append(res, line)
				}
				return res, nil
			}
			return
		}
		if line != "" {
			res = append(res, line)
		}
	}
	return
}

func GetFullPath(path string) (fullPath string, err error) {
	_user, err := user.Current()
	if err != nil {
		return
	}
	fullPath = strings.Replace(path, "~", _user.HomeDir, 1)
	fullPath, err = filepath.Abs(fullPath)
	if err != nil {
		return path, err
	}
	return
}