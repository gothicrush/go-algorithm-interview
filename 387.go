func firstUniqChar(s string) int {
    
    var m map[rune]int = make(map[rune]int)
    
    for _, item := range s {
        m[item]++
    }
    
    for index, item := range s {
        if m[item] == 1 {
            return index
        }
    }
    
    return -1
}
