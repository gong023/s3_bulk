package s3_bulk_dl

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"

	"github.com/crowdmob/goamz/s3"
)

var download_count int

type Downloader struct {
	BasePath string
	Procs    int
	Bucket   *s3.Bucket
	Contents []s3.Key
}

func (d *Downloader) execute() {
	runtime.GOMAXPROCS(d.Procs)

	done := make(chan bool)
	limit := len(d.Contents)
	for _, c := range d.Contents {
		data, err := d.Bucket.Get(c.Key)
		if err != nil {
			log.Fatal(err)
		}
		go d.create_file(c.Key, limit, data, done)
	}
	<-done
}

func (d *Downloader) create_file(s3_path string, limit int, data []byte, done chan bool) {
	full_file_path := d.BasePath + s3_path
	full_dir_path := regexp.MustCompile(".+/").FindString(full_file_path)
	if err := os.MkdirAll(full_dir_path, 0766); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(full_file_path, data, 0766); err != nil {
		log.Fatal(err)
	}
	download_count += 1
	if download_count >= limit {
		done <- true
	}
}
