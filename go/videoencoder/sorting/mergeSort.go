package sorting

// Sorts array
func MergeSort(arr *[]int) {
	mergeSort(arr, 0, len(*arr))
}

func mergeSort(arr *[]int, l int, r int) {
	if l+1 >= r {
		return
	}

	mid := (l + r) / 2

	mergeSort(arr, l, mid)
	mergeSort(arr, mid, r)

	pl := l
	pr := mid
	if pl >= pr {
		return
	}
	arrSection := make([]int, r-l)

	for i := l; i < r; i++ {
		if pl < mid && pr < r {
			if (*arr)[pl] <= (*arr)[pr] {
				arrSection[i-l] = (*arr)[pl]
				pl++
			} else {
				arrSection[i-l] = (*arr)[pr]
				pr++
			}
		} else if pl < mid {
			arrSection[i-l] = (*arr)[pl]
			pl++
		} else if pr < r {
			arrSection[i-l] = (*arr)[pr]
			pr++
		} else {

			panic("WHAT?")
		}
	}

	for i := l; i < r; i++ {
		(*arr)[i] = arrSection[i-l]
	}
}
