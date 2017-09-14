package utils

import(
	"github.com/antigloss/go/logger"
	"io/ioutil"
)


func ReadFile(path string) []uint8{
	bytesFile, err := ioutil.ReadFile(path)
	if err != nil{
		logger.Error("File Reading Failed")
	}

	return bytesFile
}
