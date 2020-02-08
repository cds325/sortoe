# sortoe 快速排序

## 优点

    1. 没有使用接口及反射，排序速度较官方sort包快
    2. 可选择使用启用 goroutine 并发快速排序，速度进一步提升

## 缺点

    只支持三种确定的类型，即 []int，[]float64，[]string

## Benchmark

    Benchmark_sortgo_10000000-12                   6         173701667 ns/op           15962 B/op        423 allocs/op
    Benchmark_sortoe_10000000-12                   2         810331950 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_10000000-12                     1        2022574000 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_1000000-12                   74          16356507 ns/op            3918 B/op        201 allocs/op
    Benchmark_sortoe_1000000-12                   16          68880738 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_1000000-12                      6         172878433 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_100000-12                   535           2145060 ns/op            2188 B/op        190 allocs/op
    Benchmark_sortoe_100000-12                   214           5514237 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_100000-12                      74          14651291 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_10000-12                   6242            192520 ns/op             997 B/op         87 allocs/op
    Benchmark_sortoe_10000-12                   2718            433712 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_10000-12                      985           1177701 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_5000-12                   12051            103199 ns/op             816 B/op         71 allocs/op
    Benchmark_sortoe_5000-12                    5001            201647 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_5000-12                      2269            545954 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_2000-12                   25068             48695 ns/op             444 B/op         37 allocs/op
    Benchmark_sortoe_2000-12                   16689             69553 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_2000-12                      6684            200828 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_1000-12                   38083             34041 ns/op             203 B/op         17 allocs/op
    Benchmark_sortoe_1000-12                   36121             33329 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_1000-12                     12516             92753 ns/op              32 B/op          1 allocs/op
    Benchmark_sortgo_500-12                    82861             14240 ns/op              16 B/op          1 allocs/op
    Benchmark_sortoe_500-12                    84129             14118 ns/op               0 B/op          0 allocs/op
    Benchmark_sort_500-12                      26862             42993 ns/op              32 B/op          1 allocs/op

    sortgo为调用 sortoe.IntsAscGo()
    sortoe为调用 sortoe.IntsAsc()
    sort为调用 sort.Sort()

Download and install

    go get github.com/cds325/sortoe
    
Import

    import "github.com/cds325/sortoe"
    
简单到不需要使用说明^_^

