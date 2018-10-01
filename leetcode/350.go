func intersect(nums1 []int, nums2 []int) []int {
    
    var result []int
    
    var m1 map[int]int = make(map[int]int, 10)
    var m2 map[int]int = make(map[int]int, 10)
    
    for _, val := range nums1 {
        if _, ok := m1[val]; !ok {
            m1[val] = 1
        } else {
            m1[val] = m1[val] + 1
        }
    }
    
    for _, val := range nums2 {
        if _, ok := m2[val]; !ok {
            m2[val] = 1
        } else {
            m2[val] = m2[val] + 1
        }
    }
    
    for val, m1times := range m1 {
        
        if m2times, ok := m2[val]; ok {
            var loopTimes int = min(m1times,m2times)
            for loopTimes > 0 {
                result = append(result, val)
                loopTimes--
            }
        }
    }
    
    return result
}

func min(i, j int) int {
    
    if i < j {
        return i
    }
    
    return j
}
