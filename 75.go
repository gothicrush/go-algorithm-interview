func sortColors(nums []int)  {
    
    var zero int = -1
    var judge int = 0
    var two int = len(nums)
    
    for judge < two {
        
        if nums[judge] == 0 {
            zero++
            swap(nums,zero,judge)
            judge++
        } else if nums[judge] == 1 {
            judge++
        } else { // nums[judge] == 2
            two--
            swap(nums,judge,two)
        }
    }
}

func swap(nums []int, i int, j int) {
    
    var temp int = nums[i]
    nums[i] = nums[j]
    nums[j] = temp
}
