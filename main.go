package main

import (
	"flag"
	"google_drive_downloader/downloader"
	"log"
)

func Usage() {
	log.Fatalln("-id id -dst dstPath")
}

func main() {

	idPtr := flag.String("id", "", "google drive id")
	dstPtr := flag.String("dst", "", "path to download data")
	flag.Parse()

	if *idPtr == "" || *dstPtr == "" {
		Usage()
	}

	downloader := downloader.NewDownloader(*idPtr)
	downloader.Start(*dstPtr)
}
