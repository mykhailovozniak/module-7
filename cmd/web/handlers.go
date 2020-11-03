package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"module-7/pkg/models"
	"net/http"
)

func IsJSON(s string) bool {
	var js map[string]interface{}

	return json.Unmarshal([]byte(s), &js) == nil

}

func (app *application) materialsHandler(res http.ResponseWriter, req *http.Request) {
	materialsList, _ := app.materials.FindAll(req.Context())
	jsonData, _ := json.Marshal(materialsList)

	res.Header().Set("Content-Type", "application/json")
	app.infoLog.Println("Successfully a list of materialsList fetched from draft")

	res.Write(jsonData)
}

func (app *application) helloHandler(res http.ResponseWriter, req *http.Request) {
	app.infoLog.Println("Hello world route")

	fmt.Fprintf(res, "Hello world")
}

func (app *application) postHandler(res http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["postId"]

	if !ok || len(keys[0]) < 1 {
		app.errorLog.Println("Url param 'postId' is missing")
		http.Error(res, "Bad request", 400)

		return
	}

	postId := keys[0]

	cachePost, hasCacheVersion := app.cache[postId]
	var post models.Post

	if !hasCacheVersion {
		resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + postId)


		body, _ := ioutil.ReadAll(resp.Body)
		app.infoLog.Println("Body from external resource: " + string(body))

		json.Unmarshal(body, &post)

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("cached", "false")

		app.cache[postId] = string(body)
		jsonData, _ := json.Marshal(post)
		res.Write(jsonData)

		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("cached", "true")

	app.infoLog.Println("isJSON", IsJSON(cachePost))

	json.Unmarshal([]byte(cachePost), &post)

	jsonData, _ := json.Marshal(post)
	res.Write(jsonData)
}
