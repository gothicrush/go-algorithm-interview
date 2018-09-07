func containsDuplicate(nums []int) bool {
    
    var m map[int]bool = make(map[int]bool, 10)
    
    for _,val := range nums {
        
        if _, ok := m[val]; ok {
            return true
        } else {
            m[val] = true
        }
    }
    
    return false
}
