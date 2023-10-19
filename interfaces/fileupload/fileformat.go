package fileupload

import (
	"path"

	"github.com/twinj/uuid"
)

func FormatFile(fn string) string {

	ext := path.Ext(fn)
	u := uuid.NewV4()

	newFileName := u.String() + ext

	return newFileName
}
