package main

import (
	"errors"
	"os"
)

// Config contains the main app preferences
type config struct {
	// FilePath contains the absolute or relative path to a file for grep
	filePath string
	// KeyString is a substring we are looking for
	keyString string
	// IgnoreCase define the behavior about char's register: if ignoreCase is "true", we should ignore
	// the register of the input
	ignoreCase string
	// regex field if set to true it will enable regex matching
	regex string
}

// mustGetConfig do read config env variables and validation of the app input
// usually in go if function has a prefix "must/Must" - it means if will panic in case of any issue,
// like here.
func mustGetConfig() config {
	// TODO: add test
	config := loadConfig()
	if err := config.validate(); err != nil {
		panic(err)
	}

	return config
}

func loadConfig() config {
	// TODO: add test
	return config{
		filePath:   os.Getenv("FILE_PATH"),
		keyString:  os.Getenv("KEY_STRING"),
		ignoreCase: os.Getenv("IGNORE_CASE"),
		regex:      os.Getenv("REGEX"),
	}
}

func (c config) validate() error {
	// TODO: add test
	if c.filePath == "" {
		return errors.New("file path is empty")
	}

	if c.keyString == "" {
		return errors.New("nothing to search")
	}

	if c.ignoreCase != "true" {
		_, boolean := os.LookupEnv("IGNORE_CASE")
		if boolean {
			return errors.New("ignorecase parameter needs to be lowercase true")
		}
	}

	return nil
}
