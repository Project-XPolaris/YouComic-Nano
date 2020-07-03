package entity

type PageFile struct {
	Path  string `json:"path"`
	Order int    `json:"order"`
	UUID string `json:"uuid"`
}
type Tag struct {
	ID   int    `json:"-"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type Book struct {
	ID    int        `json:"-"`
	UUID  string     `json:"uuid"`
	Name  string     `json:"name"`
	Path  string     `json:"path"`
	Cover string     `json:"cover"`
	Pages []PageFile `json:"pages"`
	Tags  []Tag      `json:"tags"`
}
