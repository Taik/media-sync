package vlc

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

// Upload uploads a file to a VLC endpoint.
func Upload(f io.Reader, host string) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files[]", filepath.Base(f.Name()))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, f)
	if err != nil {
		return err
	}

	err = writer.Close()

	host = fmt.Sprintf("%s/upload.json", host)
	resp, err := http.NewRequest("POST", host, body)
	if err != nil {
		return err
	}

}
