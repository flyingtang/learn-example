package main

import "errors"

// 不考虑边界问题

func Power(base float64, exponent int) (float64, error) {

	// 边界问题
	// 负面问题
	// base - 0 +
	// exponent - 0 +

	//  -  -  （+倒数）
	//  - 0
	//  - +

	// 0 - 错误
	// 0 0 没意义
	// 0 + 0

	// + - 倒数
	// + 0  1
	// + + 完美

	//if base == 0 && exponent < 0 {
	//	return 0,errors.New("base < 0 && exponent < 0 is error")
	//}
	//if base == 0 && exponent == 0 {
	//	return  0, nil
	//}

	if equal(base, 0) && exponent < 0 {
		return 0,errors.New("base < 0 && exponent < 0 is error")
	}

	var uExponent  = uint(exponent)
	if exponent < 0 {
		uExponent = uint(-exponent)
	}
	var result = powerWithUnsignedExponent(base, uExponent)
	if exponent < 0 {
		result = 1.0 / result
	}
	return  result, nil
}


func powerWithUnsignedExponent(base float64, exponent uint) float64{
	// 分子策略
	if exponent == 0 {
		return 1
	}
	if exponent == 1{
		return  base
	}

	result := powerWithUnsignedExponent(base, exponent >> 1)
	result *= result

	if exponent & 0x1 == 1{
		result *= base
	}

	return result
}

// 比较两个浮点数
func equal(num1 , num2 float64) bool{

	if num1 - num2 > -0.00000001 && num1 - num2 < 0.0000001{
		return true
	}
	return  false
}