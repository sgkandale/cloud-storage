package object

type Object struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	ParentFolderId int64     `json:"parent_folder_id"`
	ObjectType     string    `json:"object_type"`
	CreatedAt      int64     `json:"-"`
	CreatedAtStr   string    `json:"created_at"`
	Children       []*Object `json:"children"`
	OwnerId        int64     `json:"owner_id"`
	FileLink       string    `json:"file_link"`
}
