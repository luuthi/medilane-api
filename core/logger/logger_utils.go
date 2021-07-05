package logger

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func _formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	if len(arr) == 0 {
		return path
	}
	return arr[len(arr)-1]
}

// try create folder
func createFile(output string) (string, error) {
	workingDir, _ := os.Getwd()

	absPath := filepath.Join(workingDir, output)
	dirPath := filepath.Dir(absPath)

	// create new if not exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return absPath, nil
}

// split method called to only file name instead of full path
func callerPrettyfier(frame *runtime.Frame) (function string, file string) {
	//return "", fmt.Sprintf("%s:%d", _formatFilePath(frame.File), frame.Line)
	s := strings.Split(frame.Function, ".")
	funcName := frame.Function
	if len(s) >= 1 {
		funcName = s[len(s)-1]
	}
	_, filename := path.Split(frame.File)
	return funcName, fmt.Sprintf("%s:%d", filename, frame.Line)
}
