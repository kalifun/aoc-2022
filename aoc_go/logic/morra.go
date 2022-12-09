package logic

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type morraGame struct {
	buf *os.File
}

func NewMorraGame() *morraGame {
	return &morraGame{}
}

func (m *morraGame) GetStar() error {
	panic("not implemented") // TODO: Implement
}

func (m *morraGame) boot() error {
	data, err := utils.NewFileReader(consts.CaloriePath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	m.buf = data
	return nil
}

func (m *morraGame) readerFile() {
	reader := bufio.NewReader(m.buf)
	for {
		_, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
	}
}

type morraTool int

const (
	rock morraTool = iota
	paper
	scissors
	unknow
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
