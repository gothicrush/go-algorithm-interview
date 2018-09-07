func fourSumCount(A []int, B []int, C []int, D []int) int {
    
    var m1 map[int]int = make(map[int]int, 10)
    var m2 map[int]int = make(map[int]int, 10)
    
    length := len(A)
    result := 0
    
    for i := 0; i < length; i++ {
        
        for j := 0; j < length; j++ {
            sum := A[i] + B[j]
            
            if times, ok := m1[sum]; ok {
                m1[sum] = times + 1
            } else {
                m1[sum] = 1
            }
        }
        
        for j := 0; j < length; j++ {
            sum := C[i] + D[j]
            
            if times, ok := m2[sum]; ok {
                m2[sum] = times + 1
            } else {
                m2[sum] = 1
            }
        }
    }
    
    for sum, timesM1 := range m1 {
        
        sum = 0 - sum
        
        if timesM2, ok := m2[sum]; ok {
            result = result + timesM1 * timesM2
        }
    }
    
    return result
}
