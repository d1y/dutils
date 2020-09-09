// Author: d1y<chenhonzhou@gmail.com>
// fs package like nodejs `fs`
// I tried to implement most of its API
// http://nodejs.cn/api/fs.html

package fs

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
)

// ReadFile2String read file to string
//
// https://gobyexample.com/reading-files
//
// https://golangcode.com/read-a-files-contents/
//
// https://stackoverflow.com/questions/13514184/how-can-i-read-a-whole-file-into-a-string-variable
func ReadFile2String(filename string) (string, error) {
	var msg = fmt.Sprintf("read file is error, path: %v", filename)
	if !IsFile(filename) {
		return "", errors.New(msg)
	}
	bts, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New(msg)
	}
	var r = string(bts)
	return r, nil
}

// CopyDirectory copy dir
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

// Copy copy file
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

// Exists check file/dir exists
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateIfNotExists auto create files
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

// CopySymLink copy link file(...I'm not really sure)
//
// https://stackoverflow.com/questions/51779243/copy-a-folder-in-go
func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}

// IsDir 判断所给路径是否为文件夹
//
// https://m.yisu.com/zixun/143352.html
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
//
// https://m.yisu.com/zixun/143352.html
func IsFile(path string) bool {
	return !IsDir(path)
}
