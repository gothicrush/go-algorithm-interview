func removeDuplicates(nums []int) int {
    
    if length := len(nums); length < 3 {
        return length
    }
    
    var slow int = 2
    var fast int = 2
    
    for fast < len(nums) {
        
        if nums[fast] != nums[slow-2] {
            nums[slow] = nums[fast]
            slow++
        }
        
        fast++
    }
    
    return slow
}
