package _0

// 录入字符串 大写变小写，小写变大写
func Str(str string) string {
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		if str1[i] < 97 {
			str1[i] += 32
		} else {
			str1[i] -= 32
		}
	}
	return string(str1)
}


