package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	memberNum := 24
	members := generateMembers(memberNum)
	fmt.Println(members)

	shuffle(members)
	fmt.Println(members)

	tables := devideIntoTables(members)
	fmt.Println(tables)

	sortByTables(tables)
	fmt.Println(tables)
}

func generateMembers(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}

func shuffle(members []int) {
	n := len(members)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		members[i], members[j] = members[j], members[i]
	}
}

func devideIntoTables(members []int) [][]int {
	tables := make([][]int, len(members)/4)
	for i := 0; i < len(members); i += 4 {
		tables[i/4] = members[i : i+4]
	}
	return tables
}

func sortByTables(tables [][]int) {
	for _, table := range tables {
		sort.Slice(table, func(j, k int) bool {
			return table[j] < table[k]
		})
	}
}
