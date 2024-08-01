package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [Name] not blank.", func(t *testing.T) {
		book := Book{
			// Name:        "TestName",
			Description: "Test Description",
			Price:       123.00,
			CategoryID:  TEMP_UINT,
			CoverImage:  TEMP_LINK,
			BookPdf:     TEMP_LINK,
			BookEpub:    TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Name is required."))
	})

	t.Run("Check [Description] not blank.", func(t *testing.T) {
		book := Book{
			Name: "TestName",
			// Description: "Test Description",
			Price:      123.00,
			CategoryID: TEMP_UINT,
			CoverImage: TEMP_LINK,
			BookPdf:    TEMP_LINK,
			BookEpub:   TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Description is required."))
	})

	t.Run("Check [Price] not blank.", func(t *testing.T) {
		book := Book{
			Name:        "TestName",
			Description: "Test Description",
			// Price:       123.00,
			CategoryID: TEMP_UINT,
			CoverImage: TEMP_LINK,
			BookPdf:    TEMP_LINK,
			BookEpub:   TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Price is required."))
	})

	t.Run("Check [CategoryID] not blank.", func(t *testing.T) {
		book := Book{
			Name:        "TestName",
			Description: "Test Description",
			Price:       123.00,
			// CategoryID:  TEMP_UINT,
			CoverImage: TEMP_LINK,
			BookPdf:    TEMP_LINK,
			BookEpub:   TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("CategoryID is required."))
	})

	t.Run("Check [CoverImage] not blank.", func(t *testing.T) {
		book := Book{
			Name:        "TestName",
			Description: "Test Description",
			Price:       123.00,
			CategoryID:  TEMP_UINT,
			// CoverImage:  TEMP_LINK,
			BookPdf:  TEMP_LINK,
			BookEpub: TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("CoverImage is required."))
	})

	t.Run("Check [BookPdf] not blank.", func(t *testing.T) {
		book := Book{
			Name:        "TestName",
			Description: "Test Description",
			Price:       123.00,
			CategoryID:  TEMP_UINT,
			CoverImage:  TEMP_LINK,
			// BookPdf:     TEMP_LINK,
			BookEpub: TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("BookPdf is required."))
	})

	t.Run("Check [BookEpub] not blank.", func(t *testing.T) {
		book := Book{
			Name:        "TestName",
			Description: "Test Description",
			Price:       123.00,
			CategoryID:  TEMP_UINT,
			CoverImage:  TEMP_LINK,
			BookPdf:     TEMP_LINK,
			// BookEpub:    TEMP_LINK,
		}

		ok, err := govalidator.ValidateStruct(book)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("BookEpub is required."))
	})
}
