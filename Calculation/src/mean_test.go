package main

import (
	"testing"
)

func TestSD(t *testing.T) {
	// Тестовые случаи
	testCases := []struct {
		name     string
		input    []int
		expected float64
	}{
		{
			name:     "Test case 1",
			input:    []int{1, 1, 2, 2, 2},
			expected: 0.5477225575051661,
		},
		{
			name:     "Test case 2",
			input:    []int{1, 2, 3, 4, 5},
			expected: 1.5811388300841898,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := sd(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestMode(t *testing.T) {
	// Тест на пустой массив
	nums := []int{}
	if mode(nums) != 0 {
		t.Errorf("Expected mode of empty slice to be 0, but got %v", mode(nums))
	}

	// Тест на массив с одним элементом
	nums = []int{5}
	if mode(nums) != 5 {
		t.Errorf("Expected mode of [5] to be 5, but got %v", mode(nums))
	}

	// Тест на массив с несколькими одинаковыми элементами
	nums = []int{1, 1, 2, 2, 2}
	expectedMode := 2
	if mode(nums) != expectedMode {
		t.Errorf("Expected mode of [1, 1, 2, 2, 2] to be %v, but got %v", expectedMode, mode(nums))
	}

	// Тест на массив с разными элементами
	nums = []int{1, 2, 3, 4, 5}
	expectedMode = 1
	if mode(nums) != expectedMode {
		t.Errorf("Expected mode of [1, 2, 3, 4, 5] to be %v, but got %v", expectedMode, mode(nums))
	}

	// Тест на массив с отрицательными числами
	nums = []int{-1, -2, -3, -4, -5}
	expectedMode = -5
	if mode(nums) != expectedMode {
		t.Errorf("Expected mode of [-1, -2, -3, -4, -5] to be %v, but got %v", expectedMode, mode(nums))
	}

	// Тест на случай с плавающей точкой
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedMode = 1
	if mode(nums) != expectedMode {
		t.Errorf("Expected mode of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] to be %v, but got %v", expectedMode, mode(nums))
	}
}

func TestMean(t *testing.T) {
	// Тест на пустой массив
	nums := []int{}
	if mean(nums) != 0 {
		t.Errorf("Expected mean of empty slice to be 0, but got %v", mean(nums))
	}

	// Тест на один элемент
	nums = []int{5}
	if mean(nums) != 5 {
		t.Errorf("Expected mean of [5] to be 5, but got %v", mean(nums))
	}

	// Тест на несколько элементов
	nums = []int{1, 2, 3, 4, 5}
	expectedMean := 3.0
	if mean(nums) != expectedMean {
		t.Errorf("Expected mean of [1, 2, 3, 4, 5] to be %v, but got %v", expectedMean, mean(nums))
	}

	// Тест на отрицательные числа
	nums = []int{-1, -2, -3, -4, -5}
	expectedMean = -3
	if mean(nums) != expectedMean {
		t.Errorf("Expected mean of [-1, -2, -3, -4, -5] to be %v, but got %v", expectedMean, mean(nums))
	}

	// Тест на случай с плавающей точкой
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedMean = 5.5
	if mean(nums) != expectedMean {
		t.Errorf("Expected mean of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] to be %v, but got %v", expectedMean, mean(nums))
	}
}

func TestMedian(t *testing.T) {
	// Тест на пустой массив
	nums := []int{}
	if median(nums) != 0 {
		t.Errorf("Expected median of empty slice to be 0, but got %v", median(nums))
	}

	// Тест на один элемент
	nums = []int{5}
	if median(nums) != 5 {
		t.Errorf("Expected median of [5] to be 5, but got %v", median(nums))
	}

	// Тест на массив с четным количеством элементов
	nums = []int{1, 2, 3, 4}
	expectedMedian := 2.5
	if median(nums) != expectedMedian {
		t.Errorf("Expected median of [1, 2, 3, 4] to be %v, but got %v", expectedMedian, median(nums))
	}

	// Тест на массив с нечетным количеством элементов
	nums = []int{1, 2, 3, 4, 5}
	expectedMedian = 3
	if median(nums) != expectedMedian {
		t.Errorf("Expected median of [1, 2, 3, 4, 5] to be %v, but got %v", expectedMedian, median(nums))
	}

	// Тест на массив с отрицательными числами
	nums = []int{-1, -2, -3, -4, -5}
	expectedMedian = -3
	if median(nums) != expectedMedian {
		t.Errorf("Expected median of [-1, -2, -3, -4, -5] to be %v, but got %v", expectedMedian, median(nums))
	}

	// Тест на случай с плавающей точкой
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedMedian = 5.5
	if median(nums) != expectedMedian {
		t.Errorf("Expected median of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] to be %v, but got %v", expectedMedian, median(nums))
	}
}
