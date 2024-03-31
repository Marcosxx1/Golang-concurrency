package producerconsumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfBreads = 10

var breadsMade, breadsFailed, total int

type Producer struct {
	data chan BreadOrder
	quit chan chan error
}

type BreadOrder struct {
	breadNumber int
	message     string
	success     bool
}

// Close method closes the producer, signaling that it should stop producing bread orders
// It send a signal through the 'quit' channel and waits for a response on a new channel 'ch'
// Once it recieves a response, it returns the error recivied from the 'quit' channel
//
// Reciever:
//
//	p *Producer: A pointer to a 'Producer' type
//	Any type of 'Producer' now has access to 'Close()' function.
func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makeBread(breadNumber int) *BreadOrder {
	breadNumber++
	if breadNumber <= NumberOfBreads {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recieved order number #%d!\n", breadNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd <= 5 {
			breadsFailed++
		} else {
			breadsMade++
		}

		total++

		fmt.Printf("Making bread #%d. It will take %d seconds....\n", breadNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for bread number #%d!", breadNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The baker quit for bread number #%d!", breadNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Bread order #%d is ready!", breadNumber)
		}

		return &BreadOrder{
			breadNumber: breadNumber,
			message:     msg,
			success:     success,
		}

	}

	return &BreadOrder{
		breadNumber: breadNumber,
	}
}

func bakery(breadMaker *Producer) {
	// keep track of which bread we are making
	var currentBreadBeingMade = 0
	// run forever or until we recieve a quit notification

	for {
		currentBread := makeBread(currentBreadBeingMade)
		if currentBread != nil {
			currentBreadBeingMade = currentBread.breadNumber

			select {
			case breadMaker.data <- *currentBread:

			case quitChan := <-breadMaker.quit:
				close(breadMaker.data)
				close(quitChan)
				return
			}
		}
		// try to make a bread
		// decision
	}
}

func ProducerConsumer() {
	// seed the radim number generator
	seed := int64(42)
	/* r :=  */ rand.New(rand.NewSource(seed))

	// print out a message
	color.Cyan("The bakery os open for busines!")
	color.Cyan("-------------------------------")

	// create a producer
	breadJob := &Producer{
		data: make(chan BreadOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background

	go bakery(breadJob)

	// create and run consumer

	for i := range breadJob.data {
		if i.breadNumber <= NumberOfBreads {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.breadNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is mad =( ")
			}
		} else {
			color.Cyan("Done making bread.")
			err := breadJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	//print out the ending message
	color.Cyan("-----------------")
	color.Cyan("Done for the day.")

	color.Cyan("We've made %d breads, but failed to make %dm with %d attempts in total.", breadsMade, breadsFailed, total)

	switch {
	case breadsFailed > 9:
		color.Red("Everithing went bad!!!")
	case breadsFailed >= 6:
		color.Red("Kinda bad day")
	case breadsFailed >= 4:
		color.Yellow("Ok day, I guess")
	case breadsFailed >= 2:
		color.Yellow("T'was a pretty good day!")
	default:
		color.Green("T'was a great day!")
	}
}
