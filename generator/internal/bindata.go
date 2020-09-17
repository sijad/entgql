// Package internal Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// template/base_models_go.tmpl
// template/base_resolvers_go.tmpl
// template/base_sdl.tmpl
// template/node_models_go.tmpl
// template/node_resolvers_go.tmpl
// template/node_sdl.tmpl
package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templateBase_models_goTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x4d\x0a\x02\x31\x10\x44\xe1\xb5\x7d\x8a\x5a\xcc\x42\xc1\xf1\x00\x82\x07\xf0\x18\x8d\xd3\x0c\xc1\xfc\x99\xc4\x55\x5b\x77\x17\xe3\xf6\xf1\xf1\xaa\x3e\x9e\xba\x1b\x52\xd9\x2c\x8a\x84\x54\x4b\x1b\x38\x8a\xfb\x8a\xa6\x79\x37\x2c\x21\xd5\x33\x16\x8d\x41\x3b\xae\x37\x5c\xee\xd3\x74\x52\x0e\xee\xff\x4e\xc2\xfd\x07\xf1\xc1\xeb\x5d\x86\x91\xf3\x60\x79\xc3\x4a\xca\x49\xbe\x01\x00\x00\xff\xff\x19\x7a\x60\x66\x68\x00\x00\x00")

func templateBase_models_goTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateBase_models_goTmpl,
		"template/base_models_go.tmpl",
	)
}

func templateBase_models_goTmpl() (*asset, error) {
	bytes, err := templateBase_models_goTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/base_models_go.tmpl", size: 104, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateBase_resolvers_goTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\xc1\x09\x02\x31\x10\x85\xe1\xb3\x53\xc5\x3b\xec\x41\xc1\xb5\x00\xc1\x02\x2c\x63\xd0\x21\x06\xdd\xcd\x98\xc4\xd3\xf8\x7a\x17\xb3\xd7\x9f\x8f\xdf\xf5\xf6\xd4\x64\x48\x55\xfd\x21\x92\x17\x2f\xb5\x63\x2f\x11\x33\xaa\xae\xc9\x30\xe5\xc5\x8f\x98\xf4\x95\xb5\xe1\x7c\xc1\xe9\x3a\x4c\x23\x65\x17\xb1\x75\x12\x11\x7f\x88\x2f\xde\x9f\xd2\x8d\x1c\x07\x5b\xef\x98\x49\x39\x88\xfc\x02\x00\x00\xff\xff\xd8\x63\xe0\x02\x69\x00\x00\x00")

func templateBase_resolvers_goTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateBase_resolvers_goTmpl,
		"template/base_resolvers_go.tmpl",
	)
}

