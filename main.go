package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	simulator_mainGame_goroutine(10000000)
	// 0.61188126
}

type Reel struct {
	reel1  []int
	reel2  []int
	reel3  []int
	reel4  []int
	reel5  []int
	reelLength []int

}

// constructor of cReel
func (reel *Reel) Init(reel1  []int, reel2  []int ,reel3  []int, reel4  []int,  reel5  []int) {
	reel.reel1 = reel1
	reel.reel2 = reel2
	reel.reel3 = reel3
	reel.reel4 = reel4
	reel.reel5 = reel5
	reel.reelLength = []int{len(reel.reel1), len(reel.reel2), len(reel.reel3), len(reel.reel4), len(reel.reel5)}

}

// generating random number based on len(reel)
//生成5個[0,reelLength)結束的不重複的隨機數
func generateRandomNumber(reelLength []int) []int {

		//存放結果的slice
		nums := make([]int, 0)
		//隨機數生成器,加入時間戳保證每次生成的隨機數不一樣
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i:=0; i < 5; i++ {
			end := reelLength[i]
			//生成隨機數
			//rand.Seed(time.Now().UnixNano())
			num := rand.Intn(end)
			nums = append(nums, num)
		}
	return nums
}

func index(reel *Reel) [][]int{
	index := generateRandomNumber(reel.reelLength)
	// indexList := make([]int, 0)

	mm := make([][]int, 0)
	iterator := 0
	var length, x1, x2, x3 int
	for _, i := range index{
		if iterator == 0 {
			length = len(reel.reel1)
		} else if iterator == 1 {
			length = len(reel.reel2)
		}	else if iterator == 2 {
			length = len(reel.reel3)
		}	else if iterator == 3 {
			length = len(reel.reel4)
		}	else {
			length = len(reel.reel5)
		}

		if i + 1 > (length - 1) {
			x3 = 0
		} else {
			x3 = i + 1
		}
		if i - 1 < 0 {
			x1 = length - 1
		} else {
			x1 = i - 1
		}
		x2 = i

		m := make([]int, 0)
		m = append(m, x1, x2, x3)
		mm= append(mm, m)

		iterator++
	}

	return mm
}

func spin(reel *Reel) [][]int{
	reelList := make([][]int, 0)
	indexArray := index(reel)
	var n []int
	//reelList[0] = []int{reel.reel1[indexArray[0][0]], reel.reel1[indexArray[0][1]], reel.reel1[indexArray[0][2]]}
	//reelList[1] = []int{reel.reel2[indexArray[1][0]], reel.reel2[indexArray[1][1]], reel.reel2[indexArray[1][2]]}
	//reelList[2] = []int{reel.reel3[indexArray[2][0]], reel.reel3[indexArray[2][1]], reel.reel3[indexArray[2][2]]}
	//reelList[3] = []int{reel.reel4[indexArray[3][0]], reel.reel4[indexArray[3][1]], reel.reel4[indexArray[3][2]]}
	//reelList[4] = []int{reel.reel5[indexArray[4][0]], reel.reel5[indexArray[4][1]], reel.reel5[indexArray[4][2]]}
	for i := 0; i < 5; i++{
		switch {
		case i == 0:
			n = reel.reel1
		case i == 1:
			n = reel.reel2
		case i == 2:
			n = reel.reel3
		case i == 3:
			n = reel.reel4
		case i == 4:
			n = reel.reel5
		default:
			n = []int{0}
		}
		reelListi := make([]int, 0)
		reelListi = append(reelListi, n[indexArray[i][0]], n[indexArray[i][1]], n[indexArray[i][2]])
		reelList = append(reelList, reelListi)
	}

	return reelList

}

func getSymbol(reelList [][]int, payLine [][]int, line int) []int{

	symbolInPayline := make([]int, 5)
	for i := 0; i < 5; i++{
		symbolInPayline[i] = reelList[i][payLine[line][i]]
	}
	return symbolInPayline

}

