package zip

import "time"

type ZipFile struct {
	FileHeaders []LocalFileHeader
}

type LocalFileHeader struct {
	LastModTime        time.Time
	FileHeader         string
	FileName           string
	ExtraField         []byte
	Data               []byte
	UncompressedSize   uint32
	CompressedSize     uint32
	Crc                uint32
	VersionToExtract   uint16
	GeneralPurposeFlag uint16
	CompressionMethod  uint16
	FileNameLength     uint16
	ExtraFieldLength   uint16
}