func templateBase_resolvers_goTmpl() (*asset, error) {
	bytes, err := templateBase_resolvers_goTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/base_resolvers_go.tmpl", size: 105, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateBase_sdlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x59\xd1\x6e\xdb\x46\x10\x7c\xd7\x57\x6c\x04\x14\x70\x80\x46\x6e\x82\xa2\x28\xdc\x87\x42\xb5\xec\x82\x40\xe2\x04\xb5\x8d\x3c\x04\x41\x7d\x3a\x2e\xc9\x6b\x8e\x7b\xf2\xdd\xd2\xa9\x52\x18\xc8\x67\xb4\xbf\x97\x2f\x29\xee\x48\x5a\x94\x42\x59\x14\x25\xb4\xca\x93\x20\xdd\xcc\xcd\xce\x0e\x17\x1b\x7a\x38\x1c\x5e\x65\x08\xd7\x8a\x18\x9c\x14\x5a\x58\xe0\xf9\x0c\xc1\xe2\xcc\xa2\x43\x62\x07\x64\xe8\x59\x62\x85\x64\x65\x48\x68\x28\xc8\xa9\x94\x30\x06\x2a\x72\xb4\x4a\xc2\x9d\xd0\x05\xba\xd1\x70\x38\x1c\x54\x0c\x9e\x6d\x30\x18\x86\x7f\xf5\x77\x63\x9a\x87\xaf\xc6\x30\x33\xfe\x32\x45\xc0\x2a\x47\x10\x0e\x62\x74\xd2\xaa\x29\xc6\x30\x9d\x03\x67\x08\xd1\xe5\x6b\xf8\xf1\x87\xef\x9e\x83\x63\x41\xb1\xb0\xf1\x08\x5e\x89\x39\x18\x0b\xb9\x98\x03\x19\x8f\x96\xba\x88\x11\x44\x20\xf9\x64\x08\x9b\xf7\x5f\xa9\x1c\xab\xcb\x0a\x52\x77\x68\x9d\xd0\x7a\xee\x3f\xdf\x16\x08\x2a\x46\x62\x95\x28\xb4\xe5\xe5\x89\xa2\xf2\xea\x77\xbf\x9d\x9f\xc2\xf7\xcf\x5f\xbc\x78\x7f\x94\x31\xcf\xdc\xc9\xf1\x31\x1b\xa3\xdd\x48\x21\x27\x23\x63\xd3\xe3\x8c\x73\x7d\x6c\x13\xe9\x0f\x3d\x5d\xaa\xf8\x3a\x9a\x54\x37\x6a\x23\x85\xb7\xca\x57\x28\x40\x1a\x22\x0c\xd6\x01\x67\x82\x41\x0a\x82\x29\x42\xe1\x30\x86\xc4\x58\xb0\xe8\x8a\x5c\x51\x0a\x33\x91\x2a\x0a\xc0\x26\xef\x69\x61\x9d\xb1\x25\x33\x81\x99\xfe\x81\x92\xe1\xa3\xe2\x0c\x04\xa4\xda\x4c\x9b\x75\xdd\x44\x93\x9b\x80\x55\xc4\x68\x13\x21\x11\x2e\x4c\x8c\xf0\xd7\x00\x20\x28\x5b\x05\x2c\x8c\x18\xc1\x69\x43\x97\x22\xb8\x13\x56\x99\xc2\xc1\x4c\x0b\x89\x0e\x38\xb3\xa6\x48\x33\x53\x70\xe8\x8f\x9b\x3b\xc6\x1c\xd8\xd4\x14\xbe\x6d\xca\x81\x53\x94\x6a\x2c\x03\x11\x94\x00\xa8\xf8\x04\xa2\xc9\x93\xc1\x7d\xa8\x21\xa2\xc4\xd8\xbc\xb4\x47\x4c\x3d\xdd\xa2\xec\x55\xbf\x02\x41\xc8\xe2\x1b\x91\xa2\x47\xd6\x95\xbc\xcd\x90\x1e\x80\x94\x7a\x1f\x3f\x0a\x1b\xbb\x6f\x41\x58\xf4\x02\x2d\x42\x6e\x2c\x82\x62\xcc\xdd\xcf\xa5\x90\x4c\xb8\x0b\xfc\x93\x3d\xd7\x09\xfc\x62\x8c\x46\x41\x4f\x06\xed\x84\x53\x21\x3f\x74\x62\x7c\x63\xf1\xce\xfb\xb4\x2d\xab\x37\x51\x86\xd6\x7a\x13\xa5\x21\x56\xf4\x60\x99\x63\x61\xb9\x6c\xfc\x09\x5c\xb2\x55\x94\xae\x61\x5c\x14\xfe\x28\x21\x52\xbc\x42\x77\x5f\x65\x35\x51\x9a\x31\x40\xea\xde\x8b\x54\x28\x72\x5c\x1d\x84\x44\xa1\x8e\xdd\x08\xc6\x5a\x57\x9f\x83\x23\xd2\xe4\xd3\xf0\xd4\x54\x51\xd4\x26\x55\x52\x68\xf8\xf2\xf9\x6f\x41\xf1\xe8\xcb\xe7\x7f\xca\x1c\xce\x8a\x9a\xe9\xbc\xbc\xa9\xea\x60\xe4\x80\x0a\xad\xe1\x48\x25\x70\xc3\xb6\xc0\x1b\xf0\xf1\x99\xa1\xf4\x71\x8c\x9f\xfa\x47\x5d\xb9\xf0\xa4\x2f\xce\x25\x42\xbb\xd5\x83\x75\xca\xdc\x45\xa1\xf5\x43\x07\x2a\xbb\xce\x6e\x0b\xa1\x7d\x71\x21\xb3\x35\x66\x29\x9e\xe8\x8f\x5c\x99\x15\x9f\x2f\x0c\x97\xbf\x3c\x0e\x26\xc3\x67\xad\xf8\xa8\x9c\x4e\xe1\x49\x5a\x86\x6b\xe5\xb8\xd6\x4c\x27\xf0\xae\x84\x3d\x79\xdf\xb8\x58\x75\x01\x93\xe1\xa8\x05\xff\x12\x9d\x7f\x58\xc5\x2a\xb2\xa9\x5a\xa3\x73\x57\x99\xa0\x15\xcd\x0b\xa8\xb1\x5d\x8a\xaf\x69\x5e\xdb\x76\x0f\x7e\xb5\x28\x42\xb6\x1e\x57\x93\x96\xc7\x5a\x04\x2d\x11\x74\xd3\xd4\x20\x5b\x27\xeb\xd4\x10\xfb\x84\xaf\xb0\xb8\x32\xee\x47\x52\x38\x7c\xe6\x90\x9c\x62\x75\x87\x0f\xf1\x2a\x5b\xe2\x56\xb8\x26\x06\xcb\x8c\xca\x92\x74\x2b\xce\xd0\xc1\x56\xda\x2e\x12\x15\xad\x15\x19\x2d\x7e\xda\x41\x6f\xcb\x05\x0d\xc5\xeb\xef\xb8\xf4\xc3\xcb\x95\x73\x61\x1b\x3b\xc2\xd0\x73\x6f\x15\x67\xeb\x44\x87\x13\xdb\x13\x93\xe1\xcb\x75\xdc\x1d\xc5\xb6\x78\xb1\x90\xdb\xc1\xee\x6e\xca\xdb\x1d\xbf\xec\x70\xd1\x19\xc5\x3d\x1c\x47\x8a\x1f\xf5\x1b\x29\xee\xe5\xf6\x59\x3b\x6f\x27\x91\x2d\x1e\xd4\x32\x3b\xf8\xdc\x45\x71\xbb\xcb\x67\x1b\x2f\x79\x25\x58\x66\xb8\xfa\x4c\xce\x04\x33\x5a\xfa\xda\x0d\x18\x13\x14\x14\xa3\x75\xd2\xaf\x0d\x47\xbf\x3f\x85\xbc\x62\x10\x34\xaf\x17\x25\x99\x09\xbf\x57\xa3\xfd\x09\x04\xcc\xd0\x4a\xf4\x5b\xb8\x4a\x09\x8e\xbe\x59\x01\xe0\x6d\x81\x24\x11\x4c\x02\x9f\xd0\x9a\xb0\x0b\x7b\xe6\x07\x0a\x57\xcf\x65\xf5\x61\xad\x41\x81\xf1\xf0\x4b\x20\xc3\x2f\xbf\xae\xa2\x4b\x07\x96\xba\xfb\xff\xf6\xa0\x43\x5e\x37\xb7\xe3\x40\xea\xa9\x1a\xd2\x56\xd2\xe6\x4d\x32\x22\xde\xc7\x1a\x19\x11\x1f\xfa\x0e\x19\xf9\xff\xf5\xf6\x5d\x20\x17\xe0\xee\xdb\x63\x44\xdc\x7f\x75\x6c\x82\xb7\xdb\x1b\x17\x52\xf7\xb0\x34\x2e\xc8\xfa\x6c\x8c\x6b\xd0\xbb\xac\x8b\x9e\x72\x73\xaa\xc3\x0b\x93\x3d\xc4\xda\xf3\x1c\x7a\xae\xcb\xd7\x39\x7d\x83\xdd\x40\x77\x4f\xb6\x07\xf5\x8f\xf6\x12\x7a\xbb\x6c\x37\xd4\xee\x21\xdc\x0d\xb6\x3e\xe9\x5e\x07\xdf\x25\xde\x81\x73\x73\xbe\xcf\xb5\x11\x7b\x09\x78\x20\x3a\xf4\x84\x07\x91\xfd\x23\xde\x84\x77\xcf\x78\x40\xf5\x0f\xf9\x32\x7c\xbb\x94\x37\x05\xef\x21\xe6\x4d\xba\x3e\x39\x5f\x8b\xdf\x25\xe8\x25\x69\x87\x49\x7e\x1d\x4d\xf6\x32\xc9\xaf\xa3\xc9\xa1\xe7\xbc\x7c\x4d\xdd\x7b\x92\x2f\xd0\x5b\x4c\xf2\xeb\x68\xb2\xc3\x24\x2f\xd1\x9b\xbb\x78\xa5\x72\xdc\x47\x17\x3d\xcf\xa1\x77\xb1\xfc\xf3\x46\xdf\x2e\x36\xd0\xdd\xbb\xe8\x41\xfd\xbb\xb8\x84\xde\x6e\x52\x35\xd4\xee\x61\x50\x35\xd8\xfa\xcc\xa9\x75\xf0\x5d\xc6\x54\xe0\xdc\x9c\xef\x2a\x2d\xfb\x88\x78\x45\x75\xe8\x29\x5f\x46\xf6\x08\x7a\x4d\xb0\xd9\xdc\x31\xcd\xf7\x61\xec\x98\xe6\xff\xad\xa9\xf7\x83\xc1\xbf\x01\x00\x00\xff\xff\x33\xeb\x39\xae\xc0\x1d\x00\x00")

