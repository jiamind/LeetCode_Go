package main


func minimumTotal(triangle [][]int) int {

	dp := make([]int,len(triangle[len(triangle) - 1]) + 1)

	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + Min(dp[j],dp[j+1]) ;
		}
	}

	return dp[0]
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}


