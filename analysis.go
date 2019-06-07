package fileanalysis

import (
	`github.com/flywithbug/file`
)

func CurrentDirectory() (string, error) {
	return file.CurrentDir()
}

func AbsPath(dir string,p string)string  {
	return file.AbsPath(dir,p)
}