func templateBase_sdlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateBase_sdlTmpl,
		"template/base_sdl.tmpl",
	)
}

func templateBase_sdlTmpl() (*asset, error) {
	bytes, err := templateBase_sdlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/base_sdl.tmpl", size: 7616, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateNode_models_goTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\xd6\x55\x50\xc9\xcb\x4f\x49\x55\xb0\xb2\x55\xd0\x53\xd0\xad\xad\xe5\x82\x08\x25\xe6\x42\x84\xfc\x40\x0c\x90\x30\x57\x49\x65\x41\xaa\x42\x75\x35\x58\xaa\xb6\x56\xa1\xb8\xa4\xa8\x34\xb9\x44\xa1\x9a\x8b\x53\x2b\x35\xaf\x44\x0f\x2e\xc1\x55\xcb\xc5\x95\x56\x9a\x97\xac\x01\x17\xd1\x54\xf0\x2c\xf6\xcb\x4f\x49\xd5\xd0\x54\xa8\xae\xe5\xe2\x02\x04\x00\x00\xff\xff\xbc\xb7\x12\xae\x73\x00\x00\x00")

func templateNode_models_goTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNode_models_goTmpl,
		"template/node_models_go.tmpl",
	)
}

func templateNode_models_goTmpl() (*asset, error) {
	bytes, err := templateNode_models_goTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/node_models_go.tmpl", size: 115, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateNode_resolvers_goTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\x4f\x6f\x9b\x4c\x10\xc6\xcf\xf0\x29\xe6\xb5\xa2\x08\x2c\x42\xa4\x1c\x23\xe5\xf2\xda\xae\xe4\x43\x4d\x5b\x45\xea\xa1\xaa\x2a\x0c\x83\xbb\x29\xec\xe2\x61\x69\xe2\xae\xf7\xbb\x57\xb3\x60\xfc\x27\x6e\x72\x48\x2e\x66\x98\x67\x66\x77\xf6\xd9\x9f\x31\xe6\x0a\x2e\xa4\xca\x11\x6e\xef\x20\x86\x2b\x6b\xfd\x2e\x95\x56\x5d\x6a\xc1\xc1\x90\x16\xf9\xfd\xa6\x76\xc2\x48\x48\x3d\x1a\x04\x51\x40\x3c\x9f\x3a\x8d\x53\x9e\x31\x07\xa5\x3b\x65\x0b\x2b\xc5\x41\xdf\x83\x32\x77\xc5\x7e\xd1\xca\x0c\x02\x82\xf1\xba\x45\xda\x7c\xc1\x46\x95\xbf\x91\x42\x30\xc6\x8d\x61\x6d\x90\xe9\x27\xc8\x94\xd4\xf8\xa4\xe3\x49\xf7\x8c\x40\xe4\x30\x6c\x63\x6d\x08\xc1\xb8\x52\x39\x96\xf1\xd0\x16\x01\x12\x29\x5e\xc8\xf7\xd6\x3c\x08\xc5\x33\xa9\x27\xa5\x40\xa9\xf7\x55\xf1\x67\xde\x35\x08\x7d\x6f\x1d\x7f\xfd\x89\x84\x01\x4b\x2a\xc7\xf8\x53\x9a\xfd\x4a\x57\x5c\x32\x9f\x06\x22\x0f\x43\xdf\xf7\x08\x9b\x08\x7e\xf0\x62\xeb\x38\x91\xe5\x86\x47\x0b\x7d\x4f\x14\x40\xd8\xc0\xdd\x1d\x48\x51\xf2\x7e\x1e\xa1\x6e\x49\xf2\x6b\xc4\x3f\xbe\x67\x7d\xaf\xe2\xbe\xcb\x93\x29\xcd\xfe\x9c\xb7\x6e\x11\xeb\xef\x9a\xab\xae\xf5\x75\x8b\x60\x0b\x75\xd9\x52\x5a\x8a\x3f\xff\xb4\x2b\x2d\x34\x12\x8c\x1b\x4d\x42\xae\x22\x58\x62\xa1\x08\xf7\xef\x85\xa0\x46\xc3\x58\x48\x1d\x41\x99\x0e\xe1\x23\x3b\x02\xa7\xce\x3a\x9f\xe6\xb2\x6e\x75\x04\x8a\x72\xa4\xff\x37\xf0\xed\xfb\x49\x51\xd2\x09\x67\x2e\x66\xa2\xa4\xc4\x4c\x0b\x25\x0f\xaf\xe8\xfa\x1a\xee\x93\x69\xe2\x7b\x75\x2a\x45\x16\x14\x95\x8e\x67\x2c\x16\xc1\x48\x2a\x0d\xa2\xaa\x4b\xac\x50\x6a\xcc\x47\x61\xc8\xa6\x18\x43\xa9\x5c\x21\x5c\x38\xca\xba\x3b\x9b\xe5\x2b\x6c\x76\xb4\xe2\xa2\xa7\xf8\x02\x63\xa6\xc4\xc1\xbc\x27\x56\x11\x0b\x1f\x6f\x12\x7e\x24\x37\x89\xa3\x71\x70\x7a\x6f\x6d\xa9\x1e\x91\x26\x69\x83\x1f\xd8\x23\x6b\x8f\xec\xef\x16\x85\x2d\xb4\x75\x7d\x54\x75\xfe\x16\xd4\xf2\xe1\x99\x9b\x47\x0e\xe1\xe2\x19\xbb\x8e\x39\x24\xe2\x93\xa8\xe5\x43\x07\xec\x8b\x5b\x87\x27\x6c\x72\xf3\x7f\x2f\xb3\xd9\xe7\x2e\x4f\x07\x71\x78\xf6\x71\xcf\xe7\x8e\x4a\xf7\x27\x2e\x1b\x7c\x9b\x6f\x59\x5a\x21\x73\xfb\x26\x0b\xdf\x0b\xef\xfe\xa0\xaf\xf0\xdd\x57\x9d\x03\xbc\x97\xde\x87\x70\x63\xdc\x47\xd2\x1e\x46\x7f\x03\x00\x00\xff\xff\xdb\xed\xec\x97\xb1\x05\x00\x00")

