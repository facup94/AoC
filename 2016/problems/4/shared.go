package day4

import (
	"strconv"
	"strings"
)

type room struct {
	encryptedName string
	sectorID      int
	checksum      []rune
}

func roomFromInput(s string) room {
	sectorIDIdx := strings.LastIndex(s, "-") + 1
	checksumIdx := strings.Index(s, "[") + 1

	sectorID, _ := strconv.Atoi(s[sectorIDIdx : checksumIdx-1])
	checksum := []rune(s[checksumIdx : len(s)-1])

	return room{
		encryptedName: s[:sectorIDIdx-1],
		sectorID:      sectorID,
		checksum:      checksum,
	}
}
