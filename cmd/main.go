package main

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func main() {

}

func TestExample(t *testing.T) {

	cute.NewTestBuilder().
		Title("Title").
		Description("some_description").
		Create().
		RequestBuilder(
			cute.WithURI("https://jsonplaceholder.typicode.com/posts/1/comments"),
			cute.WithMethod(http.MethodGet),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusOK).
		AssertBody(
			json.Equal("$[0].email", "Eliseo@gardner.biz"),
			json.Present("$[1].name"),
		).
		ExecuteTest(context.Background(), t)
}
