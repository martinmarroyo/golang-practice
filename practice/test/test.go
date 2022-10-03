package main

import "fmt"

func main() {
	// Create a slice of numbers
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println("index:", i)
		fmt.Println("value:", num)
	}

	// Create and iterate over a map
	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k,v := range kvs{
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i,c := range "This is a test"{
		fmt.Println(i, "This is a test"[i:i+1], c)
	}

	test := "testing"
	fmt.Println(test[0:1])
}