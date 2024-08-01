package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestReviewValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [Rating] not blank.", func(t *testing.T) {
		review := Review{
			// Rating: 4.5,
			Detail: "Test Detail",
			BookID: TEMP_UINT,
			UserID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Rating is required."))
	})

	t.Run("Check [Detail] not blank.", func(t *testing.T) {
		review := Review{
			Rating: 4.5,
			// Detail: "Test Detail",
			BookID: TEMP_UINT,
			UserID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Detail is required."))
	})

	t.Run("Check [BookID] not blank.", func(t *testing.T) {
		review := Review{
			Rating: 4.5,
			Detail: "Test Detail",
			// BookID: TEMP_UINT,
			UserID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("BookID is required."))
	})

	t.Run("Check [UserID] not blank.", func(t *testing.T) {
		review := Review{
			Rating: 4.5,
			Detail: "Test Detail",
			BookID: TEMP_UINT,
			// UserID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("UserID is required."))
	})
}
