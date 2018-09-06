func isAnagram(s string, t string) bool {
    
    var rs []rune = []rune(s)
    var rt []rune = []rune(t)
    
    var m1 map[rune]int = make(map[rune]int,10)
    var m2 map[rune]int = make(map[rune]int,10)
    
    for _, word := range rs {
        if val, ok := m1[word]; !ok {
            m1[word] = 1
        } else {
            m1[word] = val + 1
        }
    }
    
    for _, word := range rt {
        if val, ok := m2[word]; !ok {
            m2[word] = 1
        } else {
            m2[word] = val + 1
        }
    }
    
    for word, times1 := range m1 {
        
        if len(m1) != len(m2) {
            return false
        }
        
        if times2, ok := m2[word]; !ok || times1 != times2 {
            return false
        }
    }

    return true
}
