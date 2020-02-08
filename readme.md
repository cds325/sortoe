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
    
Useage
    
    以整型切片（[]int）为例
    
    sortoe包:
        func IntsAsc(arr []int)
            升序快速排序
        func IntsAscAndNoRepeat(arr []int) []int
            升序快速排序并去除重复元素，使其唯一
        func IntsAscByInsertion(arr []int)
            升序插入排序
        func IntsAscGo(arr []int, depth ...int)
            启用 goroutine 并发的升序快速排序
        func IntsAscPercent(arr []int) int
            判断升序占升序降序总数百分比，可用于判断是否基本有序
        func IntsDesc(arr []int)
            降序快速排序
        func IntsDescAndNoRepeat(arr []int) []int
            降序快速排序并去除重复元素，使其唯一
        func IntsDescByInsertion(arr []int)
            降序插入排序
        func IntsDescGo(arr []int, depth ...int)
            启用 goroutine 并发的升序快速排序
        func IntsIsAsc(arr []int) bool
            判断是否是完全升序
        func IntsIsAscAndUnique(arr []int) bool
            判断是否是完全升序且所有元素唯一
        func IntsIsDesc(arr []int) bool
            判断是否是完全降序
        func IntsIsDescAndUnique(arr []int) bool
            判断是否是完全降序且所有元素唯一
        func IntsIsFlat(arr []int) bool
            判断是否是所有元素相等
        func IntsReverse(arr []int)
            使数组反转
            
        func AscIntsCount(arr []int, x int, noCheck ...bool) int
            统计升序数组中某个元素的个数
        func AscIntsDelete(arr []int, old int, count ...int) ([]int, int)
            升序数组中删除值为 old 的元素
        func AscIntsDeleteNoCheck(arr []int, old int, count ...int) ([]int, int)
            升序数组中删除值为 old 的元素, 过程中不检验是否升序
        func AscIntsHas(arr []int, x int, noCheck ...bool) bool
            判断升序数组中是否含有值为 x 的元素
        func AscIntsInsert(arr []int, elems ...int) []int
            升序数组中插入元素, 并返回新的升序数组
        func AscIntsInsertNoCheck(arr []int, elems ...int) []int
            升序数组中插入元素, 并返回新的升序数组, 过程中不检验是否升序
        func AscIntsNoRepeat(arr []int, noCheck ...bool) []int
            去除升序数组中重复元素, 使元素唯一
        func AscIntsReplace(arr []int, old, new int, count ...int) int
            升序数组中替换元素, 并返回新的升序数组
        func AscIntsReplaceNoCheck(arr []int, old, new int, count ...int) int
            升序数组中替换元素, 并返回新的升序数组, 过程中不检验是否升序
        func AscIntsSearch(arr []int, x int, noCheck ...bool) (int, int, bool)
            升序数组中查找值为 x 的元素, 返回起始终止索引和true, 未找到返回查找结束位置和false
            
        func DescIntsCount(arr []int, x int, noCheck ...bool) int
        func DescIntsDelete(arr []int, old int, count ...int) ([]int, int)
        func DescIntsDeleteNoCheck(arr []int, old int, count ...int) ([]int, int)
        func DescIntsHas(arr []int, x int, noCheck ...bool) bool
        func DescIntsInsert(arr []int, elems ...int) []int
        func DescIntsInsertNoCheck(arr []int, elems ...int) []int
        func DescIntsNoRepeat(arr []int, noCheck ...bool) []int
        func DescIntsReplace(arr []int, old, new int, count ...int) int
        func DescIntsReplaceNoCheck(arr []int, old, new int, count ...int) int
        func DescIntsSearch(arr []int, x int, noCheck ...bool) (int, int, bool)
        
        func BitsCount(a int) int
            统计整数 a 中有效 bit 位数
        func EstimateDepth(lenth int, numCPU ...int) (depth int)
            估算启用 goroutine 快速排序的递归深度
            
    sortedu包:
        func UniqueAscIntsDelete(arr []int, elem int, noCheck ...bool) (newarr []int, existing bool)
            升序且元素唯一的数组中删除元素
        func UniqueAscIntsInsert(arr []int, elem int, noCheck ...bool) (newarr []int, existing bool)
            升序且元素唯一的数组中插入元素, 返回新数组仍保持升序且元素唯一
        func UniqueAscIntsReplace(arr []int, old, new int, noCheck ...bool) (newarr []int, oldExisting, newExisting bool)
            升序且元素唯一的数组中替换元素, 返回新数组仍保持升序且元素唯一
        func UniqueDescIntsDelete(arr []int, elem int, noCheck ...bool) (newarr []int, existing bool)
        func UniqueDescIntsInsert(arr []int, elem int, noCheck ...bool) (newarr []int, existing bool)
        func UniqueDescIntsReplace(arr []int, old, new int, noCheck ...bool) (newarr []int, oldExisting, newExisting bool)



简单到不需要使用说明^_^

