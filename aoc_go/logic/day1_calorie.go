package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/kalifun/aco-2022/entity/consts"
	"github.com/kalifun/aco-2022/repo/utils"
)

// calorie
type calorie struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// calorieCount
type calorieCount struct {
	limit  int
	counts []int
}

// SortTopN
//
//	@receiver c
//	@param count
func (c *calorieCount) SortTopN(count int) {
	if len(c.counts) < c.limit {
		c.counts = append(c.counts, count)
		return
	}

	c.counts = append(c.counts, count)
	sort.Ints(c.counts)
	c.counts = c.counts[1:]
}

// Sum
//
//	@receiver c
//	@return int
func (c *calorieCount) Sum() int {
	var sum int
	for _, count := range c.counts {
		sum += count
	}
	return sum
}

// NewCalorie
//
//	@return *calorie
func NewCalorie() *calorie {
	return &calorie{}
}

func (c *calorie) GetStar() error {
	defer c.buf.Close()
	err := c.boot()
	if err != nil {
		return err
	}
	c.collect()
	log.Printf(consts.CalorieAnswer+consts.Answer, c.answer1, c.answer2)
	return nil
}

// boot
//
//	@receiver c
//	@return error
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
//
//	@receiver c
func (c *calorie) collect() {
	reader := bufio.NewReader(c.buf)
	var max, next int
	counter := calorieCount{
		limit:  3,
		counts: []int{},
	}

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			if next > max {
				max = next
			}
			counter.SortTopN(next)
			break
		}
		if len(line) == 0 {
			if next > max {
				max = next
			}
			counter.SortTopN(next)
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
	c.answer2 = counter.Sum()
}
