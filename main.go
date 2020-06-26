package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println("Hello World", "OPOPO")
	//fmt.Println(rand.Intn(100))

	// Reel Strips
	reel1 :=[] int {5, 5, 5, 1, 100, 101}
	reel2 :=[] int {5, 5, 5}
	reel3 :=[] int {5, 3, 5}
	reel4 :=[] int {5, 5, 5, 1, 1, 30, 36, 31}
	reel5 :=[] int {5, 3, 5, 3, 1, 1, 3, 42, 44}

	instanceReel := new(Reel)
	instanceReel.Init(reel1, reel2, reel3, reel4, reel5)
	// fmt.Println(instanceReel.reelLength)

	result := spin(instanceReel)
	fmt.Println(result, result[0][1])

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

	// fmt.Println(payLine[0])
	// fmt.Println(getSymbol(result, payLine, 2))
	// fmt.Println(findKeyConnection(result, payLine, 0, 3))
	// test := findKeyConnection(result, payLine, 0, 3)
	// fmt.Println(test)
	// fmt.Println(prizeMapping_single(test[0], test[1], payTable))
	// fmt.Println(prizeMapping(result, 25, payLine, payTable))
	RTP := simulator_mainGame(instanceReel, 25, payLine, payTable, 1000000)
	fmt.Println(RTP)
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
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i:=0; i < 5; i++ {
			end := reelLength[i]
			//生成隨機數
			num := r.Intn(end)
			nums = append(nums, num)
		}
	return nums
}

func index(reel *Reel) map[int][]int{
	index := generateRandomNumber(reel.reelLength)
	// indexList := make([]int, 0)
	m := map[int][]int{}
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

		m[iterator] = []int{x1, x2, x3}

		iterator++
	}

	return m
}

func spin(reel *Reel) map[int][]int{
	reelList := map[int][]int{}
	indexArray := index(reel)

	reelList[0] = []int{reel.reel1[indexArray[0][0]], reel.reel1[indexArray[0][1]], reel.reel1[indexArray[0][2]]}
	reelList[1] = []int{reel.reel2[indexArray[1][0]], reel.reel2[indexArray[1][1]], reel.reel2[indexArray[1][2]]}
	reelList[2] = []int{reel.reel3[indexArray[2][0]], reel.reel3[indexArray[2][1]], reel.reel3[indexArray[2][2]]}
	reelList[3] = []int{reel.reel4[indexArray[3][0]], reel.reel4[indexArray[3][1]], reel.reel4[indexArray[3][2]]}
	reelList[4] = []int{reel.reel5[indexArray[4][0]], reel.reel5[indexArray[4][1]], reel.reel5[indexArray[4][2]]}

	return reelList

}

func getSymbol(reelList map[int][]int, payLine [][]int, line int) []int{

	symbolInPayline := make([]int, 5)
	for i := 0; i < 5; i++{
		symbolInPayline[i] = reelList[i][payLine[line][i]]
	}
	return symbolInPayline

}

func findKeyConnection(reelList map[int][]int, payLine [][]int, line int, wild int) []int{
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
	}

	return prize

}

func prizeMapping(reelList map[int][]int, payline int, payLine [][]int, paytable  map[int][]int) int{
		totalPrize := 0
		for n := 0; n < payline; n++{
			connection := findKeyConnection(reelList, payLine, n, 3)

			prize := prizeMapping_single(connection[0], connection[1], paytable)
			totalPrize += prize
		}
	return totalPrize
}

func simulator_mainGame(reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) int{
	sum := 0
	for i := 0; i < t; i++{
		reelSpin := spin(reel)
		totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
		sum += totalOdds
	}
	return sum / (payline * t)
}