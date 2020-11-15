package downloader

import (
	"io"
	"log"
	"net/http"
	"os"
)

type Downloader struct {
	id string
}

func NewDownloader(id string) *Downloader {
	downloader := Downloader{id: id}
	return &downloader
}
func (d *Downloader) getResponse() (*http.Response, error) {

	url := "https://drive.google.com/uc?export=download&confirm=CONFIRM&id=" + d.id

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil

}

func (d *Downloader) save(res *http.Response, dstPath string) error {

	output, err := os.Create(dstPath)

	if err != nil {
		return err
	}

	defer output.Close()
	n, err := io.Copy(output, res.Body)

	if err != nil {
		return err
	}

	log.Printf("Bytes : %d\n", n)
	return nil

}

func (d *Downloader) Start(dstPath string) {

	res, err := d.getResponse()

	if err != nil {
		log.Fatalf("Error getting response : %s\n", err.Error())
	}

	defer res.Body.Close()
	d.save(res, dstPath)

}
