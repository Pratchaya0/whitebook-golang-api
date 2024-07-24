package requests

type BookCreateDto struct {
	Name        string  `form:"name"`
	Description string  `form:"description"`
	Price       float64 `form:"price"`
	CategoryID  uint    `form:"category_id"`
	Genres      []uint  `form:"genres"`
}
