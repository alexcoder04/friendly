package friendly

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Downloads data from a url and returns it as a bytes array.
func GetFromUrlBytes(dlUrl string) ([]byte, error) {
	_, err := url.Parse(dlUrl)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.Get(dlUrl)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

// Downloads data from a url and returns it as a string.
func GetFromUrlString(dlUrl string) (string, error) {
	data, err := GetFromUrlBytes(dlUrl)
	return string(data), err
}

// Downloads file from a url to given location.
func DownloadFile(dlUrl string, path string) error {
	res, err := http.Get(dlUrl)
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
