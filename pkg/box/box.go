// @File:     box.go
// @Created:  2020-03-21 03:19:10
// @Modified: 2020-03-23 23:24:38
// @Author:   Antonio Escalera
// @Commiter: Antonio Escalera
// @Mail:     aj@angelofdeauth.host
// @Copy:     Copyright © 2020 Antonio Escalera <aj@angelofdeauth.host>

//go:generate go run generator.go

// Package box provides the EmbedBox type and interfaces for storing and interacting with embedded files.
package box

import (
	"fmt"

	"github.com/klauspost/compress/zstd"
)

// EmbedDecoder is a reader that caches decompressors.
var EmbedDecoder, _ = zstd.NewReader(nil)

// EmbedBox is the container type for embedded paths.
// The files within are stored as CompressedFiles.
type EmbedBox struct {
	Path    string
	Storage []CompressedFile
}

// CompressedFile is the container type for embedded files.
// The file contents are stored as zstandard compressed byte arrays.
type CompressedFile struct {
	Name              string
	CompressedBytes   []byte
	PreCompressedSize int
}

// decompress a file's CompressedBytes buffer
func (c *CompressedFile) decompress() ([]byte, error) {
	return EmbedDecoder.DecodeAll(c.CompressedBytes, make([]byte, 0, c.PreCompressedSize))
}

// Create new box for embed files
func newEmbedBox() *EmbedBox {
	return &EmbedBox{Path: "", Storage: make([]CompressedFile, 0)}
}

// Add a file to box
func (e *EmbedBox) Add(name string, content []byte) {
	e.Storage = append(e.Storage, CompressedFile{Name: name, CompressedBytes: content, PreCompressedSize: len(content)})
}

// Get file's content
func (e *EmbedBox) Get(file string) ([]byte, error) {

	// verify e contains file, and index file
	cfidx, err := e.Index(file)
	if err != nil {
		return nil, err
	}

	// decompress file content
	c, err := e.Storage[cfidx].decompress()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Has returns true if the file exists in an EmbedBox, and false otherwise.
func (e *EmbedBox) Has(file string) bool {
	for _, v := range e.Storage {
		if v.Name == file {
			return true
		}
	}
	return false
}

// Index returns the array index of the named CompressedFile in the EmbedBox.
func (e *EmbedBox) Index(file string) (int, error) {
	for i, v := range e.Storage {
		if v.Name == file {
			return i, nil
		}
	}
	return 0, fmt.Errorf("File %s not found in EmbedBox", file)
}
