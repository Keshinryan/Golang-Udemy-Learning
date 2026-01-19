package belajar_golang_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContest(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)

}

func TestContextWithValue(t *testing.T) {
	ContextA := context.Background()
	ContextB := context.WithValue(ContextA, "name", "Eko")
	ContextC := context.WithValue(ContextB, "name", "Budi")
	ContextD := context.WithValue(ContextC, "address", "Indonesia")	
	contextE := context.WithValue(ContextD, "name", "Joko")
	fmt.Println(ContextA)
	fmt.Println(ContextB)
	fmt.Println(ContextC)
	fmt.Println(ContextD)
	fmt.Println(contextE)
	fmt.Println(ContextD.Value("name"))
	fmt.Println(ContextD.Value("address"))
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter:=1
		for {
			select{
				case <-ctx.Done():
					return
				default: 
					destination<-counter
					counter++
					time.Sleep(1*time.Second)
			}
		}	
	}()

	return destination
}


func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutines",runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutines",runtime.NumGoroutine())
	
	for n := range destination {
		fmt.Println("Counter",n)
		if n == 10 {
			break
		}
	}
	cancel()
	
	time.Sleep(2 *  time.Second)
	fmt.Println("Total Goroutines",runtime.NumGoroutine())
}

func TestContextWithCancelWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutines",runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutines",runtime.NumGoroutine())
	
	for n := range destination {
		fmt.Println("Counter",n)
	}
	
	time.Sleep(2 *  time.Second)
	fmt.Println("Total Goroutines",runtime.NumGoroutine())
}

func TestContextWithCancelWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutines",runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutines",runtime.NumGoroutine())
	
	for n := range destination {
		fmt.Println("Counter",n)
	}
	
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutines",runtime.NumGoroutine())
}