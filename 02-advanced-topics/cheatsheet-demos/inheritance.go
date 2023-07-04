package main

import "fmt"

//base class
type Discount struct {
	percent float32
}

//class containing embedding
type PremiumDiscount struct {
	Discount   //embedding
	additional float32
}

func (d *Discount) Calculate(originalPrice float64) float64 {
	return originalPrice - originalPrice*float64(d.percent)/100
}

func (pd *PremiumDiscount) CalculateAdditional(originalPrice float64) float64 {
	//percent belongs to embedded Discount struct
	//additional belongs to PremiumDiscount struct
	return originalPrice - (originalPrice*float64(pd.percent+pd.additional))/100
}

///alternative way of accessing base classes parameters : via full path
// func (pd *PremiumDiscount) CalculateAdditional(originalPrice float64) float64 {
// 	//percent uses full path
// 	return originalPrice - originalPrice*float64(pd.Discount.percent+pd.additional)/100
// }

func main() {
	price := 100.00
	b := Discount{percent: 25}
	d := PremiumDiscount{Discount: Discount{percent: 25}, additional: 10}
	fmt.Println(b.Calculate(price))
	fmt.Println(d.CalculateAdditional(price))
}
