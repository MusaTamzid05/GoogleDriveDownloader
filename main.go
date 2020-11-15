package main

import "google_drive_downloader/downloader"

func main() {

	downloader := downloader.NewDownloader("1AjXlwT19u1WlxjfEimtIaILIKCrfeIAV")
	downloader.Start("data.txt")
}
