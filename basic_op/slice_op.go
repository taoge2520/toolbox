package basic_op

//返回一个只保存非空字符串的切片。
//在调用期间修改基础数组。
func Nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//删除slice中间的某个元素并保存原有的元素顺序
func Remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
