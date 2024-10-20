package main

import (
	"fmt"
	"os"
)

func main() {
	var rows, seats int
	fmt.Println("Enter the number of rows:")
	fmt.Scan(&rows)

	fmt.Println("Enter the number of seats in each row:")
	fmt.Scan(&seats)

	// Create a 2D slice for seating arrangement
	cinema := make([][]string, rows)
	for i := range cinema {
		cinema[i] = make([]string, seats)
		for j := range cinema[i] {
			cinema[i][j] = "S" // Initialize all seats to "S"
		}
	}

	promptChoice(rows, seats, cinema)
}

func promptChoice(rows, seats int, cinema [][]string) {
	var option int
	fmt.Println("1. Show the seats")
	fmt.Println("2. Buy a ticket")
	fmt.Println("3. Statistics")
	fmt.Println("0. Exit")

	fmt.Scanln(&option)
	fmt.Println()
	switch option {
	case 1:
		showSeats(rows, seats, cinema)
	case 2:
		purchaseTicket(rows, seats, cinema)
	case 3:
		showStatistics(rows, seats, cinema)
	case 0:
		exitProgram()
	default:
		promptChoice(rows, seats, cinema)
	}
}

func showSeats(rows, seats int, cinema [][]string) {
	// Display the initial seating arrangement
	fmt.Println("Cinema:")
	fmt.Print("  ")
	for i := 1; i <= seats; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	for i := 0; i < rows; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < seats; j++ {
			fmt.Printf("%s ", cinema[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	promptChoice(rows, seats, cinema)
}

func purchaseTicket(rows, seats int, cinema [][]string) {
	var userRow, userSeat, ticketPrice int
	// Ask the user for row and seat number
	fmt.Println("Enter a row number:")
	fmt.Scan(&userRow)

	fmt.Println("Enter a seat number in that row:")
	fmt.Scan(&userSeat)

	defer func() {
		onPanic := recover() // we catch the out-of-bounds input

		if onPanic != nil {
			fmt.Println("Wrong input!")
			promptChoice(rows, seats, cinema) // re-initiate gracefully the control flow
		}
	}()

	// Calculate ticket price
	totalSeats := rows * seats
	frontHalf := rows / 2

	if totalSeats <= 60 || userRow <= frontHalf {
		ticketPrice = 10
	} else {
		ticketPrice = 8
	}

	// Mark the chosen seat as "B"
	if cinema[userRow-1][userSeat-1] == "B" {
		fmt.Println("That ticket has already been purchased")
		purchaseTicket(rows, seats, cinema)
	} else {
		fmt.Printf("Ticket price: $%d\n", ticketPrice)
		cinema[userRow-1][userSeat-1] = "B"
	}
	fmt.Println()
	promptChoice(rows, seats, cinema)
}

func showStatistics(rows, seats int, cinema [][]string) {
	var totalTickets, purchased, currentIncome, totalIncome int
	totalTickets = rows * seats
	frontHalf := rows / 2

	for i := 0; i < rows; i++ {
		for j := 0; j < seats; j++ {
			if cinema[i][j] == "B" && i+1 <= frontHalf {
				currentIncome += 10
				purchased++
			} else if cinema[i][j] == "B" && i+1 > frontHalf {
				currentIncome += 8
				purchased++
			}

			if totalTickets <= 60 || (cinema[i][j] == "B" && i <= frontHalf) {
				totalIncome = totalTickets * 10
			} else {
				totalIncome = (10 * seats * frontHalf) + (8 * seats * (rows - frontHalf))
			}
		}
	}

	fmt.Printf("Number of purchased tickets: %d", purchased)
	fmt.Printf("\nPercentage: %.2f%%", float64(purchased)/float64(totalTickets)*100.00)
	fmt.Printf("\nCurrent income: $%d", currentIncome)
	fmt.Printf("\nTotal income: $%d", totalIncome)
	fmt.Println()
	promptChoice(rows, seats, cinema)
}

func exitProgram() {
	os.Exit(0)
}
