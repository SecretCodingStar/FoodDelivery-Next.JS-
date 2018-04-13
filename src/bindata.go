// Code generated by go-bindata.
// sources:
// data/fruits.json
// data/herbs.json
// data/ingredient_corpus.txt
// data/ingredient_densities.json
// data/top_5k.txt
// data/vegetables.json
// DO NOT EDIT!

package meanrecipe

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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataFruitsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x53\x4b\x72\x24\x3b\x08\xdc\xfb\x14\x15\x5e\xbf\x1b\xbd\x98\x05\x2d\x61\x15\x5d\x48\xc8\x48\xb8\xa3\x6e\x3f\x61\x37\x2a\xd3\xb3\xcb\x24\x50\xf2\x4b\xfd\xff\xb6\x6d\xef\xd0\x3b\xe3\xfb\x7f\x4f\xa8\x94\x64\x3a\xf9\x92\x04\x59\x9e\xe4\x06\x0d\x1a\x38\x46\xe6\xad\x63\xef\xa8\x1e\x20\xbe\xa1\xea\xe9\x8c\x21\x1d\xff\xf2\x64\xaa\xd0\xe6\x8a\x88\xe4\x4d\x14\x5a\xc1\x15\x31\x8c\x4f\xe4\x1c\xd8\x62\x40\x11\xf2\x87\x1a\xb9\x42\x82\x06\x7a\x6e\x15\x59\xda\x15\x99\xc0\x62\xdd\x15\xd3\x8e\x4a\x55\x4e\xf8\xa5\x4b\x2c\xed\xc4\xf4\x32\x40\x62\xac\xd8\x26\xb5\xf5\x98\xc5\x72\x28\x9f\x24\x49\xb3\x55\x5b\x21\xb6\x96\x2c\x59\xbd\x5d\x4a\x71\xce\x0c\x75\xac\xfe\x32\x4c\x17\xcf\x0a\x45\x5a\x18\x26\x9b\x12\x78\x1a\x96\xd2\xf9\x12\x40\xce\xa8\xa1\xd6\x07\xd2\x5d\x7c\xa2\x0f\x2a\x4f\x50\xe4\x4e\x5b\x48\x2a\x22\x23\x2e\xb3\x28\xac\xa5\xfc\xc0\x50\xb9\x18\x7c\xb9\xdc\x2e\x0d\xcf\x8c\x0f\x67\x96\x0e\x8e\x22\x77\x48\x47\x78\x78\x87\x7a\x33\x76\x6c\x77\xbb\x79\x81\x83\x1e\xb4\x85\xbc\xc3\xea\xa7\x81\x13\xc6\xba\xb6\xc1\x54\xfd\x05\x4b\x48\x38\xd3\x8e\x1e\xaf\xd0\x32\xe8\x75\x91\x0a\xad\xb8\x15\xab\x45\xb7\x35\x4c\x33\xe4\x5d\x57\x12\xa6\x2f\x8f\x45\xa3\x75\xf8\xf6\xcc\xc2\x1d\x96\x3f\x3a\x8c\x41\x2f\x57\xe9\x08\x69\xbf\xa0\x2e\xa4\x83\xea\x35\x46\xdf\xcf\x01\x4c\xc3\x19\x35\x0c\x9f\xa9\xb3\x55\x47\x52\xb1\x28\xb4\xcb\x02\xdf\x81\xab\x0b\xd3\xce\xb8\xfd\x0c\x38\x26\xa2\x4b\x7f\x1a\xb5\xe4\xe9\x0a\x34\xa8\x2d\x5c\x6f\x36\xe1\x62\xa3\x87\x65\x28\xe6\x17\x03\xaa\xa4\x23\x7e\x92\x01\x0c\x1c\xbd\x32\x60\x0e\xab\xbe\x83\x31\x41\xe3\xf1\xc6\x54\x78\x84\xe4\x09\x15\x94\x78\x35\x3e\xbf\xb7\xfa\xbb\x79\x2b\xfc\x72\xfa\x07\x4c\xd4\x67\xed\xb7\x3f\x7f\x03\x00\x00\xff\xff\x44\x98\xa6\xbb\x69\x04\x00\x00")

func dataFruitsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataFruitsJson,
		"data/fruits.json",
	)
}

