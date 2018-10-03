package queue_test

import (
	"sync"
	"testing"
	"learn-example/28-copy/datadog/internal/queue"
	"fmt"
	"time"
)

type buf struct {
	mu sync.RWMutex
	b string
}

func (b *buf) Write(s string){
	b.mu.Lock()
	defer b.mu.Unlock()

	b.b += s
}

func (b * buf)String() string{
	b.mu.RLock()
	defer  b.mu.RUnlock()

	return b.b
}

func TestQueue(t *testing.T){
	var b buf
	q := queue.New(1,1)
	black := make(chan struct{})

	go func() {
		time.Sleep(time.Microsecond* 20)
		b.Write("2")
		black <- struct{}{}
	}()


	if err := q.Push(func() {
		b.Write("1")
		<-black
	}); err != nil {
		t.Fatal(err)
	}

	if err := q.Push(func() {
		b.Write("3")

	}); err != nil {
		t.Fatal(err)
	}

	q.Wait()

	if b.String() != "123"{
		t.Fatal(b.String() + " != 123")
	}else {
		fmt.Println(b.String())
	}
}

func TestQueue_Close(t *testing.T) {
	q := queue.New(1, 1)
	black := make(chan struct{})

	if err := q.Push(func() {
		fmt.Println("我阻塞了...")
		<-black
		fmt.Println("我阻塞了解开了...")
	}); err != nil {
		t.Fatal(err)
	}

	go func() {
		time.Sleep(time.Microsecond*200)
		black <- struct{}{}
	}()

	q.Close()

	if err := q.Push(func() {

	}); err ==nil {
		t.Fatal("expecting an error")
	} else if err != queue.ErrorQueueClosed  {
		t.Fatal(err)
	}
}