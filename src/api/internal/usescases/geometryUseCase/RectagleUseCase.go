package geometryUseCase

import "GolanApiTemplate/src/api/internal/model"

type rectangle struct {
	r model.RectagleModel
}

func (rec *rectangle) area() float64 {
	return rec.r.With * rec.r.Height
}
