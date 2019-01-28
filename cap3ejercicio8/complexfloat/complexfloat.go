package complexfloat
complexfloat

import "math/big"


type ComplexFloat struct {
	Re   *big.Float
	Im   *big.Float
	Prec uint
}



func (c *ComplexFloat) AbsCompare(n *big.Float) int {
	re2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Re)
	im2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Im, c.Im)
	n2 := big.NewFloat(0).SetPrec(c.Prec).Mul(n, n)
	return big.NewFloat(0).SetPrec(c.Prec).Add(re2, im2).Cmp(n2)
}


func (c *ComplexFloat) Square() *ComplexFloat {
	re2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Re)
	im2 := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Im, c.Im)
	reim := big.NewFloat(0).SetPrec(c.Prec).Mul(c.Re, c.Im)

	c.Re.Sub(re2, im2)
	c.Im.Add(reim, reim)
	return c
}


func (c *ComplexFloat) Add(z *ComplexFloat) *ComplexFloat {
	c.Re.Add(c.Re, z.Re)
	c.Im.Add(c.Im, z.Im)
	return c
}