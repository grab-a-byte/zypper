package zip

import (
	"encoding/binary"
	"io"
	"time"
	"zypper/bytesutil"
)

var readBytes = bytesutil.ReadBytes

// Data is parsed, TODO: Parse all extra Field info and store on struct, maybe methods to access data?
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
	header.ExtraField = readBytes(reader, int(header.ExtraFieldLength))
	if header.CompressedSize != 0 {
		header.Data = readBytes(reader, int(header.CompressedSize))
	}
	return header
}

func ReadZip(file io.Reader) ZipFile {
	zipFile := ZipFile{}
	for {
		start, err := bytesutil.ReadBytesSafe(file, 4)
		if err != nil {
			break
		}
		if binary.LittleEndian.Uint32(start) == LocalFileHeaderSignature {
			header := parseLocalFileHeader(file)
			zipFile.FileHeaders = append(zipFile.FileHeaders, header)
		}
	}
	return zipFile
}
