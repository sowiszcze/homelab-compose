package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

const projectReadmeTmplFile = ".templates/project-readme.mdt"
const manifestFile = "manifest.yml"

type EnvVariable struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
}

type Container struct {
	Name           string        `yaml:"name"`
	Title          string        `yaml:"title"`
	Icon           string        `yaml:"icon"`
	Color          string        `yaml:"color"`
	Webpage        string        `yaml:"webpage"`
	Repository     string        `yaml:"repository"`
	DockerImage    string        `yaml:"docker_image"`
	ComposeExample string        `yaml:"compose_example"`
	Description    string        `yaml:"description"`
	Variables      []EnvVariable `yaml:"variables,flow"`
	Tags           []string      `yaml:"tags,flow"`
}

type Manifest struct {
	Name        string      `yaml:"name"`
	Status      string      `yaml:"status"`
	LastUpdated string      `yaml:"updated"`
	Containers  []Container `yaml:"containers,flow"`
}

func applyTemplate(tmpl *template.Template, data []byte, wr io.Writer) (err error) {
	var manifest Manifest
	errYaml := yaml.Unmarshal(data, &manifest)
	if errYaml != nil {
		return errYaml
	}
	errTemplate := tmpl.Execute(wr, manifest)
	if errTemplate != nil {
		return errTemplate
	}
	return
}

func main() {
	tmpl := template.Must(template.ParseFiles(projectReadmeTmplFile))

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var errs []error

	for _, e := range entries {
		log.Printf("Checking if %s is a directory", e.Name())
		if e.IsDir() {
			path := filepath.Join(".", e.Name(), manifestFile)
			log.Printf("Checking if manifest file %s exists", path)
			if _, errStat := os.Stat(path); errStat == nil {
				if data, errRead := os.ReadFile(path); errRead == nil {
					mdFilePath := filepath.Join(".", e.Name(), "README.md")
					if mdFile, errCreate := os.Create(mdFilePath); errCreate == nil {
						errApply := applyTemplate(tmpl, data, mdFile)
						if errApply != nil {
							log.Printf("Error applying template from manifest %s", path)
							errs = append(errs, errApply)
						}
					} else {
						log.Printf("Could not open output file %s", mdFilePath)
						errs = append(errs, errCreate)
					}
				} else {
					log.Printf("Error reading file %s, skipping", path)
					errs = append(errs, errRead)
				}
			} else if !errors.Is(errStat, os.ErrNotExist) {
				log.Printf("Error checking file %s, skipping", path)
				errs = append(errs, errStat)
			}
		}
	}

	if len(errs) > 0 {
		log.Fatal(errs)
	}
}
