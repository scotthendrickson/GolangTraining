package main

import (
	"net/http"
	"text/template"
	"encoding/json"
	"strings"
	"fmt"
	"math/rand"
	"time"
)

type Context struct {
	Input  string
	Output string
	Color int
}

func main() {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", myHomeFunc)
	myMux.HandleFunc("/translate", translationFunc)
	http.ListenAndServe(":8080", myMux)
}

func myHomeFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	tmpl, err := template.New("anyNameForTemplate").Parse(doc)
	if err == nil {
		tmpl.Execute(w, nil)
	}
}

func randomColor () int {
	//var color = '#' + (Math.random() * 0xFFFFFF << 0).toString(16);
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(999999)
}

func morseTranslator (input string) (output string) {
	morseCode := map[string]string{
		".-.-.- " : ".",
		"--..-- " : ",",
		"---... " : ":",
		"..--.. " : "?",
		".----. " : "'",
		"-.-.-- " : "!",
		"-.--. " : "(",
		"-.--.- " : ")",
		".-... " : "&",
		"-.-.-. " : ";",
		"-...- " : "=",
		".-.-. " : "+",
		"..--.- " : "_",
		".-..-." : `"`,
		".--.-. " : "@",
		"...-..- " : "$",
		".- " : "a",
		"-... " : "b",
		"-.-. " : "c",
		"-.. " : "d",
		". " : "e",
		"..-. " : "f",
		"--. " : "g",
		".... " : "h",
		".. " : "i",
		".--- " : "j",
		"-.- " : "k",
		".-.. " : "l",
		"-- " : "m",
		"-. " : "n",
		"--- " : "o",
		".--. " : "p",
		"--.- " : "q",
		".-. " : "r",
		"... " : "s",
		"- " : "t",
		"..- " : "u",
		"...- " : "v",
		".-- " : "w",
		"-..- " : "x",
		"-.-- " : "y",
		"--.. " : "z",
		"----- " : "0",
		".---- " : "1",
		"..--- " : "2",
		"...-- " : "3",
		"....- " : "4",
		"..... " : "5",
		"-.... " : "6",
		"--... " : "7",
		"---.. " : "8",
		"----. " : "9",
		" " : " ",
	}
	switch firstChar := string(input[0]); firstChar {
	case ".", "-":
		var temp string
		for i := 0; i < len(input); i++ {
			temp += string(input[i])
			if string(input[i]) == " " {
				for key, value := range morseCode {
					if temp == key {
						output += value
						temp = ""
						continue
					}
				}
			}

		}
	default:
		for i := 0; i < len(input); i++ {
			for key, value := range morseCode {
				if string(input[i]) == value {
					output += key
					continue
				}
			}
		}

	}
	return output
}

func translationFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content Type", "text/html")
	r.ParseForm()
	input := strings.ToLower(strings.Join(r.Form["code"], ""))
	var output string = morseTranslator(input)
	color := randomColor()

	tmpl, err := template.New("anyNameForTemplate").Parse(docEr)
	if err == nil {
		context := Context{input, output, color}
		tmpl.Execute(w, context)
	} else {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Nope")
	}
}

const doc = `
<!DOCTYPE html>
<html>
<head lang="en">
	<meta charset="UTF-8">
	<title>Morse Code Translator</title>
</head>
<body>
	<h1 style="color:blue">Hello User!</h1>
	<div>
		<div >
			<form action="/translate" method="post">
				<textarea  name="code" rows="8" cols="40" placeholder="Enter your text or Morse Code here:" ></textarea>
				<br>
				<input type="submit" value="Translate">
			</form>
		</div>
		<div >
			<textarea name="output" rows="8" cols="40" placeholder="Your translation will appear here"  readonly></textarea>
		</div>
	</div>
</body>
</html>
`
const docEr = `
<!DOCTYPE html>
<html>
<head lang="en">
	<meta charset="UTF-8">
	<title>First Template</title>
</head>
<body style="background:#{{.Color}}">
	<h1 style="color:blue">Hello User!</h1>
	<div>
		<div >
			<form action="/translate" method="post">
				<textarea  name="code" rows="8" cols="40" placeholder="Enter your text or Morse Code here:" >{{.Input}}</textarea>
				<input type="submit" value="input">
			</form>
		</div>
		<div >
			<textarea name="output" rows="8" cols="40" placeholder="Your translation will appear here"  readonly>{{.Output}}</textarea>
		</div>
	</div>
</body>
</html>
`