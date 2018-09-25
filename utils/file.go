package utils

import (
	"os"
	"bufio"
	//"strings"
	"io"
	"os/user"
	"strings"
	"path/filepath"
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
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

func _getMd5(f io.Reader) ([]byte, error) {
	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		return nil, err
	}
	return md5Hash.Sum(nil), nil
}

func GetMd5(f io.Reader) string {
	if _md5, err := _getMd5(f); err != nil {
		return ""
	} else {
		return fmt.Sprintf("%x", _md5)
	}
}

func GetMd5FromPath(path string) (md5string string, err error) {
	if !IsFile(path) {
		return "", errors.New(path + " is not a file.")
	}
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()

	_md5, err := _getMd5(f)
	if err != nil {
		return
	}

	md5string = fmt.Sprintf("%x", _md5)
	return
}

// Judge if file or dir exists?
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
	//if err != nil {
	//	if os.IsExist(err) {
	//		return true
	//	}
	//	return false
	//}
	//return true
}

// Judge if given path is a dir. False if not exist.
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Judge if given path is a file. False if not exist.
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}