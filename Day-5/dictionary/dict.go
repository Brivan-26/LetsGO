package dictionary

import "errors"
type Dictionary map[string]string

var (
	ErrNotFound = errors.New("The word doesn't exist!")
	ErrWordExists = errors.New("The word already exists!")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, found := d[key]

	if !found {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, found := d[key]

	if found {
		return ErrWordExists
	}else {
		d[key] = value
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, found := d[key]
	if !found {
		return ErrNotFound
	}

	d[key] = newValue
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}