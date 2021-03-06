package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"os"
	"path/filepath"
	"time"
)

// Acquisition defines the basic properties we want to store.
type Acquisition struct {
	UUID         string
	Date         string
	Time         string
	ComputerName string
	ComputerUser string
	Platform     string
	Folder       string
	Storage      string
	Autoruns     string
	Memory       string
}

func (a *Acquisition) Initialize() {
	uuid, _ := uuid.NewV4()
	a.UUID = uuid.String()
	// Get the time in UTC.
	currentTime := time.Now().UTC()
	// Extract the date through Go's idiotic formatting.
	a.Date = currentTime.Format("2006-02-01")
	// Extract the time through Go's idiotic formatting.
	a.Time = currentTime.Format("15:04:05")
	// Get the computer name.
	a.ComputerName = getComputerName()
	// Get the current user.
	// TODO: for some reason, this doesn't work on Windows 10.
	a.ComputerUser = getUserName()
	// Get details on the operating system.
	a.Platform = getOperatingSystem()

	// This is some spaghetti code to generate the folder name for the current
	// acquisition. It will just try to append a number until it finds a
	// combination that has not been used yet.
	baseFolder := fmt.Sprintf("%s_%s", a.Date, a.ComputerName)
	baseStorage := filepath.Join(getCwd(), "acquisitions")
	tmpFolder := baseFolder
	tmpStorage := filepath.Join(baseStorage, baseFolder)
	counter := 1
	for {
		if _, err := os.Stat(tmpStorage); os.IsNotExist(err) {
			break
		} else {
			tmpFolder = fmt.Sprintf("%s_%d", baseFolder, counter)
			tmpStorage = filepath.Join(baseStorage, tmpFolder)
			counter++
		}
	}

	// Proceeds creating all the required subfolders.
	a.Folder = tmpFolder
	a.Storage = tmpStorage
	a.Autoruns = filepath.Join(a.Storage, "autoruns")
	a.Memory = filepath.Join(a.Storage, "memory")

	err := os.MkdirAll(a.Storage, 0755)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(a.Autoruns, 0755)
	if err != nil {
		panic(err)
	}
	err = os.Mkdir(a.Memory, 0755)
	if err != nil {
		panic(err)
	}
}
