package main

func allPossibleRoutes(str string) [8]string {
	bSlice := []byte(str)
	res := [8]string{}
	for i := 0; i < 8; i++ {
		if i/4 == 0 {
			bSlice[i] = (bSlice[i]-'0'+1)%10 + '0'
			res[i] = string(bSlice)
			if bSlice[i] == '0' {
				bSlice[i] = '9'
			} else {
				bSlice[i] = (bSlice[i]-'0'-1)%10 + '0'
			}
		} else {
			if bSlice[i%4] == '0' {
				bSlice[i%4] = '9'
				res[i] = string(bSlice)
				bSlice[i%4] = '0'
			} else {
				bSlice[i%4] = (bSlice[i%4]-'0'-1)%10 + '0'
				res[i] = string(bSlice)
				bSlice[i%4] = (bSlice[i%4]-'0'+1)%10 + '0'
			}
		}
	}
	return res
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	visited := make(map[string]bool)
	for _, str := range deadends {
		visited[str] = true
	}
	q := make([]string, 0)
	q = append(q, target)
	visited[target] = true
	res := 0
	for len(q) > 0 {
		length := len(q)
		for t := 0; t < length; t++ {
			cur := q[0]
			q = q[1:]
			if cur == "0000" {
				return res
			}
			next := allPossibleRoutes(cur)
			for _, nextStr := range next {
				_, isVisited := visited[nextStr]
				if !isVisited {
					visited[nextStr] = true
					q = append(q, nextStr)
				}
			}
		}
		res++
	}
	return -1
}
