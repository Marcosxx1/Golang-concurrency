package bankbalance

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func BankBalance() {
	// variable for bank balance
	var bankbalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("Initial account balance: R$%d.00", bankbalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 10},
		{Source: "Harvesting", Amount: 100},
	}

	wg.Add(len(incomes))

	//loop through 52 weeks and print out how much is made; keep a running total
	for index, income := range incomes {

		go func(index int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {

				balance.Lock()
				temporaryBalace := bankbalance
				temporaryBalace += income.Amount
				bankbalance = temporaryBalace
				balance.Unlock()

				fmt.Printf("One week %d, you earned R$%d.00 from %s\n", week, bankbalance, income.Source)
			}

		}(index, income)
	}

	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance: R$%d.00", bankbalance)
	fmt.Println()
}
