package utils

import (
	"fmt"
	"testing"
	"os"
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

func TestGetMd5(t *testing.T) {
	cp, _ := GetFullPath(".")

	f, _ := os.Open(cp + "/../testdata/ips2.txt")
	defer f.Close()

	_md5_2 := GetMd5(f)
	fmt.Println(_md5_2)

	_md5_2 = GetMd5(f)
	fmt.Println(_md5_2)

	_md5_3, _ := _getMd5(f)
	fmt.Printf("%x", _md5_3)
	fmt.Println()

	f2, _ := os.Open(cp + "/../testdata/ips2.txt")
	defer f2.Close()

	_md5, _ := _getMd5(f2)
	fmt.Printf("%x", _md5)
	fmt.Println()

}

func TestGetMd5FromPath(t *testing.T) {
	cp, _ := GetFullPath(".")

	fmt.Println(GetMd5FromPath(cp + "/../testdata/ips2.txt"))
	fmt.Println(GetMd5FromPath(cp + "/../testdata/ips2.txt"))
}

func TestIsDir(t *testing.T) {
	cp, _ := GetFullPath(".")

	fmt.Println(IsDir(cp))
	fmt.Println(IsDir(cp + "/../"))
	fmt.Println(IsDir(cp + "/../tmp"))
	fmt.Println(IsDir(cp + "/../tmp/testdir"))
	fmt.Println(IsDir(cp + "/../tmp/config.yml"))
}

func TestExist(t *testing.T) {
	cp, _ := GetFullPath(".")

	fmt.Println(Exist(cp))
	fmt.Println(Exist(cp + "/../tmp_not_exist"))
	fmt.Println(Exist(cp + "/../tmp/config.yml"))
}

func TestIsFile(t *testing.T) {
	cp, _ := GetFullPath(".")

	fmt.Println(IsFile(cp))
	fmt.Println(IsFile(cp + "/../tmp"))
	fmt.Println(IsFile(cp + "/../tmp/config.yml"))
}