func findKeyConnection(reelList [][]int, payLine [][]int, line int, wild int) []int{
	wild = 3
	reelSymbol := getSymbol(reelList, payLine, line)
	connectSymbol := -1
	connectLine := 0
	for k := 0; k < 5;k++{
		if reelSymbol[k] != wild{
			connectSymbol = reelSymbol[k]
			break
		}
	}
	for j := 0; j < 5; j++{
		if reelSymbol[j] == connectSymbol || reelSymbol[j] == wild{
			connectLine ++
		}else{
			break
		}

	}
	result := []int{connectSymbol, connectLine}

	return result
}

func prizeMapping_single(connectSymbol int, connectLine int, paytable  map[int][]int) int{

	var prize int

	if connectLine < 2{
		prize = 0
	}else{

		prize = paytable[connectSymbol][connectLine]
		// fmt.Println(prize)
	}

	return prize

}

func prizeMapping(reelList [][]int, payline int, payLine [][]int, paytable  map[int][]int) int{
		totalPrize := 0
		for n := 0; n < payline; n++{
			connection := findKeyConnection(reelList, payLine, n, 3)

			prize := prizeMapping_single(connection[0], connection[1], paytable)
			totalPrize += prize
		}
	return totalPrize
}

func simulator_mainGame(reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) float32{
	sum := 0
	for i := 0; i < t; i++{
		reelSpin := spin(reel)
		totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
		sum += totalOdds
	}

	return float32(sum)  / float32(payline * t)
}

