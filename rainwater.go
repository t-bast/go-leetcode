package leetcode

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	// Remove leaky borders.
	start := 0
	for {
		if start == len(height)-1 {
			return 0
		}

		if height[start] > height[start+1] {
			break
		}

		start++
	}

	end := len(height) - 1
	for {
		if end == start {
			return 0
		}

		if height[end-1] < height[end] {
			break
		}

		end--
	}

	if end-start <= 1 {
		return 0
	}

	height = height[start : end+1]

	// Find the two biggest heights: it's then easy to compute the water
	// between those two.
	var i1, h1, i2, h2 int
	for i := 0; i < len(height); i++ {
		cur := height[i]
		if cur > h2 {
			h2, i2 = cur, i
		}
		if cur > h1 {
			h2, i2 = h1, i1
			h1, i1 = cur, i
		}
	}

	// Compute the water between the two biggest heights.
	if i1 > i2 {
		i1, i2 = i2, i1
	}

	volume := 0
	surface := h1
	if h2 < h1 {
		surface = h2
	}

	for i := i1 + 1; i < i2; i++ {
		volume += surface - height[i]
	}

	// Recursively call on the rest of the map.
	volume += trap(height[:i1+1])
	volume += trap(height[i2:])

	return volume
}
