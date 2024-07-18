package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestBookPreviewImageValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	tmpBookId := uint(1)

// 	t.Run("check [BookPreviewImageUrl] not blank", func(t *testing.T) {
// 		bookPreviewImage := BookPreviewImage{
// 			BookPreviewImageUrl:    "BookPreviewImageUrl",
// 			BookPreviewImageBookId: &tmpBookId,
// 		}

// 		// TODO: loop test 3 case blank

// 		ok, err := govalidator.ValidateStruct(bookPreviewImage)
// 		g.Expect(ok).To(BeTrue())
// 		g.Expect(err).To(BeNil())
// 		// g.Expect(err.Error()).To(Equal(""))
// 	})

// 	t.Run("check [BookId] not blank", func(t *testing.T) {
// 		bookPreviewImage := BookPreviewImage{
// 			BookPreviewImageUrl: "BookPreviewImageUrl",
// 			// BookId:              &tmpBookId,
// 		}

// 		ok, err := govalidator.ValidateStruct(bookPreviewImage)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please select book"))
// 	})
// }
