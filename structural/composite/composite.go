package main

type Watt struct {

}

type Currency struct {
	value int
}

type Equipmenter interface {
	power() Watt
	netPrice() Currency
	discountPrice() Currency
	CompositeEquipmenter
}

type CompositeEquipmenter interface {
	Add(Equipmenter)
	Remove(Equipmenter)
}


type FloppyDisk struct {
	name string
}

type CompositeEquipment struct {
	name string
	equipments []Equipmenter
}

func (equipment CompositeEquipment) netPrice() Currency {
	total := Currency {value: 0}
	for i := 0; i < len(equipment.equipments); i++ {
		total.value += equipment.equipments[i].netPrice().value;
	}
	return total
}


type Chassis struct {
	name string
	equipments []Equipmenter
}

type Cabinet struct {
	name string
	equipments []Equipmenter
}

func (cabinet *Cabinet) add(equipment Equipmenter) {
	cabinet.equipments = append(cabinet.equipments, equipment)
}

func main() {
	cabinet := new(Cabinet)
	chassis := new(Chassis)

	cabinet.add(chassis)
	chassis.add(new(FloppyDisk))
}