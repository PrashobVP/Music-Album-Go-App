package album

import (
    "gorm.io/gorm"
)

// Repository interface for album data.
type Repository interface {
    GetAll() ([]Album, error)
    GetByID(id string) (*Album, error)
    Create(album Album) error
	Update(album Album) error
	Delete(id string) error
}

// GormRepository struct for GORM implementation.
type GormRepository struct {
    db *gorm.DB
}

// NewGormRepository creates a new GormRepository.
func NewGormRepository(db *gorm.DB) *GormRepository {
    return &GormRepository{db}
}

// GetAll retrieves all albums from the database.
func (r *GormRepository) GetAll() ([]Album, error) {
    var albums []Album
    if err := r.db.Find(&albums).Error; err != nil {
        return nil, err
    }
    return albums, nil
}

// GetByID retrieves an album by ID.
func (r *GormRepository) GetByID(id string) (*Album, error) {
    var album Album
    if err := r.db.First(&album, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &album, nil
}

// Create inserts a new album into the database.
func (r *GormRepository) Create(album Album) error {
    return r.db.Create(&album).Error
}

// Update updates an existing album in the database.
func (r *GormRepository) Update(album Album) error {
    // Check if the album exists
    var existingAlbum Album
    if err := r.db.First(&existingAlbum, "id = ?", album.ID).Error; err != nil {
        return err // Return error if album is not found
    }

    // Update the album
    return r.db.Save(&album).Error
}

// Delete deletes an album by its ID from the database.
func (r *GormRepository) Delete(id string) error {
    var album Album
    if err := r.db.Where("id = ?", id).Delete(&album).Error; err != nil {
        return err
    }
    return nil
}


