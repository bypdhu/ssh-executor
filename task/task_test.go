package task

import (
	"testing"
	"fmt"

	"github.com/bypdhu/ssh-executor/common"
)

func TestDefaultTask(t *testing.T) {
	a := DefaultTask(common.MODULE_SHELL.String())
	fmt.Println(a)

	b := DefaultTask(common.MODULE_COPY.String())
	fmt.Println(b)

}
