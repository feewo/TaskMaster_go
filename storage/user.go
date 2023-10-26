package storage

import (
	"sync"
	"taskmaster/entity"
)

type ItemMx struct {
	myx   sync.RWMutex
	iter  uint32
	users map[uint32]entity.User
}

var itemMx ItemMx

func init() {
	itemMx = ItemMx{
		users: make(map[uint32]entity.User),
	}
}

var users []entity.User
var iter uint32

func init() {
	users = make([]entity.User, 0)
	iter = 0
}
func UserCreate(item entity.User) *entity.User {
	itemMx.myx.Lock()
	defer itemMx.myx.Unlock()
	itemMx.iter++
	item.Id = itemMx.iter
	itemMx.users[itemMx.iter] = item
	return &item
}

func UserGetAll() []entity.User {
	itemMx.myx.RLock()
	defer itemMx.myx.RUnlock()
	lst := make([]entity.User, len(itemMx.users))
	iter := 0
	for key := range itemMx.users {
		lst[iter] = itemMx.users[key]
		iter++
	}
	return lst
}

func UserGet(uid uint32) *entity.User {
	itemMx.myx.RLock()
	defer itemMx.myx.RUnlock()
	if el, ok := itemMx.users[uid]; ok {
		return &el
	}
	return nil
}

func UserDelete(id uint32) *entity.User {
	itemMx.myx.Lock()
	defer itemMx.myx.Unlock()
	delete(itemMx.users, id)
	return nil
}

func UserUpdate(item entity.User, id uint32) *entity.User {
	itemMx.myx.Lock()
	defer itemMx.myx.Unlock()
	item.Id = itemMx.iter
	itemMx.users[itemMx.iter] = item
	return &item
}
