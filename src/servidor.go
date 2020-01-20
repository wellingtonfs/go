package main

import(
	"fmt"
	"net/http"
	"log"
	"io/ioutil"
	//"time"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func Tempo(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "%s", time.Now().Format("02-01-2006 15:04:05"))
	//title := r.URL.Path[len("/edit/"):]
	title := "/tempokkkk"
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    fmt.Fprintf(w, "<h1>Editing %s</h1>"+
        "<form action=\"/save/%s\" method=\"POST\">"+
        "<textarea name=\"body\">%s</textarea><br>"+
        "<input type=\"submit\" value=\"Save\">"+
        "</form>",
        p.Title, p.Title, p.Body)
}

func main(){
	http.HandleFunc("/", Tempo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}