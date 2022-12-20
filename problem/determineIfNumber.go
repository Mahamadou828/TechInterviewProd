package problem

import "strconv"

type State int

const (
	BEGIN State = iota
	NEGATIVE1
	DIGIT1
	DOT
	DIGIT2
	E
	NEGATIVE2
	DIGIT3
)

var mapStateValidator = map[State]func(x string) bool{
	BEGIN:     func(x string) bool { return true },
	DIGIT1:    func(x string) bool { return isDigit(x) },
	NEGATIVE1: func(x string) bool { return x == "-" },
	DOT:       func(x string) bool { return x == "." },
	DIGIT2:    func(x string) bool { return isDigit(x) },
	E:         func(x string) bool { return x == "e" },
	NEGATIVE2: func(x string) bool { return x == "-" },
	DIGIT3:    func(x string) bool { return isDigit(x) },
}

var mapNextState = map[State][]State{
	BEGIN:     {NEGATIVE1, DIGIT1},
	NEGATIVE1: {DIGIT1, DOT},
	DIGIT1:    {DIGIT1, DOT, E},
	DOT:       {DIGIT2},
	DIGIT2:    {DIGIT2, E},
	NEGATIVE2: {DIGIT3},
	DIGIT3:    {DIGIT3},
}

var mapFinalState = map[State]bool{
	DIGIT1: true,
	DIGIT2: true,
	DIGIT3: true,
}

func DetermineIfNumber(str string) bool {
	state := BEGIN

	for _, c := range str {
		var found bool

		for _, nextState := range mapNextState[state] {
			if mapStateValidator[nextState](string(c)) {
				state = nextState
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	_, ok := mapFinalState[state]
	return ok
}

func isDigit(x string) bool {
	if _, err := strconv.Atoi(x); err == nil {
		return true
	}
	return false
}
