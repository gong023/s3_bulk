## s3_bulk

Tiny package to download files concurrently from s3.

## usage

Export your AWS_ACCESS_KEY_ID and AWS_SECRET_KEY.
```bash
shell > export AWS_ACCESS_KEY_ID="xxx"
shell > export AWS_SECRET_KEY="xxxxxx"
```

Download package
```bash
shell > go get github.com/gong023/s3_bulk
```

Write script like below and execute it by `go run`.
```go
package main

import (
        "log"
        "runtime"

        "github.com/crowdmob/goamz/aws"
        "github.com/crowdmob/goamz/s3"
        "github.com/gong023/s3_bulk"
)

func main() {
        auth, err := aws.EnvAuth()
        if err != nil {
                log.Fatal(err)
        }

        // create client with appropriate region
        client := s3.New(auth, aws.APNortheast)
        bucket := client.Bucket("your bucket")
        resp, err := bucket.List("", "", "", 1)
        if err != nil {
                log.Fatal(err)
        }

        d := s3_bulk.Downloader{
                "/BasePath",      // set basepath to put files
                runtime.NumCPU(), // set concurrent process
                bucket,           // downloader needs bucket as client
                resp.Contents,    // set file list to download
        }
        d.Execute()
}
```

You will get files in `resp.Contents`.
