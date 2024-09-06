// Последовательность чисел называется арифметической, если она состоит как
// минимум из двух элементов, а разница между каждыми двумя соседними
// элементами одинакова. Более формально, последовательность s является
// арифметической, если и только если s[i + 1]- s[i] == s[1]-s[0] для всех допустимых i.
// Вам дан массив из n целых чисел, nums и два массива из m целых чисел каждый, l и
// r, представляющий m диапазонов запросов, где i-й запрос является диапазоном
// [l[i], r[i]]. Все массивы имеют индекс 0.
// Верните список булевых элементов ответа, где answer [i] является true, если
// подмассив nums [l [i]], nums [l [i] + 1], ..., nums [r [i]] можно переупорядочить для
// формирования арифметической последовательности, и false в противном случае
// Constraints:
// n == nums.length
// m == l.length
// m == r.length
// 2 <= n <= 500
// 1 <= m <= 500
// 0 <= l[i] < r[i] < n
// -105 <= nums[i] <= 105
// https://leetcode.com/problems/arithmetic-subarrays/

package main

import (
	"fmt"
	"sort"
)

func checkArithmeticSubarrays(nums, l, r []int) []bool {
	res := make([]bool, len(l))
	for i := 0; i < len(l); i++ {
		curSl := append([]int{}, nums[l[i]:r[i]+1]...)
		res[i] = Check(curSl)
	}
	return res
}

func Check(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	sort.Ints(nums)
	diff := nums[1] - nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != diff {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkArithmeticSubarrays([]int{4, 6, 5, 9, 3, 7}, []int{0, 0, 2}, []int{2, 3, 5}))
}