func templateNode_resolvers_goTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNode_resolvers_goTmpl,
		"template/node_resolvers_go.tmpl",
	)
}

func templateNode_resolvers_goTmpl() (*asset, error) {
	bytes, err := templateNode_resolvers_goTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/node_resolvers_go.tmpl", size: 1457, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templateNode_sdlTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\xc1\x6e\xe3\x36\x10\x3d\xd7\x5f\x31\x16\x7c\x48\x16\x59\x2f\xb0\x47\x03\x3d\x64\xed\x04\x30\xd0\xd8\xdb\x24\x8b\x1e\x8a\xa0\x66\xc4\x91\x4d\x94\x26\x5d\x92\x4a\xe2\x2a\xfa\xf7\x62\x28\x89\x94\x15\x25\x31\xda\xdc\x7a\x32\xcd\xa1\x66\xde\x7b\x7c\xe4\xb0\x28\x3e\xc3\x48\x69\x8e\x30\xf9\x19\xc6\xf0\xb9\x2c\x07\xd5\x14\xdb\x56\x53\x0b\x1a\x84\x69\xc1\x6f\xf7\x3b\x1f\x48\xe6\xca\x25\x21\x20\x32\x18\xcf\x67\x3e\x46\x53\x3f\x15\x45\x6b\x69\x13\x79\x06\x9b\x32\xc9\x0c\xa5\xac\xbf\x43\xc5\xfd\x07\x83\xa2\xf8\xf2\x09\x2e\x54\xbe\xb5\xf0\xe9\x8b\x9f\x42\x95\x6f\xa1\x28\x3c\x92\xb2\x5c\x1a\x8e\xe6\xdb\x1e\x8a\x01\xc0\xe2\xfc\xf6\xc7\xf5\xf9\x2f\x03\x80\xf9\xec\x8f\xf3\x9b\x69\x35\x98\x5d\xdc\x4c\x7d\x4e\xc3\xd4\x1a\x61\x24\xa8\xb4\xa7\x36\x9e\x2b\x8e\x4f\x68\xcb\x72\x00\x94\x51\x54\xa4\x9e\x21\xdf\xed\xd0\x4c\x99\xc5\xb2\xac\xf3\xbc\x16\x0d\xc9\x51\xf1\xb2\x1c\x34\x80\x89\x56\x00\xec\x88\x63\x00\x0c\x62\xbb\x93\xb8\x45\xe5\x2c\x2c\x48\x5f\x42\x2e\xf8\x04\xe6\xb3\x61\x1b\x66\x16\x61\x5e\x0a\x94\xdc\x46\x45\x95\x76\x30\xca\xc6\x37\xa8\xac\x70\xe2\x01\xa1\x21\x90\x35\x10\x53\xb6\x45\x29\xfe\xa6\xa1\xd4\x8f\x15\xda\x4b\x61\xac\x2b\xcb\x49\xb5\xb0\x47\xf8\xa2\x88\xb9\x17\x42\x4a\x76\x2f\xb1\x2c\x87\x45\x51\x71\x8b\x2c\x0f\x47\x35\x5e\x8c\x78\x2f\xf8\x1a\x1b\xb8\x23\x5c\xd4\x8e\x19\xa1\x2f\x3a\x6e\xed\xb2\xc8\x40\x1b\x0a\x5c\x7d\x5d\xd2\xcf\xf2\xeb\x32\x70\xc1\x63\xb9\x60\x17\x3d\x8e\x97\x3b\x27\xb4\x62\xb2\x8b\x5e\xda\x5a\xab\x24\x49\xae\x91\x71\x0b\x4c\x71\x40\x45\x4c\x2d\xec\xd8\x5a\x28\x46\x5f\x82\xdb\x18\x9d\xaf\x37\xc0\xc0\xa2\x03\x9d\xc1\x2a\x16\x5a\x8d\x93\x24\x39\x1e\xe3\x09\xcb\x1c\x9a\x09\xdc\x38\x23\xd4\xfa\x0c\xee\x31\xd3\x06\xe3\xff\x8c\x96\x4d\x60\xae\xdc\x19\x48\x16\x86\x8f\x1b\xa4\x55\xb1\xec\x6f\x34\x31\x57\xbb\xdc\x9d\x81\xae\x4c\x3f\x81\xdf\x63\xbc\x3e\x08\xc3\xbb\xd3\xf6\x57\x53\xad\x14\xa6\xc4\xa9\x77\xff\xca\xc1\x20\x49\x92\x73\x4f\xaf\xb2\xe7\x0a\x90\xaf\x11\x04\x69\x80\x90\x86\xcf\x3d\xe9\x43\x27\xd3\x36\x7b\xf7\xfa\x14\x69\x6e\xac\x36\x90\x69\x03\xb9\xf5\x09\xa2\x9e\xb5\x62\xd5\x92\x86\xfa\xb0\xfa\xf2\x76\x83\x07\xe5\x99\xf3\x95\xe9\x02\xd0\x59\x35\xe4\x6b\xac\x33\x90\xbd\x26\x11\xc1\x30\x10\x88\x40\xc1\x69\x60\x20\x85\x0d\xfb\xd6\x64\x7e\x60\x32\x47\x7b\x48\x84\xec\xda\x55\x2a\x72\x6a\xb2\x10\x02\x0b\x8f\x1b\x91\x6e\xa8\x92\x63\x42\x59\x0f\xed\x00\xb8\xe2\x8d\x08\x04\x41\xf0\x5e\x0d\x7c\xaa\x6a\xe3\xa2\x8a\xc3\xbb\x5a\x8c\xb9\xca\xb4\xd9\xb2\xc0\xe3\x95\x24\x3b\xb6\x46\x5a\x3a\x81\xef\xf5\xa8\x25\x66\xaa\x73\xe5\x51\x33\x29\x0f\x00\xee\x75\x4e\x41\xc9\x61\x8d\x0e\x32\xa3\xb7\x7d\x7b\x0c\xe0\xb4\x63\x72\x4a\x59\xbc\x17\xbd\xc8\x5e\xb0\xa9\x41\xe6\x30\x64\xfc\xce\xf6\x52\x33\xde\xc8\x45\xb5\xf1\x89\xa5\x0e\x2c\x1d\x8a\x55\x2a\x05\x2a\x77\x95\x3b\x8f\x7c\xce\x57\xe0\x36\xcc\xc1\x23\xb3\xb0\x33\xfa\x41\x70\xe4\x8d\xcd\xb6\xf5\x22\x10\x95\xbf\x73\x95\x6e\xe8\x56\xe1\x5e\xd4\x5c\xe5\x16\xf9\x18\xae\xd8\x1e\xee\x91\xdc\xc5\xe1\x7e\x0f\x0c\xaa\x0a\x24\x94\x33\x2c\xfd\x33\xa4\xb1\x8d\xdf\x3a\x08\x1a\xe7\xf5\x1b\x2f\xa0\x4b\x3d\x4d\x5f\xc3\x6d\x84\x0d\x69\xe3\xb9\x57\xd5\xa9\xef\xbd\x8e\x5a\xce\xf4\xa2\xcd\x50\xe2\xff\x40\x34\xee\x69\xfe\x57\xd1\xa0\xc9\x13\xe6\xa8\x45\xce\xeb\xee\xd8\x74\x57\x7f\x0b\xc6\xfe\xea\xf9\xc7\x2c\xf1\xa2\x0c\xad\x35\xbc\x3a\xca\xf2\x52\x48\x87\xe6\xc8\x46\x4b\x42\x8e\xb2\xf1\x0f\x25\xfe\xca\x11\x4e\xba\x7d\xf7\xf4\x63\xfa\x6e\x0b\x52\xf7\x7e\x06\x38\x5f\xcc\xda\xb7\x45\x24\x37\xbc\x1b\x00\x2c\xaf\xdf\x08\x2e\x96\xb7\xaf\x46\xcb\x41\x57\xb6\x46\xb1\x77\x94\xf9\xe0\x77\xc6\x6b\x9d\xba\xe9\x4e\x15\xc8\xce\xc5\x13\x37\xf7\x0d\xb3\x1e\xe5\xb8\x4a\x8e\x58\xa8\x73\x58\x8f\x2a\xd4\x31\x58\xb4\xe9\xaf\x39\x9a\x7d\x78\xb5\x3e\x39\x6a\x69\xfe\x4a\xa8\x02\xc5\x9b\x20\x4f\xba\x69\x4f\x5b\xb0\x3f\xd8\xbe\xea\xdd\xcd\xfc\xb6\x6f\xef\x7a\x78\x00\x37\x58\xdb\xc1\xe3\x7d\xf0\x92\xd2\x4b\xfb\xff\xab\x87\x9a\xea\xbe\xd3\xde\x23\x08\xcf\xb0\x93\xb9\x61\x14\xfc\x80\x17\x5b\xf7\xbc\x75\x1e\x6c\xaa\xef\xbd\xa6\xba\xcf\xb5\xc6\x45\x8d\xe7\xfa\x8c\x14\x62\xde\xa2\x87\x87\xe4\xc4\x3b\x7a\xd2\x7b\x76\x48\xf9\xfe\x66\x1e\x6e\xe0\x17\x79\xfa\x8e\x06\xe5\xe9\xef\x6f\x04\xff\x9f\x00\x00\x00\xff\xff\x0a\x30\x32\xdb\x43\x0e\x00\x00")

