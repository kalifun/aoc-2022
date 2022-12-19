package logic

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/kalifun/aco-2022/entity/consts"
	"github.com/kalifun/aco-2022/repo/utils"
)

type tuningTrouble struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// NewTuningTrouble
//
//	@return *tuningTrouble
func NewTuningTrouble() *tuningTrouble {
	return &tuningTrouble{}
}

// GetStar
//
//	@receiver t
//	@return error
func (t *tuningTrouble) GetStar() error {
	err := t.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	t.collect()
	log.Printf(consts.TuningTroubleAnswer+consts.Answer, t.answer1, t.answer2)
	return nil
}

// boot
//
//	@receiver t
//	@return error
func (t *tuningTrouble) boot() error {
	data, err := utils.NewFileReader(consts.TuningTroublePath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	t.buf = data
	return nil
}

// cs  TODO
type cs struct {
	limit int
	data  []string
	index int
}

// NewCS
//
//	@param limit
//	@return *cs
func NewCS(limit int) *cs {
	return &cs{
		limit: limit,
		data:  []string{},
	}
}

// Connect
//
//	@receiver c
//	@param line
//	@return bool
func (c *cs) Connect(line string) bool {
	for _, v := range line {
		if c.parseStr(v) {
			return true
		}
	}
	return false
}

// parseStr
//
//	@receiver c
//	@param char
//	@return bool
func (c *cs) parseStr(char rune) bool {
	c.index++
	for idx, val := range c.data {
		if val == string(char) {
			moveData := c.data[idx+1:]
			moveData = append(moveData, val)
			c.data = moveData
			return false
		}
	}
	c.data = append(c.data, string(char))
	return len(c.data) == c.limit
}

// GetIndex
//
//	@receiver c
//	@return int
func (c *cs) GetIndex() int {
	return c.index
}

// collect
//
//	@receiver t
func (t *tuningTrouble) collect() {
	reader := bufio.NewReader(t.buf)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if string(line) != "" {
			cs1 := NewCS(4)
			ok := cs1.Connect(string(line))
			if ok {
				t.answer1 = cs1.GetIndex()
			}

			cs2 := NewCS(14)
			ok2 := cs2.Connect(string(line))
			if ok2 {
				t.answer2 = cs2.GetIndex()
			}
		}

	}
}
