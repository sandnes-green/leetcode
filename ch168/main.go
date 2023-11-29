// 168. Excel表列名称
// 简单
// 给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
// 示例 1：

// 输入：columnNumber = 1
// 输出："A"
// 示例 2：

// 输入：columnNumber = 28
// 输出："AB"
// 示例 3：

// 输入：columnNumber = 701
// 输出："ZY"
// 示例 4：

// 输入：columnNumber = 2147483647
// 输出："FXSHRXW"

// 提示：

// 1 <= columnNumber <= 231 - 1

// 题解：26进制
package main

import "fmt"

func convertToTitle(columnNumber int) string {
	ans := []byte{}
	for columnNumber > 0 {
		columnNumber--
		ans = append(ans, 'A'+byte(columnNumber%26))
		columnNumber = columnNumber / 26
	}

	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return string(ans)
}

func main() {
	fmt.Println(convertToTitle(2147483647))
}
