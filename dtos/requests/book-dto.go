package requests

type BookCreateDto struct {
	Name        string  `form:"name" valid:"required~Name is required."`
	Description string  `form:"description" valid:"required~Description is required."`
	Price       float64 `form:"price" valid:"required~Price is required."`
	CategoryID  uint    `form:"category_id" valid:"required~CategoryID is required."`
	Genres      []uint  `form:"genres" valid:"required~Genres is required."`
}

type BookUpdateDto struct {
	ID          uint    `form:"id" valid:"required~ID is required."`
	Name        string  `form:"name" valid:"required~Name is required."`
	Description string  `form:"description" valid:"required~Description is required."`
	Price       float64 `form:"price" valid:"required~Price is required."`
	CategoryID  uint    `form:"category_id" valid:"required~CategoryID is required."`
	Genres      []uint  `form:"genres" valid:"required~Genres is required."`
}
