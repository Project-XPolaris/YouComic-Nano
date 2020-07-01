package datasource

import (
	"YouComic-Nano/config"
	"YouComic-Nano/entity"
	"fmt"
	"github.com/sirupsen/logrus"
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
	if (page - 1)*pageSize >= len(source) {
		return []entity.Tag{},len(source),nil
	}
	end := page * pageSize
	if end > len(source) {
		end = len(source)
	}

	return source[(page-1)*pageSize : end], len(source), nil
}
