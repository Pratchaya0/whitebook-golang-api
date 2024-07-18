package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestBookValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	bookCategoryId := uint(1)

// 	t.Run("check [BookName] not blank", func(t *testing.T) {
// 		book := Book{
// 			// BookName:          "BookName",
// 			BookDescription:   "BookDescription",
// 			BookPrice:         "BookPrice",
// 			BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input name"))
// 	})

// 	t.Run("check [BookDescription] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName: "BookName",
// 			// BookDescription:   "BookDescription",
// 			BookPrice:         "BookPrice",
// 			BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input description"))
// 	})

// 	t.Run("check [BookPrice] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:        "BookName",
// 			BookDescription: "BookDescription",
// 			// BookPrice:         "BookPrice",
// 			BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input price"))
// 	})

// 	t.Run("check [BookWriter] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:        "BookName",
// 			BookDescription: "BookDescription",
// 			BookPrice:       "BookPrice",
// 			// BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input writer"))
// 	})

// 	t.Run("check [BookPublisher] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:        "BookName",
// 			BookDescription: "BookDescription",
// 			BookPrice:       "BookPrice",
// 			BookWriter:      "BookWriter",
// 			// BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input publisher"))
// 	})

// 	t.Run("check [BookIsOnSale] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:        "BookName",
// 			BookDescription: "BookDescription",
// 			BookPrice:       "BookPrice",
// 			BookWriter:      "BookWriter",
// 			BookPublisher:   "BookPublisher",
// 			// BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input is on sale"))
// 	})

// 	t.Run("check [BookCoverImageUrl] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:        "BookName",
// 			BookDescription: "BookDescription",
// 			BookPrice:       "BookPrice",
// 			BookWriter:      "BookWriter",
// 			BookPublisher:   "BookPublisher",
// 			BookIsOnSale:    true,
// 			// BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl: "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input cover image"))
// 	})

// 	t.Run("check [BookUrl] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:          "BookName",
// 			BookDescription:   "BookDescription",
// 			BookPrice:         "BookPrice",
// 			BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			// BookUrl:           "BookUrl",

// 			BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input book"))
// 	})

// 	t.Run("check [BookCategoryId] not blank", func(t *testing.T) {
// 		book := Book{
// 			BookName:          "BookName",
// 			BookDescription:   "BookDescription",
// 			BookPrice:         "BookPrice",
// 			BookWriter:        "BookWriter",
// 			BookPublisher:     "BookPublisher",
// 			BookIsOnSale:      true,
// 			BookCoverImageUrl: "BookCoverImageUrl",
// 			BookUrl:           "BookUrl",

// 			// BookCategoryId: &bookCategoryId,
// 		}

// 		ok, err := govalidator.ValidateStruct(book)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please select category"))
// 	})
// }
