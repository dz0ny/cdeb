package main

import (
	"archive/tar"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/blakesmith/ar"
)

func md5file(filePath string) (string, error) {
	var result []byte
	in, err := os.Open(filePath)
	if err != nil {
		return string(result), err
	}
	defer in.Close()
	digest := md5.New()
	io.Copy(digest, in)
	return fmt.Sprintf("%x  %s", digest.Sum(result), filePath), nil
}

func allToTar(location string, tarball *tar.Writer) error {

	return filepath.Walk(location, func(filePath string, fileInfo os.FileInfo, _ error) error {

		if fileInfo.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(location, filePath)
		if err != nil {
			return err
		}
		hdr := tar.Header{
			Name:     relPath,
			Size:     fileInfo.Size(),
			Mode:     int64(fileInfo.Mode()),
			ModTime:  fileInfo.ModTime(),
			Typeflag: tar.TypeReg,
		}

		if err = tarball.WriteHeader(&hdr); err != nil {
			return fmt.Errorf("cannot write header of file to tarball: %v", err)
		}

		in, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("cannot open file: %v", err)
		}

		if _, err = io.Copy(tarball, in); err != nil {
			return fmt.Errorf("cannot write data of file to tarball: %v", err)
		}

		return err
	})

}

func md5data(location string) []string {
	var filesmd5 []string

	filepath.Walk(location, func(path string, fileInfo os.FileInfo, _ error) error {
		if fileInfo.IsDir() {
			return nil
		}
		relPath, err := filepath.Rel(location, path)
		if err != nil {
			return err
		}
		data, err := md5file(relPath)
		filesmd5 = append(filesmd5, data)
		return err
	})
	return filesmd5
}

func addArFile(now time.Time, w *ar.Writer, name string, body []byte) error {
	hdr := ar.Header{
		Name:    name,
		Size:    int64(len(body)),
		Mode:    0644,
		ModTime: now,
	}
	if err := w.WriteHeader(&hdr); err != nil {
		return fmt.Errorf("cannot write file header: %v", err)
	}
	_, err := w.Write(body)
	return err
}
