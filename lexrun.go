package main

import (
	"fmt"
	"io"
	"time"
	
	 u "github.com/ardeshir/version"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}))

	svc := lexruntimeservice.New(sess)

	pr, pw := io.Pipe()

	go func() {
		for i := 0; i < 5; i++ {
			pw.Write([]byte("hello there "))
			time.Sleep(10 * time.Millisecond)
		}
		pw.Close()
	}()

	req, resp := svc.PostContentRequest(&lexruntimeservice.PostContentInput{
		BotAlias:    aws.String("alias"),
		BotName:     aws.String("TestBotName"),
		ContentType: aws.String("text/plain; charset=utf-8"),
		UserId:      aws.String("user"),
		InputStream: aws.ReadSeekCloser(pr),
	})

	err := req.Send()

	fmt.Println(resp, err)


  if debugTrue() {
    u.V(version)
  }

}

// Function to check env variable DEFAULT_DEBUG bool
func debugTrue() bool {
    
     if os.Getenv("DEFAULT_DEBUG") != "" {
        return true
     }  
     return false 
}