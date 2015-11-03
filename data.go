package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"time"
)

func createData(now time.Time, path string) (dataTarGz []byte, err error) {
	buf := &bytes.Buffer{}
	compress := gzip.NewWriter(buf)
	tarball := tar.NewWriter(compress)

	if err := allToTar(path, tarball); err != nil {
		return nil, fmt.Errorf("adding files to data.tar.gz: %v", err)
	}

	if err := tarball.Close(); err != nil {
		return nil, fmt.Errorf("closing data.tar.gz: %v", err)
	}
	if err := compress.Close(); err != nil {
		return nil, fmt.Errorf("closing data.tar.gz: %v", err)
	}
	return buf.Bytes(), nil
}
