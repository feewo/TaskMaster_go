package storage

import (
	"sync"
	"taskmaster/entity"
)

type TaskMx struct {
	myx      sync.RWMutex
	iterTask uint32
	tasks    map[uint32]entity.Task
}

var taskMx TaskMx

func init() {
	taskMx = TaskMx{
		tasks: make(map[uint32]entity.Task),
	}
}

var tasks []entity.Task
var iterTask uint32

func init() {
	tasks = make([]entity.Task, 0)
	iterTask = 0
}
func TaskCreate(task entity.Task) *entity.Task {
	taskMx.myx.Lock()
	defer taskMx.myx.Unlock()
	taskMx.iterTask++
	task.Id = taskMx.iterTask
	taskMx.tasks[taskMx.iterTask] = task
	return &task
}

func TaskGetAll() []entity.Task {
	taskMx.myx.RLock()
	defer taskMx.myx.RUnlock()
	lst := make([]entity.Task, len(taskMx.tasks))
	iterTask := 0
	for key := range taskMx.tasks {
		lst[iterTask] = taskMx.tasks[key]
		iterTask++
	}
	return lst
}

func TaskGet(uid uint32) *entity.Task {
	taskMx.myx.RLock()
	defer taskMx.myx.RUnlock()
	if el, ok := taskMx.tasks[uid]; ok {
		return &el
	}
	return nil
}

func TaskDelete(id uint32) *entity.Task {
	taskMx.myx.Lock()
	defer taskMx.myx.Unlock()
	delete(taskMx.tasks, id)
	return nil
}

func TaskUpdate(task entity.Task, id uint32) *entity.Task {
	taskMx.myx.Lock()
	defer taskMx.myx.Unlock()
	task.Id = taskMx.iterTask
	taskMx.tasks[taskMx.iterTask] = task
	return &task
}
