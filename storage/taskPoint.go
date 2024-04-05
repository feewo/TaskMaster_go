package storage

import (
	"taskmaster/db"
	"taskmaster/entity"
)

func TaskPointCreate(taskpoint entity.TaskPoint) *entity.TaskPoint {
	db.DB().Create(&taskpoint)
	return &taskpoint
}
func TaskPointGetAll() []*entity.TaskPoint {
	var taskpoints []*entity.TaskPoint
	db.DB().Find(&taskpoints)
	return taskpoints
}
func TaskPointGet(id uint32) *entity.TaskPoint {
	var taskpoint entity.TaskPoint
	db.DB().Table(taskpoint.TableName()).Where("ID = ?", id).Find(&taskpoint)
	return &taskpoint
}

func TaskPointDelete(id uint32) *entity.TaskPoint {
	var taskpoint entity.TaskPoint
	db.DB().Table(taskpoint.TableName()).Where("ID = ?", id).Delete(&taskpoint)
	return &taskpoint
}

func TaskPointUpdate(taskPoint entity.TaskPoint, id uint) *entity.TaskPoint {
	// db.DB().Table("taskpoints").Where("ID = ?", id).Updates(taskPointMap)
	// updatedTaskPoint := &entity.TaskPoint{}
	// if pid, ok := taskPointMap["ID"]; ok {
	// 	db.DB().Table("taskpoint").Where("ID = ?", pid).First(updatedTaskPoint)
	// } else {
	// 	db.DB().Table("taskpoint").Where("ID = ?", id).First(updatedTaskPoint)
	// }
	// return updatedTaskPoint
	db.DB().Table(taskPoint.TableName()).Where("ID = ?", id).Updates(taskPoint)
	updatedTaskPoint := &entity.TaskPoint{}
	if updatedTaskPoint.ID == 0 {
		db.DB().Table(taskPoint.TableName()).Where("ID = ?", id).First(updatedTaskPoint)
	} else {
		db.DB().Table(taskPoint.TableName()).Where("ID = ?", taskPoint.ID).First(updatedTaskPoint)
	}
	return updatedTaskPoint
}

// type TaskPointMx struct {
// 	myx           sync.RWMutex
// 	iterTaskPoint uint32
// 	tasksPoint    map[uint32]entity.TaskPoint
// }

// var taskPointMx TaskPointMx

// func init() {
// 	taskPointMx = TaskPointMx{
// 		tasksPoint: make(map[uint32]entity.TaskPoint),
// 	}
// }

// var tasksPoint []entity.TaskPoint
// var iterTaskPoint uint32

// func init() {
// 	tasksPoint = make([]entity.TaskPoint, 0)
// 	iterTaskPoint = 0
// }
// func TaskPointCreate(taskPoint entity.TaskPoint) *entity.TaskPoint {
// 	taskPointMx.myx.Lock()
// 	defer taskPointMx.myx.Unlock()
// 	taskPointMx.iterTaskPoint++
// 	taskPoint.Id = taskPointMx.iterTaskPoint
// 	taskPointMx.tasksPoint[taskPointMx.iterTaskPoint] = taskPoint
// 	return &taskPoint
// }

// func TaskPointGetAll() []entity.TaskPoint {
// 	taskPointMx.myx.RLock()
// 	defer taskPointMx.myx.RUnlock()
// 	lst := make([]entity.TaskPoint, len(taskPointMx.tasksPoint))
// 	iterTaskPoint := 0
// 	for key := range taskPointMx.tasksPoint {
// 		lst[iterTaskPoint] = taskPointMx.tasksPoint[key]
// 		iterTaskPoint++
// 	}
// 	return lst
// }

// func TaskPointGet(uid uint32) *entity.TaskPoint {
// 	taskPointMx.myx.RLock()
// 	defer taskPointMx.myx.RUnlock()
// 	if el, ok := taskPointMx.tasksPoint[uid]; ok {
// 		return &el
// 	}
// 	return nil
// }

// func TaskPointDelete(id uint32) *entity.TaskPoint {
// 	taskPointMx.myx.Lock()
// 	defer taskPointMx.myx.Unlock()
// 	delete(taskPointMx.tasksPoint, id)
// 	return nil
// }

// func TaskPointUpdate(taskPoint entity.TaskPoint, id uint32) *entity.TaskPoint {
// 	taskPointMx.myx.Lock()
// 	defer taskPointMx.myx.Unlock()
// 	taskPoint.Id = id
// 	taskPointMx.tasksPoint[id] = taskPoint
// 	return &taskPoint
// }
