package main

import (
	"fmt"
)

func main() {
	martix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Println(martix)
	ret := spiralOrder(martix)
	fmt.Println(ret)
}

func spiralOrder(matrix [][]int) []int {
	list := []int{}

	if len(matrix) <= 0 || len(matrix[0]) <= 0 {
		return list
	}

	row := len(matrix)
	col := len(matrix[0])

	n := row * col
	dir := "right"
	i, j := 0, 0
	count := 0
	visited := map[string]bool{}

	for count < n {
		list = append(list, matrix[i][j])
		count++
		visited[rap(i, j)] = true
		switch dir {
		case "right":
			_, found := visited[rap(i, j+1)]
			if j+1 < col && !found {
				j++
			} else {
				dir = "down"
				i++
			}
		case "down":
			_, found := visited[rap(i+1, j)]
			if i+1 < row && !found {
				i++
			} else {
				dir = "left"
				j--
			}
		case "left":
			_, found := visited[rap(i, j-1)]
			if j-1 >= 0 && !found {
				j--
			} else {
				dir = "up"
				i--
			}
		case "up":
			_, found := visited[rap(i-1, j)]
			if i-1 >= 0 && !found {
				i--
			} else {
				dir = "right"
				j++
			}
		default:
		}
	}

	return list
}

func rap(i int, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}
