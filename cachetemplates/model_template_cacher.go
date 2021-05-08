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

func NewCacher(templatesFolder string) (map[TemplatePath]HTML, error) {
	i := strings.Index(templatesFolder, "/")
	folder := templatesFolder

	if i != len(folder) {
		folder = folder + "/"
	}

	files, err := getFileNames(folder)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, errors.Errorf("cacher could not find any files in %s", templatesFolder)

	}

	res := make(map[TemplatePath]HTML, len(files))

	for _, file := range files {
		log.Println(file.Name())

		data, err := getTemplateHTML(folder + file.Name())
		if err != nil {
			return nil, errors.WithMessagef(err, "could not get template HTML for %s", file.Name())
		}

		res[TemplatePath(file.Name())] = data
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

func getFileNames(folder string) ([]fs.FileInfo, error) {
	return ioutil.ReadDir(folder)
}
