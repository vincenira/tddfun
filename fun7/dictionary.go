package fun7

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryError("Could not find the word you were looking for")
	ErrWordExists       = DictionaryError("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryError("cannot perform operation on word because it does not exist")
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

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = newDefinition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

// as for Add and update we can implement a logic when the key provided, does not exist
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}
