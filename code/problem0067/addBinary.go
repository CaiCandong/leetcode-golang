package problem0067

func addBinary(a string, b string) string {
	var length int
	if len(a) > len(b) {
		length = len(a)
	} else {
		length = len(b)
	}
	aa := str2ints(a, length)
	bb := str2ints(b, length)
	carry := 0
	res := make([]int, length+1)
	for i := length - 1; i >= 0; i-- {
		res[i+1] = (aa[i] + bb[i] + carry) % 2
		carry = (aa[i] + bb[i] + carry) / 2
	}
	if carry == 1 {
		res[0] = 1
	} else {
		res = res[1:]
	}
	return Ints2str(res)
}
func str2ints(str string, l int) []int {
	res := make([]int, l-len(str))
	for _, b := range str {
		res = append(res, int(b-'0'))
	}
	return res
}
func Ints2str(nums []int) string {
	bytes := make([]byte, len(nums))
	for i := range bytes {
		bytes[i] = byte(nums[i]) + '0'
	}
	return string(bytes)
}
func addBinary2(a, b string) string {
	var result []byte
	lena, lenb := len(a), len(b)
	if lena > lenb {
		result = make([]byte, lena+1)
	} else {
		result = make([]byte, lenb+1)
	}

	var carryBit uint8
	i, j := lena-1, lenb-1
	for t := len(result) - 1; t >= 0; t-- {
		sum := carryBit
		carryBit = 0
		if i > 0 {
			sum += (a[i] - '0')
			i--
		}
		if j > 0 {
			sum += (b[i] - '0')
			j--
		}
		result[t] = sum + '0'
	}
	if result[0] == '0' {
		return string(result[1:])
	}
	return string(result)

}
