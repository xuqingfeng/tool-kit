package pkg

import (
	"os"
)

const (
	dirMode = 0755
)

func CreateFile(name string) error {

	_, err := os.Create(name)
	return err
}

func CreateFileWithContent(name, content string) error {

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = f.WriteString(content)
	defer f.Close()

	return err
}

func CreateDir(name string) error {

	return os.Mkdir(name, dirMode)
}

func CreateSoftLink(oldname, newname string) error {

	return os.Symlink(oldname, newname)
}
