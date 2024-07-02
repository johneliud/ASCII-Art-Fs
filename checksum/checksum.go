package checksum

import (
	"crypto/sha256"
	"encoding/hex"
)

/*
CheckFileValidity calculates the SHA256 hash of the provided banner file and returns a string representing the SHA256 hash of the banner file data in hexadecimal format.
*/
func CheckFileValidity(bannerFileData []byte) string {
	hasher := sha256.New()
	hasher.Write(bannerFileData)
	hashInBytes := hasher.Sum(nil)
	fileHash := hex.EncodeToString(hashInBytes)
	return fileHash
}
