package main

func moveZeroes(nums []int)  {
	length := len(nums)
	for i:= 0 ; i < length ; i++ {
		if nums[length-i-1] == 0 {
			copy(nums[length-i-1:], nums[length-i:])
			nums[length-1] = 0
		}
	}
}