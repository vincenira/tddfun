package fun7

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	/*
	   refactor: from a simple function to a method
	   func Search(dict map[string]string, word string) string {
	   	return dict[word]
	   }
	*/
	// ok is a boolean which indicates if the key was found successfully or not
	if !ok {
		return "", errors.New("Could not find the word you were looking for")
	}
	return definition, nil
}
