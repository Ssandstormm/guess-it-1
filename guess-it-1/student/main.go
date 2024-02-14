package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var k []float64
	for reader.Scan() {
		t := reader.Text()
		n, err := strconv.Atoi(t)
		if err != nil {
			fmt.Print(err)
			return
		}
		if len(k) > 2 && (n > int(Average(k))*2 || n < int(Average(k)*0.3)) {
			n = int(Average(k))
		}
		k = append(k, float64(n))
		Answer(k)
	}
}

func Average(num []float64) float64 {
	var sum float64
	for _, number := range num {
		sum += number
	}
	sampleCount := float64(len(num))
	Aver := sum / sampleCount
	return math.Round(Aver)
}

func Variance(n []float64, avg float64) float64 {
	var sum float64
	for _, i := range n {
		sum += math.Pow(i-avg, 2)
	}
	sd := sum / float64(len(n))
	return sd
}

func Std(n []float64, avg float64) float64 {
	return math.Sqrt(float64(Variance(n, avg)))
}

func Answer(num []float64) {
	avg := Average(num)
	div := Std(num, avg)
	upper := avg + div*1.28
	lower := avg - div*1.28
	fmt.Println(lower, upper)
}
