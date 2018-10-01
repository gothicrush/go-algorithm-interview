func containsNearbyDuplicate(nums []int, k int) bool {
    
    m := make(map[int]bool, 10)
    
    for i := 0; i < len(nums); i++ {
        
        if _, ok := m[nums[i]]; ok {
            return true
        }
        
        m[nums[i]] = true
        
        if len(m) == k+1 {
            delete(m, nums[i-k])
        }
    }
    
    return false
}
