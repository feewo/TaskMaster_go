package storage

import (
	"taskmaster/db"
	"taskmaster/entity"
	"time"

	"github.com/google/uuid"
)

func UserCreate(user entity.User) *entity.User {
	db.DB().Create(&user)
	var result entity.User
	db.DB().Select("id", "login", "surname", "name", "patronymic", "email", "role").Where("ID = ?", user.ID).Find(&result)
	return &result
}
func UserGetAll() []*entity.User {
	var users []*entity.User
	db.DB().Select("id", "login", "surname", "name", "patronymic", "email", "role").Find(&users)
	return users
}

func UserGet(id uint) *entity.User {
	var user entity.User
	db.DB().Select("id", "login", "surname", "name", "patronymic", "email", "role").Where("ID = ?", id).Find(&user)
	return &user
}

func UserDelete(id uint) *entity.User {
	var user entity.User
	db.DB().Where("ID = ?", id).Delete(&user)
	return &user
}

func UserUpdate(user entity.User, id uint) *entity.User {
	db.DB().Where("ID = ?", id).Updates(&user)
	updatedUser := UserGet(id)
	return updatedUser
}

func UserAuth(usr entity.User) *entity.Token {
	token := entity.Token{
		UserID:  usr.ID,
		Token:   uuid.NewString(),
		Expired: time.Now().Add(1 * time.Hour),
	}
	db.DB().Create(&token)
	return &token
}

func UserAuthDelete(token []string) *entity.Token {
	var TableToken entity.Token
	db.DB().Where("token = ?", token).Delete(&TableToken)
	return &TableToken
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
