package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

// Project https://github.com/modrinth/labrinth/blob/master/src/models/projects.rs
type Project struct {
	ID               string      `json:"id"`
	Slug             string      `json:"slug"`
	ProjectType      string      `json:"project_type"`
	Team             string      `json:"team"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	Body             string      `json:"body"`
	Published        time.Time   `json:"published"`
	Updated          time.Time   `json:"updated"`
	Approved         time.Time   `json:"approved"`
	Status           string      `json:"status"`
	RequestedStatus  interface{} `json:"requested_status"`
	ModeratorMessage interface{} `json:"moderator_message"`
	License          struct {
		ID   string      `json:"id"`
		Name string      `json:"name"`
		URL  interface{} `json:"url"`
	} `json:"license"`
	ClientSide           string        `json:"client_side"`
	ServerSide           string        `json:"server_side"`
	Downloads            int           `json:"downloads"`
	Followers            int           `json:"followers"`
	Categories           []string      `json:"categories"`
	AdditionalCategories []interface{} `json:"additional_categories"`
	GameVersions         []string      `json:"game_versions"`
	Loaders              []string      `json:"loaders"`
	Versions             []string      `json:"versions"`
	IconURL              string        `json:"icon_url"`
	IssuesURL            string        `json:"issues_url"`
	SourceURL            string        `json:"source_url"`
	WikiURL              string        `json:"wiki_url"`
	DiscordURL           string        `json:"discord_url"`
	DonationUrls         []interface{} `json:"donation_urls"`
	Gallery              []interface{} `json:"gallery"`
	FlameAnvilProject    interface{}   `json:"flame_anvil_project"`
	FlameAnvilUser       interface{}   `json:"flame_anvil_user"`
	Color                int           `json:"color"`
}

func FetchProject(name string) *Project {
	res, err := http.Get("https://api.modrinth.com/v2/project/" + name)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	project := &Project{}
	if err = json.Unmarshal(body, project); err != nil {
		log.Fatalln(err)
	}
	return project
}
