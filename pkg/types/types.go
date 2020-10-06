package types

// Money ...
type Money int64

// Currency ...
type Currency string

// PaymentCategory ...
type PaymentCategory string

// PaymentStatus ...
type PaymentStatus string

// Statuses
const (
	PaymentStatusOK         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

// PAN ...
type PAN string

// Card ...
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

// Phone ...
type Phone string

// Account ...
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

// Payment ...
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

// Favorite ...
type Favorite struct {
	ID        string
	AccountID int64
	Name      string
	Amount    Money
	Category  PaymentCategory
}
