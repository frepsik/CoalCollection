package storage

import (
	"CoalCollection/domain"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JSONStorage struct {
	path string
}

func NewJSONStorage(path string) JSONStorage {
	return JSONStorage{
		path: path,
	}
}

// Helper - предназначенный для чтения из файла и декодирования данных
func (js JSONStorage) loadJSON(destination any) error {
	file, err := os.Open(js.path)
	if err != nil {
		return fmt.Errorf("%w: open file %q: %w", ErrStorage, js.path, err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(destination); err != nil {
		return fmt.Errorf("%w: decode json file %q: %w", ErrStorage, js.path, err)
	}

	return nil
}

// Метод для получения списка типов TypeMiner из Json
func (js JSONStorage) LoadMinerTypes() (map[domain.MinerTypeName]domain.MinerType, error) {
	var minersJSON []MinerTypeJSON
	if err := js.loadJSON(&minersJSON); err != nil {
		return nil, err
	}
	result := make(map[domain.MinerTypeName]domain.MinerType, len(minersJSON))

	for _, minerJSON := range minersJSON {
		minerInterval, err := time.ParseDuration(minerJSON.Interval)
		minerTypeName := domain.MinerTypeName(minerJSON.TypeName)
		if err != nil {
			return nil, fmt.Errorf("%w: error parse minerInterval %w", ErrStorage, err)
		}

		minerType, err := domain.NewMinerType(
			minerTypeName,
			minerJSON.Name,
			domain.Coal(minerJSON.Cost),
			minerJSON.Energy,
			domain.Coal(minerJSON.Extraction),
			minerInterval,
			domain.Coal(minerJSON.Growth),
		)
		if err != nil {
			return nil, fmt.Errorf("%w: error create minerType", err)
		}

		if _, exists := result[minerTypeName]; exists {
			return nil, ErrDuplicateMinerType
		}

		result[minerTypeName] = minerType
	}

	return result, nil
}

// Метод для получения Equipments из Json
func (js JSONStorage) LoadEquipments() (map[domain.EquipmentTypeName]domain.EquipmentType, error) {
	var equipmentsJSON []EquipmentJSON

	if err := js.loadJSON(&equipmentsJSON); err != nil {
		return nil, err
	}

	result := make(map[domain.EquipmentTypeName]domain.EquipmentType, len(equipmentsJSON))

	for _, equipmentJSON := range equipmentsJSON {

		equipmentType := domain.EquipmentTypeName(equipmentJSON.EquipmentType)
		equipment, err := domain.NewEquipmentType(
			equipmentType,
			equipmentJSON.Name,
			domain.Coal(equipmentJSON.Cost),
		)
		if err != nil {
			return nil, fmt.Errorf("%w: error create equipment", err)
		}
		if _, existsts := result[equipmentType]; existsts {
			return nil, ErrDuplicateEquipmentType
		}
		result[equipmentType] = equipment
	}

	return result, nil
}
