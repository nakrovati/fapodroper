package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/fapodroper/internal/logger"
	"github.com/gocolly/colly/v2"
)

func DownloadImages(username string, start, end int) {
	if start > end {
		logger.ErrorLog.Fatal("end should be greater than or equal to start")
	}

	baseURL := "https://fapodrop.com/"

	if !userExists(baseURL + username) {
		logger.ErrorLog.Fatalf("user %s not found\n", username)
	}

	downloadDirectory := filepath.Join("images", username)

	c := colly.NewCollector()

	c.OnHTML("div.col-12.px-0 img.media-img.mx-auto.d-block", func(e *colly.HTMLElement) {
		imageSrc := e.Attr("src")
		downloadImage(baseURL, imageSrc, downloadDirectory)
	})

	for i := start; i < end; i++ {
		imageID := fmt.Sprintf("%04d", i)
		imageURL, err := url.JoinPath(baseURL, username, "media", imageID)
		if err != nil {
			logger.ErrorLog.Fatal(err)
		}

		err = c.Visit(imageURL)
		if err != nil {
			logger.ErrorLog.Fatalf("ERROR: image page not found: %v", err)
		}

		if i%10 == 0 {
			time.Sleep(2 * time.Second)
		}
	}
}

func downloadImage(baseURL, imageSrc, downloadDirectory string) {
	fullURL := baseURL + imageSrc
	fileName := filepath.Base(imageSrc)
	filePath := filepath.Join(downloadDirectory, fileName)

	response, err := http.Get(fullURL)
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}
	defer response.Body.Close()

	err = os.MkdirAll(downloadDirectory, os.ModePerm)
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	logger.InfoLog.Printf("The image has been downloaded and saved to %s\n", filePath)
}

func userExists(url string) bool {
	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		logger.InfoLog.Println(err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.InfoLog.Println(err)
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
