package components

func checkRows(b *NewBoard, mark string) bool {
	ret := true
	for i := 0; i < (b.size * b.size); i = i + b.size {
		ret = true
		for j := i; j < (i + b.size); j++ {
			if b.values[j] != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}

func checkColumns(b *NewBoard, mark string) bool {
	ret := true
	for i := 0; i < b.size; i++ {
		ret = true
		for j := i; j < (b.size * b.size); j = j + b.size {
			if b.values[j] != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}
func checkDiagonal(b *NewBoard, mark string) bool {
	ret := true
	for i := 0; i < b.size; i++ {
		if b.values[b.size*i+i] != mark {
			ret = false
		}
	}
	if ret {
		return ret
	}
	ret = true
	for i := 0; i < b.size; i++ {
		if b.values[(b.size*i)+(b.size-1-i)] != mark {
			ret = false
		}
	}
	return ret
}
func checkFull(b *NewBoard) bool {
	for i := 0; i < b.size*b.size; i++ {
		if b.values[i] == " " {
			return false
		}
	}
	return true
}

func (b *NewBoard) Analyse(mark string) (bool, string) {

	if checkRows(b, mark) {
		return true, mark
	} else if checkColumns(b, mark) {
		return true, mark
	} else if checkDiagonal(b, mark) {
		return true, mark
	} else if checkFull(b) {
		return true, "Draw"
	}
	return false, ""
}
