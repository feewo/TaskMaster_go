package entity

import (
	"taskmaster/db"
)

func init() {
	db.Add(MigrateUser)
	db.Add(MigrateTask)
	db.Add(MigrateTaskPoint)
	db.Add(MigrateToken)
}
