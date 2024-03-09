package main

var mapping map[byte][]byte = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func backtrackLetterCom(digits string, curStr string, res *[]string) {
	if digits == "" {
        *res = append(*res, curStr)
        return
	}
    possibleChar := mapping[digits[0]]
    length := len(possibleChar)
    for i := 0; i < length; i++ {
        backtrackLetterCom(digits[1:], curStr + string(possibleChar[i]), res)
    }
}
func letterCombinations(digits string) []string {
    res := []string {}
    if digits == "" {return res}
    backtrackLetterCom(digits, "", &res)
    return res;
}
