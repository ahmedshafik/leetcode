func removeDuplicates(nums []int) int {
    left := 1

    if len(nums) == 0{
        return 0
    }

for right:=1 ; right < len(nums); right++ {    
    	if nums[right] != nums[right - 1] {
        nums[left] = nums[right]
        left++
      }
      
    }
    return left
}