package argf

import (
	"io"
	"os"
)

func Argf() (io.Reader, error) {
	return ArgfFrom(os.Args[1:])
}

func ArgfFrom(filenames []string) (io.Reader, error) {
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

		return io.MultiReader(files...), nil
	} else {
		return os.Stdin, nil
	}
}
