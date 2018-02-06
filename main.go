package main

import(

  "fmt"
  "log"
  "os"

  u "github.com/ardeshir/version"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/aws/session"
)


var ( 
 debug bool = false
 version string = "0.0.0"
 )


func main() {

sess, err := session.NewSession( &aws.Config{
	Region: aws.String("us-east-1"),
})

if err != nil  {
	log.Fatal(err)
}

s3Svc := s3.New(sess)
results, err := s3Svc.ListBuckets(nil)
if err != nil {
	log.Fatal("Unable to get bucket list")
}

fmt.Println("Buckets: ")
for _, b := range results.Buckets {
	fmt.Printf("Bucket: %s \n", aws.StringValue(b.Name))
}

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
