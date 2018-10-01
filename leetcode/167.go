func twoSum(numbers []int, target int) []int {
    
    var left int = 0
    var right int = len(numbers)-1
    
    var result []int
    
    for left < right {
        
        sum := numbers[left] + numbers[right] 
        
        if sum == target {
            result = []int{ left+1, right+1 }
            break
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    
    return result
}
