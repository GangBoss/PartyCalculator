package engine

type Transaction struct {
	from, to, amount int
}

func findMasks(deltas []int) []bool {
	n := len(deltas)
	pwn := 1 << n
	masks := make([]bool, pwn)
	for i := 0; i < pwn; i++ {
		s := 0
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				s += deltas[j]
			}
		}
		if s == 0 {
			masks[i] = true
		}
	}
	return masks
}

func getMaxSubmasks(masks []bool) []int {
	pwn := len(masks)
	dp := make([]int, pwn)
	prevmask := make([]int, pwn)
	for m := 0; m < pwn; m++ {
		for s := m; s != 0; s = (s - 1) % m {
			if masks[s] {
				f := m ^ s
				if dp[m] < dp[f]+1 {
					dp[m] = dp[f] + 1
					prevmask[m] = s
				}
			}
		}
	}
	return prevmask
}

func getAnswer(deltas []int, prevmask []int) []Transaction {
	transfers := make([]Transaction, 0)
	n := len(deltas)
	pwn := 1 << len(deltas)
	for i := pwn - 1; i != 0; i ^= prevmask[i] {
		var positive []int
		var negative []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				if deltas[j] > 0 {
					positive = append(positive, j)
				} else if deltas[j] < 0 {
					negative = append(negative, j)
				}
			}
		}
		for len(negative) != 0 {
			a, b := positive[len(positive)-1], negative[len(negative)-1]
			positive = positive[:len(positive)-1]
			negative = negative[:len(negative)-1]
			transfers = append(transfers, Transaction{b, a, deltas[a]})
			deltas[a] += deltas[b]
			deltas[b] = 0
			if deltas[a] > 0 {
				positive = append(positive, a)
			} else if deltas[a] < 0 {
				negative = append(negative, a)
			}
		}
	}
	return transfers
}

func Calculate(arrayDelta []int) []Transaction {
	//n := len(arrayDelta)
	previousMasks := getMaxSubmasks(findMasks(arrayDelta))
	return getAnswer(append([]int(nil), arrayDelta...), previousMasks)
}
