package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPenetrate(t *testing.T) {
	arrayDelta := []int{50, -20, -30, -41, -9, 15, 5, 22, 8, 0, -25, 25, -12, 9, 3, -4, -3, 1, 5, 1}
	transactionResult := Calculate(arrayDelta)

	assertTransaction(t, arrayDelta, transactionResult, 13)
}

func TestSimpleTest(t *testing.T) {
	arrayDelta := []int{12, -9, -3}
	transactionResult := Calculate(arrayDelta)

	assertTransaction(t, arrayDelta, transactionResult, 2)
}

func assertTransaction(t *testing.T, deltas []int, result []Transaction, bestCount int) {
	assert.Equal(t, bestCount, len(result))
	for _, transaction := range result {
		deltas[transaction.from] += transaction.amount
		deltas[transaction.to] -= transaction.amount
	}
	assert.EqualValues(t, deltas, make([]int, len(deltas)))
}
