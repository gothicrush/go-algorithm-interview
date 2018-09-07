func numberOfBoomerangs(points [][]int) int {
    
    result := 0
    
    for i := 0; i < len(points); i++ {
    
        m := make(map[int]int, 10)
        
        for j := 0; j < len(points); j++ {
            
            if i != j {
                dis := distance(points[i],points[j])
                m[dis]++
                // if times, ok := m[dis]; ok {
                //     m[dis] = times + 1
                // } else {
                //     m[dis] = 1
                // }
            }
        }
        
        for _, times := range m {
            if times >= 2 {
                result += times * (times-1)
            }
        }
    }
    
    return result
}

func distance(a, b []int) int {
    return (a[0] - b[0]) * (a[0] - b[0]) + (a[1] - b[1]) * (a[1] - b[1])
}
