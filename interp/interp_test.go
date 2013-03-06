// Copyright 2012 Sonia Keys
// License MIT: http://www.opensource.org/licenses/MIT

package interp_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/soniakeys/meeus/common"
	"github.com/soniakeys/meeus/interp"
)

func ExampleLen3_InterpolateN() {
	// Example 3.a, p. 25.
	d3, err := interp.NewLen3(7, 9, []float64{
		.884226,
		.877366,
		.870531,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	n := 4.35 / 24
	y, err := d3.InterpolateN(n, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6f\n", y)
	// Output:
	// 0.876125
}

func ExampleLen3_InterpolateX() {
	// Example 3.a, p. 25.
	d3, err := interp.NewLen3(7, 9, []float64{
		.884226,
		.877366,
		.870531,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	x := 8 + common.NewTime(false, 4, 21, 0).Day() // 8th day at 4:21
	y, err := d3.InterpolateX(x, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.6f\n", y)
	// Output:
	// 0.876125
}

func ExampleLen3_Extremum() {
	// Example 3.b, p. 26.
	d3, err := interp.NewLen3(12, 20, []float64{
		1.3814294,
		1.3812213,
		1.3812453,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	x, y, err := d3.Extremum()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("distance:  %.7f\n", y)
	fmt.Printf("date:     %.4f\n", x)
	// Output:
	// distance:  1.3812030
	// date:     17.5864
}

func ExampleLen3_Zero() {
	// Example 3.c, p. 26.
	x1 := 26.
	x3 := 28.
	// the y unit doesn't matter.  working in degrees is fine
	yTable := []float64{
		common.DMSToDeg(true, 0, 28, 13.4),
		common.DMSToDeg(false, 0, 6, 46.3),
		common.DMSToDeg(false, 0, 38, 23.2),
	}
	d3, err := interp.NewLen3(x1, x3, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := d3.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.5f\n", x)
	// Output:
	// 26.79873
}

func ExampleLen3_Zero_strong() {
	// Example 3.d, p. 27.
	x1 := -1.
	x3 := 1.
	yTable := []float64{-2, 3, 2}
	d3, err := interp.NewLen3(x1, x3, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	x, err := d3.Zero(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.12f\n", x)
	// Output:
	// -0.720759220056
}

func ExampleLen5_InterpolateX() {
	// Example 3.e, p. 28.
	x1 := 27.
	x5 := 29.
	// work in radians to get answer in radians
	yTable := []float64{
		common.NewAngle(false, 0, 54, 36.125).Rad(),
		common.NewAngle(false, 0, 54, 24.606).Rad(),
		common.NewAngle(false, 0, 54, 15.486).Rad(),
		common.NewAngle(false, 0, 54, 08.694).Rad(),
		common.NewAngle(false, 0, 54, 04.133).Rad(),
	}
	x := 28 + (3+20./60)/24
	d5, err := interp.NewLen5(x1, x5, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	y, err := d5.InterpolateX(x, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	// radians easy to format
	fmt.Printf("%.3d", common.NewFmtAngle(y))
	// Output:
	// 54′13″.369
}

func ExampleLen5_Zero() {
	// Exercise, p. 30.
	x1 := 25.
	x5 := 29.
	yTable := []float64{
		common.DMSToDeg(true, 1, 11, 21.23),
		common.DMSToDeg(true, 0, 28, 12.31),
		common.DMSToDeg(false, 0, 16, 07.02),
		common.DMSToDeg(false, 1, 01, 00.13),
		common.DMSToDeg(false, 1, 45, 46.33),
	}
	d5, err := interp.NewLen5(x1, x5, yTable)
	if err != nil {
		fmt.Println(err)
		return
	}
	z, err := d5.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("1988 January %.6f\n", z)
	zInt, zFrac := math.Modf(z)
	fmt.Printf("1988 January %d at %s TD\n", int(zInt),
		common.NewFmtTime(zFrac*24*3600))

	// compare result to that from just three central values
	d3, err := interp.NewLen3(26, 28, yTable[1:4])
	if err != nil {
		fmt.Println(err)
		return
	}
	z3, err := d3.Zero(false)
	if err != nil {
		fmt.Println(err)
		return
	}
	dz := z - z3
	fmt.Printf("%.6f day\n", dz)
	fmt.Printf("%.1f minute\n", dz*24*60)
	// Output:
	// 1988 January 26.638587
	// 1988 January 26 at 15ʰ19ᵐ34ˢ TD
	// 0.000753 day
	// 1.1 minute
}

func ExampleLen4Half() {
	// Example 3.f, p. 32.
	half, err := interp.Len4Half([]float64{
		common.NewRA(10, 18, 48.732).Rad(),
		common.NewRA(10, 23, 22.835).Rad(),
		common.NewRA(10, 27, 57.247).Rad(),
		common.NewRA(10, 32, 31.983).Rad(),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.3d", common.NewFmtRA(half))
	// Output:
	// 10ʰ25ᵐ40ˢ.001
}

// exercise, p. 34.
func TestLagrange(t *testing.T) {
	table := []struct{ X, Y float64 }{
		{29.43, .4913598528},
		{30.97, .5145891926},
		{27.69, .4646875083},
		{28.11, .4711658342},
		{31.58, .5236885653},
		{33.05, .5453707057},
	}
	if math.Abs(interp.Lagrange(30, table)-.5) > 1e-5 {
		t.Fatal(30)
	}
	if math.Abs(interp.Lagrange(0, table)) > 1e-4 {
		t.Fatal(0)
	}
	if math.Abs(interp.Lagrange(90, table)-1) > 1e-4 {
		t.Fatal(90)
	}
}

func ExampleLagrangePoly() {
	// Example 3.g, p, 34.
	table := []struct{ X, Y float64 }{
		{1, -6},
		{3, 6},
		{4, 9},
		{6, 15},
	}
	p := interp.LagrangePoly(table)
	for _, c := range p {
		fmt.Printf("%.0f\n", c*5)
	}
	// Output:
	// -87
	// 69
	// -13
	// 1
}
