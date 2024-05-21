package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Name  string
	Price float64
}

type Sales struct {
	Product *Product
	Volume  float64
}

func main() {
	var (
		totalProfit  float64
		userProducts []*Product
		userSales    []*Sales
	)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите название продукта и его цену через пробел (или введите '-' чтобы завершить): ")
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "-" {
			break
		}

		inputParts := strings.Split(userInput, " ")
		if len(inputParts) != 2 {
			fmt.Println("Пожалуйста, введите название и цену товара через пробел.")
			continue
		}

		price, err := strconv.ParseFloat(inputParts[1], 64)
		if err != nil {
			fmt.Println("Ошибка при чтении цены товара:", err)
			continue
		}

		product := &Product{Name: inputParts[0], Price: price}
		userProducts = append(userProducts, product)
	}

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Просмотр списка блюд")
		fmt.Println("2. Новая продажа")
		fmt.Println("3. Посмотреть итоговую прибыль")
		fmt.Println("4. Просмотр списка продаж")
		fmt.Println("5. Добавить блюдо")
		fmt.Println("6. Выход")

		var menuChoice int
		fmt.Scanln(&menuChoice)

		switch menuChoice {
		case 1:
			SeeUserProducts(userProducts)
		case 2:
			NewSell(userProducts, &userSales, &totalProfit)
		case 3:
			SayTotalProfit(totalProfit)
		case 4:
			ListSales(userSales)
		case 5:
			userProducts = AddDish(userProducts)
		case 6:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Некорректный выбор. Пожалуйста, выберите от 1 до 5")
		}
	}
}

func SeeUserProducts(userProducts []*Product) {
	fmt.Println("Ваше меню:")
	for _, p := range userProducts {
		fmt.Printf("%s - Цена: %.2f\n", p.Name, p.Price)
	}
}

func NewSell(userProducts []*Product, userSales *[]*Sales, totalProfit *float64) {
	fmt.Println("Выберите товар для продажи:")
	for i, p := range userProducts {
		fmt.Printf("%d: %s - Цена: %.2f\n", i+1, p.Name, p.Price)
	}

	var choice int
	fmt.Scanln(&choice)
	if choice < 1 || choice > len(userProducts) {
		fmt.Println("Товар введён неверно")
		return
	}

	selectedProduct := userProducts[choice-1]
	fmt.Printf("Выбранный товар: %s по цене %.2f\n", selectedProduct.Name, selectedProduct.Price)

	fmt.Println("Введите количество товара:")
	var volume float64
	fmt.Scanln(&volume)

	if volume <= 0 {
		fmt.Println("Некорректное количество товара")
		return
	}

	sale := &Sales{Product: selectedProduct, Volume: volume}
	*userSales = append(*userSales, sale)

	profit := selectedProduct.Price * volume
	*totalProfit += profit

	fmt.Printf("Продано %.2f единиц товара %s. Прибыль: %.2f\n", volume, selectedProduct.Name, profit)
}

func SayTotalProfit(totalProfit float64) {
	fmt.Printf("Итоговая прибыль: %.2f\n", totalProfit)
}

func ListSales(userSales []*Sales) {
	if len(userSales) == 0 {
		fmt.Println("Список продаж пуст")
		return
	}

	fmt.Println("Список продаж:")
	for _, sale := range userSales {
		fmt.Printf("Товар: %s, Количество: %.2f, Сумма: %.2f\n", sale.Product.Name, sale.Volume, sale.Product.Price*sale.Volume)
	}
}
func AddDish(userProducts []*Product) []*Product {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите название продукта и его цену через пробел (или введите '-' чтобы завершить): ")
		scanner.Scan()
		userInput := scanner.Text()

		if userInput == "-" {
			break
		}

		inputParts := strings.Split(userInput, " ")
		if len(inputParts) != 2 {
			fmt.Println("Пожалуйста, введите название и цену товара через пробел.")
			continue
		}

		price, err := strconv.ParseFloat(inputParts[1], 64)
		if err != nil {
			fmt.Println("Ошибка при чтении цены товара:", err)
			continue
		}

		product := &Product{Name: inputParts[0], Price: price}
		userProducts = append(userProducts, product)
	}
	return userProducts
}
