package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCartValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [UserID] not blank", func(t *testing.T) {
		cart := Cart{
			// UserID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(cart)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("UserID is required."))
	})
}
