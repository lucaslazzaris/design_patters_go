package main

type Equipment interface {
	netPrice() float32
	accept(EquipmentVisitor)
}

type FloppyDisk struct {
}

func (equipment *FloppyDisk) netPrice() float32 {
	return 2.0
}

func (equipment *FloppyDisk) accept(visitor EquipmentVisitor) {
	visitor.visitFloppyDisk(equipment)
}

type Card struct {
}

func (equipment *Card) netPrice() float32 {
	return 3.0
}

func (equipment *Card) accept(visitor EquipmentVisitor) {
	visitor.visitCard(equipment)
}

type Chassis struct {
}

func (equipment *Chassis) netPrice() float32 {
	return 4.0
}

func (equipment *Chassis) accept(visitor EquipmentVisitor) {
	visitor.visitChassis(equipment)
}

type Bus struct {
}

func (equipment *Bus) netPrice() float32 {
	return 10.0
}

func (equipment *Bus) accept(visitor EquipmentVisitor) {
	visitor.visitBus(equipment)
}

type EquipmentVisitor interface {
	visitFloppyDisk(*FloppyDisk)
	visitCard(*Card)
	visitChassis(*Chassis)
	visitBus(*Bus)
}

type PricingVisitor struct {
	total float32
}

func (visitor *PricingVisitor) visitFloppyDisk(floppyDisk *FloppyDisk) {
	visitor.total += floppyDisk.netPrice()
}

func (visitor *PricingVisitor) visitCard(card *Card) {
	visitor.total += card.netPrice()
}

func (visitor *PricingVisitor) visitChassis(chassis *Chassis) {
	visitor.total += chassis.netPrice()
}

func (visitor *PricingVisitor) visitBus(bus *Bus) {
	visitor.total += bus.netPrice()
}