func dataFruitsJson() (*asset, error) {
	bytes, err := dataFruitsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/fruits.json", size: 1129, mode: os.FileMode(436), modTime: time.Unix(1523619958, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataHerbsJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x55\x4d\x73\xe3\x36\x0c\xbd\xef\xaf\xe0\xe4\xdc\xfe\x08\x5b\xbb\x9b\x6d\xe3\xb4\x9e\x38\x4d\xa6\xdb\xe9\xe1\x59\x82\x25\xc4\x24\xa1\x82\xa4\x33\xca\xaf\xef\x78\x45\xca\xb4\xf7\x86\xf7\x2c\x7c\x10\x78\x80\xff\xf9\x64\xcc\xdd\xca\xf7\x64\xb9\xc5\xdd\x2f\x67\xb4\x46\x60\x3b\x9b\xdf\xc4\x4e\xa6\xc2\xcf\x03\xb8\xc6\x6b\x4c\xc6\x12\x0e\x33\xfa\xcd\x77\x0c\x6f\xae\xc9\xb5\xd8\x4e\x8a\xa9\xe8\x69\xb6\x9b\x81\xf4\x54\xc2\x34\x03\x9f\x28\xcc\xf6\x3d\xd4\x72\x6b\x6a\xaa\xe1\x96\xec\x94\x6d\x51\x86\xef\x48\xab\x14\x0d\x5b\xf8\xa8\x4b\x16\xcb\xa7\x73\x1d\xcb\xa7\x33\xff\xc2\x14\x3d\x1c\x05\xba\xfd\xa5\x49\xb5\x7f\xa3\x14\x4a\xe2\xa4\x5a\xbf\xe5\x33\xdb\x5c\xf1\x97\x11\x1f\x12\xf3\x5b\xbe\x91\x1b\x4b\xbf\xde\x60\x02\x7c\x44\xc1\x29\xc6\x34\x79\x86\x69\x45\x3b\x2c\xfc\x14\x82\x64\x9f\xdf\xd9\xed\xd3\x6c\x3e\xb0\x1f\x4c\xcf\xc2\x33\xdc\xe0\x44\x97\x2a\x37\xe4\xc4\x9b\x3d\xac\xab\x71\xaf\x28\xd5\xce\x84\x9b\x34\x5a\xaa\x99\x13\xe9\x9e\x7c\xce\xbc\x61\xe7\x65\x1c\xd8\xc2\x40\xc5\x21\x2e\x63\xdf\xc8\x69\x99\xce\x23\xf4\x4d\x14\x39\xd3\x23\xfb\x98\xad\xd4\xbf\x8b\x16\xc0\x31\xa4\x7d\xf6\xfe\x53\xa9\x87\xcf\x2d\xdc\x42\x83\xa5\x3c\xb1\x2d\x29\x5b\x9b\x3f\x7b\x92\x40\x0e\x9a\x7f\x7a\x4a\x39\xdf\x6e\xc9\xbc\xc3\x49\xca\xcf\x3b\xf8\x30\xe4\x90\xbb\x81\x43\x31\x45\x95\x8a\x20\xa1\x8a\x5e\x7c\x91\xe7\xe4\x72\x9c\x57\x91\x4e\xd3\x21\x0f\x6e\x25\x5e\x34\x77\x75\xf5\xf6\x0e\xce\x0e\x2b\x6b\xc3\xc8\x6d\xf6\x59\xb9\x76\x10\xc9\xed\x5e\x79\x0e\xa5\xa6\x08\x35\x15\x5e\x05\x1c\x84\x22\x77\xf9\x51\x0d\xdc\x38\x14\xbf\x06\x8a\x77\x4c\x0b\xe8\xe0\x24\xf7\x71\x6d\xd1\x1e\xcd\x35\xd7\x20\x04\x2e\x71\xc8\x92\x4e\x66\x94\xf7\x8b\x34\x67\x2a\x10\x75\x65\x55\xa0\x62\x79\x59\x22\x3f\x16\x9b\xbd\x87\x2b\x9d\x68\xac\x9c\xe8\x76\x63\xaa\x28\x69\x4f\xfb\x62\xba\xd2\x8d\x3f\xb8\x27\x6b\xcf\x02\x8e\x7c\x2a\xc7\x20\x79\x4e\xce\x8c\xa4\x81\xdb\xe4\x6e\xf7\xe0\x6c\x55\x71\xbf\x92\xf7\x65\x34\x5f\xc9\xa7\x5e\x89\x8e\x19\xb2\xef\x49\x55\x24\xab\xe7\x5e\x09\x91\xd4\xdc\xc3\xc2\xf7\xb0\x45\xb1\x21\xfc\x44\xce\x27\x21\xdb\x3f\xc2\xe4\x39\x64\x01\x9b\x9a\xbc\x17\x3b\xa2\xd8\x0a\xf6\xc1\xc8\xc1\x6c\xa1\xe8\x96\x01\x5e\xf8\x1d\x59\x76\x65\x57\x35\xd0\x8f\xaf\x86\xbc\x98\xc9\xf3\x48\x6a\xf6\xa4\x45\x91\x0f\x72\x2c\x3d\x78\x10\x85\xb2\xcb\x6d\xfa\xac\x4c\x9d\xd9\x70\x91\xdf\x86\xff\x4b\xa2\x8b\xb2\x36\x1c\x03\xc1\xb4\xe7\xb6\xa3\xac\x58\xbb\x2c\x9b\xef\xe5\xd7\xbe\x7a\xc3\x23\x42\x2c\x0f\x7e\xc4\x60\xb1\xaf\x05\xf4\x98\x42\x84\xe6\x8e\xaf\x55\xde\xfd\x35\xf5\x3a\x70\xa4\x6b\x2a\x4f\x36\x83\x37\xf8\x3e\x14\x90\xa2\xa3\x7e\xd9\x06\xee\x11\x45\xcd\x96\xc6\xb1\x14\xb3\x56\x7c\xb0\x3d\x1f\xd4\x9a\x6d\x06\xb6\x7c\xcd\x60\x3a\x4f\xdf\x8c\x15\xb7\xc5\xa8\x7c\x5c\x0e\x8c\xef\xaf\x3c\xb6\xa4\xe9\x74\x1b\xf8\x0b\x42\x34\xab\x70\x4b\xef\xb8\x1d\xd2\x2d\x57\xdd\x87\x67\x04\x07\x7f\xeb\x35\x37\x6c\x66\x5a\x51\xbf\x48\x8f\xfc\x4f\xec\xdc\xb6\x5b\x76\x2b\x8e\x7a\x85\x47\xa4\x4a\xe7\x5b\x19\xc7\x7a\x2b\x9f\xd0\x0d\xc9\xf3\xe5\xc6\x95\x02\x0f\x07\x2d\x2b\xb9\x83\x8d\xc5\xd2\x80\x11\xd5\x59\xdc\x21\x04\x1c\x14\xf9\x92\xef\x28\xa0\x28\xa9\x3e\x7c\xc9\xa1\x2d\xcf\x75\x50\xf6\x39\xfd\xb3\xf8\x23\xcc\x9e\x50\xee\x60\x52\x47\x5a\x14\xf4\xd7\x07\x3e\x72\x69\x2f\xf0\x97\xa4\x2f\x82\x18\xce\x0a\xe7\x03\x15\x81\xbf\x22\x60\x9f\x3f\xfe\x3b\x7d\xe4\xff\xa5\xef\xd4\xc9\x72\xb4\xbf\x93\x52\x18\x8e\x05\x84\x78\xf7\xe9\xdf\xff\x03\x00\x00\xff\xff\xa8\xee\x99\x9d\x46\x08\x00\x00")

func dataHerbsJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataHerbsJson,
		"data/herbs.json",
	)
}

