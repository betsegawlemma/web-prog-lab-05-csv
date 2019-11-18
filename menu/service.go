package menu

import "github.com/betsegawlemma/webprogcsv/entity"

// CategoryService specifies food menu category related services
type CategoryService interface {
	Categories() ([]entity.Category, error)
	StoreCategores(categoies []entity.Category) error
}
