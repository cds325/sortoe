package sortoe

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscIntsReplace(arr []int,old,new int,count ...int) int {
	return ascIntsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscIntsReplaceNoCheck(arr []int,old,new int,count ...int) int {
	return ascIntsReplace(arr,old,new,true,count ...)
}

func ascIntsReplace(arr []int,old,new int,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := AscIntsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new < old {
			for ; i>0 && arr[i-1]>new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new > old {
			for ; j<len(arr)-1 && arr[j+1]<new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescIntsReplace(arr []int,old,new int,count ...int) int {
	return descIntsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescIntsReplaceNoCheck(arr []int,old,new int,count ...int) int {
	return descIntsReplace(arr,old,new,true,count ...)
}

func descIntsReplace(arr []int,old,new int,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := DescIntsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new > old {
			for ; i>0 && arr[i-1]<new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new < old {
			for ; j<len(arr)-1 && arr[j+1]>new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}


// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscFloatsReplace(arr []float64,old,new float64,count ...int) int {
	return ascFloatsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscFloatsReplaceNoCheck(arr []float64,old,new float64,count ...int) int {
	return ascFloatsReplace(arr,old,new,true,count ...)
}

func ascFloatsReplace(arr []float64,old,new float64,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := AscFloatsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new < old {
			for ; i>0 && arr[i-1]>new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new > old {
			for ; j<len(arr)-1 && arr[j+1]<new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescFloatsReplace(arr []float64,old,new float64,count ...int) int {
	return descFloatsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescFloatsReplaceNoCheck(arr []float64,old,new float64,count ...int) int {
	return descFloatsReplace(arr,old,new,true,count ...)
}

func descFloatsReplace(arr []float64,old,new float64,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := DescFloatsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new > old {
			for ; i>0 && arr[i-1]<new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new < old {
			for ; j<len(arr)-1 && arr[j+1]>new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscStringsReplace(arr []string,old,new string,count ...int) int {
	return ascStringsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func AscStringsReplaceNoCheck(arr []string,old,new string,count ...int) int {
	return ascStringsReplace(arr,old,new,true,count ...)
}

func ascStringsReplace(arr []string,old,new string,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := AscStringsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new < old {
			for ; i>0 && arr[i-1]>new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new > old {
			for ; j<len(arr)-1 && arr[j+1]<new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescStringsReplace(arr []string,old,new string,count ...int) int {
	return descStringsReplace(arr,old,new,false,count ...)
}

// 将有序切片arr中值为old的元素替换为new, 最多替换count个元素
// 返回实际被替换的元素个数
func DescStringsReplaceNoCheck(arr []string,old,new string,count ...int) int {
	return descStringsReplace(arr,old,new,true,count ...)
}

func descStringsReplace(arr []string,old,new string,noCheck bool,count ...int) int {
	if new == old {
		return 0
	}
	i,j,ok := DescStringsSearch(arr,old,noCheck)
	if ok {
		n := j-i+1
		if len(count)>0 && count[0] >0 && count[0]<n {
			n = count[0]
			j = i+n-1
		}
		if new > old {
			for ; i>0 && arr[i-1]<new; i-- {
				arr[j] = arr[i-1]
				j--
			}
		}else if new < old {
			for ; j<len(arr)-1 && arr[j+1]>new; j++ {
				arr[i] = arr[j+1]
				i++
			}
		}
		for ; i <= j; i++ {
			arr[i] = new
		}
		return n
	}
	return 0
}












