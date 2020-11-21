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

	info := bindataFileInfo{name: "0_init.down.sql", size: 263, mode: os.FileMode(436), modTime: time.Unix(1605802263, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __0_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x5f\x6f\x9b\x3c\x14\xc6\xef\xf9\x14\xe7\x2e\xa0\x97\xbe\xa2\x95\x26\x4d\xaa\x5a\xc9\xa3\xae\xca\x42\x4c\x6a\x60\x6b\xae\x2c\x96\xb8\x09\x2b\x31\x11\x38\xab\xb4\x4f\x3f\xd9\x98\x84\x2c\x61\x4d\xba\x29\xdc\xc1\xf9\xf3\xf8\xfc\x9e\x83\x7d\x8a\x51\x82\x21\x41\x9f\x42\x0c\xc1\x3d\x90\x28\x01\xfc\x14\xc4\x49\x0c\xaf\x59\x51\x70\x59\x5b\xb6\x05\x00\x90\xcf\xa0\x7d\x62\x4c\x03\x14\xc2\x98\x06\x23\x44\x27\x30\xc4\x13\x57\xa7\x88\x6c\xc9\x4d\xca\x17\x44\xfd\x07\x44\xed\x2b\xcf\x73\x00\x74\x57\x92\x86\xa1\xd5\xb6\xf0\x23\x12\x27\x14\x05\x24\x81\xe9\x82\x4f\x5f\x58\x23\xc6\x54\x0b\x26\xcb\x92\xd5\x8b\xb2\x92\xe0\x3f\x60\x7f\x08\x36\x14\x5c\xcc\xe5\xc2\x56\x51\x07\x6e\xe1\x03\x38\x8d\xe2\xb7\xac\xc8\xc4\x54\x8b\x92\x74\x84\x69\xe0\xdb\x97\x1f\x5d\xb8\x72\x36\x8a\x70\x87\xef\x51\x1a\x26\xe0\xfd\xef\xf5\xab\x9b\x3e\x4c\x94\x92\x09\x3e\xcf\x64\xfe\x83\xb7\xe2\xad\xc6\xed\x8d\xea\x61\x84\xa7\x15\xcf\x24\x9f\xb1\x4c\x82\xcc\x97\xbc\x96\xd9\x72\x25\x7f\x42\x67\xd4\x8d\xb0\x28\x5f\x6d\x53\xb5\x5e\xcd\x4e\xaa\xb2\x9c\x6b\xcb\xb2\x8c\x45\x29\x09\x1e\x53\x0c\x01\xb9\xc3\x4f\x90\x92\x47\xf6\x15\x85\x21\x4e\x62\x46\xd0\x08\x43\x44\x5a\xbf\xa0\xc1\x74\x6d\x59\x77\x34\x1a\x43\x32\x19\x6b\x63\x8d\xa9\x29\x09\x92\xeb\xb6\xa5\x8e\xa9\x2f\x80\x62\xc0\x24\x1d\x81\x3d\x98\xf1\x55\x59\xe7\x72\xe0\x0e\x64\x95\x89\xfa\x99\x57\x83\xee\x29\xfa\x17\x85\x99\xca\x13\x16\xa6\x29\xd4\x29\xca\x8b\xce\xb3\xe1\x41\xf1\x3d\xa6\x98\xf8\x38\xde\xce\x97\xcf\x0c\xcf\x6c\x59\xae\x45\x53\xdf\x63\xff\xdf\xb8\xd5\x2e\x9f\x51\xb9\x55\xf6\x43\xe3\xc9\xdb\x30\x5a\x78\xfb\x30\xde\xa0\xc1\x9e\xab\x72\xf9\x3e\x1c\xad\x74\xf9\x0f\x78\x1e\x0d\xf4\x6c\x44\xcb\x15\xaf\xd8\xf7\x72\x5d\x89\xac\x38\xdf\x8a\x69\xd5\x3a\x9f\x8b\xfe\xfa\xcd\x0d\xe3\xc2\xc5\x05\x78\x37\xff\xc1\xe5\xcd\xc5\x19\xf6\x73\x2d\x72\xe3\x55\xf3\x0f\xef\x1f\x6d\x87\xfa\x76\x92\x80\x80\xed\xb9\x70\xe9\xbc\x07\x3d\x2b\x72\xf1\xf2\xdb\x3f\xde\x07\x3f\x17\x4c\x65\x28\x70\xfd\xb4\x77\xba\x77\xc9\xaf\xa5\xaa\x3e\xb5\xb8\x3b\x52\xef\x95\x19\x8d\x31\x65\x9f\xa3\x94\x12\x14\xb2\x30\x20\xc3\x78\x7b\x81\x1e\x98\x16\x6c\x3d\x89\x6b\xce\x74\x1c\x34\x65\xcf\x21\x58\x7f\xc4\xd5\x78\xaa\xed\xdc\x35\x51\x29\xeb\xe0\x81\x98\xc1\x6c\x36\x74\xbf\x4e\x07\x77\x62\x47\x41\x52\x4a\x7b\x70\xb6\x53\x69\x28\xea\xd5\x85\x0e\x9d\xe6\xc3\x86\xd3\xaf\x00\x00\x00\xff\xff\x17\x39\x33\x25\x5f\x08\x00\x00")

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

	info := bindataFileInfo{name: "0_init.up.sql", size: 2143, mode: os.FileMode(436), modTime: time.Unix(1605933970, 0)}
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
