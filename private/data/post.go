package data

import (
	"context"
	"fmt"

	v1 "github.com/jace996/app-layout/api/post/v1"
	"github.com/jace996/app-layout/private/biz"
	"github.com/jace996/go-eventbus"
	kitgorm "github.com/jace996/multiapp/pkg/gorm"
	sgorm "github.com/jace996/saas/gorm"
	"gorm.io/gorm"
)

type PostRepo struct {
	*kitgorm.Repo[biz.Post, string, v1.ListPostRequest]
}

func NewPostRepo(dbProvider sgorm.DbProvider, eventbus *eventbus.EventBus) biz.PostRepo {
	res := &PostRepo{}
	res.Repo = kitgorm.NewRepo[biz.Post, string, v1.ListPostRequest](dbProvider, eventbus, res)
	return res
}

func (c *PostRepo) GetDb(ctx context.Context) *gorm.DB {
	return GetDb(ctx, c.Repo.DbProvider)
}

// BuildDetailScope preload relations
func (c *PostRepo) BuildDetailScope(withDetail bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

// BuildFilterScope filter
func (c *PostRepo) BuildFilterScope(q *v1.ListPostRequest) func(db *gorm.DB) *gorm.DB {
	search := q.Search
	filter := q.Filter
	return func(db *gorm.DB) *gorm.DB {
		ret := db

		if search != "" {
			ret = ret.Where("name like ?", fmt.Sprintf("%%%v%%", search))
		}
		if filter == nil {
			return ret
		}
		if filter.Id != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`id`", filter.Id))
		}
		if filter.Name != nil {
			ret = ret.Scopes(kitgorm.BuildStringFilter("`name`", filter.Name))
		}
		return ret
	}
}

func (c *PostRepo) DefaultSorting() []string {
	return []string{"created_at"}
}
