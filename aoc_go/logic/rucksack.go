package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type rucksack struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewRuclSack() *rucksack {
	return &rucksack{}
}

func (r *rucksack) GetStar() error {
	err := r.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	r.collect()
	log.Printf(consts.RucksackAnswer+consts.Answer, r.answer1, r.answer2)
	return nil
}

func (r *rucksack) boot() error {
	data, err := utils.NewFileReader(consts.RucksackPath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	r.buf = data
	return nil
}

func (r *rucksack) collect() {
	reader := bufio.NewReader(r.buf)
	var part1Sum, part2Sum int32
	var part2Bags []string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		part1Sum += theSameWord(string(line))

		if len(part2Bags) < 3 {
			part2Bags = append(part2Bags, string(line))
		}

		if len(part2Bags) == 3 {
			c := NewCategory(part2Bags)
			part2Sum += c.getBadges()
			part2Bags = []string{}
		}
	}
	r.answer1 = part1Sum
	r.answer2 = part2Sum
}

// theSameWord
//  @param str
//  @return int32
func theSameWord(str string) int32 {
	len := len(str)
	data := []rune(str)
	keysMap := make(map[rune]struct{})
	for i := 0; i < len/2; i++ {
		val := data[i]
		keysMap[getSite(val)] = struct{}{}
	}

	for i := len / 2; i < len; i++ {
		site := getSite(data[i])
		if _, ok := keysMap[site]; ok {
			return site
		}
	}
	return 0
}

// getSite
//  @param w
//  @return rune
func getSite(w rune) rune {
	yes := unicode.IsUpper(w)
	if yes {
		w -= 64
		w += 26
	} else {
		w -= 96
	}
	return w
}

type category struct {
	bags        []string
	fisrtBagMap map[rune]struct{}
	sameMap     map[rune]int
}

func NewCategory(data []string) *category {
	return &category{
		bags:        data,
		fisrtBagMap: make(map[rune]struct{}),
		sameMap:     make(map[rune]int),
	}
}

func (c *category) getBadges() int32 {
	if len(c.bags) == 3 {
		c.fisrtBagMap, c.sameMap = getRuneMap(c.bags[0])
		c.parseTag(c.bags[1])
		c.parseTag(c.bags[2])
	}

	fmt.Println(c.bags)
	fmt.Println(c.sameMap)
	// fmt.Println(c.bags)
	for k, v := range c.sameMap {
		if v == 3 {
			return k
		}
	}
	return 0
}

func (c *category) parseTag(str string) {
	data := []rune(str)
	keys := make(map[rune]struct{})
	for i := 0; i < len(str); i++ {
		site := getSite(data[i])
		if _, ok := keys[site]; !ok {
			keys[site] = struct{}{}
			if _, ok := c.fisrtBagMap[site]; ok {
				if v, ok := c.sameMap[site]; ok {
					c.sameMap[site] = v + 1
				}
			}
		}

	}
}

func getRuneMap(str string) (map[rune]struct{}, map[rune]int) {
	keysMap := make(map[rune]struct{})
	keys := make(map[rune]int)
	data := []rune(str)
	for i := 0; i < len(str); i++ {
		site := getSite(data[i])
		if _, ok := keysMap[site]; !ok {
			keysMap[site] = struct{}{}
		}

		if _, ok := keys[site]; !ok {
			keys[site] = 1
		}
	}
	return keysMap, keys
}
