package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/steevehook/expenses/printer"
)

const salami = "Salami"

// expenses struct literal
var expenses = map[int]struct {
	Title string
	Price float64
}{}
var id int

func main() {
	fmt.Println("Welcome to Expenses app")
	printer.Print()

	createExpense("Milk 1L", 2.5)
	createExpense(salami, 5)

	fmt.Println(getExpense(0))
	fmt.Println(getExpense(1))
	fmt.Println(getExpense(2)) // zero value
	fmt.Println(getAllExpenses())

	editExpense(1, "", 7.23)
	fmt.Println(getExpense(1))

	deleteExpense(0)
	deleteExpense(2) // does not error
	fmt.Println(getAllExpenses())
}

func getExpense(id int) struct {
	Title string
	Price float64
} {
	expense, found := expenses[id]
	if !found {
		log.Println("could not find expense with id:", id)
	}
	return expense
}

func getAllExpenses() []struct {
	Title string
	Price float64
} {
	res := make([]struct {
		Title string
		Price float64
	}, 0, len(expenses))
	for _, expense := range expenses {
		res = append(res, expense)
	}
	return res
}

func createExpense(title string, price float64) {
	expenses[id] = struct {
		Title string
		Price float64
	}{
		Title: title,
		Price: price,
	}
	id++
}

func editExpense(id int, title string, price float64) {
	expense := getExpense(id)
	if strings.TrimSpace(title) != "" {
		expense.Title = title
	}
	if price > 0 {
		expense.Price = price
	}
	expenses[id] = expense // we need to update original with copy we received
}

func deleteExpense(id int) {
	delete(expenses, id)
}
