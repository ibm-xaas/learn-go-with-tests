package main

// Dictionary ...
type Dictionary map[string]string

// Search ...
func (d Dictionary) Search(word string) string {
	return d[word]
}

func main() {}
