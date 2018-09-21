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