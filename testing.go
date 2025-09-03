package main

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
// 	mergedArr := make([]int, 0, len(nums1)+len(nums2))
// 	for len(nums1) > 0 && len(nums2) > 0 {
// 		if nums1[0] < nums2[0] {
// 			mergedArr = append(mergedArr, nums1[0])
// 			nums1 = nums1[1:]
// 		} else {
// 			mergedArr = append(mergedArr, nums2[0])
// 			nums2 = nums2[1:]
// 		}
// 	}
//
// 	for len(nums1) > 0 {
// 		mergedArr = append(mergedArr, nums1[0])
// 		nums1 = nums1[1:]
// 	}
//
// 	for len(nums2) > 0 {
// 		mergedArr = append(mergedArr, nums2[0])
// 		nums2 = nums2[1:]
// 	}
//
// 	if len(mergedArr)%2 != 0 {
// 		return float64(mergedArr[len(mergedArr)/2])
// 	} else {
// 		firstElm := len(mergedArr) / 2
// 		secondElm := firstElm - 1
//
// 		fmt.Println(firstElm, secondElm, cap(mergedArr), mergedArr)
//
// 		return (float64(mergedArr[firstElm]) + float64(mergedArr[secondElm])) / 2.0
// 	}
// }

// func main() {
// 	// nums1 := []int{1, 2}
// 	// nums2 := []int{3, 4}
// 	//
// 	// fmt.Println(findMedianSortedArrays(nums1, nums2))
//
// 	// x := 123
// 	// var arr []int
// 	// var res int
// 	// count := 0
// 	// for x != 0 {
// 	// 	lastDig := x % 10
// 	// 	x /= 10
// 	// 	arr = append(arr, lastDig)
// 	// 	count++
// 	// }
// 	// for _, v := range arr {
// 	// 	count--
// 	// 	res += v * int(math.Pow(float64(10), float64(count)))
// 	// }
//
// 	/* fmt.Println(res) */
// }
