package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/blakesmith/ar"
)

func createDumbDeb(debPath, rootPath, controlPath string) error {
	now := time.Now()
	deb, err := os.Create(debPath)
	if err != nil {
		return err
	}
	defer deb.Close()

	dataTarGz, err := createData(now, rootPath)
	if err != nil {
		return err
	}
	rootFiles := md5data(rootPath)
	md5sums := []byte(strings.Join(rootFiles, "\n"))
	controlTarGz, err := createControl(now, controlPath, md5sums)
	if err != nil {
		return err
	}

	w := ar.NewWriter(deb)
	if err := w.WriteGlobalHeader(); err != nil {
		return fmt.Errorf("cannot write ar header to deb file: %v", err)
	}

	if err := addArFile(now, w, "debian-binary", []byte("2.0\n")); err != nil {
		return fmt.Errorf("cannot pack debian-binary: %v", err)
	}
	if err := addArFile(now, w, "control.tar.gz", controlTarGz); err != nil {
		return fmt.Errorf("cannot add control.tar.gz to deb: %v", err)
	}
	if err := addArFile(now, w, "data.tar.gz", dataTarGz); err != nil {
		return fmt.Errorf("cannot add data.tar.gz to deb: %v", err)
	}
	return nil
}
