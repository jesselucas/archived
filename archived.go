package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ProgressWriter struct {
	Total uint64
}

func (pw *ProgressWriter) Write(p []byte) (int, error) {
	n := len(p)
	pw.Total += uint64(n)
	//pw.PrintProgress()
	return n, nil
}

func (pw ProgressWriter) String() string {
	//fmt.Printf("\r%s", strings.Repeat(" ", 35))
	return fmt.Sprintf("\rDownload progress: %d bytes", wc.Total)
}

type Archived struct {
	Path     string
	Progress *ProgressWriter
}

func (a *Archived) Download(u, filepath string) error {
	// Use .tmp until file is completely downloaded
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(u)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	fmt.Print("\n")

	// Remove .tmp extension
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}

	return nil
}

func (a *Archived) Save(f *os.File) error {
	return nil
}
