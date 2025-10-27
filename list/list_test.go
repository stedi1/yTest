package list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMin(t *testing.T) {
	// варианты последовательностей
	// слайсы начинаются с ожидаемого результата
	lists := map[string][]int{
		"three": {-50, 5, -50, 67},
		"four":  {-7, -7, 1, 43, 100},
		"two":   {3, 3, 4},
		"one":   {10, 10},
		"empty": {0},
	}
	// подтесты name = TestMin/three и т.д
	for name, list := range lists {
		t.Run(name, func(t *testing.T) {
			if Min(list[1:]) != list[0] {
				t.Errorf("get %d != want %d", Min(list[1:]), list[0])
			}
		})
	}
}

func TestMul(t *testing.T) {
	listA := [][]int{
		{},
		{3},
		{-5, 6, 10},
	}
	listB := [][]int{
		{},
		{7},
		{21, 0, 4},
	}
	// ожидаемые слайсы results[i] = list[i]*listB[i]
	results := [][]int{
		{},
		{21},
		{-105, 0, 40},
	}
	for i, list := range listA {
		c, err := Mul(list, listB[i])
		require.NoError(t, err)            // проверяем, что нет ошибки
		require.Len(t, c, len(results[i])) // проверяем длину результата
		// проверка значений
		for j, v := range c {
			require.Equal(t, results[i][j], v)
		}
	}
	// проверяем на возврат ошибки
	_, err := Mul([]int{22}, []int{33, 71})
	require.Error(t, err)
}

func TestGCD(t *testing.T) {
	pairs := [][]int{
		{0, 5, 0},
		{7, -350, 0},
		{1460, 124, 4},
		{198, 42, 6},
		{1386, 4494, 42},
		{1470, 1575, 105},
	}
	for _, pair := range pairs {
		gcd := GCD(pair[0], pair[1])
		if gcd != pair[2] {
			t.Errorf("НОД(%d, %d): %d != %d\n", pair[0], pair[1], gcd, pair[2])
		}
	}
}
