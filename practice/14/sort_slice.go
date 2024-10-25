package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d,%s,%d", employee.Name, employee.Age, employee.ID, employee.SSN)
}

type SortByAge []Employee

func (sortIntf SortByAge) Len() int               { return len(sortIntf) }
func (sortInft SortByAge) Swap(i int, j int)      { sortInft[i], sortInft[j] = sortInft[j], sortInft[i] }
func (sortInft SortByAge) Less(i int, j int) bool { return sortInft[i].Age < sortInft[j].Age }

func sortSlice() {
	var employees = []Employee{
		{"Graham", "231", 235643, 31},
		{"John", "3434", 245643, 42},
		{"Michael", "8934", 32432, 17},
		{"Jenny", "24334", 32444, 26},
	}

	fmt.Println("Before:", employees)
	sort.Sort(SortByAge(employees))
	fmt.Println("After:", employees)

	sort.Slice(employees, func(i int, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	fmt.Println("Adjust:", employees)
}
