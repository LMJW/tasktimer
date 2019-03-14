package main

import (
	"fmt"
	"sync"
	"time"
)

// Task contains the condition for the task
type Task struct {
	Name     string
	Content  string
	Duration time.Duration
	tick     time.Duration
	counter  int
	pause    bool
	mux      sync.Mutex

	stopchan chan bool
	Ticker   *time.Ticker
}

// NewTask create a Task struct
func NewTask(name string, content string, duration time.Duration, tick time.Duration) *Task {
	return &Task{
		Name:     name,
		Content:  content,
		Duration: duration,
		tick:     tick,
		counter:  int(duration) / int(tick),
		pause:    false,

		stopchan: make(chan bool),
		Ticker:   time.NewTicker(tick),
	}
}

// Start the task
func (t *Task) Start() {
	for {
		select {
		case <-t.stopchan:
			fmt.Println("task done!")
			t.Stop()
			return

		default:
			tt := <-t.Ticker.C

			fmt.Println("time is ticking")
			if !t.pause && t.counter > 0 {
				fmt.Println("current time: ", tt)
				t.counter--
			} else if t.pause {
				continue
			} else {
				t.stopchan <- true
			}
		}
	}
}

// Pause the task
func (t *Task) Pause(s bool) {
	t.mux.Lock()
	defer t.mux.Unlock()
	t.pause = s
}

// Stop the task
func (t *Task) Stop() {
	fmt.Println("stopped")
}

// Extend the current task timer
func (t *Task) Extend(duration time.Duration) {
	et := int(duration) / int(t.tick)
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.counter <= 0 {
		t.counter += et
		t.Start()
	} else {
		t.counter += et
	}
}
