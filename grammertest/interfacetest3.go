package main

import "fmt"

type Exchanger interface {
	Exchange()
}

type StringPair struct {
	first, second string
}

func (stringPair *StringPair) Exchange() {
	stringPair.first, stringPair.second = stringPair.second, stringPair.first
}

type PointPair [2]int

func (pointPair *PointPair) Exchange() {
	pointPair[0], pointPair[1] = pointPair[1], pointPair[0]
}

func BatchExchange(pairs ...Exchanger) {
	for _, pair := range pairs {
		pair.Exchange()
	}
}

func main() {
	stringpair := StringPair{"1", "2"}
	pointpair := PointPair{1, 2}

	fmt.Println("Before ", stringpair, pointpair)

	BatchExchange(&stringpair, &pointpair)

	fmt.Println("After ", stringpair, pointpair)
}
