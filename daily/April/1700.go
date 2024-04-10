package main

func countStudents(students []int, sandwiches []int) int {
	zeroCount := 0
	oneCount := 0
	for _, student := range students {
		if student == 1 {
			oneCount++
		} else {
			zeroCount++
		}
	}
	amountSandwiches := len(sandwiches)
	for i, sandwich := range sandwiches {
		if sandwich == 1 && oneCount > 0 {
			oneCount--
		} else if sandwich == 0 && zeroCount > 0 {
			zeroCount--
		} else {
			return amountSandwiches - i
		}
	}
	return 0
}

// func countStudents(students []int, sandwiches []int) int {
// 	haveEaten := make([]bool, len(students))
// 	stackTop := 0
// 	updated := true
// 	for updated && stackTop < len(sandwiches) {
// 		updated = false
// 		for i := range students {
// 			if !haveEaten[i] && students[i] == sandwiches[stackTop] {
//                 stackTop++
//                 haveEaten[i] = true
//                 updated = true
// 			}
// 		}
// 	}
//     return len(sandwiches) - stackTop
// }
