// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// 0_init.down.sql
// 0_init.up.sql
package migrations

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

// Mode return file modify time
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

var __0_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4f\xcc\xc9\x49\x2d\x89\x2f\xcd\xcb\x2c\x89\xcf\xc9\xcc\xcb\x2e\xb6\xe6\xc2\xa7\x2e\xbf\x20\xb5\x28\x3e\x2b\xbf\xb4\x28\x2f\x31\x87\x54\xf5\xf8\x55\xa6\xa4\x16\xe4\x17\x67\x96\x10\x30\xaf\xa4\x28\x31\xaf\x38\x2d\xb5\x08\xaf\xaa\x62\x6b\x2e\xa8\x74\x64\x00\xb2\x6c\xa8\x9f\x67\x88\x35\x17\x20\x00\x00\xff\xff\x7a\x13\xa2\x53\x07\x01\x00\x00")

func _0_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__0_initDownSql,
		"0_init.down.sql",
	)
}

func _0_initDownSql() (*asset, error) {
	bytes, err := _0_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0_init.down.sql", size: 263, mode: os.FileMode(436), modTime: time.Unix(1605779088, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\xdf\x6e\x9b\x30\x18\xc5\xef\x79\x8a\x73\x17\xd0\xc8\x44\x2a\x4d\x9a\x14\xb5\x92\x47\x1c\xd5\x0b\x71\x52\x03\x5b\x73\x85\x58\x43\x3b\x36\x62\x22\x42\x54\x69\x4f\x3f\xd9\x40\x42\x9a\xb1\xa6\xd9\x14\xee\xb0\xbf\xf3\xfd\xf9\x1d\xdb\xae\xa0\x24\xa0\x08\xc8\x27\x8f\x82\x8d\xc1\x67\x01\xe8\x3d\xf3\x03\x1f\xcf\x71\x96\x25\xe5\xc6\x30\x0d\x00\x48\x97\x68\x3e\x9f\x0a\x46\x3c\xcc\x05\x9b\x12\xb1\xc0\x84\x2e\x6c\x1d\x22\xe3\x55\x52\x87\x7c\x21\xc2\xbd\x25\xc2\xbc\x72\x1c\x0b\xd0\x59\x79\xe8\x79\x55\xdc\xb7\x38\x8b\xe5\x83\x0e\xe5\xe1\x94\x0a\xe6\x9a\x83\x8f\x36\xae\xac\x5d\x1c\x46\x74\x4c\x42\x2f\x80\xf3\xde\xa9\x34\x0f\x45\x12\x97\xc9\x32\x8a\x4b\x94\xe9\x2a\xd9\x94\xf1\x6a\x5d\xfe\x42\x2b\xf7\x4e\x23\xf3\x67\xd3\xaa\x54\xdb\xf5\xf2\x0c\x95\x7b\x4b\xdd\x09\xcc\xa6\xcd\x1b\xd5\xc5\xe1\x16\xb2\x44\x3e\x95\xdf\x4d\x35\xb1\x85\x1b\x7c\x80\x65\x58\x43\xc3\x30\x6a\x9c\x21\x67\x77\x21\x05\xe3\x23\x7a\x8f\x90\xdf\x45\x5f\x89\xe7\xd1\xc0\x8f\x38\x99\x52\xcc\x78\xc3\x16\x55\x86\xa1\x61\x8c\xc4\x6c\x8e\x60\x31\xd7\x26\xd4\x06\x84\x9c\x05\xc3\x26\xa5\xde\x53\x2b\x20\x3e\x28\x0f\xa7\x30\x7b\xcb\x64\x9d\x6f\xd2\xb2\x67\xf7\xca\x22\x96\x9b\xc7\xa4\xe8\xb5\xbb\xe8\x36\x35\xaa\x95\x6f\x30\xb7\x12\xea\x10\xc6\x03\xb4\xbe\x1d\x4a\x41\xc7\x54\x50\xee\x52\x7f\x3f\x5f\xba\xac\xc9\xc5\xab\x7c\x2b\x2b\x7d\x87\xe9\xff\x62\x74\xe3\x4b\x5d\x45\x5b\x56\x7b\xf2\x3a\x8c\x06\xde\x31\x8c\x57\x68\x44\x8f\x45\xbe\x3a\x0f\x47\x53\x3a\xff\x0f\x3c\x4f\x06\x7a\x31\xa2\xf9\x3a\x29\xa2\x1f\xf9\xb6\x90\x71\x76\xb9\x23\xa6\xab\x6e\xd2\x27\xd9\xad\xdf\xbd\x2b\x36\xfa\x7d\x38\xd7\xef\x30\xb8\xee\x5f\xe0\x7c\x6e\x65\x5a\x7b\x55\xdd\xe1\xe3\xd6\x0e\xa8\xef\x27\x61\x1c\xa6\x63\x63\x60\x9d\x83\x3e\xca\x52\xf9\xf3\xc5\x1d\xef\x82\x9f\xca\x48\x45\x28\x70\xdd\xb4\x0f\xb2\xb7\xc9\x6f\x4b\xa5\x7e\xab\xb8\x3d\x52\xe7\x93\x39\x9b\x53\x11\x7d\x9e\x85\x82\x13\x2f\xf2\x18\x9f\xf8\xfb\x07\xf4\x0f\xd3\xc2\xd4\x93\xd8\x75\x4f\xa7\x41\x53\xf6\xbc\x84\xf5\x17\x50\xda\x4d\xed\xe3\xa1\x7b\xaa\x64\xd7\x5e\xc5\xb7\x4d\xa8\x13\xdd\x49\x58\x54\x89\x23\x1c\xfb\x39\x34\x06\xf5\x6b\xa3\xc5\xa3\x5a\xd8\x91\xf9\x1d\x00\x00\xff\xff\x98\xa8\xfc\x93\xfd\x07\x00\x00")

func _0_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__0_initUpSql,
		"0_init.up.sql",
	)
}

func _0_initUpSql() (*asset, error) {
	bytes, err := _0_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "0_init.up.sql", size: 2045, mode: os.FileMode(436), modTime: time.Unix(1605778961, 0)}
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
	"0_init.down.sql": _0_initDownSql,
	"0_init.up.sql":   _0_initUpSql,
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
	"0_init.down.sql": &bintree{_0_initDownSql, map[string]*bintree{}},
	"0_init.up.sql":   &bintree{_0_initUpSql, map[string]*bintree{}},
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
