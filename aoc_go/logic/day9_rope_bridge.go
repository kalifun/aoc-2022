package logic

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/kalifun/aco-2022/entity/consts"
	"github.com/kalifun/aco-2022/repo/utils"
)

type ropeBridge struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

func NewRopeBridge() *ropeBridge {
	return &ropeBridge{}
}

// GetStar
//
//	@receiver t
//	@return error
func (r *ropeBridge) GetStar() error {
	err := r.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.RopeBridge)
		return err
	}
	r.collect()
	log.Printf(consts.RopeBridgeAnswer+consts.Answer, r.answer1, r.answer2)
	return nil
}

// boot
//
//	@receiver t
//	@return error
func (r *ropeBridge) boot() error {
	data, err := utils.NewFileReader(consts.RopeBridge)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.RopeBridge)
		return err
	}
	r.buf = data
	return nil
}

func (r *ropeBridge) collect() {
	reader := bufio.NewReader(r.buf)
	b := NewBridge()
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		b.ReadLine(string(line))
	}
	r.answer1 = b.Part1Answer()
	// t.answer2 = tMap.Part2Answer()
}

type bridge struct {
	source pts
	tail   pts
	move
	tailMap map[pts]struct{}
}

func NewBridge() *bridge {
	return &bridge{
		tailMap: map[pts]struct{}{},
	}
}

func (b *bridge) ReadLine(str string) {
	line := strings.Split(str, " ")
	forward := line[0]
	num, _ := strconv.Atoi(line[1])
	switch forward {
	case "L":
		b.move = left{}
	case "R":
		b.move = right{}
	case "U":
		b.move = top{}
	case "D":
		b.move = down{}
	}

	for i := 0; i < num; i++ {
		b.move.Run(&b.source)
		b.tail_element()
	}
}

func (b *bridge) tail_element() {
	xdelta := b.source.x - b.tail.x
	ydelta := b.source.y - b.tail.y
	if math.Abs(float64(xdelta)) > 1.0 || math.Abs(float64(ydelta)) > 1.0 {
		b.tail.x += sign(xdelta)
		b.tail.y += sign(ydelta)
	}
	pt := pts{x: b.tail.x, y: b.tail.y}
	if _, ok := b.tailMap[pt]; !ok {
		b.tailMap[pt] = struct{}{}
	}
}

func sign(x int) int {
	return btou(x > 0) - btou(x < 0)
}

func btou(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (b *bridge) Part1Answer() interface{} {
	return len(b.tailMap)
}

type pts struct {
	x int
	y int
}

type move interface {
	Run(pts *pts)
}

type top struct{}

func (t top) Run(pts *pts) {
	pts.y -= 1
}

type down struct{}

func (t down) Run(pts *pts) {
	pts.y += 1
}

type left struct{}

func (t left) Run(pts *pts) {
	pts.x -= 1
}

type right struct{}

func (t right) Run(pts *pts) {
	pts.x += 1
}
