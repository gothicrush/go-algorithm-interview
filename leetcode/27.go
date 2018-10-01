func removeElement(nums []int, val int) int {
    
    var slow int = -1
    var fast int = 0
    var count int = 0
    
    for fast < len(nums) {
        
        if nums[fast] != val {
            slow++
            count++
            nums[slow] = nums[fast]
        }
        
        fast++
    }
    
    return count
}
