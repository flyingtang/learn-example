package _1_字符串数组

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1:
// 输入: ["flower","flow","flight"]
// 输出: "fl"
// 示例 2:
// 输入: ["dog","racecar","car"]
// 输出: ""
// 解释: 输入不存在公共前缀。
// 说明:
// 所有输入只包含小写字母 a-z 。
// 思路

func LongestCommonPrefix(s []string) string {
	slength := len(s)
	if slength < 2 {
		return ""
	}

	first := s[0]
	wlength := len(first)
	index := 0
	for ; index < wlength; index++ {
		tw := first[index]

		for j := 1; j < slength; j++ {
			if index >= len(s[j]) || tw != s[j][index] {
				return first[:index]
			}
		}
	}
	return ""
}
