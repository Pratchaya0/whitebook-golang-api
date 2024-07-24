package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestPaymentInfoValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	t.Run("check [PaymentInfoName] not blank", func(t *testing.T) {
// 		paymentInfo := PaymentInfo{
// 			// PaymentInfoName:     "PaymentInfoName",
// 			PaymentInfoCode:     "PaymentInfoCode",
// 			PaymentInfoImageUrl: "PaymentInfoImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(paymentInfo)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input name"))
// 	})

// 	t.Run("check [PaymentInfoCode] not blank", func(t *testing.T) {
// 		paymentInfo := PaymentInfo{
// 			PaymentInfoName: "PaymentInfoName",
// 			// PaymentInfoCode:     "PaymentInfoCode",
// 			PaymentInfoImageUrl: "PaymentInfoImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(paymentInfo)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input code"))
// 	})
// }
