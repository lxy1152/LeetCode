package main

// 5704. 好子数组的最大分数
/*
5704. 好子数组的最大分数 显示英文描述
通过的用户数59
尝试过的用户数84
用户总通过次数60
用户总提交次数99
题目难度Hard
给你一个整数数组 nums （下标从 0 开始）和一个整数 k 。

一个子数组 (i, j) 的 分数 定义为 min(nums[i], nums[i+1], ..., nums[j]) * (j - i + 1) 。一个 好 子数组的两个端点下标需要满足 i <= k <= j 。

请你返回 好 子数组的最大可能 分数 。



示例 1：

输入：nums = [1,4,3,7,4,5], k = 3
输出：15
解释：最优子数组的左右端点下标是 (1, 5) ，分数为 min(4,3,7,4,5) * (5-1+1) = 3 * 5 = 15 。
示例 2：

输入：nums = [5,5,4,5,4,1,1,1], k = 0
输出：20
解释：最优子数组的左右端点下标是 (0, 4) ，分数为 min(5,5,4,5,4) * (4-0+1) = 4 * 5 = 20 。
*/

func maximumScore(nums []int, k int) int {
	left, right := k, k
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		for j := i; j < len(nums); j++ {
			dp[i][i] = nums[i]
			if j >= 1 {
				if i < len(nums)-1 || (i == len(nums)-1 && j != i) {
					dp[i][j] = min7(dp[i][j-1], nums[j])
				}
			}
		}
	}

	res := -1
	for ; left >= 0; left-- {
		for right = k; right < len(nums); right++ {
			temp := right - left + 1
			max := func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}
			cur := dp[left][right] * temp
			res = max(res, cur)

		}
	}
	return res
}

