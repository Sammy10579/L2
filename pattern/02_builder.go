package main

import "fmt"

/*
Описание:
	Паттерн 'строитель' это порождающий паттерн, необходим при конструировании сложного объекта,
	для которого может потребоваться много действий.
	Вместо того что бы создавать конструктор с огромным количеством параметров,
	которые могут быть еще и опциональными, в подобных случаях
	следует использовать шаблон проектирования 'строитель'
Преимущества:
	1. Позволяет создавать продукты пошагово.
	2. Позволяет использовать один и тот же код для создания различных продуктов.
	3. Изолирует сложный код сборки продукта от его основной бизнес-логики.
Недостатки:
	1. Усложняет код программы из-за введения дополнительных классов.
	2. Клиент будет привязан к конкретным классам строителей,
		так как в интерфейсе директора может не быть метода получения результата.
*/

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

type Factory struct {
	Collector Collector
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

func (collector *HpCollector) SetCore() {
	collector.Core = 6
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 32
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

func (pc *Computer) Print() {
	fmt.Printf("%s Core:[%d] Mem:[%d] GraphicCard:[%d] Monitor:[%d]", pc.Brand, pc.Core, pc.Memory, pc.GraphicCard, pc.Monitor)
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

//Основная функция для манипулирования шагами по производству обьектов
func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()
}

func main() {
	asusCollector := GetCollector("asus")
	hpCollector := GetCollector("hp")

	factory := NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)
	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(asusCollector)
	pc := factory.CreateComputer()
	pc.Print()
}
