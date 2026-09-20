package domain

type Equipment struct {
	Name      string
	Cost      Coal
	Purchased bool
}

func NewEquipment(equipmentName string, costEquipment Coal) *Equipment {
	return &Equipment{
		Name: equipmentName,
		Cost: costEquipment,
	}
}

type EquipmentType string

const (
	Pickaxe     EquipmentType = "pickaxe"
	Ventilation EquipmentType = "ventilation"
	Trolleys    EquipmentType = "trolleys"
)
