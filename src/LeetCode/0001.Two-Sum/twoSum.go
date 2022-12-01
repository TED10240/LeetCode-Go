package main

func twoSum(num []int, target int) []int {

	m := make(map[int]int)
	for i := 0; i < len(num); i++ {
		another := target - num[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[num[i]] = i
	}
	return nil
}
