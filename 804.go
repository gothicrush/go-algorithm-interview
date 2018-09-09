func uniqueMorseRepresentations(words []string) int {
    
    var alphabetArr []string = []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    
    var alphabetMap map[rune]string = make(map[rune]string, 10)
    
    var i rune = 97
    for i <= 122 {
        alphabetMap[i] = alphabetArr[i-97]
        i++
    }
         
    var m map[string]bool = make(map[string]bool, 10)
    
    for _, word := range words {
        str := morse(word, alphabetMap)
        m[str] = true
    }
    
    return len(m)
}

func morse(word string, m map[rune]string) string {
    
    var res string = ""
    
    for _, alphabet := range word {
        res += m[alphabet]
    }
    
    return res
}
