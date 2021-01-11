package algocasts

/*
 本质上，先根据分桶来先做了一轮排序，然后再在内部做插入排序，
 再按遍历桶的顺序把有效的序列置入输入数组中，从而达成排序的目的
*/
func bucketSortRewrite(arr []int) {
	if len(arr) == 0 {
		return
	}

	bucketSize := 5

	maxValue := arr[0]
	minValue := arr[0]

	for _, item := range arr {
		maxValue = max(maxValue, item)
		minValue = min(minValue, item)
	}

	bucketCount := len(arr) / bucketSize
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0, bucketSize)
	}

	for _, item := range arr {
		// 计算桶的算法，这只是其中一种，还有其他的实现
		bc := (item - minValue) / (maxValue - minValue + 1.0) * bucketCount
		buckets[bc] = append(buckets[bc], item)
	}

	idx := 0
	for _, bucket := range buckets {
		insertsort(bucket)
		for _, num := range bucket {
			arr[idx] = num
			idx++
		}
	}
}
