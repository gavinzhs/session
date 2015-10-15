package main

import (
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

const BALANCE = 100

type SimpleAccount struct {
	balance int
}

func NewSimpleAccount(b int) *SimpleAccount {
	return &SimpleAccount{balance: b}
}

func (s *SimpleAccount) setBalance(balance int) {
	s.addSomelatency()
	s.balance = balance
}

func (s *SimpleAccount) addSomelatency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}

func (s *SimpleAccount) WithDraw(num uint) {
	if s.balance < int(num) {
		log.Println("账号余额不足!")
	}
	s.setBalance(s.balance - int(num))
}

func (s *SimpleAccount) SaveMenoy(num uint) {
	s.setBalance(s.balance + int(num))
}

func (s *SimpleAccount) Balance() int {
	return s.balance
}

func TestAccount(t *testing.T) {
	bank := NewBank(NewSimpleAccount(BALANCE))
	bank.SaveMenoy(100)
	if bank.Balance() != 200 {
		t.Fatalf("余额应为200,实际为%d", bank.Balance())
	}
}

type LockSimpleAccount struct {
	lock    sync.Mutex
	account *SimpleAccount
}

func NewLockSimpleAccount(acc *SimpleAccount) *LockSimpleAccount {
	return &LockSimpleAccount{account: acc}
}

func (l *LockSimpleAccount) WithDraw(num uint) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.account.WithDraw(num)
}

func (l *LockSimpleAccount) SaveMenoy(num uint) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.account.SaveMenoy(num)
}

func (l *LockSimpleAccount) Balance() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.account.Balance()
}

type ConcurrentAccount struct {
	account   *SimpleAccount
	withDraw  chan uint
	saveMenoy chan uint
	balance   chan chan int
}

func NewConcurrentAccount(acc *SimpleAccount) *ConcurrentAccount {
	return &ConcurrentAccount{
		account:   acc,
		withDraw:  make(chan uint),
		saveMenoy: make(chan uint),
		balance:   make(chan chan int),
	}
}

func (c *ConcurrentAccount) WithDraw(num uint) {
	c.withDraw <- num
}

func (c *ConcurrentAccount) SaveMenoy(num uint) {
	c.saveMenoy <- num
}

func (c *ConcurrentAccount) Balance() int {
	ch := make(chan int)
	c.balance <- ch
	return <-ch
}

func (c *ConcurrentAccount) listen() {
	go func() {
		for {
			select {
			case withDraw := <-c.withDraw:
				c.account.WithDraw(withDraw)
			case saveMenoy := <-c.saveMenoy:
				c.account.SaveMenoy(saveMenoy)
			case ch := <-c.balance:
				ch <- c.account.Balance()
			}
		}
	}()
}

func TestAccountGo(t *testing.T) {
	//如何保证结束异步之后，再来进行主线程       核心就是阻塞所有
	//    bank := NewBank(NewLockSimpleAccount(NewSimpleAccount(BALANCE)))
	concurrent := NewConcurrentAccount(NewSimpleAccount(BALANCE))
	concurrent.listen()

	done := make(chan bool)

	go func() {
		concurrent.WithDraw(10)
		done <- true
	}()

	go func() {
		concurrent.WithDraw(10)
		done <- true
	}()

	log.Println(<-done)

	<-done

	log.Println(concurrent.Balance())
}
