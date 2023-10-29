package storage

import (
	"sync"
	"taskmaster/entity"
)

type TaskPointMx struct {
	myx           sync.RWMutex
	iterTaskPoint uint32
	tasksPoint    map[uint32]entity.TaskPoint
}

var taskPointMx TaskPointMx

func init() {
	taskPointMx = TaskPointMx{
		tasksPoint: make(map[uint32]entity.TaskPoint),
	}
}

var tasksPoint []entity.TaskPoint
var iterTaskPoint uint32

func init() {
	tasksPoint = make([]entity.TaskPoint, 0)
	iterTaskPoint = 0
}
func TaskPointCreate(taskPoint entity.TaskPoint) *entity.TaskPoint {
	taskPointMx.myx.Lock()
	defer taskPointMx.myx.Unlock()
	taskPointMx.iterTaskPoint++
	taskPoint.Id = taskPointMx.iterTaskPoint
	taskPointMx.tasksPoint[taskPointMx.iterTaskPoint] = taskPoint
	return &taskPoint
}

func TaskPointGetAll() []entity.TaskPoint {
	taskPointMx.myx.RLock()
	defer taskPointMx.myx.RUnlock()
	lst := make([]entity.TaskPoint, len(taskPointMx.tasksPoint))
	iterTaskPoint := 0
	for key := range taskPointMx.tasksPoint {
		lst[iterTaskPoint] = taskPointMx.tasksPoint[key]
		iterTaskPoint++
	}
	return lst
}

func TaskPointGet(uid uint32) *entity.TaskPoint {
	taskPointMx.myx.RLock()
	defer taskPointMx.myx.RUnlock()
	if el, ok := taskPointMx.tasksPoint[uid]; ok {
		return &el
	}
	return nil
}

func TaskPointDelete(id uint32) *entity.TaskPoint {
	taskPointMx.myx.Lock()
	defer taskPointMx.myx.Unlock()
	delete(taskPointMx.tasksPoint, id)
	return nil
}

func TaskPointUpdate(taskPoint entity.TaskPoint, id uint32) *entity.TaskPoint {
	taskPointMx.myx.Lock()
	defer taskPointMx.myx.Unlock()
	taskPoint.Id = id
	taskPointMx.tasksPoint[id] = taskPoint
	return &taskPoint
}
