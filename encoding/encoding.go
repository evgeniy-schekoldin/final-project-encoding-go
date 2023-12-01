package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var dockerCompose models.DockerCompose
	input, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(input, &dockerCompose); err != nil {
		return err
	}
	output, err := yaml.Marshal(dockerCompose)
	if err != nil {
		return err
	}
	os.WriteFile(j.FileOutput, output, 0644)
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var dockerCompose models.DockerCompose
	input, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(input, &dockerCompose); err != nil {
		return err
	}
	output, err := json.Marshal(dockerCompose)
	if err != nil {
		return err
	}
	os.WriteFile(y.FileOutput, output, 0644)
	return nil
}
