Keywords For Site

```golang
package main

import (
	"context"
	"log"
	"github.com/mainanick/dataforseo"
)

client := dataforseo.NewClient(nil).WithAuthToken("username", "password")

keywords, err := client.Keyword.GoogleSiteKeywords(context.TODO(), dataforseo.SiteKeywordRequest{
		Target:       "safaribooking.com",
		LocationName: "United States",
})
if err != nil {
    log.Fatalln("Error: ", err.Error())
}

log.Println("Keyword Response Returned", len(keywords.Tasks))
for _, task := range keywords.Tasks {
    for _, r := range task.Result {
        log.Println("Keyword: ", r.Keyword)
    }
}
```