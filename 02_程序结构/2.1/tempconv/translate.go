package tempconv

// 华氏转摄氏
func (f Fahrenheit) FToC() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 摄氏转华氏
func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// 绝对转摄氏
func (k Kelvin) KToC() Celsius {
	return Celsius(k + Kelvin(AbsoluteZeroC))
}

// 摄氏转绝对
func (c Celsius) CToK() Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

// 华氏转绝对
func (f Fahrenheit) FToK() Kelvin {
	return f.FToC().CToK()
}

// 绝对转华氏
func (k Kelvin) KToF() Fahrenheit {
	return k.KToC().CToF()
}
