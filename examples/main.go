package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/cds325/sortoe"
	"github.com/cds325/sortoe/sortedu"
)

var (
	intsArr []int
	floatsArr []float64
	stringsArr []string
	n int
)

func init(){
	n = 300
	intsArr = make([]int,n)
	floatsArr = make([]float64,n)
	stringsArr = make([]string,n)
	for i := 0; i<n; i++ {
		r := 100+rand.Intn(99)
		intsArr[i] = r
		floatsArr[i] = float64(r) +0.1
		stringsArr[i] = strconv.Itoa(r)
	}
	fmt.Println("the template:\n",intsArr)

	sortoe.IntsDescByInsertion([]int{1,2,3})
}


func main(){

	ints := make([]int,n)
	floats := make([]float64,n)
	strings := make([]string,n)

	// 升序插入排序
	copy(ints,intsArr[:])
	copy(floats,floatsArr[:])
	copy(strings,stringsArr[:])
	sortoe.IntsAscByInsertion(ints)
	sortoe.FloatsAscByInsertion(floats)
	sortoe.StringsAscByInsertion(strings)
	fmt.Println("IntsAscInsertion(ints):\n",ints)
	fmt.Println("FloatsAscInsertion(floats):\n",floats)
	fmt.Println("StringsAscInsertion(strings):\n",strings)

	// 降序插入排序
	copy(ints,intsArr[:])
	copy(floats,floatsArr[:])
	copy(strings,stringsArr[:])
	sortoe.IntsDescByInsertion(ints)
	sortoe.FloatsDescByInsertion(floats)
	sortoe.StringsDescByInsertion(strings)
	fmt.Println("IntsDescInsertion(ints):\n",ints)
	fmt.Println("FloatsDescInsertion(floats):\n",floats)
	fmt.Println("StringsDescInsertion(strings):\n",strings)

	// 升序快速排序
	copy(ints,intsArr[:])
	copy(floats,floatsArr[:])
	copy(strings,stringsArr[:])
	sortoe.IntsAsc(ints)
	sortoe.FloatsAsc(floats)
	sortoe.StringsAsc(strings)
	fmt.Println("IntsAsc(ints):\n",ints)
	fmt.Println("FloatsAsc(floats):\n",floats)
	fmt.Println("StringsAsc(strings):\n",strings)

	// 有序的升序切片查找
	find1,find2,ok := sortoe.AscIntsSearch(ints,145)
	fmt.Println("SearchAscInts(ints,102):\n",find1,find2,ok)
	find1,find2,ok = sortoe.AscFloatsSearch(floats,102.1)
	fmt.Println("SearchAscFloats(floats,102.1):\n",find1,find2,ok)
	find1,find2,ok = sortoe.AscStringsSearch(strings,"102")
	fmt.Println("SearchAscFloats(strings,\"102\"):\n",find1,find2,ok)
	have := sortoe.AscIntsHas(ints,100)
	fmt.Println("AscIntsHas(ints,100):\n",have)
	have = sortoe.AscFloatsHas(floats,100.1)
	fmt.Println("AscFloatsHas(floats,100.1):\n",have)
	have = sortoe.AscStringsHas(strings,"100")
	fmt.Println("AscStringsHave(strings,\"100\"):\n",have)
	count := sortoe.AscIntsCount(ints,100)
	fmt.Println("AscIntsCount(ints,100):\n",count)
	count = sortoe.AscFloatsCount(floats,100.1)
	fmt.Println("AscFloatsCount(floats,100.1):\n",count)
	count = sortoe.AscStringsCount(strings,"100")
	fmt.Println("AscStringsCount(strings,\"100\"):\n",count)

	// 有序的升序切片插入
	newints := sortoe.AscIntsInsert(ints,11,145,222)
	fmt.Println("AscIntsInsert(ints,11,145,222):\n",newints)
	newfloats := sortoe.AscFloatsInsert(floats,11.1,145.1,222.1)
	fmt.Println("AscFloatsInsert(floats,11.1,145.1,222.1):\n",newfloats)
	newstrings := sortoe.AscStringsInsert(strings,"011","145","222")
	fmt.Println("AscStringsInsert(strings,\"011\",\"145\",\"222\"):\n",newstrings)

	// 有序的升序切片更新
	cnt := sortoe.AscIntsReplace(newints,107,12,3)
	fmt.Println("AscIntsReplace(newints,107,12,3):\n",cnt,newints)
	cnt = sortoe.AscFloatsReplace(newfloats,107.1,12.1,3)
	fmt.Println("AscFloatsReplace(newfloats,107.1,12.1,3):\n",cnt,newfloats)
	cnt = sortoe.AscStringsReplace(newstrings,"107","012",3)
	fmt.Println("AscStringsReplace(newstrings,\"107\",\"012\",3):\n",cnt,newstrings)

	// 有序的升序切片删除
	newints,cnt = sortoe.AscIntsDelete(newints,12,2)
	fmt.Println("AscIntsDelete(newints,12,2):\n",cnt,newints)
	newfloats,cnt = sortoe.AscFloatsDelete(newfloats,12.1,2)
	fmt.Println("AscFloatsDelete(newfloats,12.1,2):\n",cnt,newfloats)
	newstrings,cnt = sortoe.AscStringsDelete(newstrings,"012",2)
	fmt.Println("AscStringsDelete(newstrings,\"012\",2):\n",cnt,newstrings)

	// 有序的升序切片去重
	newints = sortoe.AscIntsNoRepeat(newints)
	fmt.Println("AscIntsNoRepeat(newints):\n",newints)
	newfloats = sortoe.AscFloatsNoRepeat(newfloats)
	fmt.Println("AscFloatsNoRepeat(newfloats):\n",newfloats)
	newstrings = sortoe.AscStringsNoRepeat(newstrings)
	fmt.Println("AscStringsNoRepeat(newstrings):\n",newstrings)

	newints,_ = sortedu.UniqueAscIntsInsert(newints,33)
	fmt.Println("sortedu.UniqueAscIntsInsert(newints,33):\n",newints)
	newints,_ = sortedu.UniqueAscIntsDelete(newints,33)
	fmt.Println("sortedu.UniqueAscIntsDelete(newints,33):\n",newints)
	newints,_,_ = sortedu.UniqueAscIntsReplace(newints,33,44)
	fmt.Println("sortedu.UniqueAscIntsReplace(newints,33,44):\n",newints)

	newfloats,_ = sortedu.UniqueAscFloatsInsert(newfloats,33.1)
	fmt.Println("sortedu.UniqueAscFloatsInsert(newfloats,33.1):\n",newfloats)
	newfloats,_ = sortedu.UniqueAscFloatsDelete(newfloats,33.1)
	fmt.Println("sortedu.UniqueAscFloatsDelete(newfloats,33.1):\n",newfloats)
	newfloats,_,_ = sortedu.UniqueAscFloatsReplace(newfloats,33.1,44.1)
	fmt.Println("sortedu.UniqueAscFloatsReplace(newfloats,33.1,44.1):\n",newfloats)

	newstrings,_ = sortedu.UniqueAscStringsInsert(newstrings,"33")
	fmt.Println("sortedu.UniqueAscStringsInsert(newstrings,\"33\"):\n",newstrings)
	newstrings,_ = sortedu.UniqueAscStringsDelete(newstrings,"33")
	fmt.Println("sortedu.UniqueAscStringsDelete(newstrings,\"33\"):\n",newstrings)
	newstrings,_,_ = sortedu.UniqueAscStringsReplace(newstrings,"33","44")
	fmt.Println("sortedu.UniqueAscStringsReplace(newstrings,\"33\",\"44\"):\n",newstrings)

	// 降序快速排序
	copy(ints,intsArr[:])
	copy(floats,floatsArr[:])
	copy(strings,stringsArr[:])
	sortoe.IntsDesc(ints)
	sortoe.FloatsDesc(floats)
	sortoe.StringsDesc(strings)
	fmt.Println("IntsDesc(ints):\n",ints)
	fmt.Println("FloatsDesc(floats):\n",floats)
	fmt.Println("StringsDesc(strings):\n",strings)


	// 有序的降序切片查找
	find1,find2,ok = sortoe.DescIntsSearch(ints,195)
	fmt.Println("SearchDescInts(ints,195):\n",find1,find2,ok)
	find1,find2,ok = sortoe.DescFloatsSearch(floats,195.1)
	fmt.Println("SearchDescFloats(floats,195.1):\n",find1,find2,ok)
	find1,find2,ok = sortoe.DescStringsSearch(strings,"195")
	fmt.Println("SearchDescFloats(strings,\"195\"):\n",find1,find2,ok)
	have = sortoe.DescIntsHas(ints,198)
	fmt.Println("DescIntsHas(ints,198):\n",have)
	have = sortoe.DescFloatsHas(floats,198.1)
	fmt.Println("DescFloatsHas(floats,198.1):\n",have)
	have = sortoe.DescStringsHas(strings,"198")
	fmt.Println("DescStringsHas(strings,\"198\"):\n",have)
	count = sortoe.DescIntsCount(ints,198)
	fmt.Println("DescIntsCount(ints,198):\n",count)
	count = sortoe.DescFloatsCount(floats,198.1)
	fmt.Println("DescFloatsCount(floats,198.1):\n",count)
	count = sortoe.DescStringsCount(strings,"198")
	fmt.Println("DescStringsCount(strings,\"198\"):\n",count)

	// 有序的降序切片插入
	newints = sortoe.DescIntsInsert(ints,11,145,222)
	fmt.Println("DescIntsInsert(ints,11,145,222):\n",newints)
	newfloats = sortoe.DescFloatsInsert(floats,11.1,145.1,222.1)
	fmt.Println("DescFloatsInsert(floats,11.1,145.1,222.1):\n",newfloats)
	newstrings = sortoe.DescStringsInsert(strings,"011","145","222")
	fmt.Println("DescStringsInsert(strings,\"011\",\"145\",\"222\"):\n",newstrings)

	// 有序的降序切片更新
	cnt = sortoe.DescIntsReplace(newints,196,201,3)
	fmt.Println("DescIntsReplace(newints,196,201,3):\n",cnt,newints)
	cnt = sortoe.DescFloatsReplace(newfloats,196.1,201.1,3)
	fmt.Println("DescFloatsReplace(newfloats,196.1,201.1,3):\n",cnt,newfloats)
	cnt = sortoe.DescStringsReplace(newstrings,"196","201",3)
	fmt.Println("DescStringsReplace(newstrings,\"196\",\"201\",3):\n",cnt,newstrings)

	// 有序的降序切片删除
	newints,cnt = sortoe.DescIntsDelete(newints,201,2)
	fmt.Println("DescIntsDelete(newints,201,2):\n",cnt,newints)
	newfloats,cnt = sortoe.DescFloatsDelete(newfloats,201.1,2)
	fmt.Println("DescFloatsDelete(newfloats,201.1,2):\n",cnt,newfloats)
	newstrings,cnt = sortoe.DescStringsDelete(newstrings,"201",2)
	fmt.Println("DescStringsDelete(newstrings,\"201\",2):\n",cnt,newstrings)

	// 有序的降序切片去重
	newints = sortoe.DescIntsNoRepeat(newints)
	fmt.Println("DescIntsNoRepeat(newints):\n",newints)
	newfloats = sortoe.DescFloatsNoRepeat(newfloats)
	fmt.Println("DescFloatsNoRepeat(newfloats):\n",newfloats)
	newstrings = sortoe.DescStringsNoRepeat(newstrings)
	fmt.Println("DescStringsNoRepeat(newstrings):\n",newstrings)

	newints,_ = sortedu.UniqueDescIntsInsert(newints,33)
	fmt.Println("sortedu.UniqueDescIntsInsert(newints,33):\n",newints)
	newints,_ = sortedu.UniqueDescIntsDelete(newints,33)
	fmt.Println("sortedu.UniqueDescIntsDelete(newints,33):\n",newints)
	newints,_,_ = sortedu.UniqueDescIntsReplace(newints,33,44)
	fmt.Println("sortedu.UniqueDescIntsReplace(newints,33,44):\n",newints)

	newfloats,_ = sortedu.UniqueDescFloatsInsert(newfloats,33.1)
	fmt.Println("sortedu.UniqueDescFloatsInsert(newfloats,33.1):\n",newfloats)
	newfloats,_ = sortedu.UniqueDescFloatsDelete(newfloats,33.1)
	fmt.Println("sortedu.UniqueDescFloatsDelete(newfloats,33.1):\n",newfloats)
	newfloats,_,_ = sortedu.UniqueDescFloatsReplace(newfloats,33.1,44.1)
	fmt.Println("sortedu.UniqueDescFloatsReplace(newfloats,33.1,44.1):\n",newfloats)

	newstrings,_ = sortedu.UniqueDescStringsInsert(newstrings,"33")
	fmt.Println("sortedu.UniqueDescStringsInsert(newstrings,\"33\"):\n",newstrings)
	newstrings,_ = sortedu.UniqueDescStringsDelete(newstrings,"33")
	fmt.Println("sortedu.UniqueDescStringsDelete(newstrings,\"33\"):\n",newstrings)
	newstrings,_,_ = sortedu.UniqueDescStringsReplace(newstrings,"33","44")
	fmt.Println("sortedu.UniqueDescStringsReplace(newstrings,\"33\",\"44\"):\n",newstrings)


	// 切片反转
	sortoe.IntsReverse(ints)
	sortoe.FloatsReverse(floats)
	sortoe.StringsReverse(strings)
	fmt.Println("IntsReverse(ints):\n",ints)
	fmt.Println("FloatsReverse(floats):\n",floats)
	fmt.Println("StringsReverse(strings):\n",strings)

	// 升序、降序、相同判断, 有序百分比
	isAsc := sortoe.IntsIsAsc(ints)
	fmt.Println("IntsIsAsc(ints):\n",isAsc)
	isDesc := sortoe.IntsIsDesc(ints)
	fmt.Println("IntsIsDesc(ints):\n",isDesc)
	isSame := sortoe.IntsIsFlat(ints)
	fmt.Println("IntsIsFlat(ints):\n",isSame)
	ascPercent := sortoe.IntsAscPercent(ints)
	fmt.Println("IntsAscPercent(ints):\n",ascPercent)

	// 升序、降序、相同判断, 有序百分比
	isAsc = sortoe.FloatsIsAsc(floats)
	fmt.Println("FloatsIsAsc(floats):\n",isAsc)
	isDesc = sortoe.FloatsIsDesc(floats)
	fmt.Println("FloatsIsDesc(ints):\n",isDesc)
	isSame = sortoe.FloatsIsFlat(floats)
	fmt.Println("FloatsIsFlat(ints):\n",isSame)
	ascPercent = sortoe.FloatsAscPercent(floats)
	fmt.Println("FloatsAscPercent(floats):\n",ascPercent)

	// 升序、降序、相同判断, 有序百分比
	isAsc = sortoe.StringsIsAsc(strings)
	fmt.Println("StringsIsAsc(strings):\n",isAsc)
	isDesc = sortoe.StringsIsDesc(strings)
	fmt.Println("StringsIsDesc(strings):\n",isDesc)
	isSame = sortoe.StringsIsFlat(strings)
	fmt.Println("StringsIsFlat(strings):\n",isSame)
	ascPercent = sortoe.StringsAscPercent(strings)
	fmt.Println("StringsAscPercent(strings):\n",ascPercent)




	size := 10000000

	arr := make([]int, size)
	for i:=0;i<size;i++ {
		arr[i] = rand.Intn(100000000)
	}

	start := time.Now().UnixNano()
	sortoe.IntsAscGo(arr)
	end := time.Now().UnixNano()
	fmt.Println("    IntsAscGo:",end-start, " rand IntsAscGo ")
	fmt.Println("    IntsIsAsc:",sortoe.IntsIsAsc(arr))


	for i:=0;i<size;i++ {
		arr[i] = rand.Intn(100000000)
	}

	start = time.Now().UnixNano()
	sortoe.IntsDescGo(arr)
	end = time.Now().UnixNano()
	fmt.Println("   IntsDescGo:",end-start, " rand IntsDescGo ")
	fmt.Println("   IntsIsDesc:",sortoe.IntsIsDesc(arr))



	arr2 := make([]float64, size)
	for i:=0;i<size;i++ {
		arr2[i] = float64(rand.Intn(100000000))+0.1
	}

	start = time.Now().UnixNano()
	sortoe.FloatsAscGo(arr2)
	end = time.Now().UnixNano()
	fmt.Println("  FloatsAscGo:",end-start, " rand FloatsAscGo ")
	fmt.Println("  FloatsIsAsc:",sortoe.FloatsIsAsc(arr2))



	for i:=0;i<size;i++ {
		arr2[i] = float64(rand.Intn(100000000))+0.1
	}

	start = time.Now().UnixNano()
	sortoe.FloatsDescGo(arr2)
	end = time.Now().UnixNano()
	fmt.Println(" FloatsDescGo:",end-start, " rand FloatsDescGo ")
	fmt.Println(" FloatsIsDesc:",sortoe.FloatsIsDesc(arr2))



	arr3 := make([]string, size)
	for i:=0;i<size;i++ {
		arr3[i] = strconv.Itoa( rand.Intn(100000000) )
	}

	start = time.Now().UnixNano()
	sortoe.StringsAscGo(arr3)
	end = time.Now().UnixNano()
	fmt.Println(" StringsAscGo:",end-start, " rand StringsAscGo ")
	fmt.Println(" StringsIsAsc:",sortoe.StringsIsAsc(arr3))



	for i:=0;i<size;i++ {
		arr3[i] = strconv.Itoa( rand.Intn(100000000) )
	}

	start = time.Now().UnixNano()
	sortoe.StringsDescGo(arr3)
	end = time.Now().UnixNano()
	fmt.Println("StringsDescGo:",end-start, " rand StringsDescGo ")
	fmt.Println("StringsIsDesc:",sortoe.StringsIsDesc(arr3))

	fmt.Println( "EstimateDepth(1000,12)",sortoe.EstimateDepth(1000,12) )
	fmt.Println( "EstimateDepth(10000,12)",sortoe.EstimateDepth(10000,12) )
	fmt.Println( "EstimateDepth(100000,12)",sortoe.EstimateDepth(100000,12) )
	fmt.Println( "EstimateDepth(1000000,12)",sortoe.EstimateDepth(1000000,12) )
	fmt.Println( "EstimateDepth(10000000,12)",sortoe.EstimateDepth(10000000,12) )
	fmt.Println( "EstimateDepth(100000000,12)",sortoe.EstimateDepth(100000000,12) )
	fmt.Println( "EstimateDepth(1000000000,12)",sortoe.EstimateDepth(1000000000,12) )
	fmt.Println( "EstimateDepth(10000000000,12)",sortoe.EstimateDepth(10000000000,12) )
	fmt.Println( "EstimateDepth(100000000000,12)",sortoe.EstimateDepth(100000000000,12) )

	fmt.Println( "EstimateDepth(100000,1)",sortoe.EstimateDepth(100000,1) )
	fmt.Println( "EstimateDepth(100000,2)",sortoe.EstimateDepth(100000,2) )
	fmt.Println( "EstimateDepth(100000,4)",sortoe.EstimateDepth(100000,4) )
	fmt.Println( "EstimateDepth(100000,8)",sortoe.EstimateDepth(100000,8) )
	fmt.Println( "EstimateDepth(100000,12)",sortoe.EstimateDepth(100000,12) )
	fmt.Println( "EstimateDepth(100000,16)",sortoe.EstimateDepth(100000,16) )
	fmt.Println( "EstimateDepth(100000,20)",sortoe.EstimateDepth(100000,20) )
	fmt.Println( "EstimateDepth(100000,24)",sortoe.EstimateDepth(100000,24) )
	fmt.Println( "EstimateDepth(100000,28)",sortoe.EstimateDepth(100000,28) )
	fmt.Println( "EstimateDepth(100000,32)",sortoe.EstimateDepth(100000,32) )
	fmt.Println( "EstimateDepth(100000,36)",sortoe.EstimateDepth(100000,36) )
	fmt.Println( "EstimateDepth(100000,40)",sortoe.EstimateDepth(100000,40) )
	fmt.Println( "EstimateDepth(100000,48)",sortoe.EstimateDepth(100000,48) )
	fmt.Println( "EstimateDepth(100000,64)",sortoe.EstimateDepth(100000,64) )
	fmt.Println( "EstimateDepth(100000,128)",sortoe.EstimateDepth(100000,128) )

	fmt.Println( "EstimateDepth(10000000,1)",sortoe.EstimateDepth(10000000,1) )
	fmt.Println( "EstimateDepth(10000000,2)",sortoe.EstimateDepth(10000000,2) )
	fmt.Println( "EstimateDepth(10000000,4)",sortoe.EstimateDepth(10000000,4) )
	fmt.Println( "EstimateDepth(10000000,8)",sortoe.EstimateDepth(10000000,8) )
	fmt.Println( "EstimateDepth(10000000,12)",sortoe.EstimateDepth(10000000,12) )
	fmt.Println( "EstimateDepth(10000000,16)",sortoe.EstimateDepth(10000000,16) )
	fmt.Println( "EstimateDepth(10000000,20)",sortoe.EstimateDepth(10000000,20) )
	fmt.Println( "EstimateDepth(10000000,24)",sortoe.EstimateDepth(10000000,24) )
	fmt.Println( "EstimateDepth(10000000,28)",sortoe.EstimateDepth(10000000,28) )
	fmt.Println( "EstimateDepth(10000000,32)",sortoe.EstimateDepth(10000000,32) )
	fmt.Println( "EstimateDepth(10000000,36)",sortoe.EstimateDepth(10000000,36) )
	fmt.Println( "EstimateDepth(10000000,40)",sortoe.EstimateDepth(10000000,40) )
	fmt.Println( "EstimateDepth(10000000,48)",sortoe.EstimateDepth(10000000,48) )
	fmt.Println( "EstimateDepth(10000000,64)",sortoe.EstimateDepth(10000000,64) )
	fmt.Println( "EstimateDepth(10000000,128)",sortoe.EstimateDepth(10000000,128) )

	fmt.Println( "EstimateDepth(100000000,1)",sortoe.EstimateDepth(100000000,1) )
	fmt.Println( "EstimateDepth(100000000,2)",sortoe.EstimateDepth(100000000,2) )
	fmt.Println( "EstimateDepth(100000000,4)",sortoe.EstimateDepth(100000000,4) )
	fmt.Println( "EstimateDepth(100000000,8)",sortoe.EstimateDepth(100000000,8) )
	fmt.Println( "EstimateDepth(100000000,12)",sortoe.EstimateDepth(100000000,12) )
	fmt.Println( "EstimateDepth(100000000,16)",sortoe.EstimateDepth(100000000,16) )
	fmt.Println( "EstimateDepth(100000000,20)",sortoe.EstimateDepth(100000000,20) )
	fmt.Println( "EstimateDepth(100000000,24)",sortoe.EstimateDepth(100000000,24) )
	fmt.Println( "EstimateDepth(100000000,28)",sortoe.EstimateDepth(100000000,28) )
	fmt.Println( "EstimateDepth(100000000,32)",sortoe.EstimateDepth(100000000,32) )
	fmt.Println( "EstimateDepth(100000000,36)",sortoe.EstimateDepth(100000000,36) )
	fmt.Println( "EstimateDepth(100000000,40)",sortoe.EstimateDepth(100000000,40) )
	fmt.Println( "EstimateDepth(100000000,48)",sortoe.EstimateDepth(100000000,48) )
	fmt.Println( "EstimateDepth(100000000,64)",sortoe.EstimateDepth(100000000,64) )
	fmt.Println( "EstimateDepth(100000000,128)",sortoe.EstimateDepth(100000000,128) )
}