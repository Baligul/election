package sections

type Sections struct {
	Total      int            `json:"total,omitempty"`
	File_title string         `json:"file_title,omitempty"`
	Sections   map[int]string `json:"sections,omitempty"`
}
