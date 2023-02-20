package main

import "fmt"

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
				fmt.Printf("Skipped %s.\n", p)
			} else {
				fmt.Printf("Fetching %s from Modrinth...\n", p)
				projects.Data[p] = FetchProject(p)
			}
		}
	}
}
