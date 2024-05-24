package main

import (
	"math"
	"testing"
)

func TestSD(t *testing.T) {
	// Тест на пустой массив
	nums := []int{}
	if sd(nums) != 0 {
		t.Errorf("Expected standard deviation of empty slice to be 0, but got %v", sd(nums))
	}

	// Тест на массив с одним элементом
	nums = []int{5}
	if sd(nums) != 0 {
		t.Errorf("Expected standard deviation of [5] to be 0, but got %v", sd(nums))
	}

	// Тест на массив с несколькими одинаковыми элементами
	nums = []int{1, 1, 2, 2, 2}
	expectedSD := 0.0
	if sd(nums) != expectedSD {
		t.Errorf("Expected standard deviation of [1, 1, 2, 2, 2] to be %v, but got %v", expectedSD, sd(nums))
	}

	// Тест на массив с разными элементами
	nums = []int{1, 2, 3, 4, 5}
	expectedSD = math.Sqrt(2)
	if sd(nums) != expectedSD {
		t.Errorf("Expected standard deviation of [1, 2, 3, 4, 5] to be %v, but got %v", expectedSD, sd(nums))
	}

	// Тест на массив с отрицательными числами
	nums = []int{-1, -2, -3, -4, -5}
	expectedSD = math.Sqrt(2)
	if sd(nums) != expectedSD {
		t.Errorf("Expected standard deviation of [-1, -2, -3, -4, -5] to be %v, but got %v", expectedSD, sd(nums))
	}

	// Тест на случай с плавающей точкой
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSD = math.Sqrt(2)
	if sd(nums) != expectedSD {
		t.Errorf("Expected standard deviation of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10] to be %v, but got %v", expectedSD, sd(nums))
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
	expectedMode = -1
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
