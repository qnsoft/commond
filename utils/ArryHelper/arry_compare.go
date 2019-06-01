package arry

/*
获取int数组中最大的值和下标
*/
func Max_IntArry_One(_arry []int) (int, int) {
	//获取一个数组里最大值，并且拿到下标

	//声明一个数组5个元素
	// var arr [5]int = [...]int {6, 45, 63, 16 ,86}
	//假设第一个元素是最大值，下标为0
	maxVal := _arry[0]
	maxIndex := 0

	for i := 1; i < len(_arry); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < _arry[i] {
			maxVal = _arry[i]
			maxIndex = i
		}
	}
	return maxVal, maxIndex
}
