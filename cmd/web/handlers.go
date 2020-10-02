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
	materialsList, _ := app.materials.FindAll()
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

	cachePosts := map[string] string {
		"1": `{ "userId" : 1, "id" : 1, "title" : "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto" }`,
		"2": "{\"userId\":1,\"id\":2,\"title\":\"qui est esse\",\"body\":\"est rerum tempore vitae nsequi sint nihil reprehenderit dolor beatae ea dolores neque nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis nqui aperiam non debitis possimus qui neque nisi nulla\"}",
		"3": "{\"userId\":1,\"id\":3,\"title\":\"ea molestias quasi exercitationem repellat qui ipsa sit aut\",\"body\":\"et iusto sed quo iure\\nvoluptatem occaecati omnis eligendi aut ad\\nvoluptatem doloribus vel accusantium quis pariatur\\nmolestiae porro eius odio et labore et velit aut\"}",
	}

	cachePost, hasCacheVersion := cachePosts[postId]
	var post models.Post

	if !hasCacheVersion {
		resp, _ := http.Get("https://jsonplaceholder.typicode.com/posts/" + postId)


		body, _ := ioutil.ReadAll(resp.Body)
		app.infoLog.Println("Body from external resource: " + string(body))

		json.Unmarshal(body, &post)

		res.Header().Set("Content-Type", "application/json")
		res.Header().Set("cached", "false")

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