func simulator_mainGame_goroutine(times int){
	rand.Seed(time.Now().UnixNano())
	//fmt.Println("Hello World", "OPOPO")
	//fmt.Println(rand.Intn(100))

	// Reel Strips
	reel1 :=[] int {7, 2, 12, 10, 13, 9, 7, 12, 13, 10, 2, 6, 12, 10, 13, 12, 11, 12, 7, 12, 13, 5, 5, 5, 9, 10, 7, 13, 11, 10, 12, 7, 12, 5, 8, 10, 12, 9, 10, 7, 9, 12, 10, 13, 12, 6, 9, 10, 7, 7, 8, 9, 10, 6, 11, 12, 10, 7, 8, 12, 11, 13, 12, 8, 13, 7}
	reel2 :=[] int {7, 2, 12, 10, 13, 9, 12, 12, 13, 10, 2, 6, 8, 13, 12, 3, 11, 12, 6, 12, 13, 12, 5, 12, 9, 3, 6, 11, 12, 10, 5, 12, 5, 6, 3, 10, 9, 13, 6, 10, 11, 3, 8, 13, 12, 7, 12, 10, 12, 10, 8, 12, 9, 12, 3, 13, 8, 13, 10, 13, 12, 9, 13}
	reel3 :=[] int {7, 2, 12, 10, 8, 9, 10, 12, 13, 10, 2, 6, 8, 12, 12, 3, 11, 12, 6, 12, 13, 5, 5, 5, 9, 3, 6, 11, 12, 10, 5, 12, 13, 13, 3, 12, 9, 9, 6, 10, 10, 3, 8, 13, 12, 7, 12, 10, 12, 10, 8, 9, 12, 7, 3, 12, 10, 13, 8, 12, 10, 13, 8}
	reel4 :=[] int {7, 2, 12, 10, 8, 9, 10, 12, 13, 10, 2, 6, 8, 12, 12, 3, 11, 12, 6, 8, 13, 12, 5, 5, 9, 3, 6, 11, 12, 10, 5, 12, 11, 13, 3, 12, 9, 9, 6, 10, 10, 3, 12, 13, 8, 7, 12, 10, 12, 10, 8, 9, 12, 7, 3, 12, 10, 13, 8, 12, 10, 13, 8}
	reel5 :=[] int {7, 2, 12, 10, 8, 9, 10, 12, 13, 10, 2, 6, 8, 12, 12, 3, 11, 6, 6, 8, 13, 2, 5, 5, 9, 11, 6, 11, 12, 10, 5, 13, 11, 7, 11, 10, 9, 9, 6, 11, 10, 11, 8, 13, 8, 7, 9, 10, 12, 10, 8, 9, 12, 7, 11, 8, 10, 8, 9, 12, 8, 6, 10, 9, 7}

	instanceReel := new(Reel)
	instanceReel.Init(reel1, reel2, reel3, reel4, reel5)
	// fmt.Println(instanceReel.reelLength)

	// fmt.Println("Here", index(instanceReel))

	// result := spin(instanceReel)
	// fmt.Println(result, result[0][1])

	// Count
	// payTable, payLine
	payTable := map[int][]int{
		0: {0, 0, 0, 0, 0, 0},
		1: {0, 0, 0, 0, 0, 0},
		2: {0, 0, 0, 0, 0, 0},
		3: {0, 0, 0, 0, 0, 0},
		4: {0, 0, 0, 0, 0, 0},
		5: {0, 0, 0, 35, 70, 350},
		6: {0, 0, 0, 30, 60, 300},
		7: {0, 0, 0, 25, 50, 250},
		8: {0, 0, 0, 20, 40, 200},
		9: {0, 0, 0, 15, 30, 150},
		10: {0, 0, 0, 12, 24, 120},
		11: {0, 0, 0, 10, 20, 100},
		12: {0, 0, 0, 5, 10, 50},
		13: {0, 0, 0, 3, 6, 30}}

	payLine := [][]int{
		{1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{2, 2, 2, 2, 2},
		{0, 1, 2, 1, 0},
		{2, 1, 0, 1, 2},
		{0, 1, 1, 1, 0},
		{2, 1, 1, 1, 2},
		{1, 0, 0, 0, 1},
		{1, 2, 2, 2, 1},
		{0, 0, 1, 2, 2},
		{2, 2, 1, 0, 0},
		{1, 0, 1, 2, 1},
		{1, 2, 1, 0, 1},
		{0, 1, 0, 1, 0},
		{2, 1, 2, 1, 2},
		{1, 2, 0, 2, 1},
		{1, 0, 2, 0, 1},
		{1, 1, 0, 1, 1},
		{1, 1, 2, 1, 1},
		{0, 0, 2, 0, 0},
		{2, 2, 0, 2, 2},
		{0, 1, 0, 1, 2},
		{2, 1, 2, 1, 0},
		{1, 0, 0, 1, 2},
		{1, 2, 2, 1, 0}}

	c1 := make(chan float32, 1)
	c2 := make(chan float32, 1)
	c3 := make(chan float32, 1)
	c4 := make(chan float32, 1)

	go func (reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) {
		sum := 0
		for i := 0; i < t; i++{
			reelSpin := spin(reel)
			totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
			sum += totalOdds
		}

		c1 <- float32(sum)
	}(instanceReel, 25, payLine, payTable, times/4)
	go func (reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) {
		sum := 0
		for i := 0; i < t; i++{
			reelSpin := spin(reel)
			totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
			sum += totalOdds
		}

		c2 <- float32(sum)
	}(instanceReel, 25, payLine, payTable, times/4)
	go func (reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) {
		sum := 0
		for i := 0; i < t; i++{
			reelSpin := spin(reel)
			totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
			sum += totalOdds
		}

		c3 <- float32(sum)
	}(instanceReel, 25, payLine, payTable, times/4)
	go func (reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) {
		sum := 0
		for i := 0; i < t; i++{
			reelSpin := spin(reel)
			totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
			sum += totalOdds
		}

		c4 <- float32(sum)
	}(instanceReel, 25, payLine, payTable, times/4)

	x1 := <-  c1
	x2 := <-  c2
	x3 := <-  c3
	x4 := <-  c4

	rtp := (x1+x2+x3+x4)/(float32(times)*25)

	fmt.Println(rtp)

}
