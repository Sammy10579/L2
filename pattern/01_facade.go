package main

import (
	"errors"
	"fmt"
	"time"
)

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

/*
	Фасад позволяет скрыть сложные части подсистемы,
	давая доступ только к необходимым рычагам,
	тем самым упрощая использование системы для клиента,
Преимущества:
	* Изолирует клиентов от компонентов сложной подсистемы.
Недостатки:
	* Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
		Божественный_объект - описывающий объект, который хранит в себе «слишком много» или делает «слишком много».
*/

var (
	bank = Bank{
		Name: "Банк",
		Card: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-2",
		Balance: 7,
		Bank:    &bank,
	}
	user = User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = Shop{
		Name: "Магазин",
		Products: []Product{
			prod,
		},
	}
)

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

type Bank struct {
	Name string
	Card []Card
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

type User struct {
	Name string
	Card *Card
}

func (user *User) GetBalance() float64 {
	return user.Card.Balance
}

func (card Card) CheckBalance() error {
	println("[Карта] ЗАпрос в банк для проверки остатка")
	time.Sleep(time.Microsecond * 500)
	return card.Bank.CheckBalance(card.Name)
}

func (bank Bank) CheckBalance(cardNumber string) error {
	println(fmt.Sprintf("[Банс] Получение остатка по карте %s", cardNumber))
	time.Sleep(time.Microsecond * 300)
	for _, card := range bank.Card {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	println("[Банк] Остаток положительный")
	return nil
}

// фасад!!! здесь взаимодействие со всеми сервисами, освновной обьект фасада это магазин
func (shop Shop) Sell(user User, product string) error {
	println("[Магизин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Microsecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] Проверка - может ли [%s] купить товар! \n", user.Name)
	time.Sleep(time.Microsecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств на покупку товара!")
		}
		fmt.Printf("[Магазин] Товар- [%s] куплен! \n", prod.Name)
	}
	return nil
}

func main() {
	println("[Банк] Выпуск карт")
	bank.Card = append(bank.Card, card1, card2)

	fmt.Printf("[%s]\n", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]\n", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}
