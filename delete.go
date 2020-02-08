package sortoe


func sortedIntsDelete(arr []int,i,j int,ok bool,count ...int) ([]int,int) {
	if ok {
		n := j-i+1
		if len(count)>0 && count[0]>0 && count[0]<n {
			n = count[0]
			i = j-n+1
		}
		for ; j<len(arr)-1; j++ {
			arr[i] = arr[j+1]
			i++
		}
		return arr[0:i],n
	}
	return arr,0
}

func sortedFloatsDelete(arr []float64,i,j int,ok bool,count ...int) ([]float64,int) {
	if ok {
		n := j-i+1
		if len(count)>0 && count[0]>0 && count[0]<n {
			n = count[0]
			i = j-n+1
		}
		for ; j<len(arr)-1; j++ {
			arr[i] = arr[j+1]
			i++
		}
		return arr[0:i],n
	}
	return arr,0
}

func sortedStringsDelete(arr []string,i,j int,ok bool,count ...int) ([]string,int) {
	if ok {
		n := j-i+1
		if len(count)>0 && count[0]>0 && count[0]<n {
			n = count[0]
			i = j-n+1
		}
		for ; j<len(arr)-1; j++ {
			arr[i] = arr[j+1]
			i++
		}
		return arr[0:i],n
	}
	return arr,0
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscIntsDelete(arr []int,old int,count ...int) ([]int,int) {
	i,j,ok := AscIntsSearch(arr,old)
	return sortedIntsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscIntsDeleteNoCheck(arr []int,old int,count ...int) ([]int,int) {
	i,j,ok := AscIntsSearch(arr,old,true)
	return sortedIntsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescIntsDelete(arr []int,old int,count ...int) ([]int,int) {
	i,j,ok := DescIntsSearch(arr,old)
	return sortedIntsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescIntsDeleteNoCheck(arr []int,old int,count ...int) ([]int,int) {
	i,j,ok := DescIntsSearch(arr,old,true)
	return sortedIntsDelete(arr,i,j,ok,count ...)
}


// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscFloatsDelete(arr []float64,old float64,count ...int) ([]float64,int) {
	i,j,ok := AscFloatsSearch(arr,old)
	return sortedFloatsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscFloatsDeleteNoCheck(arr []float64,old float64,count ...int) ([]float64,int) {
	i,j,ok := AscFloatsSearch(arr,old,true)
	return sortedFloatsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescFloatsDelete(arr []float64,old float64,count ...int) ([]float64,int) {
	i,j,ok := DescFloatsSearch(arr,old)
	return sortedFloatsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescFloatsDeleteNoCheck(arr []float64,old float64,count ...int) ([]float64,int) {
	i,j,ok := DescFloatsSearch(arr,old,true)
	return sortedFloatsDelete(arr,i,j,ok,count ...)
}


// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscStringsDelete(arr []string,old string,count ...int) ([]string,int) {
	i,j,ok := AscStringsSearch(arr,old)
	return sortedStringsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func AscStringsDeleteNoCheck(arr []string,old string,count ...int) ([]string,int) {
	i,j,ok := AscStringsSearch(arr,old,true)
	return sortedStringsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescStringsDelete(arr []string,old string,count ...int) ([]string,int) {
	i,j,ok := DescStringsSearch(arr,old)
	return sortedStringsDelete(arr,i,j,ok,count ...)
}

// 将有序切片arr中值为old的元素删除, 最多删除count个元素
// 返回新的有序切片和实际被删除的元素个数
// 注意: 删除操作后会改变原切片, 如果需要保留原切片, 则应对其copy副本执行删除
func DescStringsDeleteNoCheck(arr []string,old string,count ...int) ([]string,int) {
	i,j,ok := DescStringsSearch(arr,old,true)
	return sortedStringsDelete(arr,i,j,ok,count ...)
}
