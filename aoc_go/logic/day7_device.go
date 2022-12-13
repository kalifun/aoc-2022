package logic

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/kalifun/aco-2022/aoc_go/entity/consts"
	"github.com/kalifun/aco-2022/aoc_go/repo/utils"
)

// device  TODO
type device struct {
	buf     *os.File
	answer1 interface{}
	answer2 interface{}
}

// NewDevice
//  @return *device
func NewDevice() *device {
	return &device{}
}

// GetStar
//
//	@receiver t
//	@return error
func (d *device) GetStar() error {
	err := d.boot()
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.DevicePath)
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
	data, err := utils.NewFileReader(consts.DevicePath)
	if err != nil {
		log.Fatalf("Failed to read the contents of the file, filename: %s", consts.DevicePath)
		return err
	}
	d.buf = data
	return nil
}

// collect
//  @receiver d
func (d *device) collect() {
	reader := bufio.NewReader(d.buf)
	t := NewTerminal()
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if string(line) != "" {
			t.ReadLine(string(line))
		}
	}
	d.answer1 = t.Part1Answer()
	d.answer2 = t.Part2Answer()
}

// terminal  TODO
type terminal struct {
	bucket          *bucket
	currentLocation *dir
}

// NewTerminal
//  @return *terminal
func NewTerminal() *terminal {
	return &terminal{
		bucket:          NewBucket(),
		currentLocation: nil,
	}
}

// ReadLine
//  @receiver t
//  @param line
func (t *terminal) ReadLine(line string) {
	// command
	if strings.HasPrefix(line, "$") {
		if strings.Contains(line, "cd") {
			splitLine := strings.Split(line, " ")
			t.cd(splitLine[len(splitLine)-1])
			return
		}

		if strings.Contains(line, "ls") {
			return
		}
	}

	// dir
	if strings.HasPrefix(line, "dir") {
		splitLine := strings.Split(line, " ")
		t.dir(splitLine[len(splitLine)-1])
		return
	}

	// file
	splitLine := strings.Split(line, " ")
	size, err := strconv.ParseInt(splitLine[0], 10, 64)
	if err != nil {
		return
	}
	t.file(splitLine[len(splitLine)-1], size)
}

// cd
//  @receiver t
//  @param path
func (t *terminal) cd(path string) {
	switch path {
	case "/":
		// root
		t.currentLocation = nil
	case "..":
		// jump up
		t.currentLocation = t.currentLocation.fatherDir
	default:
		// cd dir
		if t.currentLocation == nil {
			t.currentLocation = t.bucket.dirs[path]
		} else {
			t.currentLocation = t.currentLocation.dirs[path]
		}
	}
}

// dir
//  @receiver t
//  @param name
func (t *terminal) dir(name string) {
	if t.currentLocation == nil {
		t.bucket.Dir(name)
	} else {
		t.currentLocation.Dir(name)
	}
}

// file
//  @receiver t
//  @param name
//  @param size
func (t *terminal) file(name string, size int64) {
	// log.Printf("filename %s filesize %d \n", name, size)
	// 根目录
	t.bucket.Size(size)
	if t.currentLocation == nil {
		t.bucket.File(name, size)
		// t.bucket.Size(size)
		// log.Printf("root 目录的大小: %d \n", t.bucket.size)
	} else {
		t.currentLocation.File(name, size)
		// 将其记录当当前层的目录大小
		t.currentLocation.Size(size)
		// 将当前的大小也记录到上层大小
		t.fatherSize(t.currentLocation.fatherDir, size)
		// log.Printf("当前目录的大小: %d \n", t.currentLocation.size)
	}
}

// fatherSize
//  @receiver t
//  @param father
//  @param size
func (t *terminal) fatherSize(father *dir, size int64) {
	if father != nil {
		father.size += size
		if father.fatherDir != nil {
			t.fatherSize(father.fatherDir, size)
		}
	}
}

