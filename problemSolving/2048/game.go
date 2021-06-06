package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type TheGame struct {
	Map [][]int
	Win bool
}

func newGame() *TheGame {
	x := [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	aMap := &TheGame{Map: x, Win: false}
	aMap.newChar()
	return aMap
}

func (game *TheGame) PrintGame() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(game.Map[i][j], " ")
		}
		fmt.Println()
	}

}

// game mechanics

func (game *TheGame) newChar() {
	li := rand.Intn(4)
	lj := rand.Intn(4)

	if game.Map[li][lj] != 0 {
		game.newChar()
	} else {
		game.Map[li][lj] = 2
		return
	}
}

func (game *TheGame) moveUp(w http.ResponseWriter, r *http.Request) {
	for k := 0; k < 3; k++ {
		for j := 0; j < 4; j++ {
			for i := 0; i < 3; i++ {
				if game.Map[i][j] == 0 && game.Map[i+1][j] != 0 {
					game.Map[i][j] = game.Map[i+1][j]
					game.Map[i+1][j] = 0
				} else if game.Map[i][j] == game.Map[i+1][j] {
					game.Map[i][j] *= 2
					game.Map[i+1][j] = 0
					if game.Map[i][j] == 2048 {
						game.Win = true
					}
				}
			}
		}
	}
	game.newChar()
	tpl.ExecuteTemplate(w, "index.html", game)
}

func (game *TheGame) moveDown(w http.ResponseWriter, r *http.Request) {
	for k := 0; k < 3; k++ {
		for j := 0; j < 4; j++ {
			for i := 3; i > 0; i-- {
				if game.Map[i][j] == 0 && game.Map[i-1][j] != 0 {
					game.Map[i][j] = game.Map[i-1][j]
					game.Map[i-1][j] = 0
				} else if game.Map[i][j] == game.Map[i-1][j] {
					game.Map[i][j] *= 2
					game.Map[i-1][j] = 0
					if game.Map[i][j] == 2048 {
						game.Win = true
					}
				}
			}
		}
	}
	game.newChar()
	tpl.ExecuteTemplate(w, "index.html", game)
}

func (game *TheGame) moveRight(w http.ResponseWriter, r *http.Request) {
	for k := 0; k < 3; k++ {
		for i := 0; i < 4; i++ {
			for j := 3; j > 0; j-- {
				if game.Map[i][j] == 0 && game.Map[i][j-1] != 0 {
					game.Map[i][j] = game.Map[i][j-1]
					game.Map[i][j-1] = 0
				} else if game.Map[i][j] == game.Map[i][j-1] {
					game.Map[i][j] *= 2
					game.Map[i][j-1] = 0
					if game.Map[i][j] == 2048 {
						game.Win = true
					}
				}
			}
		}
	}
	game.newChar()
	tpl.ExecuteTemplate(w, "index.html", game)
}

func (game *TheGame) moveLeft(w http.ResponseWriter, r *http.Request) {
	for k := 0; k < 3; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 3; j++ {
				if game.Map[i][j] == 0 && game.Map[i][j+1] != 0 {
					game.Map[i][j] = game.Map[i][j+1]
					game.Map[i][j+1] = 0
				} else if game.Map[i][j] == game.Map[i][j+1] {
					game.Map[i][j] *= 2
					game.Map[i][j+1] = 0
					if game.Map[i][j] == 2048 {
						game.Win = true
					}
				}
			}
		}
	}
	game.newChar()
	tpl.ExecuteTemplate(w, "index.html", game)
}
