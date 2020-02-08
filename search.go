package sortoe


// SearchIntsAsc 在升序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
// SearchIntsAsc searches for x in ascending arr, returning the start position index and end position index of x and true.
// if you can't find it, the return value is where x should be inserted into a and false (in order to increase a). The return value can be len(a).
// note: if the arr is indeed ordered, make noCheck true, eliminating the order in which the arr is checked
func AscIntsSearch(arr []int, x int, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !IntsIsAsc(arr) {
		panic("AscIntsSearch param arr is not sorted asc")
	}
	l, r := 0, len(arr)-1
	if x < arr[0] {
		return 0,0,false
	}else if x > arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] < x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}

// SearchIntsDesc 在降序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescIntsSearch(arr []int, x int, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !IntsIsDesc(arr) {
		panic("DescIntsSearch param arr is not sorted desc")
	}
	l, r := 0, len(arr)-1
	if x > arr[0] {
		return 0,0,false
	}else if x < arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] > x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}


// SearchFloatsAsc 在升序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscFloatsSearch(arr []float64, x float64, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !FloatsIsAsc(arr) {
		panic("AscFloatsSearch param arr is not sorted asc")
	}
	l, r := 0, len(arr)-1
	if x < arr[0] {
		return 0,0,false
	}else if x > arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] < x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}


// SearchFloatsDesc 在降序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescFloatsSearch(arr []float64, x float64, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !FloatsIsDesc(arr) {
		panic("DescFloatsSearch param arr is not sorted desc")
	}
	l, r := 0, len(arr)-1
	if x > arr[0] {
		return 0,0,false
	}else if x < arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] > x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}


// SearchStringsAsc 在升序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscStringsSearch(arr []string, x string, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !StringsIsAsc(arr) {
		panic("AscStringsSearch param arr is not sorted asc")
	}
	l, r := 0, len(arr)-1
	if x < arr[0] {
		return 0,0,false
	}else if x > arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] < x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}

// SearchStringsDesc 在降序的arr中搜索x，返回x的起始位置索引和结束位置索引和true。
// 如果查找不到，返回值是x应该插入a的位置和false（以保证a的递增顺序），返回值可以是len(a)
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescStringsSearch(arr []string, x string, noCheck ...bool) (int,int,bool) {
	if len(arr) == 0 {
		return 0,0,false
	}else if (len(noCheck)==0 || !noCheck[0]) && !StringsIsDesc(arr) {
		panic("DescStringsSearch param arr is not sorted desc")
	}
	l, r := 0, len(arr)-1
	if x > arr[0] {
		return 0,0,false
	}else if x < arr[r] {
		return r+1,r+1,false
	}
	for l < r {
		m := int(uint(l+r) >> 1)
		if arr[m] == x {
			var i,j int
			for i=m; i>l && arr[i-1] == x; i-- {}
			for j=m; j<r && arr[j+1] == x; j++ {}
			return i,j,true
		}else if arr[m] > x {
			l = m + 1
		}else {
			r = m
		}
	}
	return l,l,false
}

// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscIntsCount(arr []int, x int, noCheck ...bool) int {
	i,j,ok := AscIntsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscIntsHas(arr []int, x int, noCheck ...bool) bool {
	_,_,ok := AscIntsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}

// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescIntsCount(arr []int, x int, noCheck ...bool) int {
	i,j,ok := DescIntsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescIntsHas(arr []int, x int, noCheck ...bool) bool {
	_,_,ok := DescIntsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}


// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscFloatsCount(arr []float64, x float64, noCheck ...bool) int {
	i,j,ok := AscFloatsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscFloatsHas(arr []float64, x float64, noCheck ...bool) bool {
	_,_,ok := AscFloatsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}

// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescFloatsCount(arr []float64, x float64, noCheck ...bool) int {
	i,j,ok := DescFloatsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescFloatsHas(arr []float64, x float64, noCheck ...bool) bool {
	_,_,ok := DescFloatsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}



// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscStringsCount(arr []string, x string, noCheck ...bool) int {
	i,j,ok := AscStringsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func AscStringsHas(arr []string, x string, noCheck ...bool) bool {
	_,_,ok := AscStringsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}

// 统计有序切片中指定的元素个数
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescStringsCount(arr []string, x string, noCheck ...bool) int {
	i,j,ok := DescStringsSearch(arr,x,noCheck...)
	if ok {
		return j-i+1
	}
	return 0
}

// 判断有序切片中是否包含指定的元素
// 注意: 如果arr确实有序, 可令noCheck为true, 省去检查arr的顺序
func DescStringsHas(arr []string, x string, noCheck ...bool) bool {
	_,_,ok := DescStringsSearch(arr,x,noCheck...)
	if ok {
		return true
	}
	return false
}
