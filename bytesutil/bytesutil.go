package bytesutil

import (
	"fmt"
	"io"
)

func ReadBytes(reader io.Reader, numBytes int) []byte {
	buffer := make([]byte, numBytes)
	read, err := reader.Read(buffer)
	if read != numBytes {
		panic(fmt.Sprintf("unable to read requested bytes, requested %d got %d", numBytes, read))
	}
	if err != nil {
		panic("error reading bytes")
	}

	return buffer
}

func ReadBytesSafe(reader io.Reader, numBytes int) ([]byte, error) {
	buffer := make([]byte, numBytes)
	read, err := reader.Read(buffer)
	if read != numBytes {
		return nil, fmt.Errorf("unable to read requested bytes, requested %d got %d", numBytes, read)
	}
	if err != nil {
		return nil, fmt.Errorf("error reading bytes")
	}

	return buffer, nil
}
