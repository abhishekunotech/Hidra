package utils

import(
	"github.com/antigloss/go/logger"
	"io/ioutil"
)

// Function that reads a file whose absolute path is passed to the function
//
// Business Logic: Reads the File as a byte slice from the path taken as input
//
// Returns a byte array
func ReadFile(path string) []uint8{
	bytesFile, err := ioutil.ReadFile(path)
	if err != nil{
		logger.Error("File Reading Failed")
	}

	return bytesFile
}
