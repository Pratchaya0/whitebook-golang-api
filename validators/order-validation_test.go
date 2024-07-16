package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestOrderValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)
	tmpId := uint(1)

	t.Run("check [OrderAmount] not blank", func(t *testing.T) {
		order := Order{
			// OrderAmount:          100,
			OrderPaymentImageUrl: "OrderPaymentImageUrl",
			OrderIsPaid:          true,
			PaymentInfoId:        &tmpId,
			UserId:               &tmpId,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input order amount"))
	})

	t.Run("check [OrderPaymentImageUrl] not blank", func(t *testing.T) {
		order := Order{
			OrderAmount: 100,
			// OrderPaymentImageUrl: "OrderPaymentImageUrl",
			OrderIsPaid:   true,
			PaymentInfoId: &tmpId,
			UserId:        &tmpId,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input slip"))
	})

	t.Run("check [OrderIsPaid] not blank", func(t *testing.T) {
		order := Order{
			OrderAmount:          100,
			OrderPaymentImageUrl: "OrderPaymentImageUrl",
			// OrderIsPaid:          true,
			PaymentInfoId: &tmpId,
			UserId:        &tmpId,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input is paid"))
	})

	t.Run("check [PaymentInfoId] not blank", func(t *testing.T) {
		order := Order{
			OrderAmount:          100,
			OrderPaymentImageUrl: "OrderPaymentImageUrl",
			OrderIsPaid:          true,
			// PaymentInfoId:        &tmpId,
			UserId: &tmpId,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input payment info id"))
	})

	t.Run("check [UserId] not blank", func(t *testing.T) {
		order := Order{
			OrderAmount:          100,
			OrderPaymentImageUrl: "OrderPaymentImageUrl",
			OrderIsPaid:          true,
			PaymentInfoId:        &tmpId,
			// UserId:               &tmpId,
		}

		ok, err := govalidator.ValidateStruct(order)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input user id"))
	})
}
