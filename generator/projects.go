package main

import (
	"log"
	"strings"
)

type Projects struct {
	Data map[string]*Project `json:"data"`
}

func FetchProjects(dataSet *DataSet, projects *Projects, custom *Projects) {
	if projects.Data == nil {
		projects.Data = make(map[string]*Project)
	}

	for _, v := range dataSet.Data {
		for _, p := range v {
			if _, ok := projects.Data[p]; ok {
				log.Printf("Skipped %s: cached project.\n", p)
				continue
			}
			if _, ok := custom.Data[p]; ok {
				log.Printf("Skipped %s: user-defined project.\n", p)
				continue
			}

			log.Printf("Fetching %s from Modrinth...\n", p)
			pj := FetchProject(p)
			if !contains(dataSet.Version, pj.GameVersions) {
				log.Fatalf("%s does not support version %s", p, dataSet.Version)
			}
			projects.Data[p] = pj
		}
	}
}

func contains(what string, where []string) bool {
	for _, v := range where {
		// ignore revisions at the end
		if strings.Index(v, what) == 0 {
			return true
		}
	}
	return false
}
