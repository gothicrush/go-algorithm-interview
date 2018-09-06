func wordPattern(pattern string, str string) bool {
    
    strSlice := strings.Split(str, " ")
    byteSlice := []byte(pattern)
    
    if len(strSlice) != len(byteSlice) {
        return false
    }
    
    mBS := make(map[byte]string, 10)
    mSB := make(map[string]byte, 10)
    
    for i := 0; i < len(strSlice); i++ {
        
        b := byteSlice[i]
        s := strSlice[i]
        
        _, okBS := mBS[b]
        _, okSB := mSB[s]
        
        if okBS != okSB {
            return false
        }
        
        if !okBS {
            mBS[b] = s
            mSB[s] = b
            continue
        } 
        
        if okBS {
            valB, _ := mSB[s]
            valS, _ := mBS[b]
            
            if valB == b && valS == s {
                continue
            } else {
                return false
            }
        }
    }
    
    return true
}
