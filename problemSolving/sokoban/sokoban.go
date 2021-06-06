package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Game struct {
	Map       [][]byte
	OgMap     []string
	ReqPoints int
	OwnPoints int
	I         int
	J         int
}

func newGame(level string) Game {
	f, err := os.Open("data/levels/MiniCosmos.txt")
	if err != nil {
		log.Panicln("Error loading map", err.Error())
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)

	var aMap []string
	var tempS string
	aReq := 0
	own := 0
	var ik int
	var jk int

	for scan.Scan() {
		tempS = scan.Text()
		aMap = append(aMap, tempS)
		if strings.Contains(tempS, ";") {
			if strings.Contains(tempS, level) {
				aMap = aMap[:len(aMap)-1]
				break
			} else {
				aMap = nil
			}
		}
	}

	var finalMap [][]byte
	for i := 0; i < len(aMap); i++ {
		finalMap = append(finalMap, []byte(aMap[i]))
	}

	for i := 0; i < len(finalMap); i++ {
		for j := 0; j < len(finalMap[i]); j++ {
			if finalMap[i][j] == '@' {
				ik = i
				jk = j
			} else if finalMap[i][j] == '.' {
				aReq++
			} else if finalMap[i][j] == '*' {
				own++
			}
		}
	}

	aBoard := Game{
		Map:       finalMap,
		OgMap:     aMap,
		ReqPoints: aReq,
		OwnPoints: own,
		I:         ik,
		J:         jk,
	}
	return aBoard
}

func (game *Game) printBoard() {
	for i := 1; i < len(game.Map); i++ {
		for j := 0; j < len(game.Map[i]); j++ {
			fmt.Printf("%c", game.Map[i][j])
		}
		fmt.Println()
	}
	fmt.Println("OGmap: ")
	for i := 1; i < len(game.OgMap); i++ {
		for j := 0; j < len(game.OgMap[i]); j++ {
			fmt.Printf("%c", game.OgMap[i][j])
		}
		fmt.Println()
	}
	fmt.Println("The owned points: ", game.OwnPoints)
	fmt.Println("The points required: ", game.ReqPoints)
}

// game mechanics

func (game *Game) moveUp() {
	if game.Map[game.I-1][game.J] == ' ' {
		game.Map[game.I-1][game.J] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.I--
		return
	}
	if game.Map[game.I-1][game.J] == '.' {
		game.Map[game.I-1][game.J] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.I--
		return
	}
	if game.Map[game.I-1][game.J] == '$' && game.I > 1 {
		if game.Map[game.I-2][game.J] == ' ' {
			game.Map[game.I-2][game.J] = '$'
			game.Map[game.I-1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.I--
		} else if game.Map[game.I-2][game.J] == '.' {
			game.Map[game.I-2][game.J] = '*'
			game.Map[game.I-1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.I--
		}
		return
	}
	if game.Map[game.I-1][game.J] == '*' && game.I > 1 {
		if game.Map[game.I-2][game.J] == ' ' {
			game.Map[game.I-2][game.J] = '$'
			game.Map[game.I-1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints--
			game.I--
		} else if game.Map[game.I-2][game.J] == '.' {
			game.Map[game.I-2][game.J] = '*'
			game.Map[game.I-1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.I--
		}
		return
	}
}

func (game *Game) moveDown() {
	if game.Map[game.I+1][game.J] == ' ' {
		game.Map[game.I+1][game.J] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.I++
		return
	}
	if game.Map[game.I+1][game.J] == '.' {
		game.Map[game.I+1][game.J] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.I++
		return
	}
	if game.Map[game.I+1][game.J] == '$' && game.I < len(game.OgMap)-1 {
		if game.Map[game.I+2][game.J] == ' ' {
			game.Map[game.I+2][game.J] = '$'
			game.Map[game.I+1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.I++
		} else if game.Map[game.I+2][game.J] == '.' {
			game.Map[game.I+2][game.J] = '*'
			game.Map[game.I+1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.I++
		}
		return
	}
	if game.Map[game.I+1][game.J] == '*' && game.I < len(game.OgMap) {
		if game.Map[game.I+2][game.J] == ' ' {
			game.Map[game.I+2][game.J] = '$'
			game.Map[game.I+1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints--
			game.I++
		} else if game.Map[game.I+2][game.J] == '.' {
			game.Map[game.I+2][game.J] = '*'
			game.Map[game.I+1][game.J] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.I++
		}
	}
}

func (game *Game) moveLeft() {
	if game.Map[game.I][game.J-1] == ' ' {
		game.Map[game.I][game.J-1] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.J--
		return
	}
	if game.Map[game.I][game.J-1] == '.' {
		game.Map[game.I][game.J-1] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.J--
		return
	}
	if game.Map[game.I][game.J-1] == '$' && game.J > 1 {
		if game.Map[game.I][game.J-2] == ' ' {
			game.Map[game.I][game.J-2] = '$'
			game.Map[game.I][game.J-1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.J--
		} else if game.Map[game.I][game.J-2] == '.' {
			game.Map[game.I][game.J-2] = '*'
			game.Map[game.I][game.J-1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.J--
		}
		return
	}
	if game.Map[game.I][game.J-1] == '*' && game.J > 1 {
		if game.Map[game.I][game.J-2] == ' ' {
			game.Map[game.I][game.J-2] = '$'
			game.Map[game.I][game.J-1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints--
			game.J--
		} else if game.Map[game.I][game.J-2] == '.' {
			game.Map[game.I][game.J-2] = '*'
			game.Map[game.I][game.J-1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.J--
		}
	}
}

func (game *Game) moveRight() {
	if game.Map[game.I][game.J+1] == ' ' {
		game.Map[game.I][game.J+1] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.J++
		return
	}
	if game.Map[game.I][game.J+1] == '.' {
		game.Map[game.I][game.J+1] = '@'
		if game.OgMap[game.I][game.J] == '.' {
			game.Map[game.I][game.J] = '.'
		} else {
			game.Map[game.I][game.J] = ' '
		}
		game.J++
		return
	}
	if game.Map[game.I][game.J+1] == '$' && game.J < len(game.OgMap[game.I])-1 {
		if game.Map[game.I][game.J+2] == ' ' {
			game.Map[game.I][game.J+2] = '$'
			game.Map[game.I][game.J+1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.J++
		} else if game.Map[game.I][game.J+2] == '.' {
			game.Map[game.I][game.J+2] = '*'
			game.Map[game.I][game.J+1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.J++
		}
		return
	}
	if game.Map[game.I][game.J+1] == '*' && game.J < len(game.OgMap[game.I])-1 {
		if game.Map[game.I][game.J+2] == ' ' {
			game.Map[game.I][game.J+2] = '$'
			game.Map[game.I][game.J+1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints--
			game.J++
		} else if game.Map[game.I][game.J+2] == '.' {
			game.Map[game.I][game.J+2] = '*'
			game.Map[game.I][game.J+1] = '@'
			if game.OgMap[game.I][game.J] == '.' {
				game.Map[game.I][game.J] = '.'
			} else {
				game.Map[game.I][game.J] = ' '
			}
			game.OwnPoints++
			game.J++
		}
	}
}
