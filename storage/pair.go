package storage

// Pair store account and password
type Pair struct {
	Account  string `json:"A"`
	Password string `json:"P"`
}

// NewPair create a new pair and return
func NewPair(account, password string) Pair {
	return Pair{Account: account, Password: password}
}
