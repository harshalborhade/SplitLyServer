package main

import "fmt"

type Transaction struct {
	transcationName string
	amount          float64
	splitDetails    map[string]float64
}

func main() {
	fmt.Println("Hello")
	//Input transaction name from the user
	fmt.Println("Enter the transaction name:")
	var transactionName string
	fmt.Scanln(&transactionName)
	//Input the number of splits from the user
	fmt.Println("Enter the number of splits:")
	var numberOfSplits int
	fmt.Scanln(&numberOfSplits)
	//Input the amount from the user
	fmt.Println("Enter the amount:")
	var amount float64
	fmt.Scanln(&amount)
	//Input the split details from the user
	fmt.Println("Enter the split details:")
	var splitDetails string
	fmt.Scanln(&splitDetails)
	//Print the transaction details
	fmt.Println("Transaction Name:", transactionName)
	fmt.Println("Number of Splits:", numberOfSplits)
	fmt.Println("Amount:", amount)
	fmt.Println("Split Details:", splitDetails)
}
