package archive

import (
	"archive/zip"
	"io"
)

func IsTar(fileName string) bool { _ = "STUB: not implemented"; return false }

func ExtractTar(reader io.Reader, dir string, stripComponents int) error {
	_ = "STUB: not implemented"
	return nil
}

func ExtractTarGZ(reader io.Reader, dir string, stripComponents int) error {
	_ = "STUB: not implemented"
	return nil
}

func IsZip(fileName string) bool { _ = "STUB: not implemented"; return false }

// http://golang.org/pkg/net/http/#DetectContentType

func ExtractZip(reader io.ReaderAt, size int64, dir string, stripComponents int) error {
	_ = "STUB: not implemented"
	return nil
}

func isFatFile(header zip.FileHeader) bool { _ = "STUB: not implemented"; return false }

// This identifies FAT files, based on the `zip` source: https://golang.org/src/archive/zip/struct.go

func stripPath(source string, stripComponents int) string { _ = "STUB: not implemented"; return "" }
