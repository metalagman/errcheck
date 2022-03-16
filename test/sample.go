package test

import "errors"

func mulfunc(i int) (int, error) {
	return i * 2, nil
}

func TestFunc() {
	var i int
	myfunc := func() error {
		return nil
	}
	myfunc()
	if true {
		i := 7
		i, _ = mulfunc(i)
	}
	i, _ = i+1, myfunc()

	errfunc := func() error {
		return errors.New("test")
	}
	go errfunc()

	defer errfunc()
}
