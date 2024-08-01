package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestOrderValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Check [RefCode] not blank.", func(t *testing.T) {
		order := Order{
			// RefCode:         "Temp RefCode",
			Amount:          123.00,
			SlipImage:       TEMP_LINK,
			UserID:          TEMP_UINT,
			PaymentMethodID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Reference code is required."))
	})

	t.Run("Check [Amount] not blank.", func(t *testing.T) {
		order := Order{
			RefCode: "Temp RefCode",
			// Amount:          123.00,
			SlipImage:       TEMP_LINK,
			UserID:          TEMP_UINT,
			PaymentMethodID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Amount is required."))
	})

	t.Run("Check [SlipImage] not blank.", func(t *testing.T) {
		order := Order{
			RefCode: "Temp RefCode",
			Amount:  123.00,
			// SlipImage:       TEMP_LINK,
			UserID:          TEMP_UINT,
			PaymentMethodID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Slip is required."))
	})

	t.Run("Check [UserID] not blank.", func(t *testing.T) {
		order := Order{
			RefCode:   "Temp RefCode",
			Amount:    123.00,
			SlipImage: TEMP_LINK,
			// UserID:          TEMP_UINT,
			PaymentMethodID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("UserID is required."))
	})

	t.Run("Check [PaymentMethodID] not blank.", func(t *testing.T) {
		order := Order{
			RefCode:   "Temp RefCode",
			Amount:    123.00,
			SlipImage: TEMP_LINK,
			UserID:    TEMP_UINT,
			// PaymentMethodID: TEMP_UINT,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("PaymentMethodID is required."))
	})
}
