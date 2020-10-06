package wallet

import (
	"fmt"
	"testing"

	"github.com/tohisroilov/wallet/pkg/types"
)

type testService struct {
	*Service
}

// newTestService ...
func newTestService() *testService {
	return &testService{Service: &Service{}}
}

type testAccount struct {
	phone    types.Phone
	balance  types.Money
	payments []struct {
		amount   types.Money
		category types.PaymentCategory
	}
}

var defaultTestAccount = testAccount{
	phone:   "+992932222272",
	balance: 10_000_00,
	payments: []struct {
		amount   types.Money
		category types.PaymentCategory
	}{
		{amount: 1_000_00, category: "auto"},
	},
}

func (s *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	account, err := s.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error = %v", err)
	}

	err = s.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposity account, error = %v", err)
	}

	// создаем слайс нужной длины, поскольку знаем размер
	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments {
		payments[i], err = s.Pay(account.ID, payment.amount, payment.category)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error = %v", err)
		}
	}
	return account, payments, nil
}

//

func TestRegisterAccount(t *testing.T) {
	testCases := []struct {
		name  string
		phone types.Phone
		err   error
	}{
		{
			name:  "no error",
			phone: "992000000001",
			err:   nil,
		},
		{
			name:  "phone already registered",
			phone: "992000000001",
			err:   ErrPhoneRegistered,
		},
	}
	svc := &Service{}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := svc.RegisterAccount(tC.phone)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestFindAccountByID(t *testing.T) {
	testCases := []struct {
		name string
		ID   int64
		err  error
	}{
		{
			name: "no error",
			ID:   1,
			err:  nil,
		},
		{
			name: "account not found",
			ID:   545,
			err:  ErrAccountNotFound,
		},
	}
	svc := newTestService()
	svc.RegisterAccount("992000068799")

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := svc.FindAccountByID(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestFindPaymentByID(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	testCases := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "payment not found",
			ID:   "545sEa",
			err:  ErrPaymentNotFound,
		},
	}

	for _, payment := range payments {
		testCases = append(testCases,
			struct {
				name string
				ID   string
				err  error
			}{
				name: "no error",
				ID:   payment.ID,
				err:  nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := s.FindPaymentByID(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestReject(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	testCases := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "payment not found",
			ID:   "545sEa",
			err:  ErrPaymentNotFound,
		},
	}

	for _, payment := range payments {
		testCases = append(testCases,
			struct {
				name string
				ID   string
				err  error
			}{
				name: "no error",
				ID:   payment.ID,
				err:  nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			err := s.Reject(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	testCases := []struct {
		name string
		ID   string
		err  error
	}{
		{
			name: "payment not found",
			ID:   "545sEa",
			err:  ErrPaymentNotFound,
		},
	}

	for _, payment := range payments {
		testCases = append(testCases,
			struct {
				name string
				ID   string
				err  error
			}{
				name: "no error",
				ID:   payment.ID,
				err:  nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := s.Repeat(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestFavoritePayment(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	testCases := []struct {
		name        string
		ID          string
		paymentName string
		err         error
	}{
		{
			name:        "payment not found",
			ID:          "545sEa",
			paymentName: "notFound",
			err:         ErrPaymentNotFound,
		},
	}

	for i, payment := range payments {
		testCases = append(testCases,
			struct {
				name        string
				ID          string
				paymentName string
				err         error
			}{
				name:        "no error",
				ID:          payment.ID,
				paymentName: fmt.Sprintf("payment #%v", i),
				err:         nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := s.FavoritePayment(tC.ID, tC.paymentName)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestFindFavoriteByID(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	for i, payment := range payments {
		paymentName := fmt.Sprintf("payment #%v", i)
		_, err := s.FavoritePayment(payment.ID, paymentName)
		if err != nil {
			t.Error(err)
			return
		}
	}

	testCases := []struct {
		name        string
		ID          string
		paymentName string
		err         error
	}{
		{
			name:        "payment not found",
			ID:          "545sEa",
			paymentName: "notFound",
			err:         ErrFavoriteNotFound,
		},
	}

	for _, favorite := range s.favorites {
		testCases = append(testCases,
			struct {
				name        string
				ID          string
				paymentName string
				err         error
			}{
				name:        "no error",
				ID:          favorite.ID,
				paymentName: favorite.Name,
				err:         nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := s.FindFavoriteByID(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}

func TestPayFromFavorite(t *testing.T) {
	s := newTestService()
	_, payments, err := s.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}

	for i, payment := range payments {
		paymentName := fmt.Sprintf("payment #%v", i)
		_, err := s.FavoritePayment(payment.ID, paymentName)
		if err != nil {
			t.Error(err)
			return
		}
	}

	testCases := []struct {
		name        string
		ID          string
		paymentName string
		err         error
	}{
		{
			name:        "payment not found",
			ID:          "545sEa",
			paymentName: "notFound",
			err:         ErrFavoriteNotFound,
		},
	}

	for _, favorite := range s.favorites {
		testCases = append(testCases,
			struct {
				name        string
				ID          string
				paymentName string
				err         error
			}{
				name:        "no error",
				ID:          favorite.ID,
				paymentName: favorite.Name,
				err:         nil,
			},
		)
	}

	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := s.PayFromFavorite(tC.ID)
			if err != tC.err {
				t.Errorf("invalid result in tc=%v, expected: %v, actual: %v", tC.name, tC.err, err)
			}
		})
	}
}
