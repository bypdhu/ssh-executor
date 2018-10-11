package result

import (
	"testing"
	"fmt"
)

func TestShowResult(t *testing.T) {
	a := []TaskResult{}
	a = append(a, ShellTaskResult{})
	fmt.Printf("%+v\n", a)

	a = append(a, CopyTaskResult{})
	fmt.Printf("%+v\n", a)

}