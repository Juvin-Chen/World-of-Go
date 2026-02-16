package main

import (
	"fmt"
)

type TaskFunc func()

func task1()
func task2()
func task3()
func project()

func main() {
	tasks := []TaskFunc{task1, task2, task3}

	for _, task := range tasks {
		task()
		fmt.Println()
	}
	project()
}
