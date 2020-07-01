package generate

import (
	"YouComic-Nano/entity"
	"github.com/google/uuid"
	"io/ioutil"
	"path"
	"sort"
	"strings"
)

var (
	AllowPageFileExtensions = []string{".jpg", ".jpeg", ".png", ".bmp"}
	MinPageFileRequire      = 3
	DefaultCoverIndex       = 0
)

func checkIsPageFile(pathLike string) bool {
	fileName := path.Base(pathLike)
	for _, allowPageFileExtension := range AllowPageFileExtensions {
		if strings.Contains(strings.ToLower(fileName), allowPageFileExtension) {
			return true
		}
	}
	return false
}

func CreateLibrary(scanPath string) ([]entity.Book, error) {
	queue := []string{scanPath}
	books := make([]entity.Book, 0)
	for len(queue) != 0 {
		var currentPath string
		currentPath, queue = queue[0], queue[1:]
		items, err := ioutil.ReadDir(currentPath)
		if err != nil {
			return nil, err
		}
		bookUUID, err := uuid.NewUUID()
		if err != nil {
			return nil, nil
		}
		book := entity.Book{
			Pages: []entity.PageFile{},
			Path:  path.Base(currentPath),
			UUID:  bookUUID.String(),
		}
		// scan for dir items
		for _, info := range items {
			if info.IsDir() {
				queue = append(queue, path.Join(currentPath, info.Name()))
				continue
			}
			fileName := info.Name()
			if checkIsPageFile(fileName) {
				book.Pages = append(book.Pages, entity.PageFile{
					Path: fileName,
				})
			}
		}

		// check it is book directory
		if len(book.Pages) > MinPageFileRequire {
			matchResult, err := MatchBookInfoWithTitle(path.Base(book.Path))
			if err != nil {
				return nil, err
			}
			if matchResult != nil {
				title, isExist := matchResult["title"]
				if isExist {
					book.Name = title
				}

				targetTagType := []string{"series", "theme", "artist", "translator"}
				for _, typeName := range targetTagType {
					if tagValue, isExist := matchResult[typeName]; isExist {
						book.Tags = append(book.Tags, entity.Tag{Type: typeName, Name: tagValue})
					}
				}
			}

			// resort page file
			sort.Slice(book.Pages, func(i, j int) bool {
				return book.Pages[i].Path < book.Pages[j].Path
			})

			// assign page order
			for order := range book.Pages {
				book.Pages[order].Order = order
			}

			// set cover
			book.Cover = book.Pages[DefaultCoverIndex].Path

			// save result
			books = append(books, book)
		}
	}
	return books, nil
}
