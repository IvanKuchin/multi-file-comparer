package compare

import (
	"log"
	"path/filepath"

	configreader "github.com/ivankuchin/multi-file-comparer/config-reader"
)

var comparer Comparer

func compareSingleFile(rootDir string, subfolders []string, file string) (result string, err error) {

	for _, folder := range subfolders[1:] {
		temp_result := comparer.Compare(filepath.Join(rootDir, subfolders[0], file), filepath.Join(rootDir, folder, file))
		if len(temp_result) > 0 {
			result = result + temp_result
		}
	}

	return
}

func compareAllFiles(cfg configreader.GroupConfig) ([]string, error) {
	result := make([]string, 0)

	for _, file := range cfg.Files {
		single_file_result, err := compareSingleFile(cfg.Rootdir, cfg.Subfolders, file)
		if err != nil {
			log.Println(err)
			return result, err
		}
		if len(single_file_result) > 0 {
			result = append(result, single_file_result)
		}
	}

	return result, nil
}

func Do(cfg configreader.GroupConfig) ([]string, error) {
	var err error
	comparer, err = NewComparer("bytes.Equal")
	if err != nil {
		log.Print(err)
		return []string{}, err
	}

	return compareAllFiles(cfg)
}
