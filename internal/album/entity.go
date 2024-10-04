package album

// Album represents data about a record album.
type Album struct {
    ID     string  `json:"id" gorm:"primaryKey"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}
