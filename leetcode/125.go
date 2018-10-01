func isPalindrome(s string) bool {
    
    exp := regexp.MustCompile(`[0-9a-zA-Z]+`)
    
    result := []byte(strings.ToUpper(sliceToString(exp.FindAllString(s,-1))))
    
    var left int = 0
    var right int = len(result)-1
    
    for left < right {
        
        if result[left] != result[right] {
            return false
        }
        
        left++
        right--
    }
    
    return true
}

func sliceToString(src []string) string {
    
    str := ""
    
    for _, item := range src {
        str = str + item
    }
    
    return str
}
