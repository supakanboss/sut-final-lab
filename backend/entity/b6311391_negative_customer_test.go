package entity

import(
	"testing"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
)

func Test_negative_customer (t *testing.T) {
	g := NewGomegaWithT(t)

	u := Customer{
		Name: "Boss",
		Email: "Cryvery@gmail.com",
		CustomerID: "254687",
	}

	ok,err := govalidator.ValidateStruct(u)

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("CustomerID: 254687 does not validate as matches(^[M L H]\\d{7}$)"))
}