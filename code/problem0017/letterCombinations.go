package problem0017

//回溯法
var m = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}
var result []string
var length int
var digit string

func letterCombinations(digits string) []string {
	result = make([]string, 0)
	length = len(digits)
	if length == 0 {
		return nil
	}
	digit = digits
	temp := make([]byte, length)
	helper(temp, 0)
	return result
}
func helper(temp []byte, p int) {
	if p == length {
		result = append(result, string(temp))
		return
	}
	for _, ch := range m[digit[p]] {
		temp[p] = ch
		helper(temp, p+1)
	}
}

var m2 = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations2(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	for i := 0; i < len(digits); i++ {
		temp := []string{}
		for j := 0; j < len(ret); j++ {
			for _, ch := range m2[digits[i]] {
				temp = append(temp, ret[j]+ch)
			}
		}
		//fmt.Println(temp)

		ret = temp
	}

	return ret
}
