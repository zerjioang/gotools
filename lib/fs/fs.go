package fs

import (
	"bufio"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"
	"os"

	"github.com/zerjioang/gotools/lib/logger"
)

var (
	pagesize           int
	errReadingFullData = errors.New("error reading full size of requested data. some bytes are missing")
	secureRandom       = rand.Reader
	// wraps the Reader object into a new buffered reader to read the files in chunks
	// and buffering them for performance.
	secureRandomReader = bufio.NewReaderSize(secureRandom, pagesize)
	// todo optimize for multiple operative systems
	Separator = "/"
	empty     = []byte("")
)

func init() {
	logger.Debug("loading fs module")
	// For optimum speed, Getpagesize returns the underlying system's memory page size.
	// Getpagesize returns the underlying system's memory page size.
	pagesize = os.Getpagesize()
}

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return fi.Mode().IsDir()
}

func Exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReadAll(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err == nil && content != nil && len(content) > 0 {
		return content
	}
	return empty
}

func WriteFile(path string, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, _ = f.WriteString(content)
	return f.Close()
}

func ReadEntropy(rd io.Reader, size int) ([]byte, error) {
	// wraps the Reader object into a new buffered reader to read the files in chunks
	// and buffering them for performance.
	reader := bufio.NewReaderSize(rd, pagesize)
	readed := make([]byte, size)
	b, err := reader.Read(readed)
	if err == nil {
		//check if all bytes are readed ok
		if size == b {
			//return success
			return readed, nil
		} else {
			//return error
			return nil, errReadingFullData
		}
	}
	//return error message
	return nil, err
}

func ReadEntropy16() ([16]byte, error) {
	secureRandomReader.Reset(secureRandom)
	var readed [16]byte
	idx := 0
	for i := 0; i < 16; i++ {
		v, err := secureRandomReader.ReadByte()
		if err == nil {
			readed[idx] = v
			idx++
		}
	}
	//check if all bytes are readed ok
	if 16 == idx {
		//return success
		return readed, nil
	} else {
		//return error
		return readed, errReadingFullData
	}
}

func ReadFile(rd io.Reader) ([]byte, int, error) {
	// wraps the Reader object into a new buffered reader to read the files in chunks
	// and buffering them for performance.
	reader := bufio.NewReaderSize(rd, pagesize)
	b, err := reader.ReadByte()
	logger.Debug(b, err)
	var data []byte
	n, err := io.ReadFull(reader, data)
	return data, n, err
}

func PageSize() int {
	return pagesize
}
