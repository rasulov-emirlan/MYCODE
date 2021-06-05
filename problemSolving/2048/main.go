package main

import (
	"html/template"
	"net/http"
)

func main() {
	webGame()
}

func webGame() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func index(w http.ResponseWriter, r *http.Request) {
	game := newGame()
	tpl.ExecuteTemplate(w, "index.html", game)
}

// func consoleGame() {
// 	reader := bufio.NewReader(os.Stdin)

// 	game := newGame()
// 	game.PrintGame()

// 	for {
// 		fmt.Print("-> ")
// 		text, _ := reader.ReadString('\n')
// 		// convert CRLF to LF
// 		text = strings.Replace(text, "\n", "", -1)

// 		if strings.Compare("up", text) == 0 {
// 			game.moveUp()
// 			game.PrintGame()
// 		} else if strings.Compare("down", text) == 0 {
// 			game.moveDown()
// 			game.PrintGame()
// 		} else if strings.Compare("right", text) == 0 {
// 			game.moveRight()
// 			game.PrintGame()
// 		} else if strings.Compare("left", text) == 0 {
// 			game.moveLeft()
// 			game.PrintGame()
// 		}
// 	}
// }
