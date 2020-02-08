package sortoe



// 判断整型切片是否已升序排序
// to determine whether the integer slice has been sorted in ascending order
func IntsIsAsc(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// 判断整型切片是否已升序排序且元素唯一
// determines whether the integer slice is sorted in ascending order and elements all unique
func IntsIsAscAndUnique(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断整型切片是否已降序排序
// to determine whether the integer slice has been sorted in descending order
func IntsIsDesc(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

// 判断整型切片是否已降序排序且元素唯一
// determines whether the integer slice is sorted in descending order and elements all unique
func IntsIsDescAndUnique(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断整型切片是否全部相同
// determine whether all elements of the integer slice are the same
func IntsIsFlat(arr []int) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] != arr[i-1] {
			return false
		}
	}
	return true
}

// 计算切片相邻元素为升序的百分比, 分母为升序加降序不含相等的
// calculate the percentage of slice adjacent elements in ascending order,
// the denominator is ascending order plus descending order without equality
func IntsAscPercent(arr []int) int {
	var asc,desc int
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			asc++
		}else if arr[i] < arr[i-1] {
			desc++
		}
	}
	return asc*100/(asc+desc)
}

// 判断浮点型切片是否已升序排序
func FloatsIsAsc(arr []float64) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// 判断浮点型切片是否已升序排序且元素唯一
func FloatsIsAscAndUnique(arr []float64) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断浮点型切片是否已降序排序
func FloatsIsDesc(arr []float64) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

// 判断浮点型切片是否已降序排序且元素唯一
func FloatsIsDescAndUnique(arr []float64) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断浮点型切片是否全部相同
func FloatsIsFlat(arr []float64) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] != arr[i-1] {
			return false
		}
	}
	return true
}

// 计算切片相邻元素为升序的百分比, 分母为升序加降序不含相等的
func FloatsAscPercent(arr []float64) int {
	var asc,desc int
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			asc++
		}else if arr[i] < arr[i-1] {
			desc++
		}
	}
	return asc*100/(asc+desc)
}

// 判断字符串切片是否已升序排序
func StringsIsAsc(arr []string) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// 判断字符串切片是否已升序排序且元素唯一
func StringsIsAscAndUnique(arr []string) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] <= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断字符串切片是否已降序排序
func StringsIsDesc(arr []string) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}

// 判断字符串切片是否已降序排序且元素唯一
func StringsIsDescAndUnique(arr []string) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] >= arr[i-1] {
			return false
		}
	}
	return true
}

// 判断字符串切片是否全部相同
func StringsIsFlat(arr []string) bool {
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] != arr[i-1] {
			return false
		}
	}
	return true
}

// 计算切片相邻元素为升序的百分比, 分母为升序加降序不含相等的
func StringsAscPercent(arr []string) int {
	var asc,desc int
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] > arr[i-1] {
			asc++
		}else if arr[i] < arr[i-1] {
			desc++
		}
	}
	return asc*100/(asc+desc)
}


