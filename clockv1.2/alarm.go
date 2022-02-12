package main

// alarmHolder holds the array of letters A L A R M
type lettersHolder [5]string

// alarm are the letter to display when alarm is set
var (
	a = lettersHolder{
		"███",
		"█ █",
		"███",
		"█ █",
		"█ █",
	}

	l = lettersHolder{
		"█  ",
		"█  ",
		"█  ",
		"█  ",
		"███",
	}

	r = lettersHolder{
		"██ ",
		"█ █",
		"██ ",
		"█ █",
		"█ █",
	}

	m = lettersHolder{
		"█ █",
		"███",
		"█ █",
		"█ █",
		"█ █",
	}

	qm = lettersHolder{
		" █ ",
		" █ ",
		" █ ",
		"   ",
		" █ ",
	}

)
