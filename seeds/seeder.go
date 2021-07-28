package main

import (
	"github.com/ad3n/resto/configs"
	"github.com/ad3n/resto/models"
	"github.com/ad3n/resto/repositories"
	"github.com/ad3n/resto/types"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	configs.Load()
}

func main() {
	menuRepository := repositories.Menu{
		Storage: configs.Db,
	}

	menu1 := models.Menu{
		Type:  types.MENU_DRINK,
		Name:  "Susu Jahe",
		Price: 5000,
	}

	menu2 := models.Menu{
		Type:  types.MENU_MAIN_COURSE,
		Name:  "Ayam Geprek",
		Price: 15000,
	}

	menu3 := models.Menu{
		Type:  types.MENU_SNACK,
		Name:  "Mendoan",
		Price: 5000,
	}

	menuRepository.Saves(&menu1, &menu2, &menu3)
}
