package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _bash_example_bash = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcc\x31\x4e\xc5\x30\x0c\x80\xe1\xb9\x3e\x85\x15\x31\xc0\x90\x97\xfd\x4d\xc0\xc4\x35\xd2\xe2\x26\x46\x49\x1c\xd9\x09\x1d\x68\xef\x8e\x10\x03\xac\xbf\x7e\x7d\x46\x03\x3d\x09\x76\xee\xb4\x47\x2e\x00\x99\x4a\x11\xbf\x46\xcb\x8f\x4f\xf8\x05\x0b\x6d\x59\xd0\xbd\xfd\x54\x3c\x44\xcb\x3b\xee\x2a\x15\x5f\xa3\x65\x07\x17\x40\x8d\xdc\xfe\x9f\x2f\x9a\x66\xa5\x36\xec\xee\xd0\x3d\x3c\x3b\x58\xfe\x44\x3c\x51\xe9\x93\xd4\x08\x96\x6d\x6a\x41\x6f\x98\xc7\xe8\x76\x0f\x21\x76\xbe\x25\x1e\x79\xae\xb7\x4d\x6a\x50\xea\x62\xa1\xab\x24\xe5\x59\x43\xfa\x05\x48\xf1\xc4\x0f\x93\xe6\xbb\x70\x1b\xa4\x18\xe4\x68\xa4\xa1\x48\xe2\x06\x17\x7c\x07\x00\x00\xff\xff\xca\x1c\xba\x45\xd0\x00\x00\x00")

func bash_example_bash() ([]byte, error) {
	return bindata_read(
		_bash_example_bash,
		"bash/example.bash",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"bash/example.bash": bash_example_bash,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"bash/example.bash": &_bintree_t{bash_example_bash, map[string]*_bintree_t{
	}},
}}