func dataHerbsJson() (*asset, error) {
	bytes, err := dataHerbsJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/herbs.json", size: 2118, mode: os.FileMode(436), modTime: time.Unix(1523619695, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataIngredient_corpusTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x56\xcb\x96\xec\x28\x0e\xdc\xeb\x2b\x58\xcd\x6e\x66\xce\x3c\xbb\xfb\x73\x64\xac\xb4\xb9\x06\xc4\x91\x20\xf3\xfa\x7e\x7d\x1f\xf1\xc8\xac\xea\x4d\x9d\x02\x93\x42\x8a\x08\x85\xf0\x27\x97\x42\xbb\xd3\x18\x3c\xed\x6e\xef\x7f\x7d\xab\xce\x63\xce\xb6\x5f\x83\xbf\xd4\xf1\xa3\x6f\x86\x5c\xd9\x95\x40\x9e\x14\x5a\x56\x8c\xb5\x9f\x2e\xae\x12\x6a\x61\xce\x4e\xdb\x81\xe2\x1e\x91\x9b\x38\x3a\x0e\xf7\xc4\x1c\x62\x44\x47\x07\x1d\xea\x4a\xc8\xfe\x84\x1d\xf5\x74\x11\xe5\x20\xe7\x19\x45\xc9\x3d\x42\x26\xa7\xe1\x61\xe1\x36\xe1\x8b\xf2\xdf\x5b\x71\x16\xdf\x55\x2d\xae\x6e\x5a\x1c\x70\x88\x8e\x4b\x0d\x9c\x31\x3a\xfa\x59\x05\x7d\x75\x1b\x66\xcc\xe8\x12\xea\x49\xfb\x0c\xaa\x09\x63\x74\x18\x23\x58\xdc\x78\x03\x3d\xb1\xb0\xa0\x45\x4f\x21\x5e\x90\xc8\x12\x87\x96\x5f\x82\x56\x3d\xd8\x4d\xd0\x53\x87\xad\xd5\x4a\x02\x07\x4a\x0c\x1e\x5e\x68\x0b\x8e\xe1\x49\x8e\x43\x84\xfe\xf3\x5e\x1d\x70\x0e\x9c\xa1\x50\x29\x34\x17\x0a\x5b\x44\x7f\xb9\xb9\xb7\x09\xbf\x26\x20\x40\xc7\xa1\xe0\x43\xce\x98\x38\xc3\x86\x57\xc8\x87\x2b\xfc\xda\x49\x20\x52\xe2\xec\x7e\xb4\xe0\x09\x2a\x27\xac\x4c\x0a\x13\x37\x28\x28\x1a\xe9\x5e\x3f\x51\xde\x11\xd4\xc0\xf5\x42\x98\xe0\x49\x07\x55\xdc\xe2\xc8\xce\x53\x24\xb9\xe1\x08\xf9\x58\x71\xa1\x9f\x73\xfe\x24\x52\x02\x8f\x22\x5c\x15\xfc\x49\xfb\x8e\xb2\xb6\x37\xa2\x07\x14\xae\xe3\x6a\x8b\x74\x72\xa6\x1b\x72\xab\x89\x0e\x98\xa7\x94\x6f\xa7\xd8\x3c\x41\xc2\x9b\x73\xc6\x60\x11\xcf\xe0\x2f\xca\xc6\x5a\x3d\x81\x85\x0e\xcc\x0c\xbe\xa5\x90\xa1\x9e\x77\xa2\x89\xe4\x2a\xb6\x73\x8a\x79\x5f\x20\xa5\xa6\xa7\x30\x27\x43\x27\x62\xae\xc2\xb0\xa1\x86\x08\x85\x3c\x1a\xa2\xe8\x39\xc3\x49\xf8\xbc\x67\xc9\x9f\x1b\x09\xb5\x2a\xbc\x58\x3c\x69\x25\xd1\x33\x08\xcd\x0c\x0b\x16\x09\x17\x82\x3f\xd9\x73\xc4\xfa\x4e\x14\x5e\x18\x73\xeb\x10\x84\x18\x56\x56\x18\x13\xe7\x5d\x21\x86\x44\x93\x89\x82\x92\x48\x31\x2f\x8c\x4a\xc8\x84\xa5\x44\x02\xb1\xcf\x2c\x98\x8f\x75\xf6\x10\xa2\xbc\x2a\x12\x0c\x1a\xb2\x82\x67\xcf\xb9\x55\xf0\x78\x53\xce\xb4\x3e\xf7\xcb\xf7\xf0\x83\xb3\x4b\x4d\x2b\xca\x0e\x9e\x25\xdb\x7f\xfe\x84\xf4\x0b\x85\xac\x5f\x16\x33\x5d\x8d\x5d\x75\xcf\x90\xc9\xa4\xd4\x93\x50\x10\x7a\x63\x38\x44\x33\x2b\x37\x58\x76\xe7\xa5\xa5\x4d\x81\xb1\x2a\x68\x09\x19\xfd\x09\x7a\xb2\x54\xca\x21\x1f\x5f\x7e\xec\x1e\x11\x2f\x52\xd0\x13\x63\x34\x69\xcc\x60\x05\xb5\x52\x3f\xb7\x51\x8c\xeb\x26\xc3\x07\xf4\x94\x90\x0a\x28\x25\x8e\x21\x23\xfc\x6a\xde\x9f\x21\x07\xd0\x2a\xf8\xda\x48\x24\x58\x7e\xac\x94\x50\x6e\xf0\x98\x39\x62\xd7\xe7\xc0\x69\xf5\x4a\xef\x5c\x05\xf5\x18\x63\xdf\xf1\x91\x9f\xa4\xf0\x86\x65\x12\xad\x95\xfd\x65\x2b\xfb\xf8\x3a\x43\x29\xd6\x09\x43\x0c\xbd\x5a\x48\x58\x22\x39\xbd\xa5\x95\x49\x4c\xc7\xd4\x4d\x50\x37\x8c\x8a\x29\x78\xb7\x20\xdc\xe5\x76\xaf\x33\x54\x72\xaf\x90\xfb\xd9\x80\xd9\x74\xb0\xe1\xed\x22\xe1\x03\x2e\xaa\xfe\x6c\x05\x6e\x3e\x9a\xd4\x8e\x83\x1d\x7d\x47\xc0\x27\x7b\xdc\x19\x94\x14\xd3\x6c\x3e\xdc\x36\x3c\x68\xc5\xb0\x64\x37\x61\xef\x39\x86\x8f\xea\xbf\x19\xc4\x77\x25\x77\xc1\xe0\x6c\xd1\xd1\x04\x2b\x4e\x47\x6c\x69\xd5\x37\xdf\xd2\x46\x02\x85\x30\xb7\xea\xa6\x61\x61\x8c\x5a\x4c\x8c\x56\xdc\x1b\x42\xc1\xbc\x08\x49\x21\x57\x38\x31\x4d\x16\xbe\xd2\xba\xc5\x46\xeb\x58\x77\x98\x42\xa8\xe0\x9b\xc8\xfd\xbe\x95\x25\x2f\x49\x0f\x13\x8d\x54\xab\x09\x6e\x00\x39\x23\x4d\x3c\x94\x68\x57\x28\x2c\x17\xd4\x26\x89\x24\x78\x30\x45\x21\xec\x21\x46\xb8\x29\x46\x7e\x0d\x21\xc0\x17\x22\x0c\xe6\xb1\xf9\x03\x23\x16\xca\xec\x7a\x9b\x92\x4e\x1b\xe3\x87\x5b\x86\xe1\x94\x5b\x81\x8d\xcc\x24\x46\xfb\xae\x99\x30\x5d\x78\x7c\x9a\xee\xd3\x0d\x7e\x00\x66\x4e\xb3\x93\xbc\xa9\x1c\xf7\xaf\x55\x42\x39\x50\x2c\x99\x05\x93\x45\x19\x62\x4b\x1c\x51\xd5\x92\xe1\xfc\x20\x6f\x93\x88\x44\xa7\xc3\x97\x96\xca\x15\x3e\x18\x99\x24\x46\x47\xd6\x26\x17\xdd\x70\x77\x92\xfb\x28\xd1\x29\xcf\x2e\x58\x35\xd1\x98\xa7\xbc\x73\x10\xd4\xb2\xe8\x30\x6b\x9e\xe6\xda\x45\xf4\xb7\xb7\xcd\x04\xcf\xb5\xbe\x8d\x42\x4d\xe3\xb3\x79\x47\x57\x4c\x0b\xd0\x22\x78\xbf\xcd\x75\x29\x6a\x8e\x95\xb5\xec\x9e\xe2\xbe\xe3\xe2\xb1\x90\x58\x02\x1f\x9d\x1c\x14\xb1\x86\x3c\xa1\x59\xd4\x0c\xb8\x67\x61\x0f\xfa\x92\xd3\x2b\xa8\xbe\x47\xcf\x49\x26\xa7\xf7\x80\x1b\xe3\x66\xb6\xab\x01\xc2\xd1\x7a\xb2\x7c\x98\xae\x28\x15\x87\xf2\x12\x61\xfc\x38\xf0\x74\xdd\x50\x31\x06\xcc\x4e\x09\x95\xbb\xad\xf5\x3b\xba\x5a\xb0\xc5\xf0\x88\xfc\x22\xf9\xa2\xaf\xaf\x94\xf7\x47\xcf\xa0\x75\x18\x6a\xa7\x0a\xb5\xa0\xe0\xd1\x74\xcc\xae\x05\x8f\x76\x81\x59\x8e\x7d\x08\x8f\x6e\xef\xae\x3e\xde\x02\xae\xb2\x54\x1b\xd7\x0a\x3b\x56\x52\x88\x44\x97\x42\x69\x62\xd9\x0e\x45\x7f\xf1\xbe\xde\x8a\x2b\xf6\x2e\xa4\x6a\xc9\xeb\x15\xd2\xe8\x2c\xc6\xda\x0b\x4e\x98\x0f\x86\x43\xf0\xb4\x09\x2e\xe8\x2f\x92\xe5\xec\x8f\xa0\xe7\x9a\x75\xc3\x07\x2c\xab\xe5\x53\xd0\xb5\x61\xef\xab\xb7\x3a\x70\xef\x27\xc6\x24\xfd\x8b\x8b\xd6\x96\x71\xc0\x60\x92\xed\x82\x5c\xef\x36\xd3\xe5\x9c\xbb\xb3\xfa\x21\x28\x3a\x8e\x62\x92\x82\x12\x5b\xfa\xb0\x5a\x51\x04\x0f\xb6\xd9\x2f\xcd\x32\xf3\x27\x29\x1c\x8c\x75\xa5\x62\xcc\x5d\x61\xcf\x74\xaf\xde\xb2\x19\xbb\x38\x7f\x7b\xbe\x35\xf7\x5a\xd4\x33\x1c\xa7\x4d\xa6\x47\x1b\x8d\xf3\x81\xbb\x1f\xe9\x6e\xf5\x79\x09\x8d\x6e\x19\x6f\xa1\xee\x44\x6b\xb8\x8d\x01\x03\xdf\xee\xd0\x82\xc7\x49\xb5\x86\xf9\x0c\x2b\x44\xf1\xdb\x33\xae\x0f\x67\x88\xc7\x3f\x26\x77\xd3\xbb\x3c\x5e\xe4\x52\xf8\x39\xc7\x18\xfc\x0b\xfe\x0d\xff\x81\xff\xc2\xff\xe0\xff\xf0\x1b\xfc\x0e\x7f\xc0\x3f\xff\x0c\x00\x00\xff\xff\xd4\xff\xcb\xb8\x5b\x0b\x00\x00")

