package main

import (
	"net/http"
	"testing"
)

func BenchmarkIsJSON(b *testing.B) {
	postString := `{"userId":1,"id":4,"title":"eum et est occaecati","body":"ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit"}`

	for n := 0; n < b.N; n++ {
		IsJSON(postString)
	}
}

func TestIsJSONFunc(t *testing.T) {
	validJSONString := `{ "userId" : 1, "id" : 1, "title" : "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto" }`

	result := IsJSON(validJSONString)

	if result != true {
		t.Errorf("Should return true as string valid json")
	}

	invalidJSONString := `some string`

	invalidJSONResult := IsJSON(invalidJSONString)

	if invalidJSONResult != false {
		t.Errorf("Should return false as string not valid json")
	}
}

func TestPostHandlerBadInput(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, _ := ts.get(t, "/post")

	if code != http.StatusBadRequest {
		t.Errorf("error")
	}
}

func TestPostHandlerCachePost(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, headers, body := ts.get(t, "/post?postId=1")

	if code != http.StatusOK {
		t.Errorf("error because status is not code 200")
	}

	cached := headers.Get("cached")

	if cached != "true" {
		t.Errorf("error because cached header is not true")
	}

	correctBody := `{"userId":1,"id":1,"title":"sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto"}`

	if string(body) != correctBody {
		t.Errorf("Body in response is not correct")
	}
}

func TestPostHandlerExternalPost(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, headers, body := ts.get(t, "/post?postId=4")

	if code != http.StatusOK {
		t.Errorf("error because status is not code 200")
	}

	cached := headers.Get("cached")

	if cached != "false" {
		t.Errorf("error because cached header is not false")
	}

	correctBody := `{"userId":1,"id":4,"title":"eum et est occaecati","body":"ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit"}`

	if string(body) != correctBody {
		t.Errorf("Body in response is not correct")
	}
}

func TestMaterials(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/materials")

	if code != http.StatusOK {
		t.Errorf("error status code should be 200")
	}

	correctBody := `[{"Name":"This is a mocked Material"}]`

	if string(body) != correctBody {
		t.Errorf("body is not correct")
	}
}
