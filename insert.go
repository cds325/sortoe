package sortoe

// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscIntsInsert(arr []int,elems ...int) []int {
	if !IntsIsAsc(arr) {
		panic("AscIntsInsert param arr is not sorted asc")
	}
	return AscIntsInsertNoCheck(arr,elems ...)
}

// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscIntsInsertNoCheck(arr []int,elems ...int) []int {
	newarr := make([]int,len(arr)+len(elems))
	IntsAsc(elems)
	x,y,z := 0,0,0
	for ; x<len(newarr); x++ {
		if y<len(arr) && z<len(elems) {
			if elems[z] < arr[y] {
				newarr[x] = elems[z]
				z++
			}else{
				newarr[x] = arr[y]
				y++
			}
		}else if y<len(arr) {
			newarr[x] = arr[y]
			y++
		}else if z<len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescIntsInsert(arr []int,elems ...int) []int {
	if !IntsIsDesc(arr) {
		panic("DescIntsInsert param arr is not sorted asc")
	}
	return DescIntsInsertNoCheck(arr,elems ...)
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescIntsInsertNoCheck(arr []int,elems ...int) []int {
	newarr := make([]int,len(arr)+len(elems))
	IntsDesc(elems)
	x,y,z := 0,0,0
	for ; x<len(newarr); x++ {
		if y<len(arr) && z<len(elems) {
			if elems[z] > arr[y] {
				newarr[x] = elems[z]
				z++
			}else{
				newarr[x] = arr[y]
				y++
			}
		}else if y<len(arr) {
			newarr[x] = arr[y]
			y++
		}else if z<len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}


// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscFloatsInsert(arr []float64,elems ...float64) []float64 {
	if !FloatsIsAsc(arr) {
		panic("AscFloatsInsert param arr is not sorted asc")
	}
	return AscFloatsInsertNoCheck(arr,elems ...)
}

// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscFloatsInsertNoCheck(arr []float64,elems ...float64) []float64 {
	newarr := make([]float64,len(arr)+len(elems))
	FloatsAsc(elems)
	x,y,z := 0,0,0
	for ; x<len(newarr); x++ {
		if y<len(arr) && z<len(elems) {
			if elems[z] < arr[y] {
				newarr[x] = elems[z]
				z++
			}else{
				newarr[x] = arr[y]
				y++
			}
		}else if y<len(arr) {
			newarr[x] = arr[y]
			y++
		}else if z<len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescFloatsInsert(arr []float64,elems ...float64) []float64 {
	if !FloatsIsDesc(arr) {
		panic("DescFloatsInsert param arr is not sorted asc")
	}
	return DescFloatsInsertNoCheck(arr,elems ...)
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescFloatsInsertNoCheck(arr []float64,elems ...float64) []float64 {
	newarr := make([]float64,len(arr)+len(elems))
	FloatsDesc(elems)
	x,y,z := 0,0,0
	for ; x<len(newarr); x++ {
		if y<len(arr) && z<len(elems) {
			if elems[z] > arr[y] {
				newarr[x] = elems[z]
				z++
			}else{
				newarr[x] = arr[y]
				y++
			}
		}else if y<len(arr) {
			newarr[x] = arr[y]
			y++
		}else if z<len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}


// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscStringsInsert(arr []string,elems ...string) []string {
	if !StringsIsAsc(arr) {
		panic("AscStringsInsert param arr is not sorted asc")
	}
	return AscStringsInsertNoCheck(arr,elems ...)
}

// 向升序切片arr插入新的元素, 返回新的升序切片
// 注意: 不会改变原切片
func AscStringsInsertNoCheck(arr []string,elems ...string) []string {
	newarr := make([]string,len(arr)+len(elems))
	StringsAsc(elems)
	x,y,z := 0,0,0
	for ; x<len(newarr); x++ {
		if y<len(arr) && z<len(elems) {
			if elems[z] < arr[y] {
				newarr[x] = elems[z]
				z++
			}else{
				newarr[x] = arr[y]
				y++
			}
		}else if y<len(arr) {
			newarr[x] = arr[y]
			y++
		}else if z<len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescStringsInsert(arr []string,elems ...string) []string {
	if !StringsIsDesc(arr) {
		panic("DescStringsInsert param arr is not sorted asc")
	}
	return DescStringsInsertNoCheck(arr,elems ...)
}

// 向降序切片arr插入新的元素, 返回新的降序切片
// 注意: 不会改变原切片
func DescStringsInsertNoCheck(arr []string,elems ...string) []string {
	newarr := make([]string, len(arr)+len(elems))
	StringsDesc(elems)
	x, y, z := 0, 0, 0
	for ; x < len(newarr); x++ {
		if y < len(arr) && z < len(elems) {
			if elems[z] > arr[y] {
				newarr[x] = elems[z]
				z++
			} else {
				newarr[x] = arr[y]
				y++
			}
		} else if y < len(arr) {
			newarr[x] = arr[y]
			y++
		} else if z < len(elems) {
			newarr[x] = elems[z]
			z++
		}
	}
	return newarr
}