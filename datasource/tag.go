package datasource

import (
	"YouComic-Nano/config"
	"YouComic-Nano/entity"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var Tags []entity.Tag

func InitTagDataSource() {
	Tags = make([]entity.Tag, 0)
	autoIncrease := 0
	for bookIdx := range config.Current.Books {
		for tagIdx, _ := range config.Current.Books[bookIdx].Tags {
			tag := config.Current.Books[bookIdx].Tags[tagIdx]
			isExist := false
			for _, existedTag := range Tags {
				if tag.Name == existedTag.Name && tag.Type == existedTag.Type {
					config.Current.Books[bookIdx].Tags[tagIdx].ID = existedTag.ID
					isExist = true
					break
				}
			}
			if !isExist {
				autoIncrease += 1
				config.Current.Books[bookIdx].Tags[tagIdx].ID = autoIncrease
				tag.ID = autoIncrease
				Tags = append(Tags, tag)
			}
		}
	}
	logrus.Info(fmt.Sprintf("init tag for %d", len(Tags)))
}

func GetTagList(page int, pageSize int) ([]entity.Tag, int, error) {
	source := Tags
	if (page-1)*pageSize >= len(source) {
		return []entity.Tag{}, len(source), nil
	}
	end := page * pageSize
	if end > len(source) {
		end = len(source)
	}

	return source[(page-1)*pageSize : end], len(source), nil
}

type TagsReader struct {
	Page     int
	PageSize int
	Filters  []DataFilter
}

func (t *TagsReader) Read() (interface{}, int, error) {
	if t.Filters == nil {
		t.Filters = make([]DataFilter, 0)
	}
	source := Tags
	for _, dataFilter := range t.Filters {
		result := make([]entity.Tag, 0)
		for _, tag := range source {
			ok, err := dataFilter.filter(tag)
			if err != nil {
				return make([]entity.Tag, 0), 0, err
			}
			if ok {
				result = append(result, tag)
			}
		}
		source = result
	}
	start := (t.Page - 1) * t.PageSize
	if start >= len(source) {
		return make([]entity.Tag, 0), 0, nil
	}
	end := t.Page * t.PageSize
	if end > len(source) {
		end = len(source)
	}
	result := source[start:end]
	return result, len(source), nil
}

type TagNameSearchFilter struct {
	Key string
}

func (t *TagNameSearchFilter) filter(data interface{}) (bool, error) {
	return strings.Contains((data.(entity.Tag)).Name, t.Key), nil
}

type TagTypeFilter struct {
	TypeNames []string
}

func (t *TagTypeFilter) filter(data interface{}) (bool, error) {
	for _, typeName := range t.TypeNames {
		if (data.(entity.Tag)).Type == typeName {
			return true, nil
		}
	}
	return false, nil
}
