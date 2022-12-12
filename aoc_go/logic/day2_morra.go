package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type morraGame struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewMorraGame() *morraGame {
	return &morraGame{}
}

func (m *morraGame) GetStar() error {
	err := m.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	m.collect()
	log.Printf(consts.MorraAnswer+consts.Answer, m.answer1, m.answer2)
	return nil
}

func (m *morraGame) boot() error {
	data, err := utils.NewFileReader(consts.MorraPath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	m.buf = data
	return nil
}

func (m *morraGame) collect() {
	reader := bufio.NewReader(m.buf)
	var part1Sum, part2Sum int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		splitLine := strings.Split(string(line), " ")
		if len(splitLine) != 2 {
			log.Printf("Please check the file, there are illegal contents: %s\n", string(line))
			break
		}

		fixed := getMorraTool(splitLine[0])
		if fixed == unknow {
			log.Printf("Please check the file, there are illegal contents: %s\n", string(line))
			break
		}
		other := getMorraTool(splitLine[1])
		if other == unknow {
			log.Printf("Please check the file, there are illegal contents: %s\n", string(line))
			break
		}

		fg := NewFixedGesture(fixed)
		part1Sum += fg.PK(other)
		part2Sum += fg.DiyRules(other)
	}
	m.answer1 = part1Sum
	m.answer2 = part2Sum
}

type morraTool int

const (
	rock morraTool = iota + 1
	paper
	scissors
	unknow = 0
)

// getMorraTool
//  @param str
//  @return morraTool
func getMorraTool(str string) morraTool {
	var res morraTool
	switch str {
	case "A":
		res = rock
	case "B":
		res = paper
	case "C":
		res = scissors
	case "X":
		res = rock
	case "Y":
		res = paper
	case "Z":
		res = scissors
	default:
		res = unknow
	}
	return res
}

// fixedGesture
type fixedGesture struct {
	gesture morraTool
}

// NewFixedGesture
//  @param g
//  @return fixedGesture
func NewFixedGesture(g morraTool) fixedGesture {
	return fixedGesture{gesture: g}
}

// PK
//  @receiver fg
//  @param other
//  @return int
func (fg fixedGesture) PK(other morraTool) int {
	if int(fg.gesture) == int(other) {
		return 3 + int(other)
	}

	// lost?
	lost := fg.lost()
	if lost == other {
		return 6 + int(other)
	}

	// win
	return 0 + int(other)

}

// lost
//  @receiver fg
//  @return morraTool
func (fg fixedGesture) lost() morraTool {
	switch fg.gesture {
	case rock:
		return paper
	case paper:
		return scissors
	case scissors:
		return rock
	default:
		return unknow
	}
}

func (fg fixedGesture) win() morraTool {
	switch fg.gesture {
	case rock:
		return scissors
	case paper:
		return rock
	case scissors:
		return paper
	default:
		return unknow
	}
}

func (fg fixedGesture) getNewTool(other morraTool) morraTool {
	switch other {
	case rock:
		// X need to lost
		return fg.win()
	case paper:
		// Y need to end the round in a draw
		return fg.gesture
	case scissors:
		// Z you win
		return fg.lost()
	default:
		return unknow
	}
}

// DiyRules
//  @receiver fg
//  @param other
//  @return int
func (fg fixedGesture) DiyRules(other morraTool) int {
	newTool := fg.getNewTool(other)
	return fg.PK(newTool)
}
