package calc

func Sum(numbers...float64) float64{ // variadic function taking slice as argument
	var sum float64
	for _,value:=range numbers{
		sum=sum+value
	}
	return sum	
}