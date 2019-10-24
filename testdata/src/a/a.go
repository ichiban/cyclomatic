package main

func main() { // want main:`complexity\(1\)`
	tooManyIfs()
	tooManyFors()
	tooManyCases()
	tooManyComms()
	tooManyBinaryOps()
}

func tooManyIfs() { // want `cyclomatic complexity of a\.tooManyIfs exceeded limit 11 > 10` tooManyIfs:`complexity\(11\)`
	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}

	if false {
	}
}

func tooManyFors() { // want `cyclomatic complexity of a\.tooManyFors exceeded limit 11 > 10` tooManyFors:`complexity\(11\)`
	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}

	for false {
	}
}

func tooManyCases() { // want `cyclomatic complexity of a\.tooManyCases exceeded limit 11 > 10` tooManyCases:`complexity\(11\)`
	switch 0 {
	case 1, 2, 3, 4, 5:
	case 6, 7, 8, 9, 10:
	}
}

func tooManyComms() { // want `cyclomatic complexity of a\.tooManyComms exceeded limit 11 > 10` tooManyComms:`complexity\(11\)`
	ch := make(chan struct{}, 1)
	defer close(ch)

	ch <- struct{}{}

	select {
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	case <-ch:
	}
}

func tooManyBinaryOps() { // want `cyclomatic complexity of a\.tooManyBinaryOps exceeded limit 11 > 10` tooManyBinaryOps:`complexity\(11\)`
	if true && true && true && true && true || true && true && true && true && true {
	}
}
