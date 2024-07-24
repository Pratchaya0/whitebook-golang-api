package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestGenreBookValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	tmpId := uint(1)

// 	t.Run("check [GenreId] not blank", func(t *testing.T) {
// 		genreBook := GenreBook{
// 			// GenreId: &tmpId,
// 			GenreBookBookId: &tmpId,
// 		}

// 		ok, err := govalidator.ValidateStruct(genreBook)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input genre id"))
// 	})

// 	t.Run("check [BookId] not blank", func(t *testing.T) {
// 		genreBook := GenreBook{
// 			GenreId: &tmpId,
// 			// BookId:  &tmpId,
// 		}

// 		ok, err := govalidator.ValidateStruct(genreBook)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input book id"))
// 	})
// }
