package main

type Watt struct {
	value int
}

type Currency struct {
	value int
}

type Equipmenter interface {
	power() Watt
	netPrice() Currency
	discountPrice() Currency
}

type CompositeEquipmenter interface {
	add(Equipmenter)
	remove(Equipmenter)
}

type FloppyDisk struct {
	name string
}

func (equipment *FloppyDisk) power() Watt {
	return Watt{value: 1}
}

func (equipment *FloppyDisk) netPrice() Currency {
	return Currency{value: 10}
}

func (equipment *FloppyDisk) discountPrice() Currency {
	return Currency{value: 9}
}
type Chassis struct {
	name string
	equipments []Equipmenter
}

func (equipment *Chassis) netPrice() Currency {
	total := Currency {value: 20}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].netPrice().value;
	}
	return total
}

func (equipment *Chassis) power() Watt {
	total := Watt {value: 2}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].power().value;
	}
	return total
}

func (equipment *Chassis) discountPrice() Currency {
	total := Currency {value: 19}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].discountPrice().value;
	}
	return total
}

func (equipment *Chassis) add(newEquipment Equipmenter) {
	equipment.equipments = append(equipment.equipments, newEquipment)
}

func (equipment *Chassis) remove(index int) {
	remainingEquipments := make([]Equipmenter, 0)
	remainingEquipments = append(remainingEquipments, equipment.equipments[:index]...)
	equipment.equipments = append(remainingEquipments, equipment.equipments[index+1:]...)
}


type Cabinet struct {
	name string
	equipments []Equipmenter
}

func (equipment *Cabinet) netPrice() Currency {
	total := Currency {value: 100}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].netPrice().value;
	}
	return total
}

func (equipment *Cabinet) power() Watt {
	total := Watt {value: 90}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].power().value;
	}
	return total
}

func (equipment *Cabinet) discountPrice() Currency {
	total := Currency {value: 99}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].discountPrice().value;
	}
	return total
}

func (equipment *Cabinet) add(newEquipment Equipmenter) {
	equipment.equipments = append(equipment.equipments, newEquipment)
}

func (equipment *Cabinet) remove(index int) {
	remainingEquipments := make([]Equipmenter, 0)
	remainingEquipments = append(remainingEquipments, equipment.equipments[:index]...)
	equipment.equipments = append(remainingEquipments, equipment.equipments[index+1:]...)
}

func main() {
	cabinet := new(Cabinet)
	chassis := new(Chassis)

	cabinet.add(chassis)
	chassis.add(new(FloppyDisk))
}