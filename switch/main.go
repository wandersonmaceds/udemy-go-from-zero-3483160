package main

func weekday(index int) string {
	switch index {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturnday"
	}
	return "Invalid day number"
}

// using case to match the var directly
func yearMonth(index int) string {
	switch {
	case index == 1:
		return "January"
	case index == 2:
		return "February"
	case index == 3:
		return "March"
	case index == 4:
		return "April"
	case index == 5:
		return "Mail"
	case index == 6:
		return "June"
	case index == 7:
		return "July"
	case index == 8:
		return "August"
	case index == 9:
		return "September"
	case index == 10:
		return "Octuber"
	case index == 11:
		return "November"
	case index == 12:
		return "December"
	}
	return "Invalid month number"
	// we could also have a defult clause
	// go has a fallthrough clause that make a case jump to the next one
}

func test_break(index int) int {
	var returned int
	switch index {
	case 1:
		println("case 1")
		returned = index + 10
	case 2:
		println("case 2")
		returned = index + 20
	}
	return returned
}

func main() {
	var day = weekday(3)
	var month = yearMonth(8)
	var calc = test_break(1)
	println(day, month, calc)
}
