package main

// Declare a holder array for holding the digital digits from 0 - 9
// and the hour, minute, second separator
type holder [5]string

// digit zero, one, two, three and so on...
var (
	zero = holder{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one = holder{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two = holder{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three = holder{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four = holder{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five = holder{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six = holder{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven = holder{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight = holder{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine = holder{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	sep = holder{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}

	a = holder{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	}

	l = holder{
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	}

	r = holder{
		"██ ",
		"█ █",
		"██ ",
		"█ █",
		"█ █",
	}

	m = holder{
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	}

	qm = holder{
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	}

	blank = holder{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	// digits array stores all digits in sub array type holder
	digits = [...]holder{zero, one, two, three, four, five, six, seven, eight, nine, sep}

	// alam array stores the letter ALARM! in subarray holder
	alarm = [...]holder{blank, a, l, a, r, m, qm, blank}
)
