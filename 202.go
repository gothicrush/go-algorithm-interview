func isHappy(n int) bool {
    
    var m map[int]bool = make(map[int]bool, 10)
    
    m[n] = true
    
    for n != 1 {
        
        n = calculate(n)
        
        if _, ok := m[n]; ok {
            return false
        }
        
        m[n] = true
    }
    
    return true
}

func calculate(number int) int {
    
    var result int = 0
    
    for number > 0 {
        
        n := number % 10
        result += n * n
        
        number /= 10
    }
    
    return result
}
