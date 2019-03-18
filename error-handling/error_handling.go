package erratum

const testVersion = 2

func open(o ResourceOpener) (Resource, error) {
	res, err := o()
	if err == nil {
		return res, err
	}

	if _, ok := err.(TransientError); ok {
		return open(o)
	}

	return res, err
}

func Use(o ResourceOpener, input string) (err error) {
	res, err := open(o)
	if err != nil {
		return err
	}
	defer res.Close()

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
			if frobErr, ok := err.(FrobError); ok {
				res.Defrob(frobErr.defrobTag)
			}

		}
	}()
	res.Frob(input)
	return err
}
