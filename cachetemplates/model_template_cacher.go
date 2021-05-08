package cachetemplates

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type TemplatePath string
type HTML []byte

const templateFolderSlash = "../templates/"

func NewCacher(templatesFolders ...string) (map[TemplatePath]HTML, error) {
	var folders []string

	if len(templatesFolders) == 0 {
		folders = append(folders, templateFolderSlash)
	} else {
		folders = templatesFolders
	}

	allFilesPaths := []string{}

	for _, folder := range folders {
		pos := strings.LastIndex(folder, "/")

		if pos != len(folder)-1 {
			folder = folder + "/"
		}

		files, err := getFolderFileNames(folder)
		if err != nil {
			return nil, err
		}

		for _, file := range files {
			allFilesPaths = append(allFilesPaths, folder+file.Name())
		}
	}

	if len(allFilesPaths) == 0 {
		return nil, errors.Errorf("cacher could not find any files in %s", templatesFolders)

	}

	res := make(map[TemplatePath]HTML, len(allFilesPaths))

	for _, file := range allFilesPaths {
		log.Println(file)

		data, err := getTemplateHTML(file)
		if err != nil {
			return nil, errors.WithMessagef(err, "could not get template HTML for %s", file)
		}

		res[TemplatePath(file)] = data
	}

	return res, nil
}

func pathExists(path string) error {
	_, errPath := os.Stat(path)
	if os.IsNotExist(errPath) {
		return errPath
	}

	return nil
}

func getTemplateHTML(filePath string) ([]byte, error) {
	if err := pathExists(filePath); err != nil {
		return nil, err
	}

	return ioutil.ReadFile(filePath)
}

func getFolderFileNames(folder string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(folder)
}
