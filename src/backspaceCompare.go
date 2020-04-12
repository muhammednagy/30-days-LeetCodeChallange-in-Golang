package main
// https://leetcode.com/explore/challenge/card/30-day-leetcoding-challenge/529/week-2/3291/

func backspaceCompare(S string, T string) bool {
	pointer1 := len(S) - 1
	pointer2 := len(T) - 1

	skips1 := 0
	skips2 := 0

	for pointer1 >= 0 || pointer2 >= 0 {
		for pointer1 >= 0 {
			if S[pointer1] == 35 {
				skips1++
				pointer1--
			} else if skips1 > 0 {
				pointer1--
				skips1--
			} else {
				break
			}
		}

		for pointer2 >= 0 {
			if T[pointer2] == 35 {
				skips2++
				pointer2--
			} else if skips2 > 0 {
				pointer2--
				skips2--
			} else {
				break
			}
		}
		if pointer1 >= 0 && pointer2 >= 0 && S[pointer1] != T[pointer2] {
			return false
		}

		if (pointer1 >= 0) != (pointer2 >= 0) {
			return false
		}

		pointer1--
		pointer2--
	}
	return true
}
