package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCartValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	tmpId := uint(1)

	t.Run("check [BookId] not blank", func(t *testing.T) {
		cart := Cart{
			// BookId: &tmpId,
			UserId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(cart)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input book id"))
	})

	t.Run("check [UserId] not blank", func(t *testing.T) {
		cart := Cart{
			CartBookId: &tmpId,
			// UserId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(cart)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input user id"))
	})

}
