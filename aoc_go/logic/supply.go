package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

// supplyStacks  TODO
type supplyStacks struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// NewSupplyStacks
//  @return *supplyStacks
func NewSupplyStacks() *supplyStacks {
	return &supplyStacks{}
}

// GetStar
//  @receiver s
//  @return error
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
//  @receiver s
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
	s.answer1 = c.GetPart1Answer()
}

// crates  TODO
type crates struct {
	stacks map[uint][]string
}

// NewCrates
//  @return *crates
func NewCrates() *crates {
	return &crates{
		stacks: make(map[uint][]string),
	}
}

// AutoWork
//  @receiver c
//  @param line
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
			if v, ok := c.stacks[uint(site)]; ok {
				c.stacks[uint(site)] = append(v, word)
			} else {
				c.stacks[uint(site)] = []string{word}
			}
		}
	}
}

// move
//  @receiver c
//  @param from
//  @param to
//  @param num
func (c *crates) move(from, to, num uint) {
	// 取出要move的内容
	var waitMoveList []string
	if fromVal, ok := c.stacks[from]; ok {
		waitMoveList = c.stacks[from][:num]
		c.stacks[from] = fromVal[num:]
	}

	utils.Reverse(waitMoveList)
	if _, ok := c.stacks[to]; ok {
		slice := make([]string, len(waitMoveList)+len(c.stacks[to]))
		copy(slice, waitMoveList)
		copy(slice[len(waitMoveList):], c.stacks[to])
		c.stacks[to] = slice
	}
}

// GetPart1Answer
//  @receiver c
//  @return interface{}
func (c *crates) GetPart1Answer() interface{} {
	var res []string
	for i := 0; i < len(c.stacks); i++ {
		if v, ok := c.stacks[uint(i+1)]; ok {
			if len(v) != 0 {
				res = append(res, v[0])
			}
		}
	}
	return strings.Join(res, "")
}
