package main

func myAtoi(s string) int {
	var maxInt int = 2147483647
	var minInt int = -2147483648
	i := 0
	strLen := len(s)
	for i < strLen && s[i] == ' ' {
		i++
	}
	res := 0
	positive := true
    if i == strLen {
        return 0
    }
	if s[i] == '-' {
		positive = false
        i++
	} else if s[i] == '+' {
        i++
    }
	for i < strLen && s[i]-'0' >= 0 && s[i]-'0' <= 9 {
		curNum := s[i] - '0'
        // How tf does adding 0 <= work but not 0 <
		if res > 0 && 0 <= (maxInt - int(curNum))/res && (maxInt - int(curNum))/res < 10 {
			return maxInt
		}
		if res < 0 && 0 <= (minInt + int(curNum))/res && (minInt + int(curNum))/res < 10 {
			return minInt
		}
        if positive {
            res = res * 10 + int(curNum)
        } else {
            res = res * 10 - int(curNum)
        }
        i++
	}
    return res
}

