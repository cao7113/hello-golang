package leetcode

import (
	"sort"
)

// https://leetcode.com/problems/squares-of-a-sorted-array/
func (s *LtCodeSuite) TestP977() {
	cases := []struct {
		nums []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	s.Run("O(n)", func() {
		for _, e := range cases {
			res := sortedSquares(e.nums)
			s.Equal(e.want, res)
		}
	})

	s.Run("just work", func() {
		for _, e := range cases {
			res := sortedSquaresJustWork(e.nums)
			s.Equal(e.want, res)
		}
	})
}

// https://leetcode.com/problems/squares-of-a-sorted-array/discuss/221922/Java-two-pointers-O(N)
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// in final, the max should in rightmost, and it should be in one of head and tail because it is in non-desc

	head, tail := 0, n-1
	for i := n - 1; i >= 0; i-- {
		a, b := nums[head], nums[tail]
		if a*a < b*b {
			result[i] = b * b
			tail--
		} else {
			result[i] = a * a
			head++
		}
	}
	return result
}

func sortedSquaresJustWork(nums []int) []int {
	var result []int
	for _, a := range nums {
		result = append(result, a*a)
	}
	sort.Ints(result) // n*lg(n)
	return result
}

// https://leetcode.com/problems/merge-sorted-array/
func (s *LtCodeSuite) TestP88() {
	type st struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}
	cases := []st{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
	}

	s.Run("in reverse order", func() {
		for _, e := range cases {
			merge(e.nums1, e.m, e.nums2, e.n)
			s.Equal(e.want, e.nums1)
		}
	})

	s.Run("just work", func() {
		//cases2 := make([]st, len(cases))
		//copy(cases2, cases)
		//
		//for _, e := range cases {
		//	mergeJustWork(e.nums1, e.m, e.nums2, e.n)
		//	s.Equal(e.want, e.nums1)
		//}
	})
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	// final: max should in the right-most
	// |---->m---->|---->n---|
	// --->i............j<----

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 || // nums2 finished in position
			(i >= 0 && j >= 0 && nums1[i] > nums2[j]) { //move item in nums1 into proper place in current round
			nums1[k] = nums1[i]
			i-- // move forward
		} else {
			nums1[k] = nums2[j] // put item in nums2 if it should
			j--
		}
	}
}

func mergeJustWork(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for i, b := range nums2 {
		nums1[m+i] = b
	}
	sort.Ints(nums1)
}
