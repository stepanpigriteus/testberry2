package utils

//Сonfig - конфиг 
type Сonfig struct {
	after      int
	before     int
	context    int
	countOnly  bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	pattern    string
	filename   string
}

//Line - структура для вывода строк
type Line struct {
	lineNum int
	text    string
}
