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

type SortByAge []Employee

func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d, %s, %d", employee.Name, employee.Age, employee.ID, employee.SSN)

}

func (sortIntf SortByAge) Len() int               { return len(sortIntf) }
func (sortIntf SortByAge) Swap(i int, j int)      { sortIntf[i], sortIntf[j] = sortIntf[j], sortIntf[i] }
func (sortIntf SortByAge) Less(i int, j int) bool { return sortIntf[i].Age < sortIntf[j].Age }

func main() {
	var employees = []Employee{
		{"Graham", "231", 234643, 31},
		{"John", "3434", 245643, 42},
		{"Michael", "8934", 32432, 17},
		{"Jenny", "24334", 32444, 26},
	}
	fmt.Println(employees)
	sort.Sort(SortByAge(employees))
	sort.Slice(employees, func(i int, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	fmt.Println(employees)
}
