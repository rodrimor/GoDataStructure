package main

func main() {
	var newEmp Emplist
	newEmp.Hire("Joe Doe", "Janitor", 1, 17.50)
	newEmp.Hire("Jonnet Cole", "Cook", 2, 19.00)
	newEmp.Hire("Lesly Lazy", "Bum", 3, 99.00)
	//	newEmp.Println()
	newEmp.Fire(3)

	newEmp.Println()
}
