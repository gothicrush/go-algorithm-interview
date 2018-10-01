func minSubArrayLen(s int, nums []int) int {
    
    var left int = 0
    var right int = -1
    
    var sum int = 0
    var minLength int = len(nums) + 1
    
    for left < len(nums) {
        
        if right + 1 < len(nums) && sum < s {
            right++
            sum = sum + nums[right]
        } else {
            sum = sum - nums[left]
            left++
        }
        
        if sum >= s {
            minLength = min(minLength,right - left + 1) 
        }
    }
    
    if minLength == len(nums) + 1 {
        return 0
    }
    
    return minLength
}

func min(i, j int) int {
    if i <= j {
        return i
    } else {
        return j
    }
}
