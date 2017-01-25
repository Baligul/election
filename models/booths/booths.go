package booths

type Booths struct {
	Total      int            `json:"total,omitempty"`
	File_title string         `json:"file_title,omitempty"`
	Booths     map[int]string `json:"booths,omitempty"`
}
