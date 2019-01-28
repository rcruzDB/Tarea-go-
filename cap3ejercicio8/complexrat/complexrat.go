package complexrat

import "math/big"


type ComplexRat struct {
	Re *big.Rat
	Im *big.Rat
}



func (c *ComplexRat) AbsCompare(n *big.Rat) int {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	n2 := big.NewRat(0, 1).Mul(n, n)
	return big.NewRat(0, 1).Add(re2, im2).Cmp(n2)
}


func (c *ComplexRat) Square() *ComplexRat {
	re2 := big.NewRat(0, 1).Mul(c.Re, c.Re)
	im2 := big.NewRat(0, 1).Mul(c.Im, c.Im)
	reim := big.NewRat(0, 1).Mul(c.Re, c.Im)

	c.Re.Sub(re2, im2)
	c.Im.Add(reim, reim)
	return c
}


func (c *ComplexRat) Add(z *ComplexRat) *ComplexRat {
	c.Re.Add(c.Re, z.Re)
	c.Im.Add(c.Im, z.Im)
	return c
}