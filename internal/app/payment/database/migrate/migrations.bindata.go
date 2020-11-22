// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// 0_init.down.sql
// 0_init.up.sql
// 1_hash.down.sql
// 1_hash.up.sql
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

var __0_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x94\x5f\x6f\x9b\x3c\x14\xc6\xef\xf9\x14\xe7\x0e\xd0\x4b\x5f\xd1\x4a\x93\x26\x55\xad\xe4\x51\x57\x65\x25\x26\x35\xb0\xb5\x57\x16\x4b\xdc\x84\x95\x98\x08\x9c\x55\xda\xa7\x9f\x6c\x4c\x42\x96\xb0\x26\xdd\x14\xee\xe0\xfc\x79\x7c\x7e\xcf\xc1\x01\xc5\x28\xc5\x90\xa2\x4f\x11\x86\xf0\x16\x48\x9c\x02\x7e\x0c\x93\x34\x81\xd7\xbc\x2c\xb9\x6c\x2c\xc7\x02\x00\x28\xa6\xd0\x3d\x09\xa6\x21\x8a\x60\x4c\xc3\x11\xa2\x4f\x70\x8f\x9f\x3c\x9d\x22\xf2\x05\x37\x29\x5f\x10\x0d\xee\x10\x75\x2e\x7c\xdf\x05\xd0\x5d\x49\x16\x45\x56\xd7\x22\x88\x49\x92\x52\x14\x92\x14\x26\x73\x3e\x79\x61\xad\x18\x53\x2d\x98\xac\x2a\xd6\xcc\xab\x5a\x42\x70\x87\x83\x7b\x70\xa0\xe4\x62\x26\xe7\x8e\x8a\xba\x70\x0d\x1f\xc0\x6d\x15\xbf\xe5\x65\x2e\x26\x5a\x94\x64\x23\x4c\xc3\xc0\x39\xff\xe8\xc1\x85\xbb\x56\x84\x1b\x7c\x8b\xb2\x28\x05\xff\x7f\x7f\x58\xdd\xf4\x61\xa2\x92\x4c\xf0\x59\x2e\x8b\x1f\xbc\x13\xef\x34\xae\xaf\x54\x0f\x23\x3c\xa9\x79\x2e\xf9\x94\xe5\x12\x64\xb1\xe0\x8d\xcc\x17\x4b\xf9\x13\x7a\xa3\xae\x85\x45\xf5\xea\x98\xaa\xd5\x72\x7a\x54\x95\xe5\x5e\x5a\x96\x65\x2c\xca\x48\xf8\x90\x61\x08\xc9\x0d\x7e\x84\x8c\x3c\xb0\xaf\x28\x8a\x70\x9a\x30\x82\x46\x18\x62\xd2\xf9\x05\x2d\xa6\x4b\xcb\xba\xa1\xf1\x18\xd2\xa7\xb1\x36\xd6\x98\x9a\x91\x30\xbd\xec\x5a\xea\x98\xfa\x02\x28\x01\x4c\xb2\x11\x38\xf6\x94\x2f\xab\xa6\x90\xb6\x67\xcb\x3a\x17\xcd\x33\xaf\x6d\xcf\x36\xee\x54\x4b\x5e\xb3\xef\xd5\xaa\x16\x79\x69\xf7\xcf\x36\xbc\x3e\xcc\xf4\x3b\x62\x8d\xda\x42\x9d\xa2\x1c\xea\x3d\x6b\x4a\x14\xdf\x62\x8a\x49\x80\x93\xcd\xd4\xc5\xd4\x50\xce\x17\xd5\x4a\xb4\xf5\x03\x4b\xf1\x37\x1e\x76\x2b\x69\x54\xae\xd5\x52\x40\xeb\xd4\xdb\x30\x3a\xa4\xbb\x30\xde\xa0\xc1\x9e\xeb\x6a\xf1\x3e\x1c\x9d\x74\xf5\x0f\x78\x1e\x0c\xf4\x64\x44\xfb\x2b\x79\xba\x15\xd3\xaa\x4d\x31\x13\xc3\xf5\xeb\x7b\xc7\x83\xb3\x33\xf0\xaf\xfe\x83\xf3\xab\xb3\x13\xec\xe7\x4a\x14\xc6\xab\xf6\xcf\xde\x3d\xda\x16\xf5\xcd\x24\x21\x01\xc7\xf7\xe0\xdc\x7d\x0f\x7a\x56\x16\xe2\xe5\xb7\x7f\x7c\x08\x7e\x21\x98\xca\x50\xe0\x86\x69\x6f\x75\xef\x93\x5f\x49\x55\x7d\x6c\x71\x7f\xa4\xc1\x8b\x34\x1e\x63\xca\x3e\xc7\x19\x25\x28\x62\x51\x48\xee\x93\xcd\xb5\xba\x67\x5a\x70\xf4\x24\x9e\x39\xd3\x61\xd0\x94\x3d\xfb\x60\xfd\x11\x57\xeb\xa9\xb6\x73\xdb\x44\xa5\xac\x83\x7b\x62\x06\xb3\xd9\xd0\xdd\x3a\x1d\xdc\x8a\x1d\x04\x49\x29\xed\xc0\xd9\x4c\xa5\xa1\xa8\x57\x0f\x7a\x74\xda\x0f\x6b\x4e\xbf\x02\x00\x00\xff\xff\x67\xad\x76\x66\x75\x08\x00\x00")

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

	info := bindataFileInfo{name: "0_init.up.sql", size: 2165, mode: os.FileMode(436), modTime: time.Unix(1605955593, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_hashDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4f\xcc\xc9\x49\x2d\x89\xcf\x2f\x48\x2d\x8a\xcf\xca\x2f\x2d\xca\x4b\xcc\xe1\x52\x50\x50\x50\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x48\x2c\xce\xb0\x06\x04\x00\x00\xff\xff\x8c\xd1\x7d\xc0\x3f\x00\x00\x00")

func _1_hashDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_hashDownSql,
		"1_hash.down.sql",
	)
}

func _1_hashDownSql() (*asset, error) {
	bytes, err := _1_hashDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_hash.down.sql", size: 63, mode: os.FileMode(436), modTime: time.Unix(1606053421, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_hashUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\x4f\xcc\xc9\x49\x2d\x89\xcf\x2f\x48\x2d\x8a\xcf\xca\x2f\x2d\xca\x4b\xcc\xe1\x52\x50\x50\x50\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\xf0\x74\x53\xf0\xf3\x0f\x51\x70\x8d\xf0\x0c\x0e\x09\x56\xc8\x48\x2c\xce\x50\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x33\xd1\xb4\x06\x04\x00\x00\xff\xff\x26\x47\x5d\x8b\x4e\x00\x00\x00")

func _1_hashUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_hashUpSql,
		"1_hash.up.sql",
	)
}

func _1_hashUpSql() (*asset, error) {
	bytes, err := _1_hashUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_hash.up.sql", size: 78, mode: os.FileMode(436), modTime: time.Unix(1606053424, 0)}
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
	"1_hash.down.sql": _1_hashDownSql,
	"1_hash.up.sql":   _1_hashUpSql,
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
	"1_hash.down.sql": &bintree{_1_hashDownSql, map[string]*bintree{}},
	"1_hash.up.sql":   &bintree{_1_hashUpSql, map[string]*bintree{}},
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
