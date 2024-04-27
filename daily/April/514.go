package main

import "math"

// Brute force with dfs and memoization
func findRotateSteps(ring string, key string) int {
	// return rFindRotateSteps(ring, key, 0)

	posOfChars := make(map[byte][]int)
	for i, c := range ring {
		posOfChars[byte(c)] = append(posOfChars[byte(c)], i)
	}

	var mem = make([][]int, len(ring))
	for i := range mem {
		mem[i] = make([]int, len(key))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}

	var dfs func(ringPos, keyPos int) int
	dfs = func(ringPos, keyPos int) int {
		if keyPos >= len(key) {
			return 0
		}

		if mem[ringPos][keyPos] != -1 {
			return mem[ringPos][keyPos]
		}

		res := math.MaxInt
		for _, idx := range posOfChars[key[keyPos]] {
            dif := int(math.Abs(float64(idx-ringPos)))
			res = min(res, dfs(idx, keyPos+1) + min(len(ring) - dif, dif))
		}
		mem[ringPos][keyPos] = res
		return res
	}

	return dfs(0, 0) + len(key)
}

// Take the closest way to find the next char
// Doesn't work since there's a chance the longer one => shorter distance in general

// func rFindRotateSteps(ring, key string, startFrom int) int {
// 	res := 0
// 	cur := startFrom
// 	for idx, c := range key {
// 		left := cur
// 		right := cur
//
// 		for ring[left] != byte(c) && ring[right] != byte(c) {
// 			left--
// 			if left < 0 {
// 				left = len(ring) - 1
// 			}
// 			right = (right + 1) % len(ring)
// 			res++
// 		}
// 		res++
//
// 		if ring[left] == ring[right] && left != right {
// 			leftLen := rFindRotateSteps(ring, key[idx+1:], left)
// 			rightLen := rFindRotateSteps(ring, key[idx+1:], right)
// 			if leftLen < rightLen {
// 				return res + leftLen
// 			} else {
// 				return res + rightLen
// 			}
// 		} else if ring[left] == byte(c) {
// 			cur = left
// 		} else {
// 			cur = right
// 		}
//
// 	}
//
// 	return res
// }
