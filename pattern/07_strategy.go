package main

import "fmt"

/*
Краткое описание:
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
Плюсы:
	1. Горячая замена алгоритмов на лету.
	2. Изолирует код и данные алгоритмов от остальных классов.
	3. Уход от наследования к делегированию.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Усложняет программу за счёт дополнительных классов.
	2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

type Strategy interface {
	Route(startPoint int, endPoint int)
}

type Navigator struct {
	Strategy
}

type RoadStrategy struct {
}

type PublicTransportStrategy struct {
}

type WalkStrategy struct {
}

func (r *RoadStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * 40 * trafficJam
	fmt.Printf("RoadStrategy Road A:[%d] to B:[%d] Avg speed:[%d], TraficJam:[%d] Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, trafficJam, total, totalTime)
}

func (r *PublicTransportStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * 40
	fmt.Printf("PublicTransportStrategy Road A:[%d] to B:[%d] Avg speed:[%d], Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

func (r *WalkStrategy) Route(startPoint int, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60
	fmt.Printf("WalkStrategy Road A:[%d] to B:[%d] Avg speed:[%d], Total:[%d] Total Time:[%d] min\n",
		startPoint, endPoint, avgSpeed, total, totalTime)
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}

var (
	start      = 10
	end        = 100
	strategies = []Strategy{
		&PublicTransportStrategy{},
		&RoadStrategy{},
		&WalkStrategy{},
	}
)

func main() {
	nav := Navigator{}
	for _, strategy := range strategies {
		nav.SetStrategy(strategy)
		nav.Route(start, end)
	}

}
