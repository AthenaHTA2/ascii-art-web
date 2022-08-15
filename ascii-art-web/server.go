package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

func BubbleConv(inputword string, chars map[int]string) string {

	var result string

	var splitbyNewLine []string

	// checks if there is no word entered//
	if len(inputword) == 0 {
		return ""
	}
	// checks if the word entered is a new line, if so it prints a new line//

	if inputword == "\\n" {
		return "\n"
	}

	// this loops through the input string"
	for a, b := 0, 0; a < len(inputword); a++ {

		// this checks if its the end of the string"
		if a+1 <= len(inputword)-1 {

			// this checks if there is a "\n" in the middle of the string and then seperates it into new strings//
			if inputword[a] == '\\' && inputword[a+1] == 'n' {

				// this seperates the string up to the "\n" being found
				splitbyNewLine = append(splitbyNewLine, inputword[b:a])

				// this starts from the next word after "/n"//
				b = a + 2
			}

			// this checks if you've reached the end of the string
		} else if a == len(inputword)-1 {

			// this adds the rest of the string //
			splitbyNewLine = append(splitbyNewLine, inputword[b:])
		}
	}

	// this loop ranges through every new line
	for _, b := range splitbyNewLine {

		var bubbleWriting []string

		// this loops through each character on the line
		for _, turntoBubbleWriting := range b {

			// this calls the character to the corresponding bubblewriting character on the map
			bubbleWriting = append(bubbleWriting, chars[int(turntoBubbleWriting)])
		}
		// this loops through the height of the the bubble character
		for linePos := 0; linePos < 8; linePos++ {

			// this loops through each individual bubble character
			for _, indCharacter := range bubbleWriting {

				// this adds the bubble character to the result
				splitLine := strings.Split(indCharacter, "\n")
				result += splitLine[linePos]

			}
			result += "\n"
		}
	}
	return result
}

func MapAscii(banner string) map[int]string {

	// this reads the banner text file where the bubble ascii characters are held
	file, err := ioutil.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}

	// this splits the bubble character text file by new individualLines
	individualLines := strings.Split(string(file), "\n")

	// this creates a map that will store the bubble characters
	compiledLines := make(map[int]string)

	// this assigns a map position to each bubble character corresponding to their ascii position
	for a, b := 32, 0; a < 127; a++ {
		c := b + 9
		temp := ""

		// this gathers all the lines from the character and adds it to the temporary string
		for b < c {
			temp += individualLines[b+1] + "\n"
			b++
		}

		// this assigns the bubble writing character to its ascii value
		compiledLines[a] = temp
	}

	return compiledLines
}

func PrintAsciiArt(input, ban string) string {
	// this reads the txt file where the bubble writing characters are held
	Bubblecharacters := MapAscii("banners/" + ban + ".txt")

	// this runs the input through the map and returns the bubble character equivalent
	output := BubbleConv(input, Bubblecharacters)
	return output
}

func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		p := ""
		t, _ := template.ParseFiles(("form.html"))
		t.Execute(w, p)
	case "POST":

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text := r.FormValue("text")
		font := r.FormValue("font")
		p := PrintAsciiArt(text, font)

		t, _ := template.ParseFiles(("form.html"))
		t.Execute(w, p)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {

	fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs)) // handling the CSS

	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
