package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPenetrate(t *testing.T) {

	arrayDelta := []int{50, -20, -30, -41, -9, 15, 5, 22, 8, 0, -25, 25, -12, 9, 3, -4, -3, 1, 5, 1}
	transactionResult := Calculate(arrayDelta)

	assertTransaction(t, arrayDelta,transactionResult,13)
}

func TestSimpleTest(t *testing.T) {

	arrayDelta := []int{12, -9, -3}
	transactionResult := Calculate(arrayDelta)

	assertTransaction(t, arrayDelta, transactionResult, 2)
}

func assertTransaction(t *testing.T,slice1 []int,result []Transaction, bestCount int) {
	arraySum := positiveArraySum(slice1)

	transactionSum := 0
	for _, transaction := range result {
		transactionSum += transaction.amount
	}

	assert.Len(t, result, bestCount, "result %v", result)
	assert.Equal(t, arraySum, transactionSum, "result %v", result)
}

func positiveArraySum(arr []int) int {
	sum := 0
	for _, value := range arr {
		if value > 0 {
			sum += value
		}
	}
	return sum
}
