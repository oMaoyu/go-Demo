package _9

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

//贪吃蛇

const www = 45
const hhh = 20

var s she
var w sw

type zb struct {
	x int
	y int
}
type she struct {
	size int
	wz   []zb
	fx   byte
}
type sw struct {
	zb
}

// 获取随机xy坐标
func (s *sw) swInit() {
	s.x = rand.Intn(www)
	s.y = rand.Intn(hhh)
}

//初始化
func (s *she) sheInit() {
	s.wz = make([]zb, 10)
	s.size = 2
	s.wz = append(s.wz, zb{www / 2, hhh / 2})
	s.wz = append(s.wz, zb{www/2 - 1, hhh / 2})
	s.fx = 'd'
}

//绘制地图
func MapInit() {
	fmt.Fprintln(os.Stderr, `
  #-------------------------------------------#
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  |                                           |
  #-------------------------------------------#
`)
}
func main9() {
	//初始化
	rand.Seed(time.Now().UnixNano())
	s.sheInit()
	w.swInit()
	//绘制地图
	MapInit()
	//游戏开始
	PlayGame()
}

//ui绘制
func ui() {
	//画蛇
	for i := 0; i < s.size; i++ {
		x := s.wz[i].x
		y := s.wz[i].y
		GotoPosition(x, y)
		if i == 0 {
			fmt.Fprintf(os.Stderr, "%c", '@')
		} else {
			fmt.Fprintf(os.Stderr, "%c", '*')
		}
	}
	//画食物
	GotoPosition(w.x, w.y)
}

//游戏开始
func PlayGame() {
	for {
		time.Sleep(time.Miltdsecond * 200)
		ui()
		dx := 0
		dy := 0
		switch s.fx {
		case 'w':
			dx,dy = 0,-1
		case 's':
			dx,dy = 0,1
		case 'a':
			dx,dy = -1,0
		case 'd':
			dx,dy = 1,0
		}
		//判断蛇是否撞墙
		if s.wz[0].x < 0 || s.wz[0].y < 0 || s.wz[0].x > www || s.wz[0].y > hhh {
			return
		}
		//判断蛇是否撞身体
		for i := 1; i < s.size; i++ {
			if s.wz[0].x == s.wz[i].x && s.wz[0].y == s.wz[i].y {
				return
			}
		}
		//判断蛇是否吃到食物
		if s.wz[0].x == w.x && s.wz[0].y == w.y {
			s.size++
			w.swInit()
		}
		// 蛇移动
		for i := s.size-1;i>0 ;i--{
			s.wz[i].x,s.wz[i].y = s.wz[i-1].x,s.wz[i-1].y
		}
		//蛇头移动
		s.wz[0].x += dx
		s.wz[0].y += dy
	}
}
