package acdc

func Add(s ...int) int {
	ans := 0
	for _, val := range s {
		ans = ans + val
	}
	return ans
}
