package thing

import "os"

type DataReader struct {
	filepath string
}

func NewDataReader(filepath string) DataReader {
	return DataReader{
		filepath: filepath,
	}
}

func (d DataReader) Read() (string, error) {
	data, err := os.ReadFile(d.filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
