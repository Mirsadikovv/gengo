package parser

import (
	"fmt"
	"testing"
)

func TestFindTexts(t *testing.T) {

	const EXP = `
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	`
	// fmt.Println(SnakeCase(EXP))
	// fmt.Println(CamelCase(EXP))
	// fmt.Println(CamelCaseOptimized(EXP))
	fmt.Println(CamelCaseUnsafe(EXP))
	// fmt.Println(CamelCaseUnsafe3(EXP))
}

func BenchmarkSnakeCase(b *testing.B) {
	const EXP = " Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing"
	for i := 0; i < b.N; i++ {
		_ = SnakeCase(EXP)
	}
}

func BenchmarkCamelCase(b *testing.B) {
	const EXP = `
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	`
	for i := 0; i < b.N; i++ {
		_ = CamelCase(EXP)
	}
}

// BenchmarkCamelCase3-32    	  514381	      2245 ns/op	    1408 B/op	       1 allocs/op
func BenchmarkCamelCase3(b *testing.B) {
	const EXP = `
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	`
	for i := 0; i < b.N; i++ {
		_ = CamelCaseUnsafe3(EXP)
	}
}

func BenchmarkSnakeCase_2(b *testing.B) {

	const EXP = `
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
			id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		`
	parser := NewParser(EXP)

	b.Log(len(parser.ParsePlaceholders()))

	for i := 0; i < b.N; i++ {
		_ = parser.ParsePlaceholders()
	}
}
func BenchmarkSnakeCase_3(b *testing.B) {

	const EXP = `
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
			id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
		`
	parser := NewParser(EXP)

	b.Log(parser.ParseTexts())

	for i := 0; i < b.N; i++ {
		_ = parser.ParseTexts()
	}
}

// BenchmarkCamelCaseOptimized-32    	  134290	      8907 ns/op	    5760 B/op	       3 allocs/op
func BenchmarkCamelCaseOptimized(b *testing.B) {
	const EXP = `
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
`
	for i := 0; i < b.N; i++ {
		_ = CamelCaseOptimized(EXP)
	}
}

// BenchmarkCamelCaseUnsafe-32    	  410362	      2871 ns/op	    5120 B/op	       2 allocs/op
// BenchmarkCamelCaseUnsafe-32    	  404954	      2857 ns/op	    5120 B/op	       2 allocs/op
func BenchmarkCamelCaseUnsafe(b *testing.B) {
	const EXP = `
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
`
	for i := 0; i < b.N; i++ {
		_ = CamelCaseUnsafe(EXP)
	}
}

// BenchmarkCamelCaseUnsafe_2-32    	  465364	      2498 ns/op	    1408 B/op	       1 allocs/op
func BenchmarkCamelCaseUnsafe_2(b *testing.B) {
	const EXP = `
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
`

	for i := 0; i < b.N; i++ {
		_ = CamelCaseUnsafe2(EXP)
	}
}
func BenchmarkCamelCaseUnsafeToken(b *testing.B) {
	const EXP = `
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
	id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY} id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}Testing_@@@testing _______testing Testing_@@@testing _______testing Testing_@@@testing _______testingTesting_@@@testing _______testingTesting_@@@testing _______testing
`

	for i := 0; i < b.N; i++ {
		_ = CamelCaseUnsafeTokens(EXP)
	}
}

func TestParsePlaceholders_ValidInput(t *testing.T) {
	const EXP = "id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	expected := []string{"{ID}", "{NAME}", "{AGE}", "{SALARY}"}
	for _, exp := range expected {
		found := false
		for _, ph := range placeholders {
			if ph.Value == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected placeholder %s not found", exp)
		}
	}
}

func TestParsePlaceholders_EmptyInput(t *testing.T) {
	const EXP = ""
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	if len(placeholders) != 0 {
		t.Errorf("Expected no placeholders, got %v", placeholders)
	}
}

func TestParsePlaceholders_NoPlaceholders(t *testing.T) {
	const EXP = "No placeholders here"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	if len(placeholders) != 0 {
		t.Errorf("Expected no placeholders, got %v", placeholders)
	}
}

func TestParsePlaceholders_MultiplePlaceholders(t *testing.T) {
	const EXP = "id:{ID}/name:{NAME}/age:${AGE}/salary:{SALARY}/city:{CITY}"
	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()
	expected := []string{"{ID}", "{NAME}", "{AGE}", "{SALARY}", "{CITY}"}

	for _, exp := range expected {
		found := false
		for _, ph := range placeholders {
			if ph.Value == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected placeholder %s not found", exp)
		}
	}
}

func TestReplaceWithTokens_ReplaceWithMap(t *testing.T) {

	const EXP = "id:{ID}/name:{NAME}/age:{AGE}/salary:{SALARY}"

	parser := NewParser(EXP)

	placeholders := parser.ParsePlaceholders()

	replacer := map[string]string{
		"ID":     "1",
		"NAME":   "John",
		"SALARY": "50000",
		"AGE":    "25",
	}

	result := ReplaceWithTokens(EXP, placeholders, replacer)
	t.Log(result)
	if result != "id:1/name:John/age:25/salary:50000" {
		t.Errorf("Expected result to be 'id:1/name:John/age:25/salary:50000', got %s", result)
	}
}