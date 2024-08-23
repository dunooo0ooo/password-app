package main

import (
	"fmt"
	"github.com/fatih/color"
	"study-go/account"
	"study-go/files"
	"study-go/output"
)

func main() {
	fmt.Println("_____ PASSWORD APP _____")
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}

}

func getMenu() int {
	var variant int

	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Printf("Выберите вариант: ")
	fmt.Scan(&variant)
	return variant
}

func findAccount(vault *account.VaultWithDb) {
	url := promtData("Введите URL для поиска")
	accounts := vault.FindAccountsByURL(url)
	if len(accounts) == 0 {
		color.Red("Аккаунты не найдены")
	}
	for _, account := range accounts {
		account.Output()
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	url := promtData("Введите URL для поиска")
	isDeleted := vault.DeleteAccountByURL(url)
	if isDeleted {
		color.Green("Удаление прошло успешно")
	} else {
		output.PrintError("Не найдено")
	}

}

func createAccount(vault *account.VaultWithDb) {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат Login или URL")
		return
	}

	vault.AddAccount(*myAccount)
}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var data string
	fmt.Scanln(&data)
	return data
}
