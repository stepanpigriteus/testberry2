package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError { // теперь customError реализует интерфейс error
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()    // возвращается тип customError с значением nil
	if err != nil { // он не равен nil
		println("error")
		return
	}
	println("ok")
}

// выведет error
