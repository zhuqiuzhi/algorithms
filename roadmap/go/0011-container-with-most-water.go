package roadmap

// Leetcode Problem: https://leetcode.com/problems/container-with-most-water/
// 面积是两根垂直线的较低高度乘上宽度(即两根垂直线的距离)的乘积
// 难点: height=[1,8,100,2,100,4,8,3,7], result=200
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	var result int
	for i < j {
		result = max(result, (j-i)*min(height[i], height[j]))
		// 谁比较低, 就移动谁的原因: 因为移动高的垂直线肯定会比目前的面积少, 因为宽度变窄了但高度只会变低或者不变
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return result
}

func maxAreaBrutForce(height []int) int {
	var result int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			result = max(result, (j-i)*min(height[i], height[j]))
		}
	}
	return result
}

func maxAreaError(height []int) int {
	i, j := 0, len(height)-1
	var result = (j - i) * min(height[i], height[j])
	for i < j {
		if height[i] < height[j] && height[i+1] > height[i] {
			i++
		} else if height[i] > height[j] && height[j-1] > height[j] {
			j--
		} else {
			i++
			j--
		}

		result = max(result, (j-i)*min(height[i], height[j]))
	}

	return result
}
