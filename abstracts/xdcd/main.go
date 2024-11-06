package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Comic struct {
	Month      int    `json:"month,string"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       int    `json:"year,string"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        int    `json:"day,string"`
}

// onth	"1"
// num	2
// link	""
// year	"2006"
// news	""
// safe_title	"Petit Trees (sketch)"
// transcript	"[[Two trees are growing on opposite sides of a sphere.]]\n{{Alt-title: 'Petit' being a reference to Le Petit Prince, which I only thought about halfway through the sketch}}"
// alt	"'Petit' being a reference to Le Petit Prince, which I only thought about halfway through the sketch"
// img	"https://imgs.xkcd.com/comics/tree_cropped_(1).jpg"
// title	"Petit Trees (sketch)"
// day	"1"

const baseUrl string = "https://www.xkcd.com/"

func getData(index int) (*Comic, error) {
	resp, err := http.Get(baseUrl + strconv.Itoa(index) + "/info.0.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var data Comic
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}

func checkErr(e error) {
	if e != nil {
		fmt.Printf("error: %v", e)
	}
}

func main() {
	var content []Comic
	startIdx := 1
	for {
		item, err := getData(startIdx)
		if err != nil {
			fmt.Printf("error: %v", err)
			break
		}
		fmt.Println(*item)
		content = append(content, *item)
		startIdx++
		if startIdx == 100 {
			break
		}
	}
	file, err := os.OpenFile("data.json", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		fmt.Printf("error while marshalling : %v", err)
		return
	}
	if _, err := file.Write([]byte(jsonData)); err != nil {
		fmt.Printf("error while writing file: %v", err)
	}
}
