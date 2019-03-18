package clock

import "fmt"

func ExampleClock_new() {
	// a new c
	clock1 := New(10, 30)
	fmt.Println(clock1.String())

	// Output: 10:30
}

func ExampleClock_Add_add() {
	// create a c
	clock := New(10, 30)

	// add 30 minutes to it
	clock = clock.Add(30)
	fmt.Println(clock.String())

	// Output: 11:00
}

func ExampleClock_Add_subtract() {
	// create a c
	clock := New(10, 30)

	// subtract an hour and a half from it
	clock = clock.Add(-90)
	fmt.Println(clock.String())

	// Output: 09:00
}

func ExampleClock_compare() {
	// a new c
	clock1 := New(10, 30)

	// a second c, same as the first
	clock2 := New(10, 30)

	// are the clocks equal?
	fmt.Println(clock2 == clock1)

	// change the second c
	clock2 = clock2.Add(30)

	// are the clocks equal now?
	fmt.Println(clock2 == clock1)

	// Output:
	// true
	// false
}
