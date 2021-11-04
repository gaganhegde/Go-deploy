package files

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v2"
)

//CreateFiles writes the files to the desired path on the local file system.
func CreateFiles(path string, file map[string]interface{}) {
	finalPath, err := homedir.Expand(path)
	if err != nil {
		fmt.Errorf("The following error has been seen", err)
	}

	fmt.Println("this is the path", finalPath)
	fmt.Println("this is the file path", path)
	for path, bits := range file {
		endpath := filepath.Join(finalPath, path)
		data, _ := yaml.Marshal(bits)
		ioutil.WriteFile(endpath, data, 0755)
	}

}
