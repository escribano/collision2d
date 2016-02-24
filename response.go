package sat

import ()

type Response struct {
	A, B               interface{}
	Overlap            float64
	OverlapN, OverlapV Vector
	AInB, BInA         bool
}

func (response Response) Clear() Response {
	response.AInB = true
	response.BInA = true
	response.Overlap = 0
	return response
}
