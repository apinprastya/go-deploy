package client

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/imroc/req"
)

type Client struct {
	version string
	folder  string
	rootUrl string
	secret  string
}

func NewClient(version, folder, rootUrl, secret string) *Client {
	return &Client{version: version, folder: folder, rootUrl: rootUrl, secret: secret}
}

func (c *Client) Run() error {
	fmt.Println("Uploading version: " + c.version)
	fmt.Println("Uploading folder: " + c.folder)
	fmt.Println("Uploading root url: " + c.rootUrl)

	if err := filepath.Walk(c.folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		c.UploadFile(path)
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (c *Client) UploadFile(filePath string) error {
	fmt.Println("Uploading file: " + filePath)
	dir, fn := filepath.Split(filePath)
	dir = strings.Replace(dir, c.folder, "", -1)
	header := req.Header{"secret": c.secret}
	f, _ := os.Open(filePath)
	_, err := req.Post(c.rootUrl+"/upload", header, req.Param{"version": c.version, "path": dir, "filename": fn},
		req.FileUpload{FieldName: "file", FileName: "file", File: f})
	if err != nil {
		return err
	}
	return nil
}
