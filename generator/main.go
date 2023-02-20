package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

type DataSet struct {
	Version string              `json:"version"`
	Data    map[string][]string `json:"data"`
}

type TemplateData struct {
	DataSet  *DataSet
	Projects *Projects
}

func main() {
	templatePath := flag.String("template", "template.md", "Specifies the template file path")
	dataPath := flag.String("data", "../data/mods.default.json", "Specifies the data file path")
	cachePath := flag.String("cache", "cache.json", "Specifies cache file")
	outPath := flag.String("out", "../README.md", "Specifies the output path")
	flag.Parse()

	var templateData string
	parseText(*templatePath, &templateData)

	dataSet := &DataSet{}
	parseJSON(*dataPath, dataSet)

	projects := &Projects{}
	parseJSONSilently(*cachePath, projects)
	FetchProjects(dataSet, projects)
	saveJSON(*cachePath, projects)

	out, err := os.Create(*outPath)
	if err != nil {
		log.Fatalln(err)
	}
	w := bufio.NewWriter(out)

	tmpl, err := template.New("generator").Parse(templateData)
	if err != nil {
		log.Fatalln(err)
	}
	if err = tmpl.Execute(w, &TemplateData{
		DataSet:  dataSet,
		Projects: projects,
	}); err != nil {
		log.Fatalln(err)
	}

	if err = w.Flush(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done.")
}

func parseText(path string, target *string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	*target = string(bytes)
}

func parseJSON(path string, target interface{}) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	if err = json.Unmarshal(bytes, target); err != nil {
		log.Fatalln(err)
	}
}

func parseJSONSilently(path string, target interface{}) {
	bytes, _ := os.ReadFile(path)
	_ = json.Unmarshal(bytes, target)
}

func saveJSON(path string, target interface{}) {
	data, err := json.Marshal(target)
	if err != nil {
		log.Fatalln(err)
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		log.Fatalln(err)
	}
}
