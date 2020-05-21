package components

func checkRows(b *Board, mark string) bool {
	ret := true
	for i := 0; i < (b.Size * b.Size); i = i + b.Size {
		ret = true
		for j := i; j < (i + b.Size); j++ {
			if b.Values[j] != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}

func checkColumns(b *Board, mark string) bool {
	ret := true
	for i := 0; i < b.Size; i++ {
		ret = true
		for j := i; j < (b.Size * b.Size); j = j + b.Size {
			if b.Values[j] != mark {
				ret = false
			}
		}
		if ret {
			return ret
		}
	}
	return ret
}
func checkDiagonal(b *Board, mark string) bool {
	ret := true
	for i := 0; i < b.Size; i++ {
		if b.Values[b.Size*i+i] != mark {
			ret = false
		}
	}
	if ret {
		return ret
	}
	ret = true
	for i := 0; i < b.Size; i++ {
		if b.Values[(b.Size*i)+(b.Size-1-i)] != mark {
			ret = false
		}
	}
	return ret
}
func checkFull(b *Board) bool {
	for i := 0; i < b.Size*b.Size; i++ {
		if b.Values[i] == " " {
			return false
		}
	}
	return true
}

func (b *Board) Analyse(mark string) (bool, string) {

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
