package main

import (
	"os"
	"strings"
	tafilesread "totala_reader/ta_files_read"
	"totala_reader/ta_files_read/object3d"
	"totala_reader/ta_files_read/scripts"
	"totala_reader/ta_files_read/texture"
)

func readAllGAFsFromDirectory(directoryName string) []*texture.GafEntry {
	pp("Reading all GAF entries from dir %s", directoryName)
	var allEntries []*texture.GafEntry
	if directoryName[len(directoryName)-1] != "/"[0] {
		directoryName += "/"
	}
	items, _ := os.ReadDir(directoryName)
	for _, item := range items {
		if item.IsDir() {
			// do nothing
		} else {
			openedFileName := directoryName + item.Name()
			r := &tafilesread.Reader{}
			r.ReadFromFile(openedFileName)
			readedGAFEntries := texture.ReadTextureFromReader(r, false)
			for _, e := range readedGAFEntries {
				allEntries = append(allEntries, e)
			}
		}
	}
	return allEntries
}

// Read script and model from different folders by file name
func loadModelAndCobByFilename(filename string) (*object3d.Object, *scripts.CobScript) {
	const folder3do = "game_files/files_3do/"
	const foldercob = "game_files/files_cob/"
	baseName := getBaseNameByFilename(filename)

	var modelInTAFormat *object3d.Object
	var scriptForModel *scripts.CobScript
	items, _ := os.ReadDir(folder3do)
	for _, item := range items {
		modelName := strings.ToLower(item.Name())
		if strings.Contains(modelName, baseName) {
			openedFileName := folder3do + item.Name()
			pp("Opening model %s", openedFileName)
			r := &tafilesread.Reader{}
			r.ReadFromFile(openedFileName)
			modelInTAFormat = object3d.ReadObjectFromReader(r, 0)
			break
		}
	}
	items, _ = os.ReadDir(foldercob)
	for _, item := range items {
		cobName := strings.ToLower(item.Name())
		if !strings.Contains(cobName, ".cob") {
			continue
		}
		if strings.Contains(cobName, baseName) {
			openedFileName := foldercob + item.Name()
			pp("Opening COB %s", openedFileName)
			r := &tafilesread.Reader{}
			r.ReadFromFile(openedFileName)
			scriptForModel = scripts.ReadCobFileFromReader(r)
			break
		}
	}
	return modelInTAFormat, scriptForModel
}

func getBaseNameByFilename(fName string) string {
	slashIndex := strings.LastIndex(fName, "/")
	dotIndex := strings.LastIndex(fName, ".")
	return strings.ToLower(fName[slashIndex+1 : dotIndex])
}
