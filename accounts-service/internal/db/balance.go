package db

import (
	"errors"
	"sync"

	event "github.com/Sam-Frost/accounts-service/internal/event-logger"
)

/*
 * balance struct to hold the user's current and locked balance
 * needs to be locked and unlock before/after access
 * avoids locking of the compelete map
 */
type balance struct {
	mu                   sync.Mutex
	sequence             int64
	availableBalanceTick int64
	lockedBalanceTick    int64
}

func (b *balance) GetUserBalance() (int64, int64) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.availableBalanceTick, b.lockedBalanceTick
}

// Increase the available user balance(Thread Safe)
func (b *balance) IncreaseAvailableBalance(amount int64, userId int32) int64 {
	event.EmitEvent(event.Event{
		EventId:  event.ADD_BALANCE,
		UserId:   userId,
		Amount:   amount,
		MarketId: 0,
		Sequence: b.sequence,
	})

	b.mu.Lock()
	defer b.mu.Unlock()

	b.availableBalanceTick += amount
	b.sequence++

	return b.availableBalanceTick
}

// Lock the available user balance(Thread Safe)
func (b *balance) LockBalance(amount int64, userId int32) (int64, int64, error) {
	event.EmitEvent(event.Event{
		EventId:  event.LOCK_BALANCE,
		UserId:   userId,
		Amount:   amount,
		MarketId: 0,
		Sequence: b.sequence,
	})

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.availableBalanceTick < amount {
		return 0, 0, errors.New("Not enough available balance to lock the amount")
	}
	b.availableBalanceTick -= amount
	b.lockedBalanceTick += amount
	b.sequence++

	return b.availableBalanceTick, b.lockedBalanceTick, nil
}

// Decrease the locked user balance(Thread Safe)
func (b *balance) DecreaseLockedBalance(amount int64, userId int32) error {
	event.EmitEvent(event.Event{
		EventId:  event.REDUCE_LOCKED_BALANCE,
		UserId:   userId,
		Amount:   amount,
		MarketId: 0,
		Sequence: b.sequence,
	})

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.lockedBalanceTick < amount {
		return errors.New("Not enough locked balance to reduce the amount")
	}
	b.lockedBalanceTick -= amount
	b.sequence++

	return nil
}
