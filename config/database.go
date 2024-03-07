package config

import (
	"fmt"
	"os"
)

var DB *gorm.DB

func DBconnect() (*gorm.DB, error) {
	dns := os.Getenv("DB_URL")
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Admin{})
	DB.AutoMigrate(&models.Address{})
	DB.AutoMigrate(&models.Product{})
	DB.AutoMigrate(&models.Brand{})
	DB.AutoMigrate(&models.Cart{})
	DB.AutoMigrate(&models.Image{})
	DB.AutoMigrate(&models.Payment{})
	DB.AutoMigrate(&models.OderDetails{})
	DB.AutoMigrate(&models.Coupon{})
	DB.AutoMigrate(&models.Wishlist{})
	DB.AutoMigrate(&models.Catogery{})
	DB.AutoMigrate(&models.RazorPay{})
	DB.AutoMigrate(&models.Oder_item{})
	DB.AutoMigrate(&models.Wallet{})
	DB.AutoMigrate(&models.WalletHistory{})

	return DB, nil

}
