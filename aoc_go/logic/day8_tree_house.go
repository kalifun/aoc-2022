package logic

import (
	"bufio"
	"fmt"
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
	tMap.FoundAnswer()
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

func (tm *treeMap) FoundAnswer() {
	fmt.Println(tm.horizontalTree() + tm.verticalTree())
}

func (tm *treeMap) horizontalTree() int {
	var count int
	for index, horizontalTree := range tm.horizontal {
		var max int
		for i := 0; i < len(horizontalTree); i++ {
			val := horizontalTree[i]
			if index == 0 || index == len(horizontalTree)-1 {
				horizontalTree[i].Checked = true
				count++
				fmt.Printf("可以识别的%d\n", val.Val)
				continue
			}
			if i == 0 {
				max = horizontalTree[i].Val
				if !val.Checked {
					horizontalTree[i].Checked = true
					count++
					fmt.Printf("可以识别的%d\n", val.Val)
				}
			}
			if val.Checked {
				// pass
				continue
			} else {
				if val.Val > max {
					max = val.Val
					horizontalTree[i].Checked = true
					count++
					fmt.Printf("可以识别的%d\n", val.Val)
				}
			}
		}

		max = 0
		for i := len(horizontalTree) - 1; i >= 0; i-- {
			val := horizontalTree[i]
			if i == len(horizontalTree)-1 {
				max = horizontalTree[i].Val
				if !val.Checked {
					horizontalTree[i].Checked = true
					count++
					fmt.Printf("可以识别的%d\n", val.Val)
				}
			}
			if val.Checked {
				// pass
				continue
			} else {
				if val.Val > max {
					max = val.Val
					horizontalTree[i].Checked = true
					count++
					fmt.Printf("可以识别的%d\n", val.Val)
				}
			}
		}
	}
	return count
}

func (tm *treeMap) verticalTree() int {
	var count int
	for _, verticalTree := range tm.vertical {
		var max int
		for i := 0; i < len(verticalTree); i++ {
			val := verticalTree[i]
			if i == 0 {
				max = verticalTree[i].Val
				if !val.Checked {
					verticalTree[i].Checked = true
					count++
					fmt.Printf("x可以识别的%d\n", val.Val)
				}
			}
			if val.Checked {
				// pass
				continue
			} else {
				if val.Val > max {
					fmt.Println(max, verticalTree[i])
					max = val.Val
					verticalTree[i].Checked = true
					count++
					fmt.Printf("xx可以识别的%d %v\n", val.Val, verticalTree)
				}
			}
		}

		max = 0
		for i := len(verticalTree) - 1; i >= 0; i-- {
			val := verticalTree[i]
			if i == len(verticalTree)-1 {
				max = verticalTree[i].Val
				if !val.Checked {
					verticalTree[i].Checked = true
					count++
					fmt.Printf("y可以识别的%d\n", val.Val)
				}
			}
			if val.Checked {
				// pass
				continue
			} else {
				if val.Val > max {
					max = val.Val
					verticalTree[i].Checked = true
					count++
					fmt.Printf("yy可以识别的%d\n", val.Val)
				}
			}
		}
	}
	return count
}
