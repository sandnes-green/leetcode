// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//

// 示例 1：

// 输入：s = "()"
// 输出：true
// 示例 2：

// 输入：s = "()[]{}"
// 输出：true
// 示例 3：

// 输入：s = "(]"
// 输出：false
// 示例 4：

// 输入：s = "([)]"
// 输出：false
// 示例 5：

// 输入：s = "{[]}"
// 输出：true
//

// 提示：

// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成
package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()))"))
	fmt.Println(isValid("["))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
}
func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	m := map[rune]rune{'(': 1, '[': 2, '{': 3, ')': -1, ']': -2, '}': -3}
	var sum rune
	var cur = []rune{}
	for _, v := range s {
		sum += m[v]
		switch {
		case m[v] > 0:
			cur = append(cur, v)
		case m[v] < 0:
			if len(cur) <= 0 {
				return false
			}
			if m[cur[len(cur)-1]]+m[v] == 0 {
				cur = cur[:len(cur)-1]
			} else {
				return false
			}
		}
	}
	//括号没有全部闭合
	if sum != 0 {
		return false
	}
	return true
}

// func isValid(s string) bool {
// 	m := map[rune]string{'(': "left", '[': "left", '{': "left", ')': "right", ']': "right", '}': "right"}
// 	m1 := map[rune]rune{'(': ')', '[': ']', '{': '}'}
// 	if s == "" {
// 		return false
// 	}
// 	if len(s) == 1 {
// 		return false
// 	}
// 	r := rune(s[len(s)-1])
// 	if m[r] == "left" {
// 		return false
// 	}
// 	cur := []rune{rune(s[0])}
// 	for _, v := range s {
// 		if m[v] == "left" {
// 			cur = append(cur, v)
// 		} else if m[v] == "right" {
// 			if len(cur) == 0 {
// 				return false
// 			}
// 			// 没有闭合
// 			if m1[cur[len(cur)-1]] != v {
// 				return false
// 			} else {
// 				cur = cur[:len(cur)-1]
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	if len(cur) != 1 {
// 		return false
// 	}
// 	return true
// }
