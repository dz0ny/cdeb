package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"time"
)

func createControl(now time.Time, path string, datamd5sums []byte) (controlTarGz []byte, err error) {
	buf := &bytes.Buffer{}
	compress := gzip.NewWriter(buf)
	tarball := tar.NewWriter(compress)

	hdr := tar.Header{
		Name:     "md5sums",
		Size:     int64(len(datamd5sums)),
		Mode:     0644,
		ModTime:  now,
		Typeflag: tar.TypeReg,
	}
	if err := tarball.WriteHeader(&hdr); err != nil {
		return nil, fmt.Errorf("cannot write header of md5sums file to control.tar.gz: %v", err)
	}
	if _, err := tarball.Write(datamd5sums); err != nil {
		return nil, fmt.Errorf("cannot write md5sums file to control.tar.gz: %v", err)
	}

	if err := allToTar(path, tarball); err != nil {
		return nil, fmt.Errorf("adding files to control.tar.gz: %v", err)
	}

	if err := tarball.Close(); err != nil {
		return nil, fmt.Errorf("closing control.tar.gz: %v", err)
	}
	if err := compress.Close(); err != nil {
		return nil, fmt.Errorf("closing control.tar.gz: %v", err)
	}
	return buf.Bytes(), nil
}
