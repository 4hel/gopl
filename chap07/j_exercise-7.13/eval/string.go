package eval

import (
	"fmt"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.2f", float64(l))
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	var retval string
	retval += c.fn + "("
	for i, a := range c.args {
		if i > 0 {
			retval += ", "
		}
		retval += a.String()
	}
	retval += ")"
	return retval
}
