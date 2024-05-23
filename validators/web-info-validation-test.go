package validators

import (
	"testing"

	. "github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestWebInfoValidateNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("check [WebInfoName] not blank", func(t *testing.T) {
		webInfo := WebInfo{
			// WebInfoName:     "WebInfoName",
			WebInfoEmail:    "WebInfoEmail",
			WebInfoPhone:    "WebInfoPhone",
			WebInfoFacebook: "WebInfoFacebook",
			WebInfoLine:     "WebInfoLine",
		}

		ok, err := govalidator.ValidateStruct(webInfo)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input Name"))
	})

	t.Run("check [WebInfoEmail] not blank", func(t *testing.T) {
		webInfo := WebInfo{
			WebInfoName: "WebInfoName",
			// WebInfoEmail:    "WebInfoEmail",
			WebInfoPhone:    "WebInfoPhone",
			WebInfoFacebook: "WebInfoFacebook",
			WebInfoLine:     "WebInfoLine",
		}

		ok, err := govalidator.ValidateStruct(webInfo)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input Email"))
	})

	t.Run("check [WebInfoPhone] not blank", func(t *testing.T) {
		webInfo := WebInfo{
			WebInfoName:  "WebInfoName",
			WebInfoEmail: "WebInfoEmail",
			// WebInfoPhone:    "WebInfoPhone",
			WebInfoFacebook: "WebInfoFacebook",
			WebInfoLine:     "WebInfoLine",
		}

		ok, err := govalidator.ValidateStruct(webInfo)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input Phone"))
	})

	t.Run("check [WebInfoFacebook] not blank", func(t *testing.T) {
		webInfo := WebInfo{
			WebInfoName:  "WebInfoName",
			WebInfoEmail: "WebInfoEmail",
			WebInfoPhone: "WebInfoPhone",
			// WebInfoFacebook: "WebInfoFacebook",
			WebInfoLine: "WebInfoLine",
		}

		ok, err := govalidator.ValidateStruct(webInfo)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input Facebook"))
	})

	t.Run("check [WebInfoLine] not blank", func(t *testing.T) {
		webInfo := WebInfo{
			WebInfoName:     "WebInfoName",
			WebInfoEmail:    "WebInfoEmail",
			WebInfoPhone:    "WebInfoPhone",
			WebInfoFacebook: "WebInfoFacebook",
			// WebInfoLine:     "WebInfoLine",
		}

		ok, err := govalidator.ValidateStruct(webInfo)
		g.Expect(ok).To(BeFalse())
		g.Expect(err).To(HaveOccurred())
		g.Expect(err.Error()).To(Equal("Please input Line"))
	})
}
