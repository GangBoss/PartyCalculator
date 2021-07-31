package engine

//import "sort"

type Transaction struct {
	from, to, amount int
}

/*type pair struct {
	Money int
	Index int
}

type ByKey []pair

func (s ByKey) Len() int {
	return len(s)
}

func (s ByKey) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByKey) Less(i, j int) bool {
	return s[i].Money < s[j].Money
}*/

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

func getMaxSubMasks(masks []bool) []int {
	pwn := len(masks)
	dp := make([]int, pwn)
	prevmask := make([]int, pwn)
	for m := 0; m < pwn; m++ {
		for s := m; s != 0; s = (s - 1) & m {
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

/*func getAnswer(g int, deltas []int, prevmask []int) []Transaction {
	transfers := make([]Transaction, 0)
	n := len(deltas)
	pwn := 1 << n
	pairs := make([]pair, 0, n)
	for i := pwn - 1; i != 0; i = i ^ prevmask[i] {
		m := prevmask[i]
		for j := 0; j < n; j++ {
			if m&(1<<j) != 0 {
				pairs = append(pairs, pair{deltas[j], j})
			}
		}
		sort.Sort(ByKey(pairs))
		prev := pairs[0].Index
		for _, cur := range pairs[1:] {
			transfers = append(transfers, Transaction{prev, cur.Index, -deltas[prev]})
			deltas[cur.Index] += deltas[prev]
			prev = cur.Index
		}
	}
	return transfers
}*/

func getAnswer(deltas []int, prevmask []int) []Transaction {
	transfers := make([]Transaction, 0)
	n := len(deltas)
	pwn := 1 << len(deltas)
	for i := pwn - 1; i != 0; i ^= prevmask[i] {
		m := prevmask[i]
		var positive []int
		var negative []int
		for j := 0; j < n; j++ {
			if m&(1<<j) != 0 {
				if deltas[j] > 0 {
					positive = append(positive, j)
				} else if deltas[j] < 0 {
					negative = append(negative, j)
				}
			}
		}
		for len(positive) != 0 {
			a, b := positive[len(positive)-1], negative[len(negative)-1]
			positive = positive[:len(positive)-1]
			negative = negative[:len(negative)-1]
			transfers = append(transfers, Transaction{b, a, deltas[a]})
			deltas[b] += deltas[a]
			deltas[a] = 0
			if deltas[b] > 0 {
				positive = append(positive, b)
			} else if deltas[b] < 0 {
				negative = append(negative, b)
			}
		}
	}
	return transfers
}

func Calculate(arrayDelta []int) []Transaction {
	previousMasks := getMaxSubMasks(findMasks(arrayDelta))
	return getAnswer(append([]int(nil), arrayDelta...), previousMasks)
}
