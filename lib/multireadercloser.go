package lib

import (
	"io"
	"os"
)

type multiReadCloser struct {
	readClosers []io.ReadCloser
}

func (mrc *multiReadCloser) Read(p []byte) (int, error) {
	for len(mrc.readClosers) > 0 {
		n, err := mrc.readClosers[0].Read(p)
		if n > 0 || err != io.EOF {
			if err == io.EOF {
				err = nil
			}
			return n, err
		}
		// Unfortunately, we have to ignore closing errors.
		mrc.readClosers[0].Close()
		mrc.readClosers = mrc.readClosers[1:]
	}
	return 0, io.EOF
}

func (mrc *multiReadCloser) Close() error {
	for _, i := range mrc.readClosers {
		err := i.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func MultiReadCloser(readClosers ...io.ReadCloser) io.ReadCloser {
	return &multiReadCloser{readClosers}
}

func MultiFile(names ...string) (io.ReadCloser, error) {
	files := make([]io.ReadCloser, len(names))

	// Open the files.
	for i, name := range names {
		file, err := os.Open(name)
		if err != nil {
			// TODO: Handle this better.
			return nil, err
		}
		files[i] = file
	}

	return MultiReadCloser(files...), nil
}
