package calc

import "math"

func Maximum(numbers ... float64) float64{
	max:=math.Inf(-1)
	for _,value:=range numbers{
		if value > max{
			max=value
		}
	}
	return max	
}