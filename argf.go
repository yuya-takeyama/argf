package argf

import (
	"io"
	"os"
)

// Argf returns io.Reader from STDIN or files from command-line arguments.
//
// If the files are multiple, they are concatinated as one io.Reader.
func Argf() (io.Reader, error) {
	return From(os.Args[1:])
}

// From returns io.Reader from STDIN or files from function argument.
func From(filenames []string) (io.Reader, error) {
	var reader io.Reader

	filelen := len(filenames)

	if filelen > 0 {
		files := make([]io.Reader, filelen)

		for i := 0; i < filelen; i++ {
			file, err := os.Open(filenames[i])
			if err != nil {
				return nil, err
			}

			files[i] = file
		}

		reader = io.MultiReader(files...)
	} else {
		reader = os.Stdin
	}

	return reader, nil
}
