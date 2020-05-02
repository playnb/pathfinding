package pathfinding

import (
	"fmt"
	"github.com/playnb/util/mathex"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

const (
	COST_STRAIGHT = 1000
	COST_DIAGONAL = 1414
)

var adjust8_ = [][3]int{
	{-1, -1, COST_DIAGONAL},
	{-1, 0, COST_STRAIGHT},
	{-1, 1, COST_DIAGONAL},
	{0, -1, COST_STRAIGHT},
	{0, 1, COST_STRAIGHT},
	{1, -1, COST_DIAGONAL},
	{1, 0, COST_STRAIGHT},
	{1, 1, COST_DIAGONAL},
}

type MapData struct {
	width  int
	height int
	data   []int
}

func (md *MapData) GetWidth() int {
	return md.width
}
func (md *MapData) GetHeight() int {
	return md.height
}
func (md *MapData) IsBlock(x, y int) bool {
	if x >= 0 && x < md.width && y >= 0 && y < md.height {
		return md.data[x+y*md.width] > 0
	}
	return true
}
func (md *MapData) GetNeighbors(x int, y int) [][3]int {
	var ret [][3]int
	for _, v := range adjust8_ {
		px := x + v[0]
		py := y + v[1]
		if md.IsBlock(px, py) {
			continue
		}
		ret = append(ret, [3]int{px, py, v[2]})
	}
	return ret
}

func loadMapData(filename string) *MapData {
	f, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer f.Close()
	img, _, _ := image.Decode(f)

	bounds := img.Bounds()
	md := &MapData{
		width:  bounds.Max.X,
		height: bounds.Max.Y,
		data:   make([]int, bounds.Max.X*bounds.Max.Y),
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a > 0 {
				md.data[x+y*md.width] = 1
			}
		}
	}
	return md
}

func saveMapData(filename string, md *MapData) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	nrgba := image.NewNRGBA(image.Rect(0, 0, md.width, md.height))
	for x := 0; x < md.width; x++ {
		for y := 0; y < md.height; y++ {
			switch md.data[x+y*md.width] {
			case 1:
				nrgba.Set(x, y, color.RGBA{0, 0, 0, 255})
			case 2:
				nrgba.Set(x, y, color.RGBA{200, 100, 0, 255})

			}
		}
	}
	err = png.Encode(file, nrgba)
	if err != nil {
		fmt.Println(err)
	}
}

func TestAStar(t *testing.T) {
	md := loadMapData("image/map.png")
	path := AStar(md, 0, 0, 70, 70, func(from *Node, to *Node) int {
		return mathex.MaxInt(mathex.AbsInt(from.X-to.X), mathex.AbsInt(from.Y-to.Y))
	})

	if path != nil {
		for _, n := range path {
			md.data[n.X+n.Y*md.width] = 2
		}
		saveMapData("image/path.png", md)
	} else {
		fmt.Println("No way...")
	}
}

func BenchmarkAStar(b *testing.B) {
	md := loadMapData("image/map.png")
	var path []*Node
	b.Run("AStar", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			path = AStar(md, 0, 0, 70, 70, func(from *Node, to *Node) int {
				return 1000 * (mathex.MaxInt(mathex.AbsInt(from.X-to.X), mathex.AbsInt(from.Y-to.Y)))
			})
		}
	})
	if path != nil {
		for _, n := range path {
			md.data[n.X+n.Y*md.width] = 2
		}
		saveMapData("image/path.png", md)
	} else {
		fmt.Println("No way...")
	}
}
