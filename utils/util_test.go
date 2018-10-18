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

func TestInsertStringSliceCopy(t *testing.T) {
    l1 := []string{"a", "b", "c"}
    l2 := []string{"d", "e"}

    l3 := InsertStringSliceCopy(l1, l2, 1)
    fmt.Printf("%s\n", l3)
    l3 = InsertStringSliceCopy(l1, l2, 0)
    fmt.Printf("%s\n", l3)
    l3 = InsertStringSliceCopy(l1, l2, 2)
    fmt.Printf("%s\n", l3)
    l3 = InsertStringSliceCopy(l1, l2, 3)
    fmt.Printf("%s\n", l3)
}

func TestInsertStringToSlice(t *testing.T) {
    l1 := []string{"a", "b", "c"}
    s := "d"

    l2 := InsertStringToSlice(l1, s, 0)
    fmt.Printf("%s\n", l2)
    l2 = InsertStringToSlice(l1, s, 1)
    fmt.Printf("%s\n", l2)
    l2 = InsertStringToSlice(l1, s, 2)
    fmt.Printf("%s\n", l2)
    l2 = InsertStringToSlice(l1, s, 3)
    fmt.Printf("%s\n", l2)

}