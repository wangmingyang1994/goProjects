package FatRateCalc

func GetBMI(weight, tall float64) (Bmi float64) {
	Bmi = weight / (tall * tall)
	return
}

func GetmanFatRate(bmi float64, age int) (FatRate float64) {
	FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*1) / 100
	return
}
func GetwomanFatRate(bmi float64, age int) (FatRate float64) {
	FatRate = (1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*0) / 100
	return
}
