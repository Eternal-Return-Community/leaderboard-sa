package services

import (
	"encoding/json"
	"erbs/src/structs"
	"fmt"
	"net/http"
)

const (
	url_path = "https://playeternalreturn.com/api/v1/posts/news?category=patchnote&page=1"
)

func Pathnotes() structs.Notes {

	resp, err := http.Get(url_path)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var path structs.Pathnotes
	err = json.NewDecoder(resp.Body).Decode(&path)
	if err != nil {
		fmt.Println(err)
	}

	return path.Articles[0]

}
