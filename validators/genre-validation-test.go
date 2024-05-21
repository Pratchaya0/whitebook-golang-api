package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestGenreValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check [GenreName] not blank", func(t *testing.T) {
		genre := Genre{}

		ok, err := govalidator.ValidateStruct(genre)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal(""))
	})

}