func templateNode_sdlTmplBytes() ([]byte, error) {
	return bindataRead(
		_templateNode_sdlTmpl,
		"template/node_sdl.tmpl",
	)
}

func templateNode_sdlTmpl() (*asset, error) {
	bytes, err := templateNode_sdlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "template/node_sdl.tmpl", size: 3651, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"template/base_models_go.tmpl":    templateBase_models_goTmpl,
	"template/base_resolvers_go.tmpl": templateBase_resolvers_goTmpl,
	"template/base_sdl.tmpl":          templateBase_sdlTmpl,
	"template/node_models_go.tmpl":    templateNode_models_goTmpl,
	"template/node_resolvers_go.tmpl": templateNode_resolvers_goTmpl,
	"template/node_sdl.tmpl":          templateNode_sdlTmpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"template": &bintree{nil, map[string]*bintree{
		"base_models_go.tmpl":    &bintree{templateBase_models_goTmpl, map[string]*bintree{}},
		"base_resolvers_go.tmpl": &bintree{templateBase_resolvers_goTmpl, map[string]*bintree{}},
		"base_sdl.tmpl":          &bintree{templateBase_sdlTmpl, map[string]*bintree{}},
		"node_models_go.tmpl":    &bintree{templateNode_models_goTmpl, map[string]*bintree{}},
		"node_resolvers_go.tmpl": &bintree{templateNode_resolvers_goTmpl, map[string]*bintree{}},
		"node_sdl.tmpl":          &bintree{templateNode_sdlTmpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
