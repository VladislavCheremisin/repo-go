package main

import (
	"math/rand"
	"sync"
	"testing"
)

type SetMutex struct {
	set map[float64]struct{}
	mu  *sync.Mutex
	mu2 *sync.RWMutex
}

func NewSet() *SetMutex {
	return &SetMutex{
		set: map[float64]struct{}{},
		mu:  &sync.Mutex{},
		mu2: &sync.RWMutex{},
	}
}
func (s *SetMutex) Get(x float64) {
	s.mu.Lock()
	s.set[x] = struct{}{}
	s.mu.Unlock()
}

// читаем в Some из карты(структуры)
func (s *SetMutex) Set(x float64) {
	s.mu.Lock()
	_ = s.set[x]
	s.mu.Unlock()
}
func (s *SetMutex) Get2(x float64) {
	s.mu2.Lock()
	s.set[x] = struct{}{}
	s.mu2.Unlock()
}

// читаем в Some из карты(структуры)
func (s *SetMutex) Set2(x float64) {
	s.mu2.Lock()
	_ = s.set[x]
	s.mu2.Unlock()
}

// задаем количество итераци
var iterations int = 1000

// sync.Mutex запись 50% чтение 50%
func BenchmarkMutex50w_50r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < iterations; k++ {
			wg.Add(1)
			if k < iterations/2 {
				go func() {
					structSet.Set(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}

//sync.RWMutex запись 50% чтение 50%
func BenchmarkRWMutex50w_50r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < iterations; k++ {
			wg.Add(1)
			if k < (iterations / 2) {
				go func() {
					structSet.Set2(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get2(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}

// sync.Mutex запись 10% чтение 90%
func BenchmarkMutex10w_90r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < iterations; k++ {
			wg.Add(1)
			if k%10 == 0 {
				go func() {
					structSet.Set(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}

//sync.RWMutex запись 10% чтение 90%
func BenchmarkRWMutex10w_90r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < 1000; k++ {
			wg.Add(1)
			if k%10 == 0 {
				go func() {
					structSet.Set2(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get2(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}

// sync.Mutex запись 90% чтение 10%
func BenchmarkMutex90w_10r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < iterations; k++ {
			wg.Add(1)
			if k%10 != 0 {
				go func() {
					structSet.Set(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}

//sync.RWMutex запись 90% чтение 10%
func BenchmarkRWMutex90w_10r(b *testing.B) {
	var wg = &sync.WaitGroup{}
	structSet := NewSet()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for k := 0; k < iterations; k++ {
			wg.Add(1)
			if k%10 != 0 {
				go func() {
					structSet.Set2(rand.Float64())
					wg.Done()
				}()
			} else {
				go func() {
					structSet.Get2(rand.Float64())
					wg.Done()
				}()
			}
		}
	}
	wg.Wait()
}
