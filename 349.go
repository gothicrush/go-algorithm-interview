func intersection(nums1 []int, nums2 []int) []int {

    var result []int
    
    var m map[int]bool = make(map[int]bool, 10)
    
    for _, v := range nums1 {
        m[v] = false
    }
    
    for _, v := range nums2 {
        
        if _, ok := m[v]; ok {
            delete(m, v)
            result = append(result, v)
        }
    }
    
    return result
}
