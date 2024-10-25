func twoSum(nums []int, target int) []int {
	flag := map[int]int{}
	for index, num := range nums {
		if ans, ok := flag[target-num]; ok {
			return []int{ans, index}
		}
		flag[num] = index
	}
	return nil
}