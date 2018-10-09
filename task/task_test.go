package task

import (
	"testing"
	"fmt"

	"github.com/bypdhu/ssh-executor/common"
)

func TestDefaultTask(t *testing.T) {
	a := DefaultTask(common.MODULE_SHELL)
	fmt.Println(a)

	b := DefaultTask(common.MODULE_COPY)
	fmt.Println(b)

}
