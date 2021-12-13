package main

import (
	"fmt"
	"math"
)

const (
	k = 1000.0
	M = 1000000.0
	G = 1000000000.0
	T = 1000000000000.0

	m = 0.001
	u = 0.000001
	n = 0.000000001
	p = 0.000000000001
)

func start() {
	//fixedPolarizationCircuit(16, 2, 2000, 1000000, 0.01, -8)
	//autoPolarizationCircuit(20, 3300, 1000000, 1000, 0.008, -6)
	//voltageDividerBiasCircuit(16, 2100000, 2400, 270000, 1500, 0.008, -4)
	res := 0
	circuit := make([]float64, 0, 7)

	fmt.Println("1: Fixed Polarization Circuit\n2: Auto Polarization Circuit\n3: Voltage Divider BiasCircuit")
	fmt.Scanf("%d", &res)

	switch res {
	case 1:
		fmt.Println("\nFixed Polarization Circuit\n ")

		temp := 0.0
		fmt.Print("vdd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("vgg: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("rd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("idss: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("vp: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fixedPolarizationCircuit(circuit[0], circuit[1], circuit[2], circuit[3], circuit[4])

	case 2:
		fmt.Println("\nAuto Polarization Circuit\n ")

		temp := 0.0
		fmt.Print("vdd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("rd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("rs: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("idss: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("vp: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		autoPolarizationCircuit(circuit[0], circuit[1], circuit[2], circuit[3], circuit[4])

	case 3:
		fmt.Println("\nVoltage Divider Bias Circuit\n ")

		temp := 0.0
		fmt.Print("vdd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("r1: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("rd: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("r2: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("rs: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("idss: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		fmt.Print("vp: ")
		fmt.Scanf("%f", &temp)
		circuit = append(circuit, temp)

		voltageDividerBiasCircuit(circuit[0], circuit[1], circuit[2], circuit[3], circuit[4], circuit[5], circuit[6])

	}

}

func printer(val float64) string {
	temp := ""

	if math.Abs(val) >= k {
		val /= 1000
		temp = "k"
	}
	if math.Abs(val) >= T {
		val /= 1000
		temp = "T"
	}
	if math.Abs(val) >= M {
		val /= 1000
		temp = "M"
	}
	if math.Abs(val) >= G {
		val /= 1000
		temp = "G"
	}
	if math.Abs(val) <= m*1000 {
		val *= 1000
		temp = "m"
	}
	if math.Abs(val) <= u*1000 {
		val *= 1000
		temp = "u"
	}
	if math.Abs(val) <= n*1000 {
		val *= 1000
		temp = "n"
	}
	if math.Abs(val) <= p*1000 {
		val *= 1000
		temp = "p"
	}

	return fmt.Sprintf(" %7.4f %s", val, temp)
}

func fixedPolarizationCircuit(vdd float64, vgg float64, rd float64, idss float64, vp float64) {
	vgs := -vgg
	id  := idss * (1 - (vgs / vp)) * (1 - (vgs / vp))
	vds := vdd - id * rd
	vd  := vds
	vg  := vgs
	vs  := 0.0

	//fmt.Println("        vdd\n	     │\n	   ║─┘\n	┌─►║   idss\n	│  ║─┐ Vp\n	Rg   │\n	│    │\n   vgg   │\n	└────┘")

	fmt.Printf("vgs = %sV\n", printer(vgs))
	fmt.Printf("id  = %sA\n", printer(id))
	fmt.Printf("vds = %sV\n", printer(vds))
	fmt.Printf("vd  = %sV\n", printer(vd))
	fmt.Printf("vg  = %sV\n", printer(vg))
	fmt.Printf("vs  = %sV\n\n", printer(vs))

	n := 1

	if math.Abs(vp) < 5 {
		n = 2
	} else if math.Abs(vp) < 3 {
		n = 4
	}

	fvp(vp, idss, n)
}

func autoPolarizationCircuit(vdd float64, rd float64, rs float64, idss float64, vp float64){
	a := (idss * rs * rs) / (vp * vp)
	b := ((2 * idss * rs) / vp) - 1
	c := idss

	sqrt := b * b - 4 * a * c

	if sqrt < 0.00000000000001 && sqrt > -0.00000000000001 {
		sqrt = 0
	}

	f := (-b + math.Sqrt(sqrt)) / (2 * a)

	if f > idss || f > (-b - math.Sqrt(sqrt)) / (2 * a) {
		f = (-b - math.Sqrt(sqrt)) / (2 * a)
	}

	id  := f
	vgs := -rs * id
	vds := vdd - id * (rs + rd)
	vs  := id * rs
	vg  := 0.0
	vd  := vdd - id * rd

	fmt.Printf("vgs = %sV\n", printer(vgs))
	fmt.Printf("id  = %sA\n", printer(id))
	fmt.Printf("vds = %sV\n", printer(vds))
	fmt.Printf("vd  = %sV\n", printer(vd))
	fmt.Printf("vg  = %sV\n", printer(vg))
	fmt.Printf("vs  = %sV\n\n", printer(vs))

	n := 1

	if math.Abs(vp) < 5 {
		n = 2
	} else if math.Abs(vp) < 3 {
		n = 4
	}

	fvp(vp, idss, n)
}

func voltageDividerBiasCircuit(vdd float64, r1 float64, rd float64, r2 float64, rs float64, idss float64, vp float64) {
	vg := (r2 * vdd) / (r1 + r2)

	a := rs * rs * idss
	b := 2 * rs * idss * vp - 2 * rs * vg * idss - vp * vp
	c := idss * vp * vp - 2 * vg * idss * vp + idss * vg * vg

	sqrt := b * b - 4 * a * c

	if sqrt < 0.00000000000001 && sqrt > -0.00000000000001 {
		sqrt = 0
	}

	f := (-b + math.Sqrt(sqrt)) / (2 * a)

	if f > idss || f > (-b - math.Sqrt(sqrt)) / (2 * a) {
		f = (-b - math.Sqrt(sqrt)) / (2 * a)
	}

	id  := f
	vgs := vg - rs * id
	vds := vdd - id * (rd + rs)
	vs  := id * rs
	vd  := vdd - id * rd

	fmt.Printf("vgs = %sV\n", printer(vgs))
	fmt.Printf("id  = %sA\n", printer(id))
	fmt.Printf("vds = %sV\n", printer(vds))
	fmt.Printf("vd  = %sV\n", printer(vd))
	fmt.Printf("vg  = %sV\n", printer(vg))
	fmt.Printf("vs  = %sV\n\n", printer(vs))

	n := 1

	if math.Abs(vp) < 3 {
		n = 4
	} else if math.Abs(vp) < 5 {
		n = 2
	}

	fvp(vp, idss, n)
}

func fvp(vp float64, idss float64, n int) {
	fmt.Println("┌─────────┬─────────────┐\n|   vp    |     id      |\n├─────────┼─────────────┤")

	for i := vp * float64(n); i <= 0; i++ {
		j := i / float64(n)
		fmt.Printf("| %7.2f | %sA |\n", j, printer(idss * (1 - (j / (vp))) * (1 - (j / (vp)))))
	}

	fmt.Println("└─────────┴─────────────┘")
}

func main() {
	start()
}
