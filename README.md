Keywords For Site

```golang
package main

import (
	"context"
	"log"

	"github.com/mainanick/dataforseo"
)

func main() {
	dataforseo.DefaultBaseURL = "https://sandbox.dataforseo.com/v3/"
	client := dataforseo.NewClient(nil).WithAuthToken("username", "password")

	keywords, err := client.Keyword.GoogleSiteKeywords(context.TODO(), dataforseo.SiteKeywordRequest{
		Target:       "github.com",
		LocationName: "United States",
	})
	if err != nil {
		log.Fatalln("Error: ", err.Error())
	}

	for _, task := range keywords.Tasks {
		log.Println("Results", len(task.Result))
		for _, r := range task.Result {
			log.Println("Keyword: ", r.Keyword)
		}
	}

}

```