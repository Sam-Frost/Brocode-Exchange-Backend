package db

import (
	"errors"
	"sync"

	event "github.com/Sam-Frost/accounts-service/internal/event-logger"
	"github.com/Sam-Frost/common"
)

type db struct {
	mu   sync.RWMutex
	data map[int32]*UserData
}

var database db = db{
	mu:   sync.RWMutex{},
	data: make(map[int32]*UserData),
}

// type SpotBalance struct {
// 	AvailableBalanceTick int64
// 	LockedBalanceTick    int64
// }

type UserData struct {
	BalanceData *balance
	// MarketMap   map[int32]SpotBalance
}

/*
 * Creates a new user in the database map(Thread Safe)
 * Acquire locks over the complete map
 */
func CreateNewUser(userId int32) {
	database.mu.Lock()
	defer database.mu.Unlock()

	database.data[userId] = &UserData{
		BalanceData: createBalance(userId),
		// MarketMap:   make(map[int32]SpotBalance),
	}
}

// Create balance object in heap(probably) and return reference
func createBalance(userId int32) *balance {
	event.EmitEvent(event.Event{
		EventId:  event.CREATE_USER_BALANCE,
		UserId:   userId,
		Amount:   0,
		MarketId: 0,
		Sequence: 0,
	})
	return &balance{
		mu:                   sync.Mutex{},
		availableBalanceTick: 0,
		lockedBalanceTick:    0,
		sequence:             1,
	}
}

// Returns the user balance *object address* from in memory db(thread safe)
func GetUserBalanceData(userId int32) (*balance, error) {
	database.mu.RLock()
	userData, ok := database.data[userId]
	database.mu.RUnlock()

	if !ok {
		return nil, errors.New(common.USER_NOT_FOUND)
	}

	return userData.BalanceData, nil
}
