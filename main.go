// THIS FILE IS PART OF Happy New Year 2025 PROJECT
// main.go - THE CORE PART OF THE PROJECT
//
// THIS PROGRAM IS FREE SOFTWARE
// YOU SHOULD HAVE RECEIVED A COPY OF MIT LICENSE
//
// Copyright (c) 2024-2025 宋子桓(Song Zihuan)

package main // new year
import (
	"fmt"
	"time"
)

type Feeling int

const (
	Happy   Feeling = 1
	Sadness Feeling = 2
)

var world *World

type World struct {
	Msg chan Feeling
}

func GodMakeTheWorld() {
	world = &World{make(chan Feeling, 10)}
}

type You struct{}

func Born() *You {
	return &You{}
}

func (u *You) FeelHappy(...any) {
	world.Msg <- Happy
}

func (u *You) FeelSad(...any) {
	world.Msg <- Sadness
}

func init() {
	GodMakeTheWorld()
}

func main() {
	go func() {
		for {
			defer func() {
				_ = recover()
				// 但是不要因此自责，只要我们努力，就能创造奇迹！
				fmt.Println("But don’t blame yourself for this, as long as we work hard, we can create miracles!")
			}()

			feeling := <-world.Msg
			switch feeling {
			case Happy:
				// 你若感到快乐，这个世界也会因此而变得美好
				fmt.Println("If you are happy, the world will become a better place.")
			case Sadness:
				// 如果你感到悲伤，世界也会因此暗淡
				panic("If you feel sad, the world will be dark too.")
			}
		}
	}()

	you := Born()

	// 在新的一年里，祝愿你天天开心。
	you.FeelHappy("In the new year, I wish you happiness every day.")
	time.Sleep(time.Hour * 24 * 365)
}
