package engine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToPenetrate(t *testing.T) {

	slice1 := []int{50, -20, -30, -41, -9, 15, 5, 22, 8, 0, -25, 25, -12, 9, 3, -4, -3, 1, 5, 1}
	result:=Calculate(slice1)
	rightTransactions := []Transaction{
		{17, 20, 3},
		{20,19,2},
		{16,19,2},
		{10,18,1},
		{13,12,15},
		{11,25,12},
		{3,9,30},
		{9,8,22},
		{2,7,20},
		{7,6,15},
		{5,1,9},
		{4,41,1},
	}

	assert.ObjectsAreEqualValues(result, rightTransactions)
}
