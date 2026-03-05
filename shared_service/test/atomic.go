package main

import (
	"log"
	"sync/atomic"
	"time"

	"golang.org/x/exp/rand"
)

type SyncMap[K comparable, V any] struct {
	data atomic.Pointer[map[K]V]
}

// NewSyncMap создает новую потокобезопасную карту
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	m := make(map[K]V, 10000)
	sm := &SyncMap[K, V]{}
	sm.data.Store(&m)
	return sm
}

// Load безопасно читает значение из карты
func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	data := m.data.Load()
	value, ok = (*data)[key] // Чтение безопасно, так как карта только заменяется
	return
}

// Store безопасно записывает значение в карту
func (m *SyncMap[K, V]) Store(key K, value V) {
	for {
		oldMap := m.data.Load() // Загружаем текущую карту
		newMap := make(map[K]V, len(*oldMap)+1)
		for k, v := range *oldMap { // Копируем данные в новую карту
			newMap[k] = v
		}
		newMap[key] = value // Добавляем новое значение

		// Атомарно пытаемся заменить старую карту новой
		if m.data.CompareAndSwap(oldMap, &newMap) {
			return
		}
	}
}

// Delete безопасно удаляет ключ из карты
func (m *SyncMap[K, V]) Delete(key K) {
	for {
		oldMap := m.data.Load() // Загружаем текущую карту
		newMap := make(map[K]V, len(*oldMap)-1)
		for k, v := range *oldMap { // Копируем все, кроме удаляемого ключа
			if k != key {
				newMap[k] = v
			}
		}

		// Атомарно пытаемся заменить старую карту новой
		if m.data.CompareAndSwap(oldMap, &newMap) {
			return
		}
	}
}

func main() {

	m := NewSyncMap[int, int]()

	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			m.Store(rand.Intn(100), rand.Int())
		}
	}()

	go func() {
		for {
			time.Sleep(5 * time.Millisecond)
			log.Println(m.Load(rand.Intn(100)))
		}
	}()

	time.Sleep(time.Hour)
}
