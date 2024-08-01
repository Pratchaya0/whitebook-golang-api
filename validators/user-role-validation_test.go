package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUserRoleValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [Name] not blank.", func(t *testing.T) {
		userRole := UserRole{
			// Name:        "Test Name",
			Description: "Test Description",
		}

		ok, err := govalidator.ValidateStruct(userRole)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Name is required."))
	})

	t.Run("Check [Description] not blank.", func(t *testing.T) {
		userRole := UserRole{
			Name: "Test Name",
			// Description: "Test Description",
		}

		ok, err := govalidator.ValidateStruct(userRole)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Description is required."))
	})
}
