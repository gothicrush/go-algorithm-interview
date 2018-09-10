func isValid(s string) bool {
    
    var stack []rune
    
    for _, item := range s {
        
        length := len(stack)
        
        if item == '{' || item == '[' || item == '(' {
            stack = append(stack, item)
        }
        
        if item == '}' {
            if length <= 0 || stack[length-1] != '{' {
                return false
            } else {
                stack = stack[:length-1]
            }
        }
        
        if item == ')' {
            if length <= 0 || stack[length-1] != '(' {
                return false
            } else {
                stack = stack[:length-1]
            }
        }
        
        if item == ']' {
            if length <= 0 || stack[length-1] != '[' {
                return false
            } else {
                stack = stack[:length-1]
            }
        } 
    }
    
    if len(stack) == 0 {
        return true
    }
    
    return false
}
