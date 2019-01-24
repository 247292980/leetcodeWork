package leetcode20

import "fmt"

/*用切片模拟栈*/
func isValid2(s string) bool {
	stack := make([]string, 0, 0)
	parenMap := make(map[string]string, 0)
	parenMap[")"] = "("
	parenMap["]"] = "["
	parenMap["}"] = "{"

	for i := 0; i < len(s); i++ {
		key := string(s[i])
		value, ok := parenMap[key]
		if !ok {
			stack = append(stack, key)
		} else if len(stack) > 0 {
			pop := stack[len(stack)-1]
			if pop != value {
				return false
			} else {
				stack = stack[0 : len(stack)-1]
			}
		} else { // 直接右括号进来干掉
			return false
		}
	}
	return len(stack) == 0

}

/*第一次打败了3.33%的人...ok，要用栈*/
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	var l1, r1 = '[', ']'
	var l2, r2 = '(', ')'
	var l3, r3 = '{', '}'
	for index, v := range s {
		if index == len(s)-1 {
			continue
		}
		if v == l1 && s[index+1] == uint8(r1) {
			return IsValid(s[0:index-1] + s[index+1:])
		}
		if v == l2 && s[index+1] == uint8(r2) {
			return IsValid(s[0:index] + s[index+2:])
		}
		if v == l3 && s[index+1] == uint8(r3) {
			return IsValid(s[0:index] + s[index+2:])
		}
	}
	if len(s) > 0 {
		return false
	}
	return true
}
