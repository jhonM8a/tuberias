package services_interface

import (
	"tuberias/dto"
)

type MetadaFile interface {
	GetMetadatByFile() dto.FileInfo
}

type SaveMetadaNoSql interface {
	SaveData(string)
}
