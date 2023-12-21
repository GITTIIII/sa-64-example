package unit

import (
	"testing"

	"github.com/GITTIIII/sa-66-example/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestMember(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`Member pattern is true`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "1234",
			Url:         "https://www.google.com/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`PhoneNumber pattern is not true`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "45678", // ผิดตรงนี้
			Password:    "1234",
			Url:         "https://www.google.com/",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PhoneNumber length is not 10 digits."))
	})

	t.Run(`Url pattern is not true`, func(t *testing.T) {
		member := entity.Member{
			PhoneNumber: "0123456789",
			Password:    "1234",
			Url:         "httpswwwgooglecom", // ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Url is invalid"))
	})
}
