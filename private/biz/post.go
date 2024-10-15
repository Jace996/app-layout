package biz

import (
	v1 "github.com/jace996/app-layout/api/post/v1"
	"github.com/jace996/multiapp/pkg/data"
	"github.com/jace996/multiapp/pkg/gorm"
)

type Post struct {
	gorm.UIDBase
	gorm.AuditedModel
	Name string
}

type PostRepo interface {
	data.Repo[Post, string, v1.ListPostRequest]
}
