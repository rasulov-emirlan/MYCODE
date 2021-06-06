package main

func main() {

}

// Console game :
// reader := bufio.NewReader(os.Stdin)

// 	game := newGame("07")
// 	game.printBoard()

// 	for {
// 		fmt.Print("-> ")
// 		text, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		text = strings.Replace(text, "\n", "", -1)

// 		if strings.Compare("up", text) == 0 {
// 			game.moveUp()
// 			game.printBoard()
// 		} else if strings.Compare("down", text) == 0 {
// 			game.moveDown()
// 			game.printBoard()
// 		} else if strings.Compare("left", text) == 0 {
// 			game.moveLeft()
// 			game.printBoard()
// 		} else if strings.Compare("right", text) == 0 {
// 			game.moveRight()
// 			game.printBoard()
// 		}
// 		if strings.Compare("end", text) == 0 {
// 			return
// 		}
// 	}
