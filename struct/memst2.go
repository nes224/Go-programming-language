package main

import "fmt"

type Employee struct {
	name   string
	salary int
}

type Project struct {
	name       string
	invest     int
	supervisor Employee
}

func main() {
	pj1 := Project{
		name:       "Skywalk",
		invest:     18000000,
		supervisor: Employee{"TesterDev", 50000},
	}

	fmt.Println("Details of Project1")
	fmt.Println("Project name : ", pj1.name)
	fmt.Println("Invesment values :", pj1.invest)
	fmt.Println("\nDetails of Supervisor")
	fmt.Println("Supervisor name : ", pj1.supervisor.name)
	fmt.Println("Supervisor salary : ", pj1.supervisor.salary)
}
