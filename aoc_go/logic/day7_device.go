package logic

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

type device struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// GetStar
//
//	@receiver t
//	@return error
func (d *device) GetStar() error {
	err := d.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	d.collect()
	log.Printf(consts.Deviceanswer+consts.Answer, d.answer1, d.answer2)
	return nil
}

// boot
//
//	@receiver t
//	@return error
func (d *device) boot() error {
	data, err := utils.NewFileReader(consts.TuningTroublePath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.CaloriePath)
		return err
	}
	d.buf = data
	return nil
}

func (d *device) collect() {
	reader := bufio.NewReader(d.buf)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

	}
}

type terminal struct {
	bucket *bucket
}

func NewTerminal() *terminal {
	return &terminal{
		bucket: NewBucket(),
	}
}

func (t *terminal) ReadLine(line string) {

}

type bucket struct {
	files map[string]int64
	dirs  map[string]dir
}

func NewBucket() *bucket {
	return &bucket{}
}

func (b *bucket) File(name string, size int64) {
	b.files[name] = size
}

func (b *bucket) Dir(name string) dir {
	folder := NewDir()
	b.dirs[name] = *folder
	return *folder
}

type dir struct {
	size  int64
	files map[string]int64
	dirs  map[string]dir
}

func NewDir() *dir {
	return &dir{
		size:  0,
		files: make(map[string]int64),
		dirs:  make(map[string]dir),
	}
}

func (d *dir) Size(size int64) {
	d.size = size
}

func (d *dir) File(name string, size int64) {
	d.files[name] = size
}

func (d *dir) Dir(name string) dir {
	folder := NewDir()
	d.dirs[name] = *folder
	return *folder
}
