func frequencySort(s string) string {
    
    
    m := make(map[rune]int, 10)
    
    for _, v := range s {
        if times, ok := m[v]; !ok {
            m[v] = 1
        } else {
            m[v] = times + 1
        }
    }
    
    var count map[int]string = make(map[int]string, 10)
    var times []int = make([]int, 10)
    
    for k, v := range m {

        times = append(times, v)
        
        str := ""
        
        for i := v; i > 0; i-- {
            str += string(k)
        }
        
        if val, ok := count[v]; !ok {
            count[v] = str
        } else {
            count[v] = val + str
        }
    }
    
    sort.Ints(times[:])

    result := ""
    
    set := make(map[int]bool,10)

    for i := len(times)-1; i >=0 ; i-- {

        if _, ok := set[times[i]]; !ok {
            result = result + count[times[i]]
            set[times[i]] = true    
        }
    }
    
    
    return result
}func frequencySort(s string) string {
    
    
    m := make(map[rune]int, 10)
    
    for _, v := range s {
        if times, ok := m[v]; !ok {
            m[v] = 1
        } else {
            m[v] = times + 1
        }
    }
    
    var count map[int]string = make(map[int]string, 10)
    var times []int = make([]int, 10)
    
    for k, v := range m {

        times = append(times, v)
        
        str := ""
        
        for i := v; i > 0; i-- {
            str += string(k)
        }
        
        if val, ok := count[v]; !ok {
            count[v] = str
        } else {
            count[v] = val + str
        }
    }
    
    sort.Ints(times[:])

    result := ""
    
    set := make(map[int]bool,10)

    for i := len(times)-1; i >=0 ; i-- {

        if _, ok := set[times[i]]; !ok {
            result = result + count[times[i]]
            set[times[i]] = true    
        }
    }
    
    
    return result
}
