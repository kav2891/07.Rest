package storage

import (
	"errors"
	"sync"
	"time"

	"tasks-api/internal/models"
)

type MemoryStorage struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}

func NewMemory() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

func (m *MemoryStorage) List() []models.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	result := make([]models.Task, 0, len(m.tasks))
	for _, t := range m.tasks {
		result = append(result, t)
	}
	return result
}

func (m *MemoryStorage) Create(task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	task.ID = m.nextID
	task.CreatedAt = time.Now().Format(time.RFC3339)
	m.tasks[m.nextID] = task
	m.nextID++

	return task, nil
}

func (m *MemoryStorage) Get(id int) (models.Task, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	task, ok := m.tasks[id]
	return task, ok
}

func (m *MemoryStorage) Update(id int, task models.Task) (models.Task, error) {
	if task.Title == "" {
		return models.Task{}, errors.New("title is required")
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	old, ok := m.tasks[id]
	if !ok {
		return models.Task{}, errors.New("not found")
	}

	task.ID = id
	task.CreatedAt = old.CreatedAt
	m.tasks[id] = task

	return task, nil
}

func (m *MemoryStorage) Delete(id int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.tasks[id]; !ok {
		return errors.New("not found")
	}

	delete(m.tasks, id)
	return nil
}
