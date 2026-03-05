package sharedutil

import (
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/rand"
)

type number interface {
	constraints.Float | constraints.Integer
}

func Join(w ...string) string {
	return strings.Join(w, "")
}

func Map[T, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, fn func(U, T) U, initial U) U {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func ReduceRight[T, U any](slice []T, fn func(U, T) U, initial U) U {
	result := initial
	for i := len(slice) - 1; i >= 0; i-- {
		result = fn(result, slice[i])
	}
	return result
}

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func Find[T any](slice []T, fn func(T) bool) (T, bool) {
	for _, v := range slice {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}

func FindIndex[T any](slice []T, fn func(T) bool) int {
	for i, v := range slice {
		if fn(v) {
			return i
		}
	}
	return -1
}

func FindLast[T any](slice []T, fn func(T) bool) (T, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i]) {
			return slice[i], true
		}
	}
	var zero T
	return zero, false
}

func FindLastIndex[T any](slice []T, fn func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i]) {
			return i
		}
	}
	return -1
}

func ForEach[T any](slice []T, fn func(T)) {
	for _, v := range slice {
		fn(v)
	}
}

func ForEachRight[T any](slice []T, fn func(T)) {
	for i := len(slice) - 1; i >= 0; i-- {
		fn(slice[i])
	}
}

func Partition[T any](slice []T, fn func(T) bool) (trueSlice []T, falseSlice []T) {
	trueSlice = make([]T, 0)
	falseSlice = make([]T, 0)
	for _, v := range slice {
		if fn(v) {
			trueSlice = append(trueSlice, v)
		} else {
			falseSlice = append(falseSlice, v)
		}
	}
	return trueSlice, falseSlice
}

func Sort[T constraints.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func SortBy[T any, U constraints.Ordered](slice []T, fn func(T) U) {
	sort.Slice(slice, func(i, j int) bool {
		return fn(slice[i]) < fn(slice[j])
	})
}

func SortByDesc[T any, U constraints.Ordered](slice []T, fn func(T) U) {
	sort.Slice(slice, func(i, j int) bool {
		return fn(slice[i]) > fn(slice[j])
	})
}

func SortByString[T constraints.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func SortByStringDesc[T constraints.Ordered](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

func GroupBy[T any, U comparable](slice []T, fn func(T) U) map[U][]T {
	result := make(map[U][]T)
	for _, v := range slice {
		key := fn(v)
		if _, ok := result[key]; !ok {
			result[key] = make([]T, 0)
		}
		result[key] = append(result[key], v)
	}
	return result
}

func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0)
	for _, subSlice := range slice {
		result = append(result, subSlice...)
	}
	return result
}

func Concat[T any](slice1 []T, slice2 []T) []T {
	result := make([]T, 0)
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
}

func SliceOf[T any](value T, length int) []T {
	result := make([]T, length)
	for i := 0; i < length; i++ {
		result[i] = value
	}
	return result
}

func SliceCopy[T any](slice []T) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	return result
}

func SliceCopyRange[T any](slice []T, start, end int) []T {
	result := make([]T, end-start)
	copy(result, slice[start:end])
	return result
}

func Index[T comparable](slice []T, value T) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Shuffle[T any](slice []T) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func Unique[T comparable](slice []T) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func Intersect[T comparable](slice1 []T, slice2 []T) []T {
	result := make([]T, 0)
	for _, v := range slice1 {
		if Contains(slice2, v) {
			result = append(result, v)
		}
	}
	return result
}

func Union[T comparable](slice1 []T, slice2 []T) []T {
	result := make([]T, 0)
	for _, v := range slice1 {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	for _, v := range slice2 {
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func Difference[T comparable](slice1 []T, slice2 []T) []T {
	result := make([]T, 0)
	for _, v := range slice1 {
		if !Contains(slice2, v) {
			result = append(result, v)
		}
	}
	return result
}

func Equal[T comparable](slice1 []T, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
}

func Max[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("slice is empty")
	}

	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		panic("slice is empty")
	}

	min := slice[0]
	for _, v := range slice[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func Sum[T number](slice []T) T {
	var result T
	for _, v := range slice {
		result += v
	}
	return result
}

func SumIf[T number](slice []T, fn func(T) bool) T {
	var result T
	for _, v := range slice {
		if fn(v) {
			result += v
		}
	}
	return result
}

func Count[T comparable](slice []T) int {
	return len(slice)
}

func CountIf[T comparable](slice []T, fn func(T) bool) int {
	var count int
	for _, v := range slice {
		if fn(v) {
			count++
		}
	}
	return count
}

func Avg[T number](slice []T) float64 {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return float64(sum) / float64(len(slice))
}

func AvgIf[T number](slice []T, fn func(T) bool) float64 {
	var sum T
	var count int
	for _, v := range slice {
		if fn(v) {
			sum += v
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func Median[T number](slice []T) T {
	sorted := make([]T, len(slice))
	copy(sorted, slice)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	} else {
		return sorted[mid]
	}
}

func Mode[T comparable](slice []T) []T {
	counts := make(map[T]int)
	for _, v := range slice {
		counts[v]++
	}
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	result := make([]T, 0)
	for k, v := range counts {
		if v == maxCount {
			result = append(result, k)
		}
	}
	return result
}

func Slice[T any, E number](start, end, step int) []E {
	result := make([]E, 0)
	for i := start; i < end; i += step {
		result = append(result, E(i))
	}
	return result
}

func SliceFrom[T any, E number](start int) []E {
	return Slice[T, E](start, 1<<63-1, 1)
}

func SliceTo[T any, E number](end int) []E {
	return Slice[T, E](0, end, 1)
}

func SliceStep[T any, E number](start, end, step int) []E {
	return Slice[T, E](start, end, step)
}

func SliceStepFrom[T any, E number](start, step int) []E {
	return Slice[T, E](start, 1<<63-1, step)
}

func SliceStepTo[T any, E number](end, step int) []E {
	return Slice[T, E](0, end, step)
}

func SliceStepFromTo[T any, E number](start, end, step int) []E {
	return Slice[T, E](start, end, step)
}

func SliceStepFromToStep[T any, E number](start, end, step, step2 int) []E {
	return Slice[T, E](start, end, step2)
}

func SliceStepFromToStepFrom[T any, E number](start, end, step, step2 int) []E {
	return Slice[T, E](start, end, step2)
}

func SliceStepFromToStepTo[T any, E number](start, end, step, step2 int) []E {
	return Slice[T, E](start, end, step2)
}

type Pair[T1, T2 any] struct {
	First  T1
	Second T2
}

func Zip[T1 any, T2 any](slice1 []T1, slice2 []T2) []Pair[T1, T2] {
	result := make([]Pair[T1, T2], 0)
	for i := 0; i < len(slice1) && i < len(slice2); i++ {
		result = append(result, Pair[T1, T2]{slice1[i], slice2[i]})
	}
	return result
}

func Unzip[T1 any, T2 any](slice []Pair[T1, T2]) (slice1 []T1, slice2 []T2) {
	slice1 = make([]T1, 0)
	slice2 = make([]T2, 0)
	for _, pair := range slice {
		slice1 = append(slice1, pair.First)
		slice2 = append(slice2, pair.Second)
	}
	return slice1, slice2
}

func Enumerate[T any](slice []T) []Pair[int, T] {
	result := make([]Pair[int, T], 0)
	for i, v := range slice {
		result = append(result, Pair[int, T]{i, v})
	}
	return result
}
