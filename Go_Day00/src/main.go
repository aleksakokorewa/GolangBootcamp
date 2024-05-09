package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
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
	printRes(numbers())
}

func numbers() (num []int) {
	var length int
	fmt.Println("Enter count of numbers:")
	_, err := fmt.Scan(&length)
	if err != nil {
		fmt.Println("You should enter a number")
		os.Exit(1)
	}
	fmt.Println("Enter numbers less 10k or more -10k:")
	num = make([]int, length)
	for i := 0; i < length; i++ {
		_, err := fmt.Scanf("%d", &num[i])
		if err != nil {
			fmt.Println("You should enter a number")
			os.Exit(1)
		}
		if num[i] < -100000 || num[i] > 100000 {
			fmt.Println("Numbers should be less 10k or more -10k")
			os.Exit(1)
		}
	}
	return num
}

func mean(num []int) float64 {
	var sum float64
	for i := 0; i < len(num); i++ {
		sum += float64(num[i])
	}
	return sum / float64(len(num))
}

func median(num []int) float64 {
	sort.Ints(num)
	var half = len(num) / 2
	if len(num)%2 == 0 {
		return mean(num[half-1 : half+1])
	} else {
		return float64(num[half])
	}
}

func mode(num []int) int {
	common := make(map[int]int)
	value, count := 0, 0
	for _, v := range num {
		common[v]++
	}
	for key, v := range common {
		if v > count || (key == count && v < value) {
			value, count = key, v
		}
	}
	return value
}

func sd(num []int) float64 {
	average := mean(num)
	var res float64
	for i := range num {
		res += math.Pow(float64(num[i])-average, 2) / float64(len(num))
	}
	return math.Sqrt(res)
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
