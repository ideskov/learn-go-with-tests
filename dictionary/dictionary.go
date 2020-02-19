package dictionary

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, exists := d[word]
	if exists {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, newDefinition string) (err error) {
	_, exists := d[word]
	if !exists {
		err = ErrWordDoesNotExist
	}
	d[word] = newDefinition
	return
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
