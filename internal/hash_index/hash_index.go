package hash_index

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const splitString = ","

type HashIndex struct {
	filename string
	index    map[string]int64
}

func NewFromFile(filePath string) (*HashIndex, error) {
	hashIndex := HashIndex{filename: filePath, index: make(map[string]int64)}

	_, err := os.Stat(filePath)

	// in case file doesn't exist return an empty index
	if err != nil {
		return &hashIndex, nil
	}

	fd, err := os.Open(filePath)
	defer fd.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		var k, v string
		var offset int64

		s := scanner.Text()

		split := strings.Split(s, splitString)

		if len(split) != 2 {
			return nil, errors.New("error parsing index file")
		}

		k = split[0]
		v = split[1]
		offset, err = strconv.ParseInt(v, 10, 64)

		if err != nil {
			return nil, errors.New("can't parse '" + v + "' as int64")
		}

		hashIndex.SetOffset(k, offset)
	}

	return &hashIndex, nil
}

func (hi *HashIndex) FlushToFile() error {
	fd, err := os.OpenFile(hi.filename, os.O_CREATE|os.O_WRONLY, 0755)
	defer fd.Close()
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(fd)
	defer writer.Flush()

	for k, v := range hi.index {
		s := fmt.Sprintf("%v,%v\n", k, v)
		nn, err := writer.Write([]byte(s))
		if err != nil || nn != len(s) {
			return errors.New("error storing index to file")
		}
	}

	return nil
}

func (hi *HashIndex) GetOffset(key string) (int64, bool) {
	value, isOk := hi.index[key]

	return value, isOk
}

func (hi *HashIndex) SetOffset(key string, offset int64) {
	hi.index[key] = offset
}
