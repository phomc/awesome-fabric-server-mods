package main

import (
	"log"
)

type Projects struct {
	Data map[string]*Project `json:"data"`
}

func FetchProjects(dataSet *DataSet, projects *Projects) {
	if projects.Data == nil {
		projects.Data = make(map[string]*Project)
	}

	for _, v := range dataSet.Data {
		for _, p := range v {
			_, ok := projects.Data[p]
			if ok {
				log.Printf("Skipped %s.\n", p)
			} else {
				log.Printf("Fetching %s from Modrinth...\n", p)
				pj := FetchProject(p)
				if !contains(dataSet.Version, pj.GameVersions) {
					log.Fatalf("%s does not support version %s", p, dataSet.Version)
				}
				projects.Data[p] = pj
			}
		}
	}
}

// why we have to reinvent the wheel in go???
func contains(what string, where []string) bool {
	for _, v := range where {
		if v == what {
			return true
		}
	}
	return false
}
