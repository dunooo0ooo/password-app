package main

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"study-go/account"
	"study-go/files"
	"study-go/output"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину",
	"4. Удалить аккаунт",
	"5. Выход",
	"Выберите вариант",
}

func main() {
	fmt.Println("_____ PASSWORD APP _____")
	vault := account.NewVault(files.NewJsonDB("data.json"))
Menu:
	for {
		variant := promtData(menuVariants...)

		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)
	}

}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promtData("Введите URL для поиска")
	accounts := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&accounts)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	name := promtData("Введите Login для поиска")
	accounts := vault.FindAccounts(name, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&accounts)

}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		color.Red("Аккаунты не найдены")
	}
	for _, account := range *accounts {
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

func promtData(data ...string) string {
	for i, line := range data {
		if i == len(data)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}
	var res string
	fmt.Scanln(&res)
	return res
}
