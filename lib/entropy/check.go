package entropy

import (
	"strconv"

	"github.com/zerjioang/gotools/lib/fs"

	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/util/str"
)

//this package will check about entropy bits available in the system
var (
	defaultFileAvailable bool
	defaultFilePath      = "/proc/sys/kernel/random/entropy_avail"
	initialEntropy       int
	isCritical           bool
)

func init() {
	//check if default file exists
	logger.Info("reading entropy available bytes")
	defaultFileAvailable = fs.Exists(defaultFilePath)
	initialEntropy = AvailableEntropy()
	isCritical = initialEntropy < 1000
}

// return first loaded available entropy measure
func InitialEntropy() int {
	return initialEntropy
}

// This command shows you how much entropy your server has collected.
// If it is rather low (<1000), you should probably install haveged.
// Otherwise cryptographic applications will block until there is enough entropy available
// which eg. could result in slow wlan speed, if your server is a Software access point.
// sudo apt-get install haveged.
func HasCriticalValue() bool {
	return isCritical
}

// this method will return available entropy in the system
// if error found, a -1 is returned
func AvailableEntropy() int {
	// read entropy form file
	content := fs.ReadAll(defaultFilePath)
	t := len(content)
	if t > 0 {
		// convert entropy data to int
		v, err := strconv.Atoi(str.UnsafeString(content[:t-1]))
		if err == nil {
			return v
		}
	}
	return -1
}
