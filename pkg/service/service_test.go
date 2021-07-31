package service

import (
	"bufio"
	"io"
	"os"
	"path"
	"testing"
)

func getFile(filePath string) (*os.File, io.Reader, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	r := bufio.NewReader(fi)
	return fi, r, nil
}

func TestSaveFile(t *testing.T) {
	rootPath := os.Getenv("ROOT_FOLDER")

	srv := NewService(rootPath)

	filePath := path.Join(rootPath, "test_file.json")

	f, r, err := getFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err := srv.SaveFile("0.0.1", "public", "index.html", r); err != nil {
		t.Error(err)
	}
	f.Close()
	if _, err := os.Stat(path.Join(rootPath, "0.0.1", "public", "index.html")); os.IsNotExist(err) {
		t.Error("file 0.0.1 not exists")
	}

	f, r, err = getFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if err := srv.SaveFile("0.0.2", "", "index.html", r); err != nil {
		t.Error(err)
	}
	f.Close()
	if _, err := os.Stat(path.Join(rootPath, "0.0.2", "index.html")); os.IsNotExist(err) {
		t.Error("file 0.0.2 not exists")
	}

	if err := srv.SetLive("0.0.1"); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(path.Join(rootPath, "live", "public", "index.html")); os.IsNotExist(err) {
		t.Error("file live 0.0.1 not exists")
	}
	if err := srv.SetLive("0.0.2"); err != nil {
		t.Error(err)
	}
	if _, err := os.Stat(path.Join(rootPath, "live", "public", "index.html")); !os.IsNotExist(err) {
		t.Error("file live 0.0.1 still exists")
	}
	if _, err := os.Stat(path.Join(rootPath, "live", "index.html")); os.IsNotExist(err) {
		t.Error("file live 0.0.2 not exists")
	}

	os.RemoveAll(path.Join(rootPath, "0.0.1"))
	os.RemoveAll(path.Join(rootPath, "0.0.2"))
	os.RemoveAll(path.Join(rootPath, "live"))
}
