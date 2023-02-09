package entity

import (
	"mime/multipart"
)

type Problem struct {
	ID       string                `form:"id"`
	Language string                `form:"language"`
	Source   *multipart.FileHeader `form:"source"`
}
