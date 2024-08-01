package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPaymentValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [Code] not blank", func(t *testing.T) {
		paymentMethod := PaymentMethod{
			// Code:         "Test Code",
			ProviderName: "Test ProviderName",
			AccountName:  "Test AccountName",
		}

		ok, err := govalidator.ValidateStruct(paymentMethod)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Code is required."))
	})

	t.Run("Check [ProviderName] not blank", func(t *testing.T) {
		paymentMethod := PaymentMethod{
			Code: "Test Code",
			// ProviderName: "Test ProviderName",
			AccountName: "Test AccountName",
		}

		ok, err := govalidator.ValidateStruct(paymentMethod)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("ProviderName is required."))
	})

	t.Run("Check [AccountName] not blank", func(t *testing.T) {
		paymentMethod := PaymentMethod{
			Code:         "Test Code",
			ProviderName: "Test ProviderName",
			// AccountName:  "Test AccountName",
		}

		ok, err := govalidator.ValidateStruct(paymentMethod)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("AccountName is required."))
	})
}
