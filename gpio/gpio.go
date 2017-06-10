package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	gpioExportFile = "/sys/class/gpio/export"

	gpio27Dir = "/sys/class/gpio/gpio27"
	gpio22Dir = "/sys/class/gpio/gpio22"

	gpio27DirectionFile = gpio27Dir + "/direction"
	gpio22DirectionFile = gpio22Dir + "/direction"

	gpio27ValueFile = gpio27Dir + "/value"
	gpio22ValueFile = gpio22Dir + "/value"
)

const (
	// Hot
	Hot = 22
	// Cold
	Cold = 27
)

// NinetyNineInputRules
const NinetyNineInputRules = `SUBSYSTEM=="input", GROUP="input", MODE="0660"`

// Toggle
func Toggle(relay int, onOff string) error {
	switch onOff {
	case "0":
		switch relay {
		case 22:
			if err := ioutil.WriteFile(gpio22ValueFile, []byte(onOff), 0644); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("relay off cold")
		case 27:
			if err := ioutil.WriteFile(gpio27ValueFile, []byte(onOff), 0644); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("relay off cold")
		}
	case "1":
		switch relay {
		case 22:
			if err := ioutil.WriteFile(gpio22ValueFile, []byte(onOff), 0644); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("relay on cold")
		case 27:
			if err := ioutil.WriteFile(gpio27ValueFile, []byte(onOff), 0644); err != nil {
				log.Fatalln(err)
			}
			fmt.Println("relay on cold")
		}
	}

	return nil
}

func init() {
	for _, i := range []string{"22", "27"} {
		if err := ioutil.WriteFile(gpioExportFile, []byte(i), 0644); err != nil {
			log.Fatalln(err)
		}
	}

	for _, i := range []string{gpio22DirectionFile, gpio27DirectionFile} {
		if err := ioutil.WriteFile(i, []byte("out"), 0644); err != nil {
			log.Fatalln(err)
		}
	}

	for _, i := range []string{gpio22Dir, gpio27Dir} {
		if err := chmodr(i, 0777); err != nil {
			log.Fatalln(err)
		}
	}
}

// chmodr
func chmodr(path string, perms int) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chmod(name, os.FileMode(perms))
		}
		return err
	})
}
