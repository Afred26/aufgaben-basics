package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	//sort.Ints(nums)
	if len(nums) == 0 {
		return 0
	} else {
		a := nums[0]
		nums1 := delete_at_index(nums, 0)
		if len(nums) > 1 {
			n := MIN(a, MinListRecursive(nums1))
			return n
		} else {
			return a
		}
		//n := nums[0]

	}

}

func MIN(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}

}

func delete_at_index(slice []int, index int) []int {

	// Append function used to append elements to a slice
	// first parameter as the slice to which the elements
	// are to be added/appended second parameter is the
	// element(s) to be appended into the slice
	// return value as a slice
	return append(slice[:index], slice[index+1:]...)
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
