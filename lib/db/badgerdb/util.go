package badgerdb

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/zerjioang/gotools/util/str"

	"github.com/zerjioang/gotools/lib/logger"
	"golang.org/x/crypto/bcrypt"
)

var (
	empty []byte
)

func Hash(data string) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(str.UnsafeBytes(data), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash user password")
		return ""
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func CompareHash(plainPwd, hash string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	return err == nil
}

func Serialize(item interface{}) []byte {
	if item != nil {
		raw, err := json.Marshal(item)
		if err == nil {
			return raw
		}
	}
	return empty
}

func Unserialize(data []byte, item interface{}) error {
	return json.Unmarshal(data, item)
}

func chownR(path string, uid, gid int, mode os.FileMode) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
			err = os.Chmod(name, mode)
		}
		return err
	})
}
