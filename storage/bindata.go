// Code generated by go-bindata.
// sources:
// init.sql
// DO NOT EDIT!

package storage

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

var _initSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xc1\x6e\xbb\x30\x0c\xc6\xef\x79\x0a\x1f\xff\x95\xfe\x6f\xd0\x53\xba\xba\x9b\xb5\x10\x2a\x70\x35\xd8\x25\x8a\xc0\x9a\x38\x14\xaa\xe0\x4d\x7b\xfc\xa9\x6c\x1d\xa3\xc2\xb7\xd8\xfa\x7d\xf1\xf7\xf9\xa1\x40\xcb\x08\x6c\x77\x0e\x81\x0e\xe0\x73\x06\xac\xa8\xe4\x12\x46\x8d\x2a\xf0\xcf\x00\x74\x2d\x2c\x6b\x47\x8f\x25\x16\x64\xdd\x5d\x7f\xc2\xfd\xc9\x39\x38\x16\x94\xd9\xa2\x86\x67\xac\xff\x1b\x80\x66\xe8\x55\x3e\x35\xfc\x2a\x31\x56\x7c\x0f\x2f\x14\x26\x2a\x49\x54\x69\x43\xd4\x1b\x45\x19\x96\x6c\xb3\x23\xbc\x10\x3f\x4d\x4f\x78\xcd\x3d\xce\xff\xee\xf1\x60\x4f\x8e\xa1\x79\x4f\x49\x7a\x0d\xda\x9d\x65\xd4\x78\xbe\x5c\xf5\xbe\x1d\xfd\x29\xf2\xeb\x4b\x2c\xb6\x68\xa3\xc6\xa5\xfb\x9a\xd1\x5e\x27\xf2\x21\xbd\xae\x4e\x2e\x69\x68\x64\x1c\xbb\xfe\x2d\x48\x4a\x43\x9a\xfc\x9a\xcd\xd6\x98\x9f\xc0\xc9\xef\xb1\x5a\x0b\x3c\xcc\x51\x19\x80\xdc\xdf\xce\x30\xb7\x37\xdb\xaf\x00\x00\x00\xff\xff\x82\x8d\xe6\x33\xb5\x01\x00\x00")

func initSqlBytes() ([]byte, error) {
	return bindataRead(
		_initSql,
		"init.sql",
	)
}

func initSql() (*asset, error) {
	bytes, err := initSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init.sql", size: 437, mode: os.FileMode(436), modTime: time.Unix(1534487688, 0)}
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
	"init.sql": initSql,
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
	"init.sql": &bintree{initSql, map[string]*bintree{}},
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