func dataIngredient_corpusTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataIngredient_corpusTxt,
		"data/ingredient_corpus.txt",
	)
}

func dataIngredient_corpusTxt() (*asset, error) {
	bytes, err := dataIngredient_corpusTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ingredient_corpus.txt", size: 2907, mode: os.FileMode(436), modTime: time.Unix(1523087669, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataIngredient_densitiesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x57\xdb\xb2\xa3\x3a\x0e\x7d\x3f\x5f\x41\xed\xe7\xae\x14\xd7\x90\xf4\xdf\x08\xa3\x80\x27\xc6\x66\x7c\xd9\xa9\xf4\xd4\xfc\xfb\x29\xdb\x92\x21\xdd\xfd\xba\x90\x65\x59\x5a\x5a\x12\xff\xfb\xa7\xaa\xbe\xa6\xe0\x3d\xda\xaf\x9f\x55\xdb\x8e\x97\xfa\x47\x82\x54\xc0\x4a\xac\x88\x0e\xbf\x7e\x56\x4d\x37\x10\x7e\x40\x7d\x7f\x19\x19\x9a\x67\xb0\x27\xeb\xf1\xce\xd6\xc6\x7b\x58\x4e\x8e\xda\xe6\x7a\xe9\xf2\x27\x8b\xb0\x9d\x3e\x74\x23\x7d\x78\xa0\x87\x93\xaf\xa1\x26\x5f\x3b\xd8\x0d\x1d\xe8\xd3\xb7\x9a\xbf\x59\x99\x6e\x3a\xb9\xeb\xf9\x21\xee\x25\x9d\x3b\xbf\x64\xe4\xb0\x63\x00\x11\xb9\x8d\x97\x5b\x42\x5e\xab\xdc\x77\x9c\x2b\xfe\x32\xb2\x7b\x67\x82\x2d\x68\x8c\xb4\x4d\xf0\x6c\xd1\x39\xa9\x97\x74\xdf\x70\xe9\x13\xb8\x49\xf5\x4c\x99\xac\xe9\x3d\xdf\xb8\xa0\x87\x49\x61\x65\xa4\x4a\x29\x68\x2e\xc3\x8f\x23\xed\x7c\xa0\x3f\x52\x6c\x84\x51\xe0\x73\x2a\x1b\xba\xec\x6d\x96\x60\x3d\x9d\xcf\xf1\xba\xa7\xdc\xaa\x3f\x8e\x7f\x83\x96\x4a\x41\x3a\x3c\x5c\xee\x09\xc3\x6f\xd8\x8d\x05\x8f\x73\xb1\x1f\x5a\xae\x34\xb8\x14\x57\x73\x6f\x08\x99\xa5\x52\xd5\x0b\x71\xfe\xfa\x59\xdd\xc8\xc3\x16\x9c\x07\x3b\xa7\x8a\x74\x9c\x16\x50\x29\xa0\x7b\x8e\x50\xc8\x19\x6d\xf5\x2d\x35\x2e\x90\xc8\xd4\x31\x0d\xce\xd8\x8d\x4b\x86\x73\xf5\x92\x1a\xff\x7a\x60\x02\xe5\x60\x93\xe2\xfc\x71\xe0\x07\x4e\x88\x8f\x94\x87\x96\xf2\x30\x81\x30\x3a\x45\x76\x25\x64\xcd\x95\x1a\x6a\x2a\x8a\x58\xa5\x78\x62\xb2\xb9\xf7\xe4\xc6\x07\xfb\xc4\x77\x34\xab\xaf\x97\x6b\x82\x16\x8b\xa8\xab\x09\x41\x3b\x2a\xe1\xc8\x37\x7a\x97\x0e\x0f\x74\x81\x00\x6b\x4d\xc6\xc6\x1b\x9d\x76\x2f\x44\x5f\xed\xc6\x83\x37\x98\x1d\xf0\x5d\x27\xb0\xb9\x95\xa0\x8c\x4d\x11\x5d\xfb\x4b\x93\xad\x10\x1c\xd5\x3c\x03\x6e\x97\x1a\xc4\xfa\x61\x64\xc0\x6f\x08\xb1\x60\xd7\x9a\x78\xb4\xc1\xdb\x68\x0d\x92\x1a\xa9\xe6\x78\x56\x63\x3d\x6a\x22\x68\xcd\xfd\xe5\x40\xc1\x4c\x48\xe9\x9e\x78\xf2\x76\x27\xaa\x11\x4f\xcb\x89\x1d\x41\x07\x5f\xe8\x7b\x2d\xcd\x2d\xcc\x19\xe7\xca\x1a\x25\xbf\xf1\x0f\x6b\x87\x0e\x36\xfc\xc3\x18\xbe\x8d\x80\xd9\x7c\x60\x02\xb4\x51\x40\xa6\x31\x5f\x03\x07\xbe\xe5\x3a\x8f\x1c\xf9\x06\x76\x01\x2b\x75\x7a\x78\x3b\x50\xb9\x8e\x8e\xfb\x2d\x05\xc3\x67\xe4\x11\xbb\x36\x44\xf0\x94\x96\x12\x1e\x17\x0e\xdc\x0e\x16\x96\x90\xca\x59\xf7\x9c\x10\x63\x9f\x49\x02\xaf\x07\x41\x1e\xd5\x64\x8d\x5f\x29\x10\xea\x08\x54\x68\xdf\xe9\x11\xcc\x55\x62\xe2\x61\x5c\x77\x64\x3c\x59\x23\x84\x51\x32\x11\xb9\xa7\x1a\xb2\xb9\x33\x61\x4f\x5d\x7e\xfb\x90\x4f\xf3\xa8\xb6\xe0\x56\x6b\xcc\x76\x98\xdc\xe9\x45\x29\x28\xe7\x8d\x60\x7d\xc8\x27\x8d\x96\xd4\x2d\x2d\x5d\x92\xd9\x49\x12\x30\x12\x77\xac\xdc\x92\xbb\xba\xa6\xf8\xbc\xd9\xb2\x55\x64\x18\x57\x24\x24\xe6\xb4\x43\x43\x9c\x9e\xc0\x4e\x28\x02\x56\xe5\xd3\x78\xff\x38\x5f\xc2\xac\xaf\x9c\x24\x7e\x62\x89\x94\x65\xf7\x05\x34\x95\x7a\xce\x51\x6e\xd0\x1d\xf7\x3d\x7d\x68\xba\x63\x38\x38\x9f\x24\xaf\xf0\x6d\x01\xab\xa4\xa0\x41\xc5\x25\x0e\x0e\x96\x3c\xbc\x58\x9d\x79\xaa\x44\xb0\xe5\x72\x46\x6d\x3a\x2e\x69\x0f\xe1\xc9\xc2\xd0\xdc\xee\xdc\x79\x94\x7c\x47\x57\xdf\x7e\xe3\xdf\x99\x12\x1d\xbf\x56\xc9\x23\x3b\x03\x2b\xde\xcb\x58\x81\xce\xa3\x75\xab\xb4\xe7\xf4\xf1\xe5\xa9\x6c\x59\x1e\xae\x74\x3b\xec\xbb\xe2\x51\x7c\x3d\x77\x6a\x99\xe7\x91\x48\x7d\xa1\x7c\x4c\x50\xa9\xd4\x6a\x74\x52\xbf\xe1\x7e\x08\x4b\x74\x7f\xe3\x29\x20\xa4\xd6\x90\x1b\x6e\xe8\xe8\x46\x0b\xd2\xc9\x1c\x45\xcf\xbc\x78\xd8\x20\x3d\xbd\x91\xe6\x52\x20\xc9\xae\x19\xd9\x60\x8f\xdd\xf8\xb6\xa9\xf6\x5d\x5b\x18\xff\xd2\x15\x5b\x37\xd7\x9a\xea\xb4\x87\x6d\x7f\xca\x54\x92\x8e\x65\x6f\x02\x0d\x3a\x55\xb8\x04\xbc\x81\x75\xeb\x06\x4a\x99\x57\x8c\xa8\xbb\x53\x8a\x93\x4c\xcc\xb9\x24\x03\xd9\xa6\x54\x39\x62\x7d\x7b\x40\xd5\x7f\x42\x16\xbf\xc8\xb1\x53\x56\x4b\xfe\x7b\x1e\xf6\x39\x80\xec\xa2\x34\xe1\x8a\xd6\xca\xac\xf5\xf5\x8d\x0d\x55\xc0\xe9\x84\x73\xb5\x84\x05\x7d\xe0\x4d\xdb\x91\x93\x19\x7c\x46\xca\xca\xb2\x4b\x8d\x5c\xdb\xb6\x63\xe6\x28\xdc\x32\x01\xe2\xe0\xab\x0f\xe8\xf4\x04\x56\x2d\x25\x37\xfc\x0b\x6c\x2c\xe8\x05\x99\x43\x67\xec\x30\x1e\x78\xdf\xd8\x11\xc4\x4a\x63\x8c\x15\x65\x47\xb0\x79\x0e\xf2\x8c\x29\xa1\x9e\x3c\x94\xa5\x0c\xdc\x7e\x7a\xf0\x78\x67\x7a\xec\x16\xde\x1f\x3b\xda\x06\x7a\x31\x74\x7b\x46\x9e\x72\xd6\xf8\x2e\xa3\x38\x4e\x6d\x12\x44\x98\x26\xea\xe1\xa6\xa5\xe2\x0a\x08\x4a\x3e\x94\x79\x65\xce\xf7\x5c\x61\x61\xac\x04\x3d\x67\xb4\xe8\xc2\x49\x42\x1c\x6d\x24\x54\xa0\x20\xc2\x36\x91\x0f\x5e\x53\x71\x59\x76\x05\xda\xa7\x75\x93\x13\xb9\x48\xbd\x24\xbb\x3b\x7b\x7d\x42\xee\xc4\x66\x2c\xa5\xc1\x67\x74\x5f\xfa\x57\xa1\xf7\x99\x53\x3d\x2b\xa2\x13\xa0\x14\x77\x75\xd9\x65\x73\x80\xa5\xdd\xc7\xe6\xd8\x7f\x9d\x4a\x2d\xdb\xf1\x30\xf9\x15\x84\x58\xa5\x96\x94\xa1\x9e\xd2\x0b\xcb\x8a\xde\xe7\x41\xc2\xed\x9e\xd5\x97\x2a\x31\x52\x32\x49\x92\x0f\xa6\x34\x1f\x3a\x9b\xb4\x2a\x53\xa0\x2f\xaa\xb0\xca\xef\x04\xb1\x24\xbb\x35\x76\x60\x5e\x85\x8a\xd8\xa4\xed\x28\xdd\x5f\x46\x2e\x78\x47\xf3\x89\x1f\xfa\x40\xad\x31\x8e\xdb\x1b\x27\x6d\x85\x6d\x0a\x76\x21\xd1\x2f\x33\xd8\x86\x25\xa4\x85\xb6\xad\x8f\x28\x52\xbe\xbb\xf1\xdc\x09\x29\x0d\xd7\xc2\xd5\xa2\xdf\xec\xe8\xa1\x4c\x48\x48\xd3\xf1\x98\xcb\x4b\x89\x43\xcc\x7a\x51\xee\xd4\x21\xbf\x68\x60\xa2\xf3\xb2\xc3\xeb\x73\xc7\x6d\xbd\xc2\x2f\x54\x6c\x5e\xdf\x3e\x56\x26\x47\xbf\x4e\x37\xc2\x04\xb1\xb9\x3e\x77\x7a\x75\x32\xe4\xb9\xc7\x0e\xef\x0c\x65\x4d\x3b\x2f\xfb\xc4\x59\xf3\x78\x20\xd2\x66\x50\x76\x11\x4b\xbf\x29\xe7\x16\x27\x68\xfc\x98\xc5\xac\x72\xcd\x48\xad\xe4\x43\x16\xd9\xfe\x5a\xfe\x03\x41\x3c\x3f\xfa\x70\x38\xa6\x36\x2f\xac\x77\xd6\x32\x67\xde\x87\xd3\xa2\xa8\xde\x3c\xc2\x87\x32\xbc\x56\xe9\xe3\x70\xc4\xb4\x88\x16\xd6\x25\xa0\x12\x36\x6c\x53\x56\x6f\xae\x06\x5b\x36\x3d\xe7\x72\xb1\xb0\xc6\xdf\x47\x0b\xe2\x99\xfb\xf8\xd6\x97\x9f\xb5\x39\x0d\xba\x91\x6e\xdf\x25\x46\x9f\x2e\xf1\xb1\xe5\x79\xbb\xa0\x02\x9f\x67\x4c\x49\x72\xee\x82\xf8\x3f\xd8\x15\x19\x4d\x06\xbc\x98\x6e\x46\x81\x73\x99\xfe\x25\xbd\x71\x8d\xaf\x0a\xb5\xda\xee\xb4\xdd\xd3\xa2\x7e\xfc\x2a\xbf\x31\x4e\xab\xea\xe3\x1b\xe7\x2e\x82\xf1\x1f\x2b\xef\xfb\x2d\xdf\xf9\xdf\x20\xb5\x01\x6a\xdb\xe1\x34\x37\x69\x6b\x6f\xc6\x9e\xd2\x97\xb3\x5a\xe0\x96\xb7\x1e\xdc\x8c\x92\xb9\xac\x35\x77\x3e\x2e\x4b\xa5\x8d\x99\x69\x28\x36\xec\xc3\xc5\x1f\x2f\xae\xd5\x3f\xff\xff\x37\x00\x00\xff\xff\x8f\xba\xf8\xf2\x1e\x10\x00\x00")

