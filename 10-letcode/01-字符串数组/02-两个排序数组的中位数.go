package _1_字符串数组

// 给定两个大小为 m 和 n 的有序数组 nums1 和 nums2 。
// 请找出这两个有序数组的中位数。要求算法的时间复杂度为 O(log (m+n)) 。
// 你可以假设 nums1 和 nums2 不同时为空。
// 示例 1:
// nums1 = [1, 3]
// nums2 = [2]
// 中位数是 2.0
// 示例 2:
// nums1 = [1, 2]
// nums2 = [3, 4]
// 中位数是 (2 + 3)/2 = 2.5

// func findMedianSortedArrays(s1, s2 []float64, k int) float64 {

// 	s1length := len(s1)
// 	s2length := len(s2)

// 	// 长度 s1 <= s2
// 	if s1length > s2length {
// 		s1, s2 = s2, s1
// 	}

// 	if s1length == 0 {
// 		return s2[k-1]
// 	}

// 	if k == 1 {
// 		return math.Min(s1[0], s2[0])
// 	}

// 	pa := math.Min(k/2, s1length), pb := k - pa
// 	if s1[pa -1 ] < s2[pb - 1] {

// 	}

// 	return 0
// }

// 归并排序
// func merge(left, right []int) (res []int) {
// 	var l, r int
// 	for l < len(left) && r < len(right) {
// 		if left[l] < right[r] {
// 			res = append(res, left[l])
// 		} else {
// 			res = append(res, right[r])
// 		}
// 	}
// 	res = append(res, left[l])
// 	res = append(res, right[r])
// }

// double findKth(int a[], int m, int b[], int n, int k)
// {
// 	//always assume that m is equal or smaller than n
// 	if (m > n)
// 		return findKth(b, n, a, m, k);
// 	if (m == 0)
// 		return b[k - 1];
// 	if (k == 1)
// 		return min(a[0], b[0]);
// 	//divide k into two parts
// 	int pa = min(k / 2, m), pb = k - pa;
// 	if (a[pa - 1] < b[pb - 1])
// 		return findKth(a + pa, m - pa, b, n, k - pa);
// 	else if (a[pa - 1] > b[pb - 1])
// 		return findKth(a, m, b + pb, n - pb, k - pb);
// 	else
// 		return a[pa - 1];
// }

// class Solution
// {
// public:
// 	double findMedianSortedArrays(int A[], int m, int B[], int n)
// 	{
// 		int total = m + n;
// 		if (total & 0x1)
// 			return findKth(A, m, B, n, total / 2 + 1);
// 		else
// 			return (findKth(A, m, B, n, total / 2)
// 					+ findKth(A, m, B, n, total / 2 + 1)) / 2;
// 	}
// }
