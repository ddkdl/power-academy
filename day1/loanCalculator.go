package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Loan Payment Calculator")

	var loanAmount float64
	var yearlyPayments float64
	var loanTerm float64

	const interestRate = 0.065

	// Read User Input
	fmt.Print("Enter Loan Amount: ")
	fmt.Scan(&loanAmount)

	fmt.Print("Enter Number of Yearly Payments: ")
	fmt.Scan(&yearlyPayments)

	fmt.Print("Enter Loan Term (in years): ")
	fmt.Scan(&loanTerm)

	// Calculate Monthly Payment Amount
	numerator := interestRate * loanAmount
	denominator := yearlyPayments * (1 - math.Pow(1+(interestRate/yearlyPayments), -1*yearlyPayments*loanTerm))

	monthlyPaymentAmount := numerator / denominator

	// Calculate Total Amount To Pay (With interest)
	totalAmountWithInterest := yearlyPayments*monthlyPaymentAmount*loanTerm - loanAmount

	fmt.Printf("Monthly payment amount: $%.2f\n", monthlyPaymentAmount)
	fmt.Printf("Total amount to be paid at the end of the loan term: $%.2f\n", totalAmountWithInterest)
}
