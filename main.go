package main

import (
	storage "CoalCollection/Storage"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Получаем переменную окружения, чтобы её задать используется следующая команда: $env:COAL_CONFIG_DIR = "Путь"
	configDir, exists := os.LookupEnv("COAL_CONFIG_DIR")
	if !exists {
		fmt.Println("Coal mine fail start")
		return
	}

	// Собираем полный путь
	minersPath := filepath.Join(configDir, "miners.json")
	equipmentsPath := filepath.Join(configDir, "equipments.json")

	minerStorage := storage.NewJSONStorage(minersPath)
	equipmentsStorage := storage.NewJSONStorage(equipmentsPath)

	minerTypes, err := minerStorage.LoadMinerTypes()
	if err != nil {
		fmt.Println(err)
	}

	equipments, err := equipmentsStorage.LoadEquipments()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(minerTypes)
	fmt.Println(equipments)

	// fmt.Println("Coal mine started")

	// rootCtx := context.Background()

	// gameService := service.NewGame(rootCtx)

	// gameService.Start()

	// time.Sleep(10 * time.Second)

	// gameService.Shutdown()

	// fmt.Println("Статус игры", gameService.StatusEnterprise())
	// fmt.Println("Конец")
}
