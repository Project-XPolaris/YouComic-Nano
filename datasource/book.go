package datasource

import (
	"YouComic-Nano/config"
	"YouComic-Nano/entity"
)

func InitDataSource() {
	InitTagDataSource()
	for idx := range config.Current.Books {
		config.Current.Books[idx].ID = idx + 1
	}
}
func GetBookList(page int, pageSize int,tag int) ([]entity.Book, int, error) {
	source := config.Current.Books
	if tag > 0 {
		result := make([]entity.Book,0)
		for _, book := range source {
			for _, bookTag := range book.Tags {
				if bookTag.ID == tag {
					result = append(result, book)
				}
			}
		}
		source = result
	}
	return source[(page-1)*pageSize : (page * pageSize)], len(source), nil
}

func GetBookById(id int) (*entity.Book, error) {
	for _, book := range config.Current.Books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func GetBookCoverCacheById(id int) (string, error) {
	book, err := GetBookById(id)
	if err != nil {
		return "", err
	}
	rawThumbnails := config.CacheConfig.GetStringMapString("thumbnail")
	return rawThumbnails[book.UUID], nil
}
