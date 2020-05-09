package db

import "github.com/zerjioang/gotools/thirdparty/echo"

//snoflake based composited id
// show flake generates 8 bytes
type DatumId [8 + 1 + 8]byte

type ItemKey struct {
	Left  []byte
	Sep   []byte
	Right []byte
}

type Policy func(context echo.Context, key string) error

type ControllerDBPolicyInterface interface {
	// method used to decode http input byte data to go struct
	CanRead(context echo.Context, key string) error
	CanUpdate(context echo.Context, key string) error
	CanDelete(context echo.Context, key string) error
	CanWrite(context echo.Context) error
	CanList(context echo.Context) error
}
