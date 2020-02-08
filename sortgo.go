package sortoe

import (
	"runtime"
	"sync"
)

// 统计比特位数
func BitsCount(a int) int {
	var c int
	for a>0 {
		c++
		a>>=1
	}
	return c
}

func EstimateDepth(lenth int,numCPU ...int) int {
	var numCpu int
	if len(numCPU)>0 && numCPU[0]>0 {
		numCpu = numCPU[0]
	}else {
		numCpu = runtime.NumCPU()
	}
	return BitsCount((numCpu-1)*BitsCount(lenth>>9)/2)
}

// 整型切片升序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// integer slice ascending quicksort, allowing goroutine accelerated sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func IntsAscGo(arr []int,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	intsAscGo(arr,Depth)
}
func intsAscGo(arr []int,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			intsAscGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			intsAscGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		IntsAsc(arr[0:l])
		IntsAsc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}

// 整型切片降序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// integer slice descending quicksort, allowing goroutine accelerated sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func IntsDescGo(arr []int,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	intsDescGo(arr,Depth)
}
func intsDescGo(arr []int,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			intsDescGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			intsDescGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		IntsDesc(arr[0:l])
		IntsDesc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}


// 浮点型切片升序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// float slice ascending quicksort, allowing goroutine accelerated sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func FloatsAscGo(arr []float64,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	floatsAscGo(arr,Depth)
}
func floatsAscGo(arr []float64,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			floatsAscGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			floatsAscGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		FloatsAsc(arr[0:l])
		FloatsAsc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}

// 浮点型切片降序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// float slice descending quicksort, allowing goroutine accelerated sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func FloatsDescGo(arr []float64,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	floatsDescGo(arr,Depth)
}
func floatsDescGo(arr []float64,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			floatsDescGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			floatsDescGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		FloatsDesc(arr[0:l])
		FloatsDesc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}



// 字符串型切片升序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// string slice ascending quicksort, allowing goroutine accelerated sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func StringsAscGo(arr []string,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	stringsAscGo(arr,Depth)
}
func stringsAscGo(arr []string,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			stringsAscGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			stringsAscGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		StringsAsc(arr[0:l])
		StringsAsc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}

// 字符串型切片降序快速排序, 允许开启 goroutine 加速排序,
// 传入 depth 指定 goroutine 递归深度, 否则自动计算可能为较合适的深度.
// string slice descending quicksort, allowing goroutine to accelerate sorting,
// pass in depth to specify the recursive depth of goroutine, otherwise the automatic calculation may be more appropriate depth.
func StringsDescGo(arr []string,depth ...int){
	var Depth int
	if len(depth)>0 && depth[0]>=0 {
		Depth = depth[0]
	}else{
		Depth = EstimateDepth(len(arr))
	}
	stringsDescGo(arr,Depth)
}
func stringsDescGo(arr []string,depth int) {
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
	if depth >0 {
		depth--
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			stringsDescGo(arr[0:l],depth)
			wg.Done()
		}()
		go func() {
			stringsDescGo(arr[r+1:],depth)
			wg.Done()
		}()
		for ;l<=r;l++ {
			arr[l] = temp
		}
		wg.Wait()
	}else{
		StringsDesc(arr[0:l])
		StringsDesc(arr[r+1:])
		for ;l<=r;l++ {
			arr[l] = temp
		}
	}
}


