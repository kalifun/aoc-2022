package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/kalifun/aco-2022/entity/consts"
	"github.com/kalifun/aco-2022/repo/utils"
)

type treeHouse struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewTreeHouse() *treeHouse {
	return &treeHouse{}
}

// GetStar
//
//	@receiver t
//	@return error
func (t *treeHouse) GetStar() error {
	err := t.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.TreeHouse)
		return err
	}
	t.collect()
	log.Printf(consts.TreeHouseAnswer+consts.Answer, t.answer1, t.answer2)
	return nil
}

// boot
//
//	@receiver t
//	@return error
func (t *treeHouse) boot() error {
	data, err := utils.NewFileReader(consts.TreeHouse)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.TreeHouse)
		return err
	}
	t.buf = data
	return nil
}

func (t *treeHouse) collect() {
	reader := bufio.NewReader(t.buf)
	tMap := NewTreeMap()
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		tMap.ReadLine(string(line))
	}
	t.answer1 = tMap.Part1Answer()
	t.answer2 = tMap.Part2Answer()
}

type treeInfo struct {
	Val     int
	Checked bool
}

type treeMap struct {
	horizontal [][]*treeInfo
	vertical   [][]*treeInfo
}

func NewTreeMap() *treeMap {
	return &treeMap{}
}

func (tm *treeMap) ReadLine(str string) {
	var init bool
	if tm.horizontal == nil && tm.vertical == nil {
		init = true
	}
	var horizontalTree []*treeInfo
	for i := 0; i < len(str); i++ {
		val, _ := strconv.Atoi(string(str[i]))
		info := treeInfo{
			Val: val,
		}

		if init {
			var verticalTree []*treeInfo
			verticalTree = append(verticalTree, &info)
			tm.vertical = append(tm.vertical, verticalTree)
		} else {
			tm.vertical[i] = append(tm.vertical[i], &info)
		}

		horizontalTree = append(horizontalTree, &info)
	}
	tm.horizontal = append(tm.horizontal, horizontalTree)
}

// Part1Answer
//
//	@receiver tm
//	@return interface{}
func (tm *treeMap) Part1Answer() interface{} {
	return visibleQuantity(tm.horizontal) + visibleQuantity(tm.vertical)
}

func (tm *treeMap) Part2Answer() interface{} {
	var sumRes int
	for index, tree := range tm.horizontal {
		for i := 0; i < len(tree); i++ {
			val := tree[i]
			top := tm.vertical[i][0:index]
			left := tree[0:i]
			right := tree[i+1:]
			down := tm.vertical[i][index+1:]
			var count []int
			count = append(count, reverse(top, val.Val))
			count = append(count, order(down, val.Val))
			count = append(count, reverse(left, val.Val))
			count = append(count, order(right, val.Val))

			if sum(count) > sumRes {
				sumRes = sum(count)
			}
		}
	}
	return sumRes
}

func sum(res []int) int {
	var sum int
	for i, v := range res {
		if i == 0 {
			sum = v
			continue
		}
		sum *= v
	}
	return sum
}

func order(trees []*treeInfo, val int) int {
	var count int
	for _, tree := range trees {
		count += 1
		if tree.Val >= val {
			break
		}
	}
	return count
}

func reverse(trees []*treeInfo, val int) int {
	var count int
	for i := len(trees) - 1; i >= 0; i-- {
		valT := trees[i]
		count += 1
		if valT.Val >= val {
			break
		}
	}
	return count
}

// visibleQuantity
//
//	@param data
//	@return int
func visibleQuantity(data [][]*treeInfo) int {
	var count int
	for index, tree := range data {
		var max int
		for i := 0; i < len(tree); i++ {
			val := tree[i]
			if i == 0 || index == 0 || index == len(tree)-1 {
				max = val.Val
				if !val.Checked {
					count++
					tree[i].Checked = true
				}
				continue
			}

			if val.Val > max {
				max = val.Val
				if val.Checked {
					continue
				}
				count++
				tree[i].Checked = true
			}
		}

		for i := len(tree) - 1; i >= 0; i-- {
			val := tree[i]
			if i == len(tree)-1 {
				max = val.Val
				if !val.Checked {
					count++
					tree[i].Checked = true
				}
				continue
			}

			if val.Val > max {
				max = val.Val
				if val.Checked {
					continue
				}
				count++
				tree[i].Checked = true
			}
		}
	}
	return count
}
