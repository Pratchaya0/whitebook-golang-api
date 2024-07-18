package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestOrderBookDetailValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)
// 	tmpId := uint(1)

// 	t.Run("check [OrderId] not blank", func(t *testing.T) {
// 		orderBookDetail := OrderBookDetail{
// 			// OrderId: &tmpId,
// 			OrderBookDetailBookId: &tmpId,
// 		}

// 		ok, err := govalidator.ValidateStruct(orderBookDetail)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input order id"))
// 	})

// 	t.Run("check [BookId] not blank", func(t *testing.T) {
// 		orderBookDetail := OrderBookDetail{
// 			OrderId: &tmpId,
// 			// BookId:  &tmpId,
// 		}

// 		ok, err := govalidator.ValidateStruct(orderBookDetail)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input book id"))
// 	})
// }
