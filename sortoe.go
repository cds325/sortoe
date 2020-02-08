package sortoe

// 整型切片升序插入排序
// integer slice ascending insertion sort
func IntsAscByInsertion(arr []int){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp<arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 整型切片降序插入排序
// integer slice descending insertion sort
func IntsDescByInsertion(arr []int){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp>arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 整型取中位数
func getIntsMedian(arr ... int) int  {
	IntsAscByInsertion(arr)
	return arr[len(arr)/2]
}
// 整型切片升序快速排序
// integer slice ascending quicksort
func IntsAsc(arr []int) {
	if len(arr) <= 80 {
		IntsAscByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp int
	temp= getIntsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] < temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] > temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	IntsAsc(arr[0:l])
	IntsAsc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 整型切片降序快速排序
// integer slice descending quicksort
func IntsDesc(arr []int) {
	if len(arr) <= 80 {
		IntsDescByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp int
	temp= getIntsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] > temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] < temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	IntsDesc(arr[0:l])
	IntsDesc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 整型切片顺序反转
// reverse the order of integer slice
func IntsReverse(arr []int){
	if len(arr)<2{
		return
	}
	for i:=0;i<=len(arr)/2;i++{
		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	}
}

// 浮点型切片升序插入排序
func FloatsAscByInsertion(arr []float64){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp<arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 浮点型切片降序插入排序
func FloatsDescByInsertion(arr []float64){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp>arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 浮点型取中位数
func getFloatsMedian(arr ... float64) float64  {
	FloatsAscByInsertion(arr)
	return arr[len(arr)/2]
}
// 浮点型切片升序快速排序
func FloatsAsc(arr []float64) {
	if len(arr) <= 80 {
		FloatsAscByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp float64
	temp= getFloatsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] < temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] > temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	FloatsAsc(arr[0:l])
	FloatsAsc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 浮点型切片降序快速排序
func FloatsDesc(arr []float64) {
	if len(arr) <= 80 {
		FloatsDescByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp float64
	temp= getFloatsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] > temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] < temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	FloatsDesc(arr[0:l])
	FloatsDesc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 浮点型切片顺序反转
func FloatsReverse(arr []float64){
	if len(arr)<2{
		return
	}
	for i:=0;i<=len(arr)/2;i++{
		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	}
}

// 字符串切片升序插入排序
func StringsAscByInsertion(arr []string){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp<arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 字符串切片降序插入排序
func StringsDescByInsertion(arr []string){
	if len(arr)<2{
		return
	}
	for i:=1;i<len(arr);i++{
		temp := arr[i]
		j:=i
		for ;j>0 && temp>arr[j-1];j--{
			arr[j] = arr[j-1]
		}
		if j<i {
			arr[j] = temp
		}
	}
}
// 字符串取中位数
func getStringsMedian(arr ... string) string  {
	StringsAscByInsertion(arr)
	return arr[len(arr)/2]
}
// 字符串切片升序快速排序
func StringsAsc(arr []string) {
	if len(arr) <= 80 {
		StringsAscByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp string
	temp= getStringsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] < temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] > temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	StringsAsc(arr[0:l])
	StringsAsc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 字符串切片降序快速排序
func StringsDesc(arr []string) {
	if len(arr) <= 80 {
		StringsDescByInsertion(arr)
		return
	}
	l, r := 0, len(arr)-1
	li, ri := l, r
	var temp string
	temp= getStringsMedian(arr[0], arr[r], arr[r/6], arr[2*r/6], arr[3*r/6], arr[4*r/6], arr[5*r/6])
	for li <= ri {
		if arr[li] > temp {
			if l<li {
				arr[l] = arr[li]
			}
			l++
			li++
		}else if arr[li] == temp{
			li++
		}else{
			for li <= ri {
				if arr[ri] < temp {
					if r>ri {
						arr[r] = arr[ri]
					}
					r--
					ri--
				}else if arr[ri] == temp {
					ri--
				}else{
					arr[l], arr[r] = arr[ri], arr[li]
					l++
					li++
					r--
					ri--
					break
				}
			}
		}
	}
	StringsDesc(arr[0:l])
	StringsDesc(arr[r+1:])
	for ;l<=r;l++ {
		arr[l] = temp
	}
}
// 字符串切片顺序反转
func StringsReverse(arr []string){
	if len(arr)<2{
		return
	}
	for i:=0;i<=len(arr)/2;i++{
		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	}
}


