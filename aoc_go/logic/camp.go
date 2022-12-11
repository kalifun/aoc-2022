package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type campCleanup struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewCampCleanup() *campCleanup {
	return &campCleanup{}
}

func (c *campCleanup) GetStar() error {
	err := c.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	c.collect()
	log.Printf(consts.CampAnswer+consts.Answer, c.answer1, c.answer2)
	return nil
}

func (c *campCleanup) boot() error {
	data, err := utils.NewFileReader(consts.CampCleanupPath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	c.buf = data
	return nil
}

func (c *campCleanup) collect() {
	reader := bufio.NewReader(c.buf)
	var part1Sum, part2Sum int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		c, err := NewCleanup(string(line))
		if err != nil {
			break
		}
		if c.FullyContains() {
			part1Sum += 1
		}
		if c.Overlapping() {
			part2Sum += 1
		}
	}
	c.answer1 = part1Sum
	c.answer2 = part2Sum
}

type cleanup struct {
	line  string
	part1 []int
	part2 []int
}

func NewCleanup(str string) (*cleanup, error) {
	c := &cleanup{
		line: str,
	}
	ok := c.checkLine()
	if !ok {
		return c, fmt.Errorf(fmt.Sprintf(consts.InvalidLine, str))
	}
	return c, nil
}

func (c *cleanup) FullyContains() bool {
	if c.part1[0] <= c.part2[0] && c.part1[1] >= c.part2[1] {
		return true
	}

	if c.part1[0] >= c.part2[0] && c.part1[1] <= c.part2[1] {
		return true
	}
	return false
}

func (c *cleanup) Overlapping() bool {
	for i := c.part2[0]; i <= c.part2[1]; i++ {
		if c.part1[0] <= i && i <= c.part1[1] {
			return true
		}
	}
	return false
}

func (c *cleanup) checkLine() bool {
	lineSplit := strings.Split(string(c.line), ",")
	if len(lineSplit) != 2 {
		log.Printf(consts.InvalidLine, c.line)
		return false
	}

	part1Ctx := strings.Split(lineSplit[0], "-")
	if len(part1Ctx) != 2 {
		log.Printf(consts.InvalidLine, c.line)
	}

	part1_1, err := strconv.Atoi(part1Ctx[0])
	if err != nil {
		log.Println(err)
		return false
	}

	part1_2, err := strconv.Atoi(part1Ctx[1])
	if err != nil {
		log.Println(err)
		return false
	}

	c.part1 = append(c.part1, part1_1)
	c.part1 = append(c.part1, part1_2)

	part2Ctx := strings.Split(lineSplit[1], "-")
	if len(part2Ctx) != 2 {
		log.Printf(consts.InvalidLine, c.line)
		return false
	}

	part2_1, err := strconv.Atoi(part2Ctx[0])
	if err != nil {
		log.Println(err)
		return false
	}

	part2_2, err := strconv.Atoi(part2Ctx[1])
	if err != nil {
		log.Println(err)
		return false
	}

	c.part2 = append(c.part2, part2_1)
	c.part2 = append(c.part2, part2_2)

	return true
}
