package sortoe

import (
	"math/rand"
	"sort"
	"testing"
)

var tmp []int

func init(){
	tmp = make([]int, 10000000)
	for i:=0;i<10000000;i++ {
		tmp[i] = rand.Intn(100000000)
	}
}

// 在当前目录命令行下执行 go test -bench=. -benchmem

func benchmark_sortoe(b *testing.B,size int){

	for i:=0;i<b.N;i++ {
		b.StopTimer()
		arr := make([]int, size)
		for i:=0;i<size;i++ {
			arr[i] = tmp[i]
		}
		b.StartTimer()
		IntsAsc(arr)
	}

}

func benchmark_sortgo(b *testing.B,size int){
	//Depth := 0
	for i:=0;i<b.N;i++ {
		b.StopTimer()
		arr := make([]int, size)
		for i:=0;i<size;i++ {
			arr[i] = tmp[i]
		}
		b.StartTimer()
		//Depth = IntsAscGo(arr)
		IntsAscGo(arr)
	}
	//fmt.Printf("[%d]",Depth)
}

func benchmark_sort(b *testing.B,size int){

	for i:=0;i<b.N;i++ {
		b.StopTimer()
		arr := make(sort.IntSlice, size)
		for i:=0;i<size;i++ {
			arr[i] = tmp[i]
		}
		b.StartTimer()
		sort.Sort(arr)
	}

}

func Benchmark_sortgo_10000000(b *testing.B ){
	benchmark_sortgo(b,10000000)
}

func Benchmark_sortoe_10000000(b *testing.B ){
	benchmark_sortoe(b,10000000)
}

func Benchmark_sort_10000000(b *testing.B ){
	benchmark_sort(b,10000000)
}

func Benchmark_sortgo_1000000(b *testing.B ){
	benchmark_sortgo(b,1000000)
}

func Benchmark_sortoe_1000000(b *testing.B ){
	benchmark_sortoe(b,1000000)
}

func Benchmark_sort_1000000(b *testing.B ){
	benchmark_sort(b,1000000)
}

func Benchmark_sortgo_100000(b *testing.B ){
	benchmark_sortgo(b,100000)
}

func Benchmark_sortoe_100000(b *testing.B ){
	benchmark_sortoe(b,100000)
}

func Benchmark_sort_100000(b *testing.B ){
	benchmark_sort(b,100000)
}

func Benchmark_sortgo_10000(b *testing.B ){
	benchmark_sortgo(b,10000)
}

func Benchmark_sortoe_10000(b *testing.B ){
	benchmark_sortoe(b,10000)
}

func Benchmark_sort_10000(b *testing.B ){
	benchmark_sort(b,10000)
}

func Benchmark_sortgo_5000(b *testing.B ){
	benchmark_sortgo(b,5000)
}

func Benchmark_sortoe_5000(b *testing.B ){
	benchmark_sortoe(b,5000)
}

func Benchmark_sort_5000(b *testing.B ){
	benchmark_sort(b,5000)
}

func Benchmark_sortgo_2000(b *testing.B ){
	benchmark_sortgo(b,2000)
}

func Benchmark_sortoe_2000(b *testing.B ){
	benchmark_sortoe(b,2000)
}

func Benchmark_sort_2000(b *testing.B ){
	benchmark_sort(b,2000)
}

func Benchmark_sortgo_1000(b *testing.B ){
	benchmark_sortgo(b,1000)
}

func Benchmark_sortoe_1000(b *testing.B ){
	benchmark_sortoe(b,1000)
}

func Benchmark_sort_1000(b *testing.B ){
	benchmark_sort(b,1000)
}

func Benchmark_sortgo_500(b *testing.B ){
	benchmark_sortgo(b,500)
}

func Benchmark_sortoe_500(b *testing.B ){
	benchmark_sortoe(b,500)
}

func Benchmark_sort_500(b *testing.B ){
	benchmark_sort(b,500)
}

//func Benchmark_sortgo_80(b *testing.B ){
//	benchmark_sortgo(b,80)
//}
//
//func Benchmark_sortoe_80(b *testing.B ){
//	benchmark_sortoe(b,80)
//}
//
//func Benchmark_sort_80(b *testing.B ){
//	benchmark_sort(b,80)
//}




func benchmark_norepeat(b *testing.B,size int){

	for i:=0;i<b.N;i++ {
		b.StopTimer()
		arr := make([]int, size)
		for i:=0;i<size;i++ {
			arr[i] = tmp[i]
		}
		IntsAsc(arr)
		b.StartTimer()
		AscIntsNoRepeat(arr)
	}

}

//func Benchmark_norepeat_10000000(b *testing.B ){
//	benchmark_norepeat(b,10000000)
//}
//

