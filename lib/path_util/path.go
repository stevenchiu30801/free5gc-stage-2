package path_util

import (
	"gofree5gc/lib/path_util/logger"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Gofree5gcPath ...
/*
 * Author: Roger Chu aka Sasuke
 *
 * This package is used to locate the root directory of gofree5gc project
 * Compatible with Windows and Linux
 *
 * Please import "gofree5gc/lib/path_util"
 *
 * Return value:
 * A string value of the relative path between the working directory and the root directory of the gofree5gc project
 *
 * Usage:
 * path_util.Gofree5gcPath("your file location starting with gofree5gc")
 *
 * Example:
 * path_util.Gofree5gcPath("gofree5gc/abcdef/abcdef.pem")
 */
func Gofree5gcPath(path string) string {
	rootCode := strings.Split(path, "/")[0]
	rootPath := ""

	pwd, err := os.Getwd()
	if err != nil {
		logger.PathLog.Fatal(err)
	}

	currentPath := filepath.Clean(pwd)
	if strings.Contains(currentPath, rootCode) {
		returnPath, ok := FindRoot(currentPath, rootCode, "go.mod")
		if ok {
			rootPath = returnPath
		} else {
			returnPath, ok := FindRoot(currentPath, rootCode, "lib")
			if ok {
				rootPath = returnPath
			}
		}
	}

	binPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		logger.PathLog.Fatal(err)
	}
	if strings.Contains(binPath, rootCode) {
		returnPath, ok := FindRoot(binPath, rootCode, "go.mod")
		if ok {
			rootPath = returnPath
		} else {
			returnPath, ok := FindRoot(binPath, rootCode, "lib")
			if ok {
				rootPath = returnPath
			}
		}
	}

	if rootPath == "" {
		_, fpath, _, _ := runtime.Caller(0)
		pkgFilePath := filepath.Clean(fpath)
		rootStringLoc := strings.LastIndex(pkgFilePath, rootCode)
		rootPath = pkgFilePath[:rootStringLoc]
		if !Exists(rootPath + filepath.Clean(path)) {
			rootPath = currentPath[:strings.LastIndex(currentPath, rootCode)]
		}
	}

	target := rootPath + filepath.Clean(path)

	location, err := filepath.Rel(currentPath, target)
	if err != nil {
		logger.PathLog.Fatal(err)
	}

	return location
}

func Exists(fpath string) bool {
	_, err := os.Stat(fpath)
	return !os.IsNotExist(err)
}

func FindRoot(path string, rootCode string, objName string) (string, bool) {
	rootPath := path
	loc := strings.LastIndex(rootPath, rootCode)
	for loc != -1 {
		rootPath = rootPath[:loc+len(rootCode)]
		if Exists(rootPath + filepath.Clean("/"+objName)) {
			return rootPath[:loc], true
		}
		rootPath = rootPath[:loc]
		loc = strings.LastIndex(rootPath, rootCode)
	}
	return "", false
}
