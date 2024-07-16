package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestReviewValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	tmpId := uint(1)

	t.Run("check [ReviewComment] not blank", func(t *testing.T) {
		review := Review{
			// ReviewComment: "ReviewComment",
			UserId:       &tmpId,
			ReviewBookId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input comment"))
	})

	t.Run("check [UserId] not blank", func(t *testing.T) {
		review := Review{
			ReviewComment: "ReviewComment",
			// UserId: &tmpId,
			ReviewBookId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input user id"))
	})

	t.Run("check [BookId] not blank", func(t *testing.T) {
		review := Review{
			ReviewComment: "ReviewComment",
			UserId:        &tmpId,
			// BookId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(review)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input book id"))
	})
}
