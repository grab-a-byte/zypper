package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"
)

type ZipFile struct {
	FileHeaders []LocalFileHeader
}

type LocalFileHeader struct {
	LastModTime        time.Time
	FileHeader         string
	FileName           string
	ExtraField         string
	Crc                uint32
	CompressedSize     uint32
	UncompressedSize   uint32
	VersionToExtract   uint16
	GeneralPurposeFlag uint16
	CompressionMethod  uint16
	FileNameLength     uint16
	ExtraFieldLength   uint16
}

func readBytes(reader io.Reader, numBytes int) []byte {
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

func parseLocalFileHeader(reader io.Reader) LocalFileHeader {
	var header LocalFileHeader
	header.VersionToExtract = binary.LittleEndian.Uint16(readBytes(reader, 2))
	header.GeneralPurposeFlag = binary.LittleEndian.Uint16(readBytes(reader, 2))
	header.CompressionMethod = binary.LittleEndian.Uint16(readBytes(reader, 2))
	header.LastModTime = time.Unix(int64(binary.LittleEndian.Uint32(readBytes(reader, 4))), 0) // Not correct, stored in extra fields for some reason
	header.Crc = binary.LittleEndian.Uint32(readBytes(reader, 4))
	header.CompressedSize = binary.LittleEndian.Uint32(readBytes(reader, 4))
	header.UncompressedSize = binary.LittleEndian.Uint32(readBytes(reader, 4))
	header.FileNameLength = binary.LittleEndian.Uint16(readBytes(reader, 2))
	header.ExtraFieldLength = binary.LittleEndian.Uint16(readBytes(reader, 2))
	header.FileName = string(readBytes(reader, int(header.FileNameLength)))
	header.ExtraField = string(readBytes(reader, int(header.ExtraFieldLength)))
	return header
}

func readZip(file io.Reader) {
	for {
		start := readBytes(file, 4)
		if binary.LittleEndian.Uint32(start) == 0x04034b50 {
			fmt.Println("Found local file header")
			header := parseLocalFileHeader(file)
			fmt.Printf("%+v \n", header)
		}
	}
}

func main() {
	file, err := os.Open("test.zip")
	if err != nil {
		panic("Unable to open test zip file")
	}
	defer file.Close()

	readZip(file)
	fmt.Println("Hello World")
}
