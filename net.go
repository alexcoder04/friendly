package friendly

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func DownloadFile(downloadUrl string, path string) error {
	_, err := url.Parse(downloadUrl)
	if err != nil {
		return err
	}

	res, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)

	return err
}
