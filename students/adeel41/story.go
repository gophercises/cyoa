package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

//Story contains all arcs found in the file
type Story struct {
	arcs []StoryArc
}

//Load loads the file and extract all arcs found in json file
func (s *Story) Load(filePath string) error {
	jsonData, err := s.getJSON(filePath)
	if err != nil {
		return err
	}

	var arcs []StoryArc
	for key, data := range jsonData {
		arc := new(StoryArc)
		arc.Load(key, data.(map[string]interface{}))
		arcs = append(arcs, *arc)
	}
	s.arcs = arcs
	return nil
}

func (s Story) getJSON(filepath string) (map[string]interface{}, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Cannot find gopher.json file")
		return nil, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Unable to read gopher.json file")
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		log.Println("Cannot parse json specified in gopher.json file")
		return nil, err
	}
	return data.(map[string]interface{}), nil
}

//GetArc finds the arc specified in the key argument and returns it
func (s Story) GetArc(key string) (*StoryArc, error) {
	for _, arc := range s.arcs {
		if arc.Identifier == key {
			return &arc, nil
		}
	}
	return nil, errors.New("Cannot find " + key + " arc")
}
