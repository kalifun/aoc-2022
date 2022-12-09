package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

// calorie
type calorie struct {
	buf       *os.File
	answer1   interface{}
	answer2   interface{}
	spareTire []int
}

// NewCalorie
//  @return *calorie
func NewCalorie() *calorie {
	return &calorie{}
}

func (c *calorie) GetStar() error {
	defer c.buf.Close()
	err := c.boot()
	if err != nil {
		return err
	}
	c.readerFile()
	log.Printf(consts.CalorieAnswer+consts.Answer, c.answer1, c.answer2)
	return nil
}

// boot
//  @receiver c
//  @return error
func (c *calorie) boot() error {
	data, err := utils.NewFileReader(consts.CaloriePath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	c.buf = data
	return nil
}

// readerFile
//  @receiver c
func (c *calorie) readerFile() {
	reader := bufio.NewReader(c.buf)
	var max, next int
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			if next > max {
				max = next
			}
			c.spareTire = append(c.spareTire, next)
			break
		}
		if len(line) == 0 {
			if next > max {
				max = next
			}
			c.spareTire = append(c.spareTire, next)
			next = 0
			continue
		}

		lineInt, err := strconv.Atoi(string(line))
		if err != nil {
			log.Printf("Failed to parse %s, not a qualified numeric type\n", string(line))
			break
		}
		next += lineInt
	}
	c.answer1 = max
	c.part2()
}

// part2
//  @receiver c
func (c *calorie) part2() {
	sort.Ints(c.spareTire)
	length := len(c.spareTire)
	if length >= 3 {
		c.answer2 = c.spareTire[length-1] + c.spareTire[length-2] + c.spareTire[length-3]
	}
}
