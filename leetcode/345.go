func reverseVowels(s string) string {
    
    var src []byte = []byte(s)
    
    var left int = moveRightToVolve(src, -1)
    var right int = moveLeftToVolve(src, len(src))
    
    for left < right {
        
        swap(src, left, right)
        
        left = moveRightToVolve(src, left)
        right = moveLeftToVolve(src, right)
    }
    
    return string(src)
}

func swap(src []byte, i int, j int) {
    
    var temp byte = src[i]
    src[i] = src[j]
    src[j] = temp
}

func moveRightToVolve(src []byte, left int) int {
    
    for left < len(src)-1 {
        
        left++

        if judge(src[left]) {
            return left
        }   
    }
    
    return len(src)
}

func moveLeftToVolve(src []byte, right int) int {
    
    for right > 0 {
        
        right--

        if judge(src[right]) {
            return right
        }
    }
            
    return -1
}

func judge(b byte) bool {
    
    return b == 'a' || b == 'e' || b == 'u' || b == 'i' || b == 'o' ||
    	   b == 'A' || b == 'E' || b == 'U' || b == 'I' || b == 'O'
}
