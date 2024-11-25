package main

import (
	"fmt"
	"time"
)

type task struct {
	name        string
	description string
	dueDate     time.Time
	isCompleted bool
}

func createTask(name string, description string, dueDate time.Time) task {
	return task{name: name, description: description, dueDate: dueDate, isCompleted: false}
}
func (t *task) markAsCompleted() {

	t.isCompleted = true
	fmt.Println("task is:", &t)
	return
}
func filterTasks(tasks []task, filter func([]task, bool) []task, isCompleted bool) []task {
	return filter(tasks, isCompleted)

}
func filter(tasks []task, condition bool) []task {
	ans := []task{}
	for _, task := range tasks {
		if task.isCompleted == condition {
			ans = append(ans, task)
		}
	}
	return ans
}
func (t *task) updateDescription(newDesc string) {
	t.description = newDesc
	return
}
func countCompletedTask(tasks ...task) int {
	ans := 0
	for _, t := range tasks {
		if t.isCompleted {
			ans++
		}
	}
	return ans
}
func main() {

	task1 := createTask("Work", "Work is important", time.Now())
	task2 := createTask("Study", "Study is also important", time.Now())
	task3 := createTask("Sleep", "Sleep is very important", time.Now())

	var tasks = []task{
		task1, task2, task3,
	}
	

	fmt.Println("Before:")
	for _, t := range tasks {
		fmt.Printf("%+v\n", t)
	}

	// Modify task1
	tasks[0].markAsCompleted()
	tasks[0].updateDescription("Updated Work Description")

	// Print updated tasks
	fmt.Println("\nAfter:")
	for _, t := range tasks {
		fmt.Printf("%+v\n", t)
	}
	completedtasks := filterTasks(tasks, filter, true)
	fmt.Println("\nCompleted tasks:", completedtasks)

	completedataska := countCompletedTask(tasks...)
	fmt.Println("\nNumber of completed tasks:", completedataska)
}
