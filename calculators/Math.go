package calculators

import "github.com/iVitaliya/logger-go"

type ICalcInt struct {
	Devision func() int
	Minus    func() int
	Modulo   func() int
	Times    func() int
	Plus     func() int
}

type ICalcFloat struct {
	Devision func() float64
	Minus    func() float64
	Plus     func() float64
}

type ICalc struct {
	Int   func(int) *ICalcInt
	Float func(float64) *ICalcFloat
}

func Calc(base interface{}) *ICalc {
	Integer := func(additional int) *ICalcInt {
		var b, ok = base.(int)
		if !ok {
			logger.LogNextLine(1)
			logger.Fatal("The base number couldn't be properly translated to an integer!")
		}

		return &ICalcInt{
			Devision: func() int {
				var res = (b / additional)

				return res
			},
			Minus: func() int {
				var res = (b - additional)

				return res
			},
			Modulo: func() int {
				var res = (b % additional)

				return res
			},
			Times: func() int {
				var res = (b * additional)

				return res
			},
			Plus: func() int {
				var res = (b + additional)

				return res
			},
		}
	}

	Float := func(additional float64) *ICalcFloat {
		var b, ok = base.(float64)
		if !ok {
			logger.LogNextLine(1)
			logger.Fatal("The base number couldn't be properly translated to an float64!")
		}

		return &ICalcFloat{
			Devision: func() float64 {
				var res = (b / additional)

				return res
			},
			Minus: func() float64 {
				var res = (b - additional)

				return res
			},
			Plus: func() float64 {
				var res = (b + additional)

				return res
			},
		}
	}

	return &ICalc{
		Int:   Integer,
		Float: Float,
	}
}
