package pkg

type User struct {
	Username string
	Password string
	Email    string
	Budgets  []Budget
	Wallet   []Account
}

type Budget struct {
	Name       string
	Amount     int
	UsedAmount int
	State      int
}

type Account struct {
	Name   string
	Amount int
}
