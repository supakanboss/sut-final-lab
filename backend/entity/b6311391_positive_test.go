package entity

import(
	"testing"
	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Customer struct{
	gorm.Model
	Name string `valid:"required"`
	Email string
	CustomerID string `valid:"matches(^[M L H]\\d{7}$)"`
}

func Test_positive (t *testing.T) {
	g := NewGomegaWithT(t)

	u := Customer{
		Name: "boss",
		Email: "Cryvery@gmail.com",
		CustomerID: "M2546875",
	}

	ok,err := govalidator.ValidateStruct(u)

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}