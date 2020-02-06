package upload

import (
	"bytes"
	"fmt"
	"github.com/pniedzwiedzinski/photoprism-cli/internal/api"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

// Command
func Command(c *cli.Context) error {
	if c.NArg() == 0 {
		return fmt.Errorf("You need to pass files: photoprism-cli upload [FILE]")
	}
	uploadTimestamp := fmt.Sprintf("%v", time.Now().Unix())

	a := api.NewAPI("192.168.1.101")

	fileDir, _ := os.Getwd()
	fileName := c.Args().Get(0)
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("files", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	r, _ := http.NewRequest("POST", "http://192.168.1.101:2342/api/v1/upload/"+uploadTimestamp, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Session-Token", a.GetToken())
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	re, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(re))

	return nil
}
