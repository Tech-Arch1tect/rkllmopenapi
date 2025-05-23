package model

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Tech-Arch1tect/rkllmopenai/config"
)

var (
	ModelList = []Model{}
)

func RefreshModelList() {
	ModelList = []Model{}
	root := config.C.StoragePath
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".rkllm") {
			ModelList = append(ModelList, Model{ModelName: d.Name(), ModelPath: path, ModelDir: filepath.Dir(path)})
		}
		return nil
	})
}
