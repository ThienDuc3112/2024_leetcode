package main

// Check answer
// Modified version of longest subsequent
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)
	dp := make([][]int, len1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len2+1)
	}
	for i, list := range dp {
		for j := range list {
            if i == 0 {
                dp[i][j] = j
                continue
            }
			if j == 0 {
                dp[i][j] = i
                continue
			}
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len1][len2]
}


// Wrong algo, for longest SUBSEQUENT, which is not correct
// func minDistance(word1 string, word2 string) int {
// 	len1 := len(word1)
// 	len2 := len(word2)
// 	dp := make([][]int, len1+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, len2+1)
// 	}
// 	for i, list := range dp {
// 		if i == 0 {
// 			continue
// 		}
// 		for j := range list {
// 			if j == 0 {
// 				continue
// 			}
// 			if word1[i-1] == word2[j-1] {
// 				dp[i][j] = dp[i-1][j-1] + 1
// 			} else {
// 				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
// 			}
// 		}
// 	}
//     for _, list := range dp {
//         for _, num := range(list) {
//             fmt.Printf("%v ", num)
//         }
//         fmt.Println()
//     }
// 	maxCommonSubstring := dp[len1][len2]
// 	fmt.Printf("%v", maxCommonSubstring)
// 	return max(len1, len2) - maxCommonSubstring
// }

