package main

/*
   这里主要是用来测试gon中得chan的使用

   这里是一个银行账号系统
   1. 定义一个账号的接口用来控制账号的基本属性
   2. 定义一个银行来包含账号，做账号的基本工作
   3. 定义一个账号实体来实际测试
*/

type Account interface {
	WithDraw(uint)
	SaveMenoy(uint)
	Balance() int
}

type Bank struct {
	account Account
}

func NewBank(acc Account) *Bank {
	return &Bank{account: acc}
}

func (b *Bank) WithDraw(num uint) {
	b.account.WithDraw(num)
}

func (b *Bank) SaveMenoy(num uint) {
	b.account.SaveMenoy(num)
}

func (b *Bank) Balance() int {
	return b.account.Balance()
}
