package sortedu

import "github.com/cds325/sortoe"

// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueAscIntsInsert(arr []int,elem int,noCheck ...bool) (newarr []int,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsAscAndUnique(arr) {
		panic("UniqueAscIntsInsert param arr is not sorted asc or not unique")
	}
	if sortoe.AscIntsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.AscIntsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueAscIntsDelete(arr []int,elem int,noCheck ...bool) (newarr []int,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsAscAndUnique(arr) {
		panic("UniqueAscIntsDelete param arr is not sorted asc or not unique")
	}
	if sortoe.AscIntsHas(arr,elem,true) {
		newarr,_ = sortoe.AscIntsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueAscIntsReplace(arr []int,old,new int,noCheck ...bool) (newarr []int,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsAscAndUnique(arr) {
		panic("UniqueAscIntsReplace param arr is not sorted asc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.AscIntsHas(arr,old,true)
	newExisting = sortoe.AscIntsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.AscIntsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.AscIntsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.AscIntsInsertNoCheck(arr,new)
	}
	return
}

// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueDescIntsInsert(arr []int,elem int,noCheck ...bool) (newarr []int,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsDescAndUnique(arr) {
		panic("UniqueDescIntsInsert param arr is not sorted desc or not unique")
	}
	if sortoe.DescIntsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.DescIntsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueDescIntsDelete(arr []int,elem int,noCheck ...bool) (newarr []int,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsDescAndUnique(arr) {
		panic("UniqueDescIntsDelete param arr is not sorted desc or not unique")
	}
	if sortoe.DescIntsHas(arr,elem,true) {
		newarr,_ = sortoe.DescIntsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueDescIntsReplace(arr []int,old,new int,noCheck ...bool) (newarr []int,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.IntsIsDescAndUnique(arr) {
		panic("UniqueDescIntsReplace param arr is not sorted desc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.DescIntsHas(arr,old,true)
	newExisting = sortoe.DescIntsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.DescIntsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.DescIntsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.DescIntsInsertNoCheck(arr,new)
	}
	return
}


// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueAscFloatsInsert(arr []float64,elem float64,noCheck ...bool) (newarr []float64,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsAscAndUnique(arr) {
		panic("UniqueAscFloatsInsert param arr is not sorted asc or not unique")
	}
	if sortoe.AscFloatsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.AscFloatsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueAscFloatsDelete(arr []float64,elem float64,noCheck ...bool) (newarr []float64,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsAscAndUnique(arr) {
		panic("UniqueAscFloatsDelete param arr is not sorted asc or not unique")
	}
	if sortoe.AscFloatsHas(arr,elem,true) {
		newarr,_ = sortoe.AscFloatsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueAscFloatsReplace(arr []float64,old,new float64,noCheck ...bool) (newarr []float64,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsAscAndUnique(arr) {
		panic("UniqueAscFloatsReplace param arr is not sorted asc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.AscFloatsHas(arr,old,true)
	newExisting = sortoe.AscFloatsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.AscFloatsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.AscFloatsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.AscFloatsInsertNoCheck(arr,new)
	}
	return
}


// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueDescFloatsInsert(arr []float64,elem float64,noCheck ...bool) (newarr []float64,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsDescAndUnique(arr) {
		panic("UniqueDescFloatsInsert param arr is not sorted desc or not unique")
	}
	if sortoe.DescFloatsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.DescFloatsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueDescFloatsDelete(arr []float64,elem float64,noCheck ...bool) (newarr []float64,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsDescAndUnique(arr) {
		panic("UniqueDescFloatsDelete param arr is not sorted desc or not unique")
	}
	if sortoe.DescFloatsHas(arr,elem,true) {
		newarr,_ = sortoe.DescFloatsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueDescFloatsReplace(arr []float64,old,new float64,noCheck ...bool) (newarr []float64,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.FloatsIsDescAndUnique(arr) {
		panic("UniqueDescFloatsReplace param arr is not sorted desc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.DescFloatsHas(arr,old,true)
	newExisting = sortoe.DescFloatsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.DescFloatsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.DescFloatsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.DescFloatsInsertNoCheck(arr,new)
	}
	return
}




// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueAscStringsInsert(arr []string,elem string,noCheck ...bool) (newarr []string,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsAscAndUnique(arr) {
		panic("UniqueAscStringsInsert param arr is not sorted asc or not unique")
	}
	if sortoe.AscStringsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.AscStringsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueAscStringsDelete(arr []string,elem string,noCheck ...bool) (newarr []string,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsAscAndUnique(arr) {
		panic("UniqueAscStringsDelete param arr is not sorted asc or not unique")
	}
	if sortoe.AscStringsHas(arr,elem,true) {
		newarr,_ = sortoe.AscStringsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueAscStringsReplace(arr []string,old,new string,noCheck ...bool) (newarr []string,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsAscAndUnique(arr) {
		panic("UniqueAscStringsReplace param arr is not sorted asc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.AscStringsHas(arr,old,true)
	newExisting = sortoe.AscStringsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.AscStringsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.AscStringsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.AscStringsInsertNoCheck(arr,new)
	}
	return
}



// 向有序且元素唯一的切片插入指定元素,
// 返回新的有序且元素唯一的切片和原切片中元素是否存在的bool值.
func UniqueDescStringsInsert(arr []string,elem string,noCheck ...bool) (newarr []string,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsDescAndUnique(arr) {
		panic("UniqueDescStringsInsert param arr is not sorted desc or not unique")
	}
	if sortoe.DescStringsHas(arr,elem,true) {
		return arr,true
	}
	return sortoe.DescStringsInsertNoCheck(arr,elem),false
}

// 将有序且元素唯一的切片删除指定元素,
// 返回新的有序且元素唯一的切片和原切片中新元素是否存在的bool值.
func UniqueDescStringsDelete(arr []string,elem string,noCheck ...bool) (newarr []string,existing bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsDescAndUnique(arr) {
		panic("UniqueDescStringsDelete param arr is not sorted desc or not unique")
	}
	if sortoe.DescStringsHas(arr,elem,true) {
		newarr,_ = sortoe.DescStringsDeleteNoCheck(arr,elem)
		return newarr,true
	}
	return arr,false
}

// 将有序且元素唯一的切片中的old替换为new,
// 返回新的有序且元素唯一的切片和原切片中old、new是否存在的bool值.
func UniqueDescStringsReplace(arr []string,old,new string,noCheck ...bool) (newarr []string,oldExisting,newExisting bool) {
	if (len(noCheck)==0 || !noCheck[0]) && !sortoe.StringsIsDescAndUnique(arr) {
		panic("UniqueDescStringsReplace param arr is not sorted Desc or not unique")
	}
	newarr = arr
	oldExisting = sortoe.DescStringsHas(arr,old,true)
	newExisting = sortoe.DescStringsHas(arr,new,true)
	if oldExisting && newExisting {
		newarr,_ = sortoe.DescStringsDeleteNoCheck(arr,old)
	}else if oldExisting {
		sortoe.DescStringsReplaceNoCheck(arr,old,new,1)
	}else if newExisting {
		// 无需操作, no op
	}else{
		newarr = sortoe.DescStringsInsertNoCheck(arr,new)
	}
	return
}