// func (t *terminal) Part1(dirs map[string]*dir, pre int) {
// 	for name, dir := range dirs {
// 		data := []string{}
// 		for i := 0; i < pre; i++ {
// 			data = append(data, "-------")
// 		}
// 		var filecount int64
// 		for _, v := range dir.files {
// 			filecount += v
// 		}
// 		tt := "False"
// 		if dir.size < 100000 {
// 			t.max += dir.size
// 			tt = "Yes"
// 		}

// 		fmt.Printf("%s [files %d]\n", strings.Join(data, ""), filecount)
// 		fmt.Printf("%s (%s %d)  %s  \n", strings.Join(data, ""), name, dir.size, tt)
// 		t.Part1(dir.dirs, pre+1)
// 	}
// }

// Part1Answer
//  @receiver t
//  @return interface{}
func (t *terminal) Part1Answer() interface{} {
	// counter
	var count int64
	for _, dir := range t.bucket.dirs {
		count += counter(dir)
	}
	return count
}

// Part2Answer
//  @receiver t
//  @return interface{}
func (t *terminal) Part2Answer() interface{} {
	// log.Printf("已使用大小 %d\n", t.bucket.size)
	total := 70000000
	unused := 30000000
	delete := unused - (total - int(t.bucket.size))
	del := NewDeleteWork(int64(delete))
	return del.MinDelSize(t.bucket.dirs)
}

// deleteWork  TODO
type deleteWork struct {
	delete int64
	option int64
}

// NewDeleteWork
//  @param del
//  @return *deleteWork
func NewDeleteWork(del int64) *deleteWork {
	return &deleteWork{
		delete: del,
	}
}

// Option
//  @receiver d
//  @param dirs
func (d *deleteWork) options(dirs map[string]*dir) {
	for _, dir := range dirs {
		if dir.size > d.delete {
			if d.option == 0 {
				d.option = dir.size
			} else {
				if d.option > dir.size {
					d.option = dir.size
				}
			}
		}
		d.options(dir.dirs)
	}
}

// MinDelSize
//  @receiver d
//  @param dirs
//  @return int64
func (d *deleteWork) MinDelSize(dirs map[string]*dir) int64 {
	d.options(dirs)
	return d.option
}

// counter
//  @param dir
//  @return int64
func counter(dir *dir) int64 {
	var size int64
	for _, v := range dir.dirs {
		if v.size < 100000 {
			size += v.size
		}
		count := counter(v)
		size += count
	}
	return size
}

// bucket  TODO
type bucket struct {
	size  int64
	files map[string]int64
	dirs  map[string]*dir
}

// NewBucket
//  @return *bucket
func NewBucket() *bucket {
	return &bucket{
		files: make(map[string]int64),
		dirs:  make(map[string]*dir),
	}
}

// File
//  @receiver b
//  @param name
//  @param size
func (b *bucket) File(name string, size int64) {
	b.files[name] = size
}

// Dir
//  @receiver b
//  @param name
//  @return *dir
func (b *bucket) Dir(name string) *dir {
	folder := NewDir(nil)
	b.dirs[name] = folder
	return folder
}

// Size
//  @receiver b
//  @param size
func (b *bucket) Size(size int64) {
	b.size += size
}

// dir  TODO
type dir struct {
	fatherDir *dir
	size      int64
	files     map[string]int64
	dirs      map[string]*dir
}

// NewDir
//  @param fatherDir
//  @return *dir
func NewDir(fatherDir *dir) *dir {
	return &dir{
		fatherDir: fatherDir,
		size:      0,
		files:     make(map[string]int64),
		dirs:      make(map[string]*dir),
	}
}

// Size
//  @receiver d
//  @param size
func (d *dir) Size(size int64) {
	d.size += size
}

// File
//  @receiver d
//  @param name
//  @param size
func (d *dir) File(name string, size int64) {
	d.files[name] = size
}

// Dir
//  @receiver d
//  @param name
//  @return *dir
func (d *dir) Dir(name string) *dir {
	folder := NewDir(d)
	d.dirs[name] = folder
	return folder
}
