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
func (js JSONStorage) LoadMiners() ([]domain.MinerType, error) {
	var minersJson []MinerTypeJson
	if err := js.loadJSON(&minersJson); err != nil {
		return nil, err
	}
	result := make([]domain.MinerType, 0, len(minersJson))

	for _, miner := range minersJson {
		minerInterval, err := time.ParseDuration(miner.Interval)
		if err != nil {
			return nil, fmt.Errorf("%w: error parse minerInterval %w", ErrStorage, err)
		}

		minerType, err := domain.NewMinerType(
			domain.MinerTypeName(miner.TypeName),
			miner.Name,
			domain.Coal(miner.Cost),
			miner.Energy,
			domain.Coal(miner.Extraction),
			minerInterval,
			domain.Coal(miner.Growth),
		)
		if err != nil {
			return nil, fmt.Errorf("%w: error parse minerType", err)
		}

		result = append(result, minerType)
	}

	return result, nil
}

// Метод для получения Equipments из Json
func (js JSONStorage) LoadEquipments() (map[domain.EquipmentType]domain.Equipment, error) {
	var equipmentsJson []EquipmentJson

	if err := js.loadJSON(&equipmentsJson); err != nil {
		return nil, err
	}

	result := make(map[domain.EquipmentType]domain.Equipment, len(equipmentsJson))

	for _, equipmentJson := range equipmentsJson {
		equipmentType := domain.EquipmentType(equipmentJson.EquipmentType)

		equipment, err := domain.NewEquipment(
			equipmentType,
			equipmentJson.Name,
			domain.Coal(equipmentJson.Cost),
		)
		if err != nil {
			return nil, fmt.Errorf("%w: error parse equipment", err)
		}
		result[equipmentType] = *equipment
	}

	return result, nil
}
