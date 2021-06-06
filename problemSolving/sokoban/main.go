package main

import (
	"html/template"
	"net/http"
)

func main() {
	game := newGame("01")
	http.HandleFunc("/", game.index)
	http.HandleFunc("/moveUp", game.moveUp)
	http.HandleFunc("/moveDown", game.moveDown)
	http.HandleFunc("/moveLeft", game.moveLeft)
	http.HandleFunc("/moveRight", game.moveRight)
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func (game *Game) index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", game)
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
