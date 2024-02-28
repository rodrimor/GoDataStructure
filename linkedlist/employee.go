package main

import (
	"fmt"
	"time"
)

// Employee type represents an employee with their details
type Employee struct {
	Name         string    `name:json`          // Employee name
	Role         string    `role:json`          // Employee role
	HourlyWage   float64   `hourly_wage:json`   // Hourly wage of the employee
	Register     int       `register:json`      // Unique employee register number
	StartingDate time.Time `starting_date:json` // Date the employee started working
	Next         *Employee // Pointer to the next employee in the linked list
}

// Emplist type represents a list of employees using a linked list
type Emplist struct {
	Head *Employee // Head of the linked list (first employee)
}

// Hire function adds a new employee to the linked list
func (emp *Emplist) Hire(name, role string, reg int, wage float64) {
	// Create a new employee object
	empToHire := &Employee{
		Name:         name,
		Role:         role,
		Register:     reg,
		HourlyWage:   wage,
		StartingDate: time.Now(),
	}

	// If the list is empty, set the new employee as the head
	if emp.Head == nil {
		emp.Head = empToHire
		return
	}

	// Traverse the list to find the last node
	current := emp.Head
	for current.Next != nil {
		current = current.Next
	}

	// Append the new employee to the end of the list
	current.Next = empToHire
}

// Fire function removes an employee from the linked list by register number
func (emp *Emplist) Fire(reg int) {
	// If the list is empty, do nothing
	if emp.Head == nil {
		return
	}

	// If the head node has the matching register number, remove it
	if emp.Head.Register == reg {
		emp.Head = emp.Head.Next
		return
	}

	// Traverse the list to find the employee to remove
	current := emp.Head
	for current.Next != nil {
		// If the next node has the matching register number, remove it
		if current.Next.Register == reg {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
	}
}

// Println function prints the details of all employees in the list
func (emp *Emplist) Println() {
	current := emp.Head

	// Keep iterating until the end of the list is reached
	for current != nil {
		// Format employee details and print them
		employeeDetails := fmt.Sprintf(`
Name %s Working since %v
Role %s
Hourly Wage %.2f
--------------------
`, current.Name, current.StartingDate.Format("02/01/2006"), current.Role, current.HourlyWage)
		fmt.Printf(employeeDetails)
		current = current.Next
	}
	fmt.Println() // Print an empty line after all employees are printed
}
