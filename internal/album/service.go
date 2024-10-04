package album

// Service interface for album operations.
type Service interface {
	GetAllAlbums() ([]Album, error)
	GetAlbumByID(id string) (*Album, error)
	CreateAlbum(album Album) error
	UpdateAlbum(album Album) error
	DeleteAlbum(id string) error
}

// albumService struct for service implementation.
type albumService struct {
	repo Repository
}

// NewAlbumService creates a new album service.
func NewAlbumService(repo Repository) Service {
	return &albumService{repo}
}

// GetAllAlbums retrieves all albums.
func (s *albumService) GetAllAlbums() ([]Album, error) {
	return s.repo.GetAll()
}

// GetAlbumByID retrieves an album by ID.
func (s *albumService) GetAlbumByID(id string) (*Album, error) {
	return s.repo.GetByID(id)
}

// CreateAlbum creates a new album.
func (s *albumService) CreateAlbum(album Album) error {
	return s.repo.Create(album)
}

// UpdateAlbum updates an existing album in the repository.
func (s *albumService) UpdateAlbum(album Album) error {
	return s.repo.Update(album)
}

// DeleteAlbum deletes an album by its ID.
func (s *albumService) DeleteAlbum(id string) error {
	return s.repo.Delete(id)
}
