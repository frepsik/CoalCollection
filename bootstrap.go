package main

import (
	storage "CoalCollection/Storage"
	"CoalCollection/domain"
	"CoalCollection/service"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type gameConfig struct {
	minerTypes map[domain.MinerTypeName]domain.MinerType
	equipments map[domain.EquipmentTypeName]domain.EquipmentType
}

// Метод подгружающий игровой конфиг
func loadGameConfig() (gameConfig, error) {
	// Получаем переменную окружения, чтобы её задать используется следующая команда: $env:COAL_CONFIG_DIR = "Путь"
	configDir, exists := os.LookupEnv("COAL_CONFIG_DIR")
	if !exists || strings.TrimSpace(configDir) == "" {
		return gameConfig{}, errors.New("COAL_CONFIG_DIR is not set")
	}

	// Собираем полный путь
	minersPath := filepath.Join(configDir, "miners.json")
	equipmentsPath := filepath.Join(configDir, "equipments.json")

	minerStorage := storage.NewJSONStorage(minersPath)
	equipmentsStorage := storage.NewJSONStorage(equipmentsPath)

	minerTypes, err := minerStorage.LoadMinerTypes()
	if err != nil {
		return gameConfig{}, err
	}

	equipments, err := equipmentsStorage.LoadEquipments()
	if err != nil {
		return gameConfig{}, err
	}

	return gameConfig{minerTypes: minerTypes, equipments: equipments}, nil
}

// Метод, который собирает основные данные и проверяет имеет ли смысл вообще дальше давать доступ к запуску игры
func buildGame(ctx context.Context) (*service.Game, error) {
	gameConfig, err := loadGameConfig()
	if err != nil {
		return nil, err
	}

	gameService := service.NewGame(ctx, gameConfig.equipments, gameConfig.minerTypes)

	return gameService, nil
}
