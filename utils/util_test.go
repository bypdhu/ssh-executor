package utils

import (
	"testing"
	"fmt"
)

func TestRemoveDupString(t *testing.T) {

	l := []string{"a", "b", "c", "d", "a", "", "a", "b", ""}

	l_new := RemoveDupString(l)

	fmt.Printf("length:%d, elements:%s", len(l_new), l_new)
}