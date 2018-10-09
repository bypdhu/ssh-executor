package utils

// 该函数总共初始化两个变量，一个长度为0的slice，一个空map。由于slice传参是按引用传递，没有创建额外的变量。
// 只是用了一个for循环，代码更简洁易懂。
// 利用了map的多返回值特性。
// 空struct不占内存空间，可谓巧妙。
func RemoveDupString(a []string) []string {
	result := make([]string, 0, len(a))
	temp := map[string]struct{}{}

	for _, item := range a {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// 内置append()函数能够在切片末尾位置添加新的项,
// 假设要在切片的前面或者中间某位置插入特定项,可以这样实现
func InsertStringSliceCopy(slice, insertion []string, index int) []string {
	result := make([]string, len(slice) + len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

// 在切片的某位置插入一个元素
func InsertStringToSlice(slice []string, insert string, index int) []string {
	result := make([]string, len(slice) + 1)
	copy(result, slice[:index])
	result[index] = insert
	copy(result[index + 1:], slice[index:])
	return result
}