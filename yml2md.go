package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

const projectReadmeTmplFilename = "project-readme.mdt"
const projectReadmeTmplPath = ".templates/project-readme.mdt"
const manifestFile = "manifest.yml"

type EnvVariable struct {
	Name        string `yaml:"name"`
	Default     string `yaml:"default_value"`
	Example     string `yaml:"example"`
	Description string `yaml:"description"`
}

type Link struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Container struct {
	Name        string        `yaml:"name"`
	Slug        string        `yaml:"slug"`
	Category    string        `yaml:"category"`
	Icon        string        `yaml:"icon"`
	ServesHttp  bool          `yaml:"serves_http"`
	Color       string        `yaml:"color"`
	Description string        `yaml:"description"`
	Links       []Link        `yaml:"links,flow"`
	Variables   []EnvVariable `yaml:"variables,flow"`
	Tags        []string      `yaml:"tags,flow"`
}

type Manifest struct {
	Name        string        `yaml:"name"`
	Status      string        `yaml:"status"`
	LastUpdated string        `yaml:"updated"`
	Variables   []EnvVariable `yaml:"variables,flow"`
	Containers  []Container   `yaml:"containers,flow"`
}

func StatusToEmoji(status string) (icon string) {
	switch status {
	case "working":
		return ":heavy_check_mark:"
	case "draft":
		return ":building_construction:"
	case "obsolete":
		return ":file_cabinet:"
	default:
		return ":grey_question:"
	}
}

func WrapAsMarkdownCode(input string) (output string) {
	if len(input) == 0 {
		return ""
	}
	return fmt.Sprintf("`%s`", input)
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
	funcMap := template.FuncMap{
		"ToTitle": cases.Title(language.English).String,
		"ToEmoji": StatusToEmoji,
		"AsCode":  WrapAsMarkdownCode,
	}
	tmpl := template.Must(template.New(projectReadmeTmplFilename).Funcs(funcMap).ParseFiles(projectReadmeTmplPath))

	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	var errs []error

	for _, e := range entries {
		if e.IsDir() && string(e.Name()[0]) != "." {
			path := filepath.Join(".", e.Name(), manifestFile)
			if _, errStat := os.Stat(path); errStat == nil {
				if data, errRead := os.ReadFile(path); errRead == nil {
					mdFilePath := filepath.Join(".", e.Name(), "README.md")
					if mdFile, errCreate := os.Create(mdFilePath); errCreate == nil {
						errApply := applyTemplate(tmpl, data, mdFile)
						if errApply != nil {
							log.Printf("Error applying template from manifest file %s", path)
							errs = append(errs, errApply)
						} else {
							log.Printf("Created README.md file from manifest file %s", path)
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
			} else {
				log.Printf("Manifest file %s deos not exist, skipping", path)
			}
		}
	}

	if len(errs) > 0 {
		log.Fatal(errs)
	}
}
