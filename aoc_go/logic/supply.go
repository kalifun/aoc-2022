package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type supplyStacks struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewSupplyStacks() *supplyStacks {
	return &supplyStacks{}
}

func (s *supplyStacks) GetStar() error {
	err := s.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	s.collect()
	log.Printf(consts.RucksackAnswer+consts.Answer, s.answer1, s.answer2)
	return nil
}

func (s *supplyStacks) boot() error {
	data, err := utils.NewFileReader(consts.SupplyStacksPath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	s.buf = data
	return nil
}

func (s *supplyStacks) collect() {
	reader := bufio.NewReader(s.buf)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(strings.Split(string(line), " "))
	}
}

type crates struct {
	stacks map[uint][]string
}

func (c *crates) Move(from, to, num uint) {

}
