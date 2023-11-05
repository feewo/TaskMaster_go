package storage

import (
	"taskmaster/db"
	"taskmaster/entity"
)

func TaskCreate(task entity.Task) *entity.Task {
	db.DB().Create(&task)
	return &task
}
func TaskGetAll() []*entity.Task {
	var tasks []*entity.Task
	db.DB().Find(&tasks)
	return make([]*entity.Task, 0)
}
func TaskGet(id uint32) *entity.Task {
	var task entity.Task
	db.DB().Table(task.TableName()).Where("tid = ?", id).Find(&task)
	return &task
}

func TaskDelete(id uint32) *entity.Task {
	var task entity.Task
	db.DB().Table(task.TableName()).Where("tid = ?", id).Find(&task)
	db.DB().Delete(&task)
	return &task
}

func TaskUpdate(task entity.Task, id uint32) *entity.Task {
	db.DB().Save(&task)
	return &task
}

// type TaskMx struct {
// 	myx      sync.RWMutex
// 	iterTask uint32
// 	tasks    map[uint32]entity.Task
// }

// var taskMx TaskMx

// func init() {
// 	taskMx = TaskMx{
// 		tasks: make(map[uint32]entity.Task),
// 	}
// }

// var tasks []entity.Task
// var iterTask uint32

// func init() {
// 	tasks = make([]entity.Task, 0)
// 	iterTask = 0
// }
// func TaskCreate(task entity.Task) *entity.Task {
// 	taskMx.myx.Lock()
// 	defer taskMx.myx.Unlock()
// 	taskMx.iterTask++
// 	task.Id = taskMx.iterTask
// 	taskMx.tasks[taskMx.iterTask] = task
// 	return &task
// }

// func TaskGetAll() []entity.Task {
// 	taskMx.myx.RLock()
// 	defer taskMx.myx.RUnlock()
// 	lst := make([]entity.Task, len(taskMx.tasks))
// 	iterTask := 0
// 	for key := range taskMx.tasks {
// 		lst[iterTask] = taskMx.tasks[key]
// 		iterTask++
// 	}
// 	return lst
// }

// func TaskGet(uid uint32) *entity.Task {
// 	taskMx.myx.RLock()
// 	defer taskMx.myx.RUnlock()
// 	if el, ok := taskMx.tasks[uid]; ok {
// 		return &el
// 	}
// 	return nil
// }

// func TaskDelete(id uint32) *entity.Task {
// 	taskMx.myx.Lock()
// 	defer taskMx.myx.Unlock()
// 	delete(taskMx.tasks, id)
// 	return nil
// }

// func TaskUpdate(task entity.Task, id uint32) *entity.Task {
// 	taskMx.myx.Lock()
// 	defer taskMx.myx.Unlock()
// 	task.Id = id
// 	taskMx.tasks[id] = task
// 	return &task
// }