package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookPreviewImageValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [PageNumber] not blank.", func(t *testing.T) {
		bookPreviewImage := BookPreviewImage{
			// PageNumber: TEMP_UINT,
			ImageLink: TEMP_LINK,
			BookID:    TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(bookPreviewImage)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("PageNumber is required."))
	})

	t.Run("Check [ImageLink] not blank.", func(t *testing.T) {
		bookPreviewImage := BookPreviewImage{
			PageNumber: TEMP_UINT,
			// ImageLink:  TEMP_LINK,
			BookID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(bookPreviewImage)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("ImageLink is required."))
	})

	t.Run("Check [BookID] not blank.", func(t *testing.T) {
		bookPreviewImage := BookPreviewImage{
			PageNumber: TEMP_UINT,
			ImageLink:  TEMP_LINK,
			// BookID:     TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(bookPreviewImage)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("BookID is required."))
	})
}
