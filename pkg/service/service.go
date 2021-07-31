package service

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"

	"github.com/zloylos/grsync"
)

type Service struct {
	rootFolder string
}

func NewService(rootFolder string) *Service {
	return &Service{rootFolder: rootFolder}
}

func (s *Service) SaveFile(version, folderPath, fileName string, r io.Reader) error {
	fullFolderPath := path.Join(s.rootFolder, version, folderPath)
	fullPath := path.Join(s.rootFolder, version, folderPath, fileName)

	if _, err := os.Stat(fullFolderPath); os.IsNotExist(err) {
		os.MkdirAll(fullFolderPath, os.ModePerm)
	}

	fo, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(fo)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			return err
		}
	}
	if err = w.Flush(); err != nil {
		return err
	}
	return nil
}

func (s *Service) SetLive(version string) error {
	fullFolderSourcePath := path.Join(s.rootFolder, version)
	fullFolderDestinationPath := path.Join(s.rootFolder, "live")

	if _, err := os.Stat(fullFolderSourcePath); os.IsNotExist(err) {
		return errors.New("version not found")
	}

	if _, err := os.Stat(fullFolderDestinationPath); os.IsNotExist(err) {
		os.MkdirAll(fullFolderDestinationPath, os.ModePerm)
	}

	task := grsync.NewTask(
		fullFolderSourcePath+"/",
		fullFolderDestinationPath+"/",
		grsync.RsyncOptions{Recursive: true, Delete: true},
	)

	if err := task.Run(); err != nil {
		return err
	}

	return nil
}
