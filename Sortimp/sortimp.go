package sortimp

import (
	"cmp"
	"slices"
	"sort"
)

type Customer struct {
	name   string
	age    int
	salary int
}

func CreateCustomer(name string, age int, salary int) *Customer {
	return &Customer{name: name, age: age, salary: salary}
}
func CmpPerson(a, b Customer) int {
	if a.name == b.name {
		return cmp.Compare(a.age, b.age)
	}
	return cmp.Compare(a.name, b.name)
}
func SortTester(customerList []Customer, attribute string, direction string) []Customer {
	newList := []Customer{}
	for _, customer := range customerList {
		newList = append(newList, customer)
	}
	switch attribute {
	case "age":
		sort.Slice(newList, func(i, j int) bool {
			return newList[i].age < newList[j].age
		})
	case "salary":
		sort.Slice(newList, func(i, j int) bool {
			return newList[i].salary < newList[j].salary
		})
	case "name":
		slices.SortFunc(newList, CmpPerson)
	}
	if direction == "desc" {
		slices.Reverse(newList)
	}
	return newList
}
