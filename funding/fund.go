package funding

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

// A regular function returning a pointer to a fund
func NewFund(initialBalance int) *Fund {
	// We can return a pointer to a new struct without worrying about
	// whether it's on the stack or heap: Go figures that out for us.
	return &Fund{
		balance: initialBalance,
	}
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}

type Transactor func(fund *Fund)

type TransactionCommand struct {
	Transactor Transactor
	Done       chan bool
}

type FundServer struct {
	commands chan TransactionCommand
	fund     *Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		commands: make(chan TransactionCommand),
		fund:     NewFund(initialBalance),
	}

	go server.loop()

	return server
}

func (s *FundServer) Transact(transactor Transactor) {
	command := TransactionCommand{
		Transactor: transactor,
		Done:       make(chan bool),
	}
	s.commands <- command
	<-command.Done
}

func (s *FundServer) Balance() int {
	var balance int

	s.Transact(func(f *Fund) {
		balance = f.Balance()
	})

	return balance
}

func (s *FundServer) Withdraw(amount int) {
	s.Transact(func(f *Fund) {
		f.Withdraw(amount)
	})
}

func (s *FundServer) loop() {
	for transaction := range s.commands {
		transaction.Transactor(s.fund)
		transaction.Done <- true
	}
}