func dataIngredient_densitiesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataIngredient_densitiesJson,
		"data/ingredient_densities.json",
	)
}

func dataIngredient_densitiesJson() (*asset, error) {
	bytes, err := dataIngredient_densitiesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/ingredient_densities.json", size: 4126, mode: os.FileMode(436), modTime: time.Unix(1523087669, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataTop_5kTxt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x57\x5b\x9a\xdc\xac\x0e\x7c\xd7\x2a\xd8\x9a\x8c\xd5\x36\x31\x20\x7e\x09\xa6\x8f\x67\xf5\xe7\x13\x17\x77\x27\x4f\x19\xb7\x31\x88\x52\xa9\xaa\xa2\x18\x2b\x68\x3b\x50\x60\x6b\xb5\x92\xc0\x81\x12\x83\x87\x37\xda\x03\xc7\xf0\x43\x8e\x43\x84\x14\xe2\x05\xaf\xc8\x4d\x80\x73\xe0\x0c\x85\x4a\xa1\xf9\xa0\xb0\x09\xbf\xb3\x1b\x1b\xd1\x71\x28\xf8\x90\x33\x26\xce\xb0\xe1\x15\xf2\xe1\x0a\xbf\x77\x12\x88\x94\x38\xbb\x3f\x2d\x78\x82\x84\xf5\x97\x5d\x22\x8c\x50\x39\x61\x65\x52\xf8\xc1\x1c\x62\x44\x28\x28\x1a\xe9\x5e\x5f\x2b\xef\x08\xca\x4d\x9c\x17\xc2\x04\x3f\x74\x50\xc5\x2d\x8e\xda\x3c\x45\x92\x1b\x8e\x90\x8f\x75\x04\xf4\x75\xce\x9f\x44\x4a\xe0\x51\x84\xab\x82\x3f\x69\xdf\x51\xd6\xcf\x1b\xd1\x0b\x0a\xd7\x71\xb4\xed\x74\x72\xa6\x1b\x72\xab\x89\x0e\x98\xab\x94\x6f\xa7\xd8\x7a\xc1\x37\xe7\x8c\xc1\x76\x3c\x83\xbf\x28\xbb\x4d\xb8\x9e\xc0\x42\x07\x66\x06\xdf\x52\xc8\x50\xcf\x3b\xd1\xc4\x71\xdd\x3b\x35\x3d\x85\x39\x19\x30\x11\x73\x15\x86\x0d\x35\x44\x28\xe4\xd1\xf0\x43\xcf\x19\x4e\xc2\x9f\x7b\x5e\xf1\x73\x02\xa1\x56\x85\x37\x8b\x27\xad\x24\x7a\x06\xa1\x59\x51\xc1\x22\xe1\x42\xf0\x27\x7b\x8e\x58\x9f\xc2\xe0\x8d\x31\xb7\xaa\xb0\xa3\x5c\xee\xaf\xd7\x31\xac\xa2\x30\x26\xce\xbb\x42\x0c\x89\x66\x4f\x0a\x4a\x22\xc5\xbc\x20\x2a\x21\x13\x96\x12\x09\xc4\x5e\xb3\x60\x3e\xd6\xda\x43\x88\xb2\x9b\x3c\x10\x0c\x1a\xb2\x82\x67\xcf\xb9\x55\xf0\x78\x53\xce\xb4\x5e\x8f\x5a\xc2\x1f\xce\x2e\x35\xad\x28\x3b\x78\x96\x6c\x7f\xf9\x13\x12\xff\xfe\xa2\x50\x8c\xf8\xb4\xa6\x93\xb1\x93\xee\x27\x64\x32\x5a\xf5\x32\x14\x84\xf6\xb5\xe9\x60\xcd\x84\xc2\x70\xda\x9d\x97\x96\x36\x05\xc6\xaa\xa0\x25\x64\xf4\x27\xe8\xc9\x52\x29\x87\x7c\x7c\x7d\xec\x5e\x11\x2f\x52\xd0\x13\x63\x34\x6e\xcc\xcd\x0a\x6a\xa5\xbe\x6e\xa3\x18\xd7\x49\x86\x10\xe8\x29\x21\x15\x50\x4a\x1c\x43\x46\xf8\x6d\xde\x9f\x21\x07\xd0\x2a\xf8\xde\x48\x24\x58\x7d\xac\x94\x50\x6e\xf0\x98\x39\x62\x27\xe8\x40\x6a\x8d\x0a\x66\xcc\xa8\xa0\x1e\x63\xec\xbf\xf8\xc8\x3f\xa4\xf0\x00\x33\x3b\xaf\x95\xfd\x65\x4f\xf6\xf2\x7d\x86\x52\x6c\x14\x06\x3b\x12\x96\x48\x4e\x6f\x69\x65\x36\xa5\xe3\xe9\xc6\x84\xf6\x3f\x27\xb6\x1b\x46\xc5\x14\xbc\x5b\x38\xee\x72\xbb\xf7\x19\x2a\xb9\x77\xc8\xfd\xb3\x80\xd9\xe8\xb0\xe1\xed\x22\xe1\x0b\x2e\xaa\xfe\x6c\x05\x6e\x3e\x9a\xd4\x0e\x86\x2d\x7d\x76\xc0\x1f\xf6\xb8\x33\x28\x29\xa6\x39\x82\xb8\x6d\x78\xd0\xda\xc3\x2a\xde\x84\xbd\xe7\x18\xfe\x61\x72\x67\x08\xce\x91\x1c\x43\xb0\xbe\xe8\x00\x2d\x72\xfa\xe6\x5b\xda\x48\xa0\x10\xe6\x56\xdd\x94\x27\x8c\x51\x8b\xb1\xcf\xae\xf1\x20\x26\x98\x17\xfe\x29\xe4\x0a\x27\xa6\x09\xfa\x77\x17\xb7\xd8\x68\x2d\xeb\x8a\x52\x08\x15\x7c\x13\xb9\x9f\x53\x59\xf2\xe2\xb0\xeb\xf4\x8b\x54\xab\xf1\x6b\x40\x36\x77\x9a\x37\x57\xa2\x5d\xa1\xb0\x5c\x50\x9b\x24\x92\xe0\xc1\x08\x84\xb0\x87\x18\xe1\xa6\x18\xf9\x3d\xfa\x0e\x5f\x90\x1b\xa0\xe3\xc7\x3f\x18\xb1\x50\x66\xd7\xe7\x92\x74\xca\x16\xbf\xdc\x12\x0c\xd8\xc8\xf4\x61\x8c\xea\x6c\xef\x7c\xa0\xff\x55\x41\x5f\x61\x8b\xe8\x2f\x37\xd6\x4d\xd5\xe9\xb2\x3e\x80\x33\xc5\xd9\x49\x9e\xe6\x8d\x3a\xd6\x53\x42\x39\x50\xac\xa8\x05\x97\xed\x32\x39\xc6\x11\x55\x49\xa1\xb4\x54\xae\xf0\xc1\xc5\x1a\x3e\x86\xae\x36\xb9\xe8\x86\xbb\x37\xb6\x9b\x85\x4e\xf2\x75\x66\xaa\x51\xc2\x84\xe3\x39\x4f\x50\xcb\x6a\x81\xc9\xef\x14\x50\x09\x9e\x6b\x7d\xa6\x5f\x8d\xb3\x73\x22\x41\x8b\xe0\xfd\xc8\xe6\xe2\xca\x34\x88\xf5\xd8\xc5\xc1\xfd\x7d\x53\x8f\x85\xc4\x8e\xf9\x30\xe0\xa0\x88\x35\xe4\x79\xd9\x05\xfa\x00\x70\x96\xff\xa2\xaf\x3a\xde\x41\xf5\x31\x91\x93\x8c\x28\x8f\x55\x0d\xe3\x98\xff\xac\x69\x1b\xb3\x68\x20\x70\xb4\x29\x2b\x9f\x8e\x56\x94\x8a\x83\x61\xdd\xf3\x1e\x69\x9d\x72\x1a\x2a\xc6\x80\xd9\x29\xa1\x72\x57\xab\x7e\x62\x67\x05\xb6\x18\x5e\x91\xdf\x24\x5f\x3c\xfa\x6e\x69\x29\xb4\x4f\x69\x18\x3a\xd9\xdb\x83\x5a\x50\xf0\x68\x3a\x3c\x69\x81\xa5\x9d\x40\x56\x63\x37\xd7\x31\xbf\x5d\xae\x87\xc3\xbb\xca\x52\xcd\x86\xcd\x3e\x2a\x29\x44\xa2\xcb\x68\x20\x56\xed\x60\xee\x97\xa4\xf5\x91\x5b\x7b\xef\x42\xaa\x56\xbc\x5e\x21\x8d\x09\x4a\x98\x0f\x86\x43\xf0\x34\x47\x16\xf4\x17\xc9\x12\xea\x57\xd0\x73\x79\xd9\x98\x73\xab\x66\x29\x0e\x74\x56\x1c\xf4\xc5\x0b\xdc\xfb\x8a\xe1\x94\xff\x88\x62\x6d\x19\xc7\xf5\x8d\x9e\x9d\x7c\x33\x4f\x38\xe3\xe0\xf4\xd5\x79\xeb\x41\x2b\x3a\x8e\x62\xc4\x82\x12\x5b\xfa\xf4\xb6\xa2\x08\x1e\x6c\x5e\x2e\xcd\x2a\xf3\x27\x29\x1c\x8c\x75\x95\x62\x1d\xbb\xc2\x9e\xe9\x9e\x33\x53\xf9\xd5\x06\xf7\x3f\xe8\x75\xe5\xeb\x22\xf3\x09\x2c\x83\xf0\x23\xb2\x74\x01\x59\x16\x34\x6c\xe0\xa3\xfe\x6c\x03\x54\xf0\x38\xa9\xd6\x30\x83\x53\x21\x8a\x93\xae\x83\xd1\xdd\x44\x67\x16\xeb\xca\x37\xdc\xc5\x04\x70\x6b\x62\x71\xc8\xa3\xec\x98\x38\x81\xc7\xaa\xad\xc0\x26\x98\xf7\x7b\x40\x09\x27\x8b\x92\xe0\x1e\xf4\x84\x1f\xde\x2f\x04\x7d\x13\x55\xf7\x84\xa2\x8d\xba\x28\xbe\x5e\xb4\x3c\xd9\x7a\xa4\xff\x35\xd4\x73\x52\x71\x70\x40\x13\x5f\x66\xae\x33\x90\x7c\x13\xdb\xc6\xcc\x92\x04\x39\x3d\x29\xc6\x25\x27\xae\xff\xd2\xb5\xbc\x5f\xee\x1b\xa4\x61\x7d\x74\x1c\x2e\x33\xef\xf1\xcb\x4d\x4c\x02\x85\xaa\xf6\xf1\x71\x5d\xdf\x46\xc9\x2b\x95\xa2\xac\x34\x2a\x33\xd0\x74\x9d\x98\x76\xff\x09\x00\x0a\xff\xb5\x90\x19\xe1\xe4\xba\x32\xc1\x14\xb5\x19\x2b\xac\x01\xe6\x32\xda\x07\xbf\xbb\xc7\xea\x3e\x4a\x3b\x5a\xc4\xfe\xfa\x8b\xac\x92\xdd\x45\x92\x29\x6a\xdf\x75\x6c\xf7\x1d\x93\x74\x44\xc4\xd1\xf7\x0d\x65\x23\xdf\x56\x96\x43\xa9\xc1\x9f\x7c\x91\x3b\x09\xc5\x12\x9f\x45\x70\x3b\x51\x6b\xbf\xe7\xa0\x80\x84\xbc\xcf\x18\x37\xcb\xc5\x0d\xd5\x3f\x31\xc8\x44\x75\xc0\x37\xc3\x5a\xe7\x4c\x42\xd1\x33\x19\xcf\xde\x0a\x17\xc6\x95\x97\x86\xa1\x7c\x01\xff\xc9\x4a\xd8\xfd\xec\x13\x05\xcd\xd2\xdc\xdb\x0a\x7f\x49\x0b\x75\x76\xbf\x83\x7c\xe2\x2f\x8d\xb8\xe9\x05\xb7\x44\x58\x7b\xbb\xbd\x34\xad\x56\x51\x7d\x0c\x6c\x0a\x82\xfc\x61\xc1\x04\x1b\x6e\xb7\x5b\x31\xed\x6f\x81\x50\x28\xc2\xea\x43\xab\x95\xa1\x4f\x6f\xdb\xba\x3a\xe7\x4c\xeb\xff\x0b\x93\x5f\xf4\x83\x85\x05\x2b\xcd\xed\x57\x86\x5d\xfe\xfb\x28\xee\xa3\xac\xc3\x1f\x0f\x14\x4c\x2e\xa1\x8d\xc2\xff\x03\x00\x00\xff\xff\xc3\xc8\x26\x7d\x08\x0d\x00\x00")

