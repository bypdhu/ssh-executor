package task

import (
	"testing"
	"github.com/bypdhu/ssh-executor/common"
	"fmt"
)

func TestDefaultTask(t *testing.T) {
	a := DefaultTask(common.MODULE_SHELL)
	fmt.Println(a)

	b := DefaultTask(common.MODULE_COPY)
	fmt.Println(b)

}
