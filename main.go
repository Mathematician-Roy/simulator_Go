package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)
var wg sync.WaitGroup

func main() {

	simulator_mainGame_goroutine(1000000, 4)
	// 0.61188126
	//reel5 := []int{5, 3, 5, 3, 1, 1, 3, 42, 44}
	//instanceReel := new(Reel)
	//instanceReel.Init(reel1, reel2, reel3, reel4, reel5)
	//RTP := simulator_mainGame(instanceReel, 25, payLine, payTable, 1000000)
	//fmt.Println(RTP)
}

type Reel struct {
	reel1  []int
	reel2  []int
	reel3  []int
	reel4  []int
	reel5  []int
	reelLength []int

}

// constructor of Reel
func (reel *Reel) Init(reel1  []int, reel2  []int ,reel3  []int, reel4  []int,  reel5  []int) {
	reel.reel1 = reel1
	reel.reel2 = reel2
	reel.reel3 = reel3
	reel.reel4 = reel4
	reel.reel5 = reel5
	reel.reelLength = []int{len(reel.reel1), len(reel.reel2), len(reel.reel3), len(reel.reel4), len(reel.reel5)}

}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
// generating random number based on len(reel)
//生成5個[0,reelLength)結束的不重複的隨機數
func generateRandomNumber(reelLength []int) []int {

		//存放結果的slice
		nums := make([]int, 0)
		//隨機數生成器,加入時間戳保證每次生成的隨機數不一樣
		// r := rand.New(rand.NewSource(time.Now().UnixNano()))
		//for i:=0; i < 5; i++ {
		//	end := reelLength[i]
		//	//生成隨機數
		//	//rand.Seed(time.Now().UnixNano())
		//	num := rand.Intn(end)
		//	nums = append(nums, num)
		var src cryptoSource
		r := rand.New(src)
		nums = []int{r.Intn(reelLength[0]), r.Intn(reelLength[1]), r.Intn(reelLength[2]), r.Intn(reelLength[3]), r.Intn(reelLength[4])}
		// }
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
		reelListi = []int{n[indexArray[i][0]], n[indexArray[i][1]], n[indexArray[i][2]]}
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

func simulator_mainGame(c chan int, reel *Reel, payline int, payLine [][]int, paytable  map[int][]int, t int) {
	defer wg.Done()
	sum := 0
	for i := 0; i < t; i++{
		reelSpin := spin(reel)
		totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
		sum += totalOdds
	}
	c <- sum
	// return  float32(sum)
	// return float32(sum)  / float32(payline * t)
}

func simulator_mainGame_goroutine(times int, max int){

	rand.Seed(time.Now().UnixNano())

	instanceReel := new(Reel)
	instanceReel.Init(reel1, reel2, reel3, reel4, reel5)

	c1 := make(chan int, max)
	//
	//go simulator_mainGame(c1, instanceReel, 25, payLine, payTable, times/4)
	//go simulator_mainGame(c1, instanceReel, 25, payLine, payTable, times/4)
	//go simulator_mainGame(c1, instanceReel, 25, payLine, payTable, times/4)
	//go simulator_mainGame(c1, instanceReel, 25, payLine, payTable, times/4)
	//
	//
	//x1, x2, x3, x4 := <- c1, <- c1, <- c1, <- c1
	for i := 0; i < max; i++ {
		wg.Add(1)
		go simulator_mainGame(c1, instanceReel, 25, payLine, payTable, times/max)
	}
	wg.Wait()
	close(c1)

	sumChannel := 0
	for item := range c1 {
		sumChannel += item
	}

	rtp := float32(sumChannel)/(float32(times)*25)

	fmt.Println(rtp)

}

//func simBody(reel *Reel, payline int, payLine [][]int, paytable map[int][]int, sum *int, group *sync.WaitGroup, mu *sync.Mutex) {
//	reelSpin := spin(reel)
//	totalOdds := prizeMapping(reelSpin, payline, payLine, paytable)
//	mu.Lock()
//	*sum += totalOdds
//	mu.Unlock()
//	group.Done()
//}
//
//func simulator_mainGame(reel *Reel, payline int, payLine [][]int, paytable map[int][]int, t int) float32 {
//	//defer timeTrack(time.Now(), "Simulator Time")
//	sum := 0
//	group := sync.WaitGroup{}
//	mu := sync.Mutex{}
//	for i := 0; i < t; i++ {
//		group.Add(1)
//		go simBody(reel, payline, payLine, paytable, &sum, &group, &mu)
//	}
//	group.Wait()
//	return float32(sum) / float32(payline * t)
//}