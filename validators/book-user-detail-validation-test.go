package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBookUserDetailValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	tmpId := uint(1)

	t.Run("check [BookUserDetailIsAvailable] not blank", func(t *testing.T) {
		bookUserDetail := BookUserDetail{
			// BookUserDetailIsAvailable: true,
			BookUserDetailBookId: &tmpId,
			UserId:               &tmpId,
		}

		ok, err := govalidator.ValidateStruct(bookUserDetail)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input book user detail is available"))
	})

	t.Run("check [BookId] not blank", func(t *testing.T) {
		bookUserDetail := BookUserDetail{
			BookUserDetailIsAvailable: true,
			// BookId: &tmpId,
			UserId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(bookUserDetail)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input book id"))
	})

	t.Run("check [UserId] not blank", func(t *testing.T) {
		bookUserDetail := BookUserDetail{
			BookUserDetailIsAvailable: true,
			BookUserDetailBookId:      &tmpId,
			// UserId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(bookUserDetail)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input user id"))
	})
}
