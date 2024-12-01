// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

// The overall run time complexity should be O(log (m+n)).

// Example 1:

// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
// Example 2:

// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

// Constraints:

// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m
	for low <= high {
		partition1 := (low + high) / 2
		partition2 := (m+n+1)/2 - partition1

		maxLeft1 := math.MinInt32
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		minRight1 := math.MaxInt32
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		maxLeft2 := math.MinInt32
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight2 := math.MaxInt32
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 1 {
				return float64(max(maxLeft1, maxLeft2))
			}
			return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
		} else if maxLeft1 > minRight2 {
			high = partition1 - 1
		} else {
			low = partition1 + 1
		}
	}

	panic("No median found")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
