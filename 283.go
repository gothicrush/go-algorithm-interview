func moveZeroes(nums []int)  {
    var slow int = 0
    var fast int = 0
    
    for fast < len(nums) {
        if (nums[fast] != 0) {
            nums[slow] = nums[fast]
            slow++
        }
        
        fast++
    }
    
    for slow < len(nums) {
        nums[slow] = 0
        slow++
    }
}
