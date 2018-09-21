package conf

import (
	"testing"
	"fmt"
	"github.com/bypdhu/ssh-executor/utils"
)

func TestLoad(t *testing.T) {
	cp, _ := utils.GetFullPath(".")
	c := Load(cp + "/../testdata/config.yml")
	fmt.Printf("%s", c)
	fmt.Printf("%s", c.original)
}
