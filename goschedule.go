package goschedule

import (
	"errors"
	"sync"
	"time"
)

type TaskID int

type task struct {
	Hour     int
	Minute   int
	Function func()
	ID       TaskID
}

type Scheduler struct {
	tasksMutex sync.Mutex
	tasks      []task
	running    bool

	idMutex sync.Mutex
	lastId  TaskID

	stopChan chan struct{}
}

func NewScheduler() Scheduler {
	return Scheduler{}
}

func (s *Scheduler) getNewID() (id TaskID) {
	s.idMutex.Lock()
	defer s.idMutex.Unlock()

	s.lastId++

	return s.lastId
}

func (s *Scheduler) run() {
	minuteTicker := time.NewTicker(1 * time.Minute)
	for {
		select {
		case <-s.stopChan:
			return
		case now := <-minuteTicker.C:
			for _, task := range s.tasks {
				if now.Minute() == task.Minute && now.Hour() == task.Hour {
					go task.Function()
				}
			}
		}
	}
}

func (s *Scheduler) RemoveTask(id TaskID) (err error) {
	s.tasksMutex.Lock()
	defer s.tasksMutex.Unlock()

	var found bool = false

	for i, task := range s.tasks {
		found = true

		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
		}
	}

	if len(s.tasks) < 1 && s.running {
		s.stopChan <- struct{}{}
		close(s.stopChan)
		s.running = false
	}

	if !found {
		return errors.New("couldn't find any task with given ID")
	}

	return nil
}

func (s *Scheduler) AddTask(hour, minute int, taskFunction func()) (id TaskID) {
	s.tasksMutex.Lock()
	defer s.tasksMutex.Unlock()

	id = s.getNewID()

	s.tasks = append(s.tasks, task{
		Hour:     hour,
		Minute:   minute,
		Function: taskFunction,
		ID:       id,
	})

	if !s.running {
		s.running = true
		s.stopChan = make(chan struct{})
		go s.run()
	}

	return id
}
