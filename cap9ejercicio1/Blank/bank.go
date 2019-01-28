package bank

type withdrawMessage struct {
	amount     int
	sufficient chan<- bool
}

var deposits = make(chan int)              
var balances = make(chan int)              
var withdraws = make(chan withdrawMessage) 


func Deposit(amount int) {
	deposits <- amount
}


func Withdraw(amount int) bool {
	sufficient := make(chan bool)
	withdraws <- withdrawMessage{amount, sufficient}
	return <-sufficient
}


func Balance() int {
	return <-balances
}

func teller() {
	var balance int 
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case msg := <-withdraws:
			ok := balance >= msg.amount
			if ok {
				balance -= msg.amount
			}
			msg.sufficient <- ok
		case balances <- balance:
		}
	}
}

func init() {
	go teller() 
}