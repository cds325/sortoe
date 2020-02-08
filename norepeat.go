package sortoe


func sortedIntsNoRepeat(arr []int) []int {
	i := 0
	for j:=1; j<len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			if j>i {
				arr[i] = arr[j]
			}
		}
	}
	return arr[:i+1]
}
func sortedFloatsNoRepeat(arr []float64) []float64 {
	i := 0
	for j:=1; j<len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			if j>i {
				arr[i] = arr[j]
			}
		}
	}
	return arr[:i+1]
}
func sortedStringsNoRepeat(arr []string) []string {
	i := 0
	for j:=1; j<len(arr); j++ {
		if arr[i] != arr[j] {
			i++
			if j>i {
				arr[i] = arr[j]
			}
		}
	}
	return arr[:i+1]
}

// 使切片arr升序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func IntsAscAndNoRepeat(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	IntsAsc(arr)
	return sortedIntsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscIntsNoRepeat(arr []int, noCheck ...bool) []int {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !IntsIsAsc(arr) {
		panic("AscIntsNoRepeat param arr is not sorted asc")
	}
	return sortedIntsNoRepeat(arr)
}

// 使切片arr降序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func IntsDescAndNoRepeat(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	IntsDesc(arr)
	return sortedIntsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescIntsNoRepeat(arr []int, noCheck ...bool) []int {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !IntsIsDesc(arr) {
		panic("DescIntsNoRepeat param arr is not sorted desc")
	}
	return sortedIntsNoRepeat(arr)
}

// 使切片arr升序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func FloatsAscAndNoRepeat(arr []float64) []float64 {
	if len(arr) == 0 {
		return arr
	}
	FloatsAsc(arr)
	return sortedFloatsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscFloatsNoRepeat(arr []float64, noCheck ...bool) []float64 {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !FloatsIsAsc(arr) {
		panic("AscFloatsNoRepeat param arr is not sorted asc")
	}
	return sortedFloatsNoRepeat(arr)
}

// 使切片arr降序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func FloatsDescAndNoRepeat(arr []float64) []float64 {
	if len(arr) == 0 {
		return arr
	}
	FloatsDesc(arr)
	return sortedFloatsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescFloatsNoRepeat(arr []float64, noCheck ...bool) []float64 {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !FloatsIsDesc(arr) {
		panic("DescFloatsNoRepeat param arr is not sorted desc")
	}
	return sortedFloatsNoRepeat(arr)
}

// 使切片arr升序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func StringsAscAndNoRepeat(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	StringsAsc(arr)
	return sortedStringsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscStringsNoRepeat(arr []string, noCheck ...bool) []string {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !StringsIsAsc(arr) {
		panic("AscStringsNoRepeat param arr is not sorted asc")
	}
	return sortedStringsNoRepeat(arr)
}

// 使切片arr降序且元素唯一
// 返回新的有序切片
// 注意: 操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
func StringsDescAndNoRepeat(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	StringsDesc(arr)
	return sortedStringsNoRepeat(arr)
}

// 使有序切片arr中元素唯一, 重复的元素只保留一个, 多余的删除
// 返回新的有序切片
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行操作
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescStringsNoRepeat(arr []string, noCheck ...bool) []string {
	if len(arr) == 0 {
		return arr
	}else if (len(noCheck)==0 || !noCheck[0]) && !StringsIsDesc(arr) {
		panic("DescStringsNoRepeat param arr is not sorted desc")
	}
	return sortedStringsNoRepeat(arr)
}












