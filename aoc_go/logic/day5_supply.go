package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kalifun/aco-2022/entity/consts"
	"github.com/kalifun/aco-2022/repo/utils"
)

// supplyStacks  TODO
type supplyStacks struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// NewSupplyStacks
//
//	@return *supplyStacks
func NewSupplyStacks() *supplyStacks {
	return &supplyStacks{}
}

// GetStar
//
//	@receiver s
//	@return error
func (s *supplyStacks) GetStar() error {
	err := s.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	s.collect()
	log.Printf(consts.SupplyAnswer+consts.Answer, s.answer1, s.answer2)
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

// collect
//
//	@receiver s
func (s *supplyStacks) collect() {
	reader := bufio.NewReader(s.buf)
	c := NewCrates()
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		c.AutoWork(string(line))
	}
	s.answer1 = getAnswer(c.part1Stacks)
	s.answer2 = getAnswer(c.part2Stacks)
}

// crates  TODO
type crates struct {
	part1Stacks map[uint][]string
	part2Stacks map[uint][]string
}

// NewCrates
//
//	@return *crates
func NewCrates() *crates {
	return &crates{
		part1Stacks: make(map[uint][]string),
		part2Stacks: make(map[uint][]string),
	}
}

// AutoWork
//
//	@receiver c
//	@param line
func (c *crates) AutoWork(line string) {
	// Parsing whether the current row is a header or a move operation
	if line == "" {
		return
	}

	if strings.HasPrefix(string(line), "move") {
		// 使用正则查询数字
		re := regexp.MustCompile(`\d[\d,]*[\.]?[\d{3}]*`)
		data := re.FindAllString(line, -1)
		if len(data) != 3 {
			return
		}
		from, err := strconv.Atoi(data[1])
		if err != nil {
			return
		}

		to, err := strconv.Atoi(data[2])
		if err != nil {
			return
		}

		move, err := strconv.Atoi(data[0])
		if err != nil {
			return
		}
		c.move(uint(from), uint(to), uint(move))
		return
	}

	for i := 0; i < len(line); i++ {
		if string(line[i]) == "[" {
			// 说明满足
			word := string(line[i+1])
			site := (i / 4) + 1
			if v, ok := c.part1Stacks[uint(site)]; ok {
				c.part1Stacks[uint(site)] = append(v, word)
			} else {
				c.part1Stacks[uint(site)] = []string{word}
			}

			if v, ok := c.part2Stacks[uint(site)]; ok {
				c.part2Stacks[uint(site)] = append(v, word)
			} else {
				c.part2Stacks[uint(site)] = []string{word}
			}
		}
	}
}

// move
//
//	@receiver c
//	@param from
//	@param to
//	@param num
func (c *crates) move(from, to, num uint) {
	dm := NewDataMigration(from, to, num)
	part1 := dm.Migration(c.part1Stacks, true)
	c.part1Stacks = part1
	part2 := dm.Migration(c.part2Stacks, false)
	c.part2Stacks = part2
}

// dataMigration  TODO
type dataMigration struct {
	from uint
	to   uint
	num  uint
}

// NewDataMigration
//
//	@param from
//	@param to
//	@param num
//	@return dataMigration
func NewDataMigration(from, to, num uint) dataMigration {
	return dataMigration{
		from: from,
		to:   to,
		num:  num,
	}
}

// Migration
//
//	@receiver d
//	@param stack
//	@param reversal
//	@return map
func (d dataMigration) Migration(stack map[uint][]string, reversal bool) map[uint][]string {
	var waitMoveList []string
	if fromVal, ok := stack[d.from]; ok {
		waitMoveList = stack[d.from][:d.num]
		stack[d.from] = fromVal[d.num:]
	}

	if reversal {
		utils.Reverse(waitMoveList)
	}
	if _, ok := stack[d.to]; ok {
		slice := make([]string, len(waitMoveList)+len(stack[d.to]))
		copy(slice, waitMoveList)
		copy(slice[len(waitMoveList):], stack[d.to])
		stack[d.to] = slice
	}

	return stack
}

// getAnswer
//
//	@param stack
//	@return interface{}
func getAnswer(stack map[uint][]string) interface{} {
	var res []string
	for i := 0; i < len(stack); i++ {
		if v, ok := stack[uint(i+1)]; ok {
			if len(v) != 0 {
				res = append(res, v[0])
			}
		}
	}
	return strings.Join(res, "")
}
