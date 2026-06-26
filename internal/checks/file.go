package checks

import (
	"fmt"
	"os"

	"github.com/allenbiji/clone-sage/internal/model"
	"github.com/allenbiji/clone-sage/internal/registry"
)

type FileCheck struct {
	Path string
}

//execute method for file check
func (f *FileCheck) Execute() error {
	info, err := os.Stat(f.Path)
	if os.IsNotExist(err){
		return fmt.Errorf("File does not exist: %s", f.Path)
	}

	if err != nil {
		return fmt.Errorf("Error accessing file %s: %w", f.Path, err)
	}

	if info.IsDir(){
		return fmt.Errorf("Expected a file, returned a directory at %s", f.Path)
	}

	return nil
}

//creates file check factory
func buildFileCheck(cfg model.CheckConfig) (registry.Check, error){
	path, ok := cfg.Options["path"]
	if(!ok || path == ""){
		return nil, fmt.Errorf("file_exists check requires a 'path' option")
	}
	return &FileCheck{
		Path: path,
	}, nil
}

//Registers file exists check in registry
func init(){
	registry.Register(model.TypeFileExists, buildFileCheck)
}