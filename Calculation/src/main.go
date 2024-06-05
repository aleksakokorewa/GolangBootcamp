package main

import (
	"flag"
	"fmt"
	"log/slog"
	"math"
	"os"
	"sort"
)

const (
	maxValue = 10000
	minValue = -10000
)

var (
	meanVar, medianVar, modeVar, sdVar bool
)

func init() {
	if len(os.Args) == 1 {
		flag.BoolVar(&meanVar, "mean", true, "Mean: true or false?\n")
		flag.BoolVar(&medianVar, "median", true, "Median: true or false?\n")
		flag.BoolVar(&modeVar, "mode", true, "Mode: true or false?\n")
		flag.BoolVar(&sdVar, "sd", true, "SD: true or false?\n")
	} else {
		flag.BoolVar(&meanVar, "mean", false, "Mean: true or false?\n")
		flag.BoolVar(&medianVar, "median", false, "Median: true or false?\n")
		flag.BoolVar(&modeVar, "mode", false, "Mode: true or false?\n")
		flag.BoolVar(&sdVar, "sd", false, "SD: true or false?\n")
	}
}

func main() {
	flag.Parse()
	res, err := numbers()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	printRes(res)
}

func numbers() ([]int, error) {
	var num []int
	var length int
	fmt.Println("Enter count of numbers:")
	_, err := fmt.Scan(&length)
	if err != nil {
		return nil, fmt.Errorf("length should be a number: %w", err)
	}
	fmt.Println("Enter numbers less 10k or more -10k:")
	num = make([]int, length)
	for i := 0; i < length; i++ {
		_, err = fmt.Scanf("%d", &num[i])
		if err != nil {
			return nil, fmt.Errorf("you should enter a number: %w", err)
		}
		if num[i] < minValue || num[i] > maxValue {
			return nil, fmt.Errorf("numbers should be less 10k or more -10k")
		}
	}
	return num, nil
}

func mean(num []int) float64 {
	if len(num) == 0 {
		return 0
	}
	var sum, count float64
	for _, v := range num {
		sum += float64(v)
		count++
	}

	return sum / float64(len(num))
}

func median(num []int) float64 {
	if len(num) == 0 {
		return 0
	}
	sort.Ints(num)
	var half = len(num) / 2
	if len(num)%2 == 0 {
		return mean(num[half-1 : half+1])
	} else {
		return float64(num[half])
	}
}

func mode(num []int) int {
	if len(num) == 0 {
		return 0
	}
	common := make(map[int]int)
	value, count := math.MaxInt, math.MinInt
	for _, v := range num {
		common[v]++
	}
	for key, v := range common {
		if v > count || (key < value && v >= count) {
			value, count = key, v
		}
	}
	return value
}

func sd(num []int) float64 {
	average := mean(num)
	var res float64
	for _, i := range num {
		res += math.Pow(float64(i)-average, 2)
	}
	result := res / float64(len(num)-1)
	return math.Sqrt(result)
}

func printRes(num []int) {
	if meanVar {
		fmt.Printf("Mean: %.2f\n", mean(num))
	}
	if medianVar {
		fmt.Printf("Median: %.2f\n", median(num))
	}
	if modeVar {
		fmt.Printf("Mode: %d\n", mode(num))
	}
	if sdVar {
		fmt.Printf("SD: %.2f\n", sd(num))
	}
}
