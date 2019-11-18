package main

import (
	"fmt"
	"sync"
)

type Account struct {
	Money  int
	Locker sync.Mutex
}

func (a *Account) SaveMoney(n int) {
	a.Locker.Lock()
	fmt.Println("before save,money=", a.Money)
	a.Money += n
	fmt.Println("after save,money=", a.Money)
	a.Locker.Unlock()

}

func (a *Account) GetMoney(n int) {
	a.Locker.Lock()
	fmt.Println("before get,money=", a.Money)
	a.Money -= n
	fmt.Println("after get,money=", a.Money)
	a.Locker.Unlock()
}

func (a *Account) PrintHistory() {
	fmt.Println("the current money=", a.Money)
}

var wg sync.WaitGroup
var locker sync.Mutex

func main() {
	a := Account{1000, locker}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			a.GetMoney(100)
			wg.Done()
		}()

	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			a.SaveMoney(100)
			wg.Done()
		}()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			a.PrintHistory()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("the final money=", a.Money)
}
