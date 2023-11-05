package storage

import (
	"taskmaster/db"
	"taskmaster/entity"
)

func UserCreate(user entity.User) *entity.User {
	db.DB().Create(&user)
	return &user
}
func UserGetAll() []*entity.User {
	var users []*entity.User
	db.DB().Find(&users)
	return make([]*entity.User, 100)
}
func UserGet(id uint32) *entity.User {
	var user entity.User
	db.DB().Table(user.TableName()).Where("tid = ?", id).Find(&user)
	return &user
}

func UserDelete(id uint32) *entity.User {
	var user entity.User
	db.DB().Table(user.TableName()).Where("tid = ?", id).Find(&user)
	db.DB().Delete(&user)
	return &user
}

func UserUpdate(user entity.User, id uint32) *entity.User {
	db.DB().Save(&user)
	return &user
}

// type UserMx struct {
// 	myx      sync.RWMutex
// 	iterUser uint32
// 	users    map[uint32]entity.User
// }

// var userMx UserMx

// func init() {
// 	userMx = UserMx{
// 		users: make(map[uint32]entity.User),
// 	}
// }

// var users []entity.User
// var iterUser uint32

// func init() {
// 	users = make([]entity.User, 0)
// 	iterUser = 0
// }
// func UserCreate(user entity.User) *entity.User {
// 	userMx.myx.Lock()
// 	defer userMx.myx.Unlock()
// 	userMx.iterUser++
// 	user.Id = userMx.iterUser
// 	userMx.users[userMx.iterUser] = user
// 	return &user
// }

// func UserGetAll() []entity.User {
// 	userMx.myx.RLock()
// 	defer userMx.myx.RUnlock()
// 	lst := make([]entity.User, len(userMx.users))
// 	iterUser := 0
// 	for key := range userMx.users {
// 		lst[iterUser] = userMx.users[key]
// 		iterUser++
// 	}
// 	return lst
// }

// func UserGet(uid uint32) *entity.User {
// 	userMx.myx.RLock()
// 	defer userMx.myx.RUnlock()
// 	if el, ok := userMx.users[uid]; ok {
// 		return &el
// 	}
// 	return nil
// }

// func UserDelete(id uint32) *entity.User {
// 	userMx.myx.Lock()
// 	defer userMx.myx.Unlock()
// 	delete(userMx.users, id)
// 	return nil
// }

// func UserUpdate(user entity.User, id uint32) *entity.User {
// 	userMx.myx.Lock()
// 	defer userMx.myx.Unlock()
// 	user.Id = id
// 	userMx.users[id] = user
// 	return &user
// }