func min7(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// 按理说上述解法应该是可以的 不过堆栈溢出了， 因为数据开的太大了

/*
单调栈解法
*/
func maximumScore1(nums []int, k int) int {
	leng := len(nums)
	stack := make([]int, leng+2)
	h := make([]int, leng+2)
	for i := 1; i < leng+1; i++ {
		h[i] = nums[i-1]
	}
	h[0] = -1
	h[leng+1] = -1
	top := 0
	left := make([]int, leng+1)
	right := make([]int, leng+1)
	for i := 1; i < leng+1; i++ {
		for h[stack[top]] >= h[i] {
			top--
		}
		left[i] = stack[top]
		top++
		stack[top] = i
	}
	top = 0
	stack[0] = leng + 1
	for i := leng; i > 0; i-- {
		for h[stack[top]] >= h[i] {
			top--
		}
		right[i] = stack[top]
		top++
		stack[top] = i
	}
	k++
	res := 0
	for i := 1; i <= leng; i++ {
		if left[i] < k && right[i] > k {
			res = max(res, (right[i]-left[i]-1)*h[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	print(maximumScore1([]int{529, 7655, 4113, 7929, 7745, 6149, 2691, 435, 7858, 978, 5467, 8761, 2881, 4175, 359, 9711, 7157, 1740, 5214, 7660, 1113, 614, 4077, 2859, 2954, 1170, 3997, 4526, 2771, 4046, 3174, 7126, 3866, 7603, 5596, 9587, 1826, 9228, 9983, 2487, 7018, 1130, 1299, 1389, 5914, 2926, 5679, 5781, 5770, 8705, 9053, 3449, 9893, 6042, 4349, 2008, 2959, 9730, 5639, 8813, 5206, 747, 460, 149, 2553, 4006, 2012, 8966, 2245, 2635, 9089, 7962, 5879, 9785, 1069, 7122, 7926, 9546, 4557, 5719, 6926, 2353, 1887, 2744, 2240, 472, 7559, 824, 4320, 8267, 7210, 4764, 8368, 3071, 8973, 7849, 8704, 2375, 5130, 5092, 8311, 6147, 7692, 6794, 8280, 4863, 4932, 4863, 946, 4243, 7307, 5061, 3752, 426, 6352, 1978, 818, 5604, 3306, 9034, 9805, 7900, 8462, 4473, 1914, 8449, 113, 8566, 7671, 9518, 693, 730, 8064, 2004, 9584, 8725, 4392, 468, 3013, 502, 6472, 4185, 6380, 471, 4258, 3424, 8435, 7470, 5109, 7228, 486, 8925, 3130, 1574, 4454, 9048, 5824, 679, 5364, 3804, 5143, 6714, 5695, 2806, 7653, 9153, 26, 3195, 2856, 4303, 6914, 270, 9011, 3325, 7835, 1770, 8668, 3395, 1756, 6702, 743, 2618, 6841, 7492, 6068, 9150, 8852, 5004, 9610, 6716, 7053, 4937, 6209, 2772, 4454, 3988, 432, 8097, 9400, 3194, 5565, 6266, 2821, 4769, 5019, 164, 443, 8620, 695, 8173, 7705, 5081, 1376, 6573, 9364, 7328, 2513, 8020, 2361, 4915, 5346, 7410, 2379, 5538, 7560, 9431, 1049, 9069, 9616, 9368, 3998, 3240, 8006, 3177, 4380, 794, 8668, 4685, 3456, 2257, 2645, 7269, 2424, 5209, 4198, 3086, 6821, 7598, 3764, 7505, 6191, 9018, 7019, 338, 5606, 1484, 1962, 1415, 8423, 6988, 9140, 8305, 6196, 2299, 3397, 8634, 3855, 2139, 4924, 7102, 4662, 6713, 1056, 5418, 5982, 940, 3149, 9245, 637, 8873, 7651, 2975, 6447, 5539, 291, 9119, 7773, 9287, 4286, 858, 9111, 5012, 9499, 1595, 4408, 918, 5742, 5134, 9475, 8427, 9556, 8422, 4825, 9747, 522, 266, 9267, 7351, 5413, 7374, 8556, 8014, 6194, 6076, 3011, 1612, 217, 6101, 7673, 5623, 6757, 5038, 2615, 3711, 8296, 7453, 408, 868, 4090, 1038, 6549, 7307, 7886, 3571, 6450, 1947, 5236, 8943, 5827, 1325, 1278, 9292, 1930, 7778, 6039, 3935, 623, 7886, 256, 1194, 833, 4219, 3641, 1047, 2543, 4442, 9242, 150, 8148, 1460, 9650, 5882, 9647, 431, 7709, 714, 6345, 6104, 3240, 6151, 6310, 7220, 216, 6452, 3271, 3368, 1247, 7870, 8492, 2124, 1932, 3524, 5515, 2911, 8798, 7242, 1019, 8154, 3942, 5654, 5994, 8323, 449, 1914, 5680, 2821, 8399, 1904, 6356, 8538, 298, 7418, 2339, 1672, 2948, 584, 4877, 2702, 36, 8044, 1528, 3815, 799, 8865, 7043, 5846, 3204, 4238, 3012, 2060, 8169, 9422, 8049, 8871, 1517, 819, 6066, 4821, 7262, 6131, 9774, 9210, 9683, 338, 3573, 3812, 4398, 5022, 1313, 9705, 792, 3644, 3397, 529, 8340, 2683, 524, 5635, 6653, 9852, 8619, 39, 6060, 4273, 3026, 7680, 6861, 4470, 8, 4094, 4720, 388, 9344, 7749, 394, 3603, 8406, 8713, 9328, 4423, 3074, 9966, 694, 4094, 3627, 6926, 3558, 4322, 1720, 3814, 7277, 5770, 7077, 9871, 7872, 4016, 1196, 7425, 8843, 8631, 2500, 142, 991, 9472, 1603, 7433, 2955, 5213, 3083, 9960, 3993, 8940, 2675, 5442, 8308, 6653, 2105, 4640, 2020, 8924, 8918, 2949, 4949, 5442, 9580, 2811, 5205, 7078, 9779, 5049, 4455, 3420, 5586, 6315, 1629, 101, 5384, 4917, 7629, 2731, 8121, 2852, 9260, 8445, 2896, 8075, 4667, 2164, 7085, 3616, 3177, 9104, 5318, 8728, 2332, 3943, 2388, 4352, 8537, 199, 1655, 3881, 7058, 3912, 7027, 5721, 9737, 6031, 3265, 2699, 8285, 7835}, 314))
}
