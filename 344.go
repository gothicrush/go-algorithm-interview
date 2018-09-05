func reverseString(s string) string {
    
    var src []byte = []byte(s)
    
    var left int = 0
    var right int = len(src) - 1
    
    for left < right {
        
        swap(src, left, right)
        
        left ++
        right --
    }
    
    return string(src)
}

func swap(src []byte, i int, j int) {
    
    var temp byte = src[i]
    src[i] = src[j]
    src[j] = temp
}
