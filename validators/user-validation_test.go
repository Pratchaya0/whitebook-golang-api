package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUserValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [Name] bot blank.", func(t *testing.T) {
		user := User{
			// Name:     "Test Name",
			Email:    "Test Email",
			Password: TEMP_BYTE_ARRAY,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Name is required."))
	})

	t.Run("Check [Email] bot blank.", func(t *testing.T) {
		user := User{
			Name: "Test Name",
			// Email:    "Test Email",
			Password: TEMP_BYTE_ARRAY,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Email is required."))
	})

	t.Run("Check [Password] bot blank.", func(t *testing.T) {
		user := User{
			Name:  "Test Name",
			Email: "Test Email",
			// Password: TEMP_BYTE_ARRAY,
		}

		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Password is required."))
	})
}
