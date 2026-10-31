package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("cannot find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(word, value string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = value
	default:
		return err
	}
	return nil

}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}

	return nil
}