package utils

import (
	"math"
	"time"
)

func FormatDate(timestamp int64) string {
	t := time.Unix(timestamp, 0).UTC()
	return t.Format("2006-01-02")
}

func GetTodayDate() string {
	return time.Now().UTC().Format("2006-01-02")
}

func CalculateAverage(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func CalculateMax(nums []float64) float64 {
	max := math.Inf(-1)
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func CalculateMin(nums []float64) float64 {
	min := math.Inf(1)
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func GetDominantCondition(conditions map[string]int) string {
	maxCount := 0
	dominant := ""
	for condition, count := range conditions {
		if count > maxCount {
			maxCount = count
			dominant = condition
		}
	}
	return dominant
}
