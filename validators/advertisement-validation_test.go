package validators

// import (
// 	"testing"

// 	. "github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	. "github.com/onsi/gomega"
// )

// func TestAdvertiseValidateNotBlank(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	t.Run("check [AdvertisementTitle] not blank", func(t *testing.T) {
// 		advertisement := Advertisement{
// 			// AdvertisementTitle:       "AdvertisementTitle",
// 			AdvertisementDescription: "AdvertisementDescription",
// 			AdvertisementHighlight:   "AdvertisementHighlight",
// 			AdvertisementImageUrl:    "AdvertisementImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(advertisement)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input title"))
// 	})

// 	t.Run("check [AdvertisementDescription] not blank", func(t *testing.T) {
// 		advertisement := Advertisement{
// 			AdvertisementTitle: "AdvertisementTitle",
// 			// AdvertisementDescription: "AdvertisementDescription",
// 			AdvertisementHighlight: "AdvertisementHighlight",
// 			AdvertisementImageUrl:  "AdvertisementImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(advertisement)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input description"))
// 	})

// 	t.Run("check [AdvertisementHighlight] not blank", func(t *testing.T) {
// 		advertisement := Advertisement{
// 			AdvertisementTitle:       "AdvertisementTitle",
// 			AdvertisementDescription: "AdvertisementDescription",
// 			// AdvertisementHighlight:  "AdvertisementHighlight",
// 			AdvertisementImageUrl: "AdvertisementImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(advertisement)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input highlight"))
// 	})

// 	t.Run("check [AdvertisementImageUrl] not blank", func(t *testing.T) {
// 		advertisement := Advertisement{
// 			AdvertisementTitle:       "AdvertisementTitle",
// 			AdvertisementDescription: "AdvertisementDescription",
// 			AdvertisementHighlight:   "AdvertisementHighlight",
// 			// AdvertisementImageUrl:    "AdvertisementImageUrl",
// 		}

// 		ok, err := govalidator.ValidateStruct(advertisement)
// 		g.Expect(ok).To(BeFalse())
// 		g.Expect(err).To(HaveOccurred())
// 		g.Expect(err.Error()).To(Equal("Please input image"))
// 	})
// }
