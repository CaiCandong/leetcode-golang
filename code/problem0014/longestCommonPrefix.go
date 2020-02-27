package problem0014

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix1(strs)
}

// 解法1:leetcode水平解法
// 设res为strs[0],与后续字符串迭代更新res
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, str := range strs {
		i := 0
		for i < len(res) && i < len(str) && res[i] == str[i] {
			i++
		}
		if i == 0 {
			return ""
		} else {
			res = str[:i]
		}
	}
	return res
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i == len(str) {
				return str
			}
			if str[i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

// 解法3: 分治算法
func longestCommonPrefix3(strs []string) string {
	switch {
	case len(strs) == 0:
		return ""
	case len(strs) == 1:
		return strs[0]
	case len(strs) == 2:
		i := 0
		for i < len(strs[0]) && i < len(strs[1]) && strs[0][i] == strs[1][i] {
			i++
		}
		return strs[0][:i]
	default:
		str1 := longestCommonPrefix3(strs[:(len(strs)+1)/2])
		str2 := longestCommonPrefix3(strs[(len(strs)+1)/2:])
		return longestCommonPrefix3([]string{str1, str2})
	}
}
