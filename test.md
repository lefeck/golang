1.两数之和

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

package main

import (
	"fmt"
)
//分析:
//首先输入给定的是一个数组,和一个target, 返回的的是两数的索引,因此,我们构造的函数如下:
//循环遍历数组中的所有元素,2次for循环取2个值
//遍历到的元素相加和为给定值,返回结果--->if判断,return
func sum(nums []int, target int) []int {
	l := len(nums)
	for i:=0; i <l; i++{
		for j:=i+1; j<l; j++ {
			if nums[i]+nums[j] == target {
				return []int {i, j}
			}
		}
	}
	return []int{}
}

func main() {
	b := sum()
	fmt.Println(b)
}