func dataTop_5kTxtBytes() ([]byte, error) {
	return bindataRead(
		_dataTop_5kTxt,
		"data/top_5k.txt",
	)
}

func dataTop_5kTxt() (*asset, error) {
	bytes, err := dataTop_5kTxtBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/top_5k.txt", size: 3336, mode: os.FileMode(436), modTime: time.Unix(1523620832, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dataVegetablesJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x54\xc1\x92\x2c\x29\x08\xbc\xbf\xaf\x30\xe6\xbc\xfb\x45\x1b\x7b\x40\x8b\x56\xa6\x14\x5d\xd0\xe9\xa8\xf9\xfa\x8d\xea\x11\x9f\xfd\x22\xea\x40\xa2\x65\x42\x92\xfa\xcf\x2f\xe7\x3e\x20\x54\x61\xa7\xff\x0d\xd0\xf4\xf1\xd7\x2b\x93\x1f\xf7\xe7\xb4\x49\x1d\x7d\xe6\x0a\x08\x70\xb7\x1d\x4c\x8a\x33\x94\x4e\x21\xd5\x73\xc1\x11\x47\x86\x09\xb4\x81\x40\x1c\x3a\xe1\xf0\x28\x91\xd8\xb6\x7e\x8f\x93\x9c\x47\xe0\x1f\xec\x81\x81\xe1\xad\x12\x0f\x4a\x79\x86\x08\xfc\x56\x91\x47\xb4\x28\x43\x38\xf7\x83\x6e\xfc\x37\x5e\x78\xb8\x86\xb3\x14\x5f\x4f\x17\x52\xbd\x0c\x49\xae\xbd\xbf\xb1\x4b\x85\xe3\x85\x75\x25\x42\xa8\x8f\x5c\x9f\x28\x7b\x26\x93\xa1\xa1\x8a\x59\xdf\x8b\x1a\xbd\xa3\xf0\xe8\x6f\x6d\x04\xf0\x1e\x22\x1a\xc8\xe0\x05\x75\x41\x81\x27\x5c\x0b\x48\xed\x16\x8f\x4c\x3b\x7d\x80\x0b\x99\xd1\x35\x6c\x6d\xe5\x30\xa3\x10\x84\x0d\xd9\x51\x09\x4a\x2d\x94\x71\x41\x39\x56\x78\xd5\xbe\xf2\x14\xce\xa5\x52\x48\xf4\x85\xb3\xff\x40\x19\xb8\x4b\x9d\xa8\xe6\x0c\x72\xb8\x28\x88\x6c\x29\xd9\x22\xa7\x90\xc1\x18\xea\x90\x88\x7d\x71\x8c\x30\x8a\xb7\x92\x0f\xa0\xb3\xce\x1f\x0f\xcc\x14\xa0\x4f\xf6\x83\xf2\x1c\x36\xc6\xd8\x6e\xf6\x89\xf8\xa0\xaf\x79\xd6\xe3\x96\x60\xee\x7a\xd0\x71\x64\x4c\x68\xb4\x0f\x21\xc5\xb9\x2f\x82\x64\x9a\xb2\x44\x2c\x6f\xe3\x88\xc4\xd1\xaa\x79\xf5\xb3\xd9\xe0\x07\xef\x12\x27\xf0\xc0\x68\x3a\x24\x14\xaf\x0e\xf8\x70\xda\x28\x4c\xb2\x54\x45\x51\xe0\x20\x63\x48\xc3\xfb\x5b\xad\x9d\xf5\x13\x32\x34\xe4\x79\xce\x27\xca\x50\xc8\x58\xdc\x1f\x17\xe8\x93\x02\x94\x29\xc8\x09\x36\xbe\x93\x0e\xc6\x6b\xab\xf3\xac\x29\x0b\xf8\xe9\xc5\x0c\x5f\xc8\x87\x55\x9c\x11\x4f\x67\x61\x1c\x05\x2d\x2e\x95\x5d\x14\x50\xb5\x04\x77\xca\x0b\xf4\x3e\xac\xa1\x4c\x05\x36\xb2\x02\x05\x2f\x0b\x39\x62\x5f\x7e\x2f\x20\x9f\x55\xa0\x4c\x34\x38\xee\xbf\x0d\x4d\x52\xeb\x5a\xd4\xfe\x87\x83\x18\xbe\xf6\x9e\x18\x9f\xee\x1b\x21\x4f\x75\x19\xc2\x94\x8e\x6b\x5b\x4a\xd4\x53\xa6\x3a\x95\xc9\x7c\x54\x05\x23\x98\xb4\x0d\x9a\xd0\x09\x06\x44\xb3\x15\x7f\x03\xa6\x66\xa0\xf7\xcb\x35\x23\x5f\x77\xa0\x11\xf7\xba\x55\xd5\x6a\x87\x6e\x47\x8f\xd2\x4e\x9a\x0b\xf7\xc0\x43\x48\x54\x7f\x43\x9b\xb5\xa4\xe1\x41\xfc\x04\x55\xb1\x80\xdd\x4b\x19\xcc\x28\xdb\xf9\x32\x3a\x78\x88\x93\x5d\xd7\x3b\xa1\x01\x72\x5e\x1d\x6a\x82\x9c\xed\x69\xd0\x93\x44\xec\xf5\x53\x86\xf6\xfb\x9d\xd3\xba\x2b\xaa\x0d\x62\xc2\xfb\xa1\xdb\x9d\xf8\xa6\xed\xcf\xc2\xb4\x8b\x3e\x11\xbb\xdb\x3b\xbe\x6b\xd3\x50\xdf\x6e\x44\x07\x59\xab\xdd\x69\x9d\x26\xec\xe9\x32\xab\xf5\x7a\x53\x14\x3f\xec\x8f\xfb\xe5\x57\x8b\x65\xcd\xe0\x09\xba\x3c\xfc\x84\x8e\xe2\x42\x42\xed\x6c\xf6\x7a\xe5\x82\xa0\x39\xf6\x99\xa8\xa3\xdb\x95\xbe\xcc\x7b\xdf\xe3\x9e\x05\xd3\xc7\xaf\x7f\xff\x0f\x00\x00\xff\xff\xd1\xfe\xa5\xd6\xd7\x06\x00\x00")

func dataVegetablesJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataVegetablesJson,
		"data/vegetables.json",
	)
}

func dataVegetablesJson() (*asset, error) {
	bytes, err := dataVegetablesJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/vegetables.json", size: 1751, mode: os.FileMode(436), modTime: time.Unix(1523619650, 0)}
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
	"data/fruits.json":               dataFruitsJson,
	"data/herbs.json":                dataHerbsJson,
	"data/ingredient_corpus.txt":     dataIngredient_corpusTxt,
	"data/ingredient_densities.json": dataIngredient_densitiesJson,
	"data/top_5k.txt":                dataTop_5kTxt,
	"data/vegetables.json":           dataVegetablesJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"fruits.json":               &bintree{dataFruitsJson, map[string]*bintree{}},
		"herbs.json":                &bintree{dataHerbsJson, map[string]*bintree{}},
		"ingredient_corpus.txt":     &bintree{dataIngredient_corpusTxt, map[string]*bintree{}},
		"ingredient_densities.json": &bintree{dataIngredient_densitiesJson, map[string]*bintree{}},
		"top_5k.txt":                &bintree{dataTop_5kTxt, map[string]*bintree{}},
		"vegetables.json":           &bintree{dataVegetablesJson, map[string]*bintree{}},
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
