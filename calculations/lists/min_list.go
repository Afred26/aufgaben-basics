package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	var n int
	for i := 0; i < len(nums); i++ {
		n = MIN(nums[i], n)

	}
	return n
}

func MINI(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}
