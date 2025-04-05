package fun7

import (
	"errors"
)

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("Could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	/*
	   refactor: from a simple function to a method
	   func Search(dict map[string]string, word string) string {
	   	return dict[word]
	   }
	*/
	// ok is a boolean which indicates if the key was found successfully or not
	//Let's extract the errors type in our function to be able to re-use it
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
