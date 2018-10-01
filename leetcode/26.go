func removeDuplicates(nums []int) int {
    
    if length := len(nums); length <= 1 {
        return length
    }
    
    var slow int = 0
    var fast int = 1
    
    for fast < len(nums) {
        
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
        
        fast++
    }
    
    return slow + 1
}
