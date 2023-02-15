package entity

import(
	"testing"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func Test_negative_name (t *testing.T) {
	g := NewGomegaWithT(t)

	u := Customer{
		Name: "",
		Email: "Cryvery@gmail.com",
		CustomerID: "M2546875",
	}

	ok,err := govalidator.ValidateStruct(u)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name: non zero value required"))
}