package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	close   bool
	lock    sync.RWMutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{initialDeposit, false, sync.RWMutex{}}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.close {
		return 0, false
	}
	a.close = true
	return a.balance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	if a.close {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.close {
		return 0, false
	}
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
