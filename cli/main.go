package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var Survey = []*survey.Question{
	{
		Name: "Source",
		Prompt: &survey.Input{
			Message: "inform a file to save:",
			Suggest: func(toComplete string) []string {
				files, _ := filepath.Glob(toComplete + "*")
				return files
			},
		},
		Validate: survey.Required,
		Transform: func(ans interface{}) (newAns interface{}) {
			v, ok := ans.(string)
			if !ok {
				return nil
			}
			f, err := os.Open(v)
			if err != nil {
				return nil
			}
			defer f.Close()
			h := sha256.New()
			if _, err := io.Copy(h, f); err != nil {
				return nil
			}
			hash := hex.EncodeToString(h.Sum(nil))
			return fmt.Sprintf("p2p://%s", hash)
		},
	},
	{
		Name: "Name",
		Prompt: &survey.Input{Message: "Song name"},
		Validate: survey.Required,
	},
	{
		Name: "Artist",
		Prompt: &survey.Input{Message: "Artist name"},
		Validate: survey.Required,
	},
	{
		Name: "Album",
		Prompt: &survey.Input{Message: "Album name"},
		Validate: survey.Required,
	},
	{
		Name: "Genre",
		Prompt: &survey.Input{Message: "Genre"},
		Validate: survey.Required,
	},
	{
		Name: "Image",
		Prompt: &survey.Input{Message: "Image URL"},
		Validate: survey.Required,
	},
	{
		Name: "Track",
		Prompt: &survey.Input{Message: "Track Number"},
		Validate: survey.Required,
	},
	{
		Name: "Total",
		Prompt: &survey.Input{Message: "Total Track Count"},
		Validate: survey.Required,
	},
	{
		Name: "Duration",
		Prompt: &survey.Input{Message: "Song duration"},
		Validate: survey.Required,
	},
}

type File struct {
	Name     string `json:"title,omitempty"`
	Artist   string `json:"artist,omitempty"`
	Album    string `json:"album,omitempty"`
	Genre    string `json:"genre,omitempty"`
	Source   string `json:"source,omitempty"`
	Image    string `json:"image,omitempty"`
	Track    int    `json:"trackNumber,omitempty"`
	Total    int    `json:"totalTrackCount,omitempty"`
	Duration int    `json:"duration,omitempty"`
	Site 	 string `json:"site"`
}

type Upload struct {
	Songs []File `json:"songs,omitempty"`
}

func main() {
	var file File
	err := survey.Ask(Survey, &file)
	if err != nil {
		log.Println("error getting data", err)
		return
	}
	fmt.Println(err, file)
	up := Upload{Songs: []File{file}}
	buf, err := json.Marshal(up)
	if err != nil {
		log.Println("error marshalling file")
		return
	}
	resp, err := http.Post("http://localhost:3000/addSong", "application/json", bytes.NewReader(buf))
	if err != nil {
		log.Println("error sending request:", err)
		return
	}
	fmt.Println("success", resp)
}
