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

func NewJsonStorage(path string) JSONStorage {
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
func (js JSONStorage) LoadEquipments() (map[domain.EquipmentType]*domain.Equipment, error) {
	var equipmentsJson []EquipmentJSON

	if err := js.loadJSON(&equipmentsJson); err != nil {
		return nil, err
	}

	result := make(map[domain.EquipmentType]*domain.Equipment, len(equipmentsJson))

	for _, equipmentJson := range equipmentsJson {
		equipmentType := domain.EquipmentType(equipmentJson.EquipmentType)

		equipment, err := domain.NewEquipment(
			equipmentType,
			equipmentJson.Name,
			domain.Coal(equipmentJson.Cost),
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
