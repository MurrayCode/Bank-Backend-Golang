package migrations

import (
	"home/murray/GoBank/helpers"
	"home/murray/GoBank/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func createAccount() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username:"Murray", Email: "Murray@bank.com"},
		{Username: "Liam", Email: "Liam@bank.com"},
	}

	for i := 0; i < len(users); i++{
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(100), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccount()
}