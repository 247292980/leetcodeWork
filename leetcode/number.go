package leetcode

import (
	"bytes"
	"fmt"
	"math"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ThreeSumClosest(nums []int, target int) int {
	fmt.Println("threeSumClosest begin")
	sort.Ints(nums)

	var ThreeSumClosest = nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := len(nums) - 1
			for l < r {
				if ThreeSumClosest == target {
					return target
				}
				s := nums[i] + nums[l] + nums[r]
				if math.Abs(float64(s-target)) < math.Abs(float64(ThreeSumClosest-target)) {
					ThreeSumClosest = s
				}
				if s > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return ThreeSumClosest
}

/*开头的时候思路是对的，但是双指针+定点的定点找错了...后模仿js重写一次*/
func ThreeSum(nums []int) [][]int {
	fmt.Println("ThreeSum begin")
	sort.Ints(nums)

	var threeSum [][]int

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			l := i + 1
			r := len(nums) - 1
			for l < r {
				s := nums[i] + nums[l] + nums[r]
				if s == 0 {
					threeSum = append(threeSum, []int{nums[l], nums[i], nums[r]})
					l++
					r--
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					for r > l && nums[r] == nums[r+1] {
						r--
					}
				} else if s > 0 {
					r--
				} else {
					l++
				}
			}
		}
	}
	return threeSum
}
func RomanToInt(s string) int {
	fmt.Println("romanToInt begin")
	m := map[rune]int{
		rune('I'): 1,
		rune('V'): 5,
		rune('X'): 10,
		rune('L'): 50,
		rune('C'): 100,
		rune('D'): 500,
		rune('M'): 1000,
	}
	runes := []rune(s)
	if len(runes) == 1 {
		return m[runes[0]]
	}
	ret := m[runes[0]]
	for i := 1; i < len(runes); i++ {
		if m[runes[i-1]] < m[runes[i]] {
			ret += m[runes[i]] - 2*m[runes[i-1]]
		} else {
			ret += m[runes[i]]
		}
	}
	return ret
}
func IntToRoman(num int) string {
	fmt.Println("intToRoman begin")

	bufStr := bytes.NewBufferString("")
	nums := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	for i := 0; i < 13; i++ {
		tmp := num / nums[i]
		if tmp > 0 {
			bufStr.Write(bytes.Repeat([]byte(romanNums[i]), tmp))
			num -= tmp * nums[i]
		} else {
			continue
		}
	}
	return bufStr.String()
}
func MaxArea(height []int) int {
	fmt.Println("maxArea begin")
	var r = len(height) - 1
	var l = 0
	var temp = 0
	var maxArea = 0

	for l < r {
		if height[l] < height[r] {
			temp = height[l] * (r - l)
			l++
		} else {
			temp = height[r] * (r - l)
			r--
		}
		if temp > maxArea {
			maxArea = temp
		}
	}
	return maxArea
}

/*第一次和官方解答一样啊。纪念下*/
func IsPalindrome(x int) bool {
	fmt.Println("isPalindrome begin")
	var temp = x
	if temp < 0 {
		return false
	}
	var reverse = 0
	for temp/10 != 0 || temp%10 != 0 {
		reverse = reverse*10 + temp%10
		temp = temp / 10
	}
	if reverse == x {
		return true
	}
	return false
}
func Reverse(x int) int {
	fmt.Println("reverse begin")
	var reverse = 0
	for x/10 != 0 || x%10 != 0 {
		reverse = reverse*10 + x%10
		x = x / 10
	}
	if reverse > (2<<30)-1 || reverse < -2<<30 {
		reverse = 0
	}
	return reverse
}
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	fmt.Println("findMedianSortedArrays begin")
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	var midPoint = len(nums) / 2

	if len(nums)%2 == 0 {
		return float64(nums[midPoint]+nums[midPoint-1]) / 2
	} else {
		return float64(nums[midPoint])
	}
}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("AddTwoNumbers begin")

	var head = ListNode{}
	var curr = &head

	var list1 = l1
	var list2 = l2

	var temp = 0
	for list1 != nil || list2 != nil {
		var x = 0
		if list1 != nil {
			x = list1.Val
			list1 = list1.Next
		}

		var y = 0
		if list2 != nil {
			y = list2.Val
			list2 = list2.Next
		}

		var sum = temp + x + y
		temp = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
	}
	if temp > 0 {
		curr.Next = &ListNode{temp, nil}
		curr = curr.Next
	}
	return head.Next
}

func TwoSum(nums []int, target int) []int {
	fmt.Println("twoSum begin")
	var temp = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		needNum := target - nums[i]
		/*判断key是否存在*/
		_, ok := temp[needNum]
		if ok {
			return []int{temp[needNum], i}
		}
		temp[nums[i]] = i
	}
	return nil
}
