package main

import "fmt"

type allType struct {
	i  int
	f  float64
	s  string
	a  [3]string
	sl []int
	m  map[string]int
}

func modifyVal(st allType) {
	st.sl[0] *= 2
	st.i *= 2
	st.f *= 2.0
	st.a[0] = st.a[0] + st.a[0]
	st.s = st.s + st.s
	st.m["a"] *= 2
}

func modifyPtr(st *allType) {
	st.sl[0] *= 2
	st.i *= 2
	st.f *= 2.0
	st.a[0] = st.a[0] + st.a[0]
	st.s = st.s + st.s
	st.m["a"] *= 2
}

func displaychanges(org, temp, temp1 allType) {
	ch := map[bool]string{true: "no change", false: "change"}

	fmt.Println("int->       value:", ch[org.i == temp.i], "\tpointer:", ch[org.i == temp1.i])
	fmt.Println("float64->   value:", ch[org.f == temp.f], "\tpointer:", ch[org.f == temp1.f])
	fmt.Println("string->    value:", ch[org.s == temp.s], "\tpointer:", ch[org.s == temp1.s])
	fmt.Println("arrays->    value:", ch[org.a[0] == temp.a[0]], "\tpointer:", ch[org.a[0] == temp1.a[0]])
	fmt.Println("slice->     value:", ch[org.sl[0] == temp.sl[0]], "\tpointer:", ch[org.sl[0] == temp1.sl[0]])
	fmt.Println("map->       value:", ch[org.m["a"] == temp.m["a"]], "\tpointer:", ch[org.m["a"] == temp1.m["a"]])

}

func main() {
	orgVal := allType{
		i:  10,
		f:  10.23,
		s:  "ayush",
		a:  [3]string{"a", "b", "c"},
		sl: []int{1, 2},
		m:  map[string]int{"a": 1, "b": 2},
	}
	temp := allType{
		i:  10,
		f:  10.23,
		s:  "ayush",
		a:  [3]string{"a", "b", "c"},
		sl: []int{1, 2},
		m:  map[string]int{"a": 1, "b": 2},
	}
	temp1 := allType{
		i:  10,
		f:  10.23,
		s:  "ayush",
		a:  [3]string{"a", "b", "c"},
		sl: []int{1, 2},
		m:  map[string]int{"a": 1, "b": 2},
	}
	modifyVal(temp)
	modifyPtr(&temp1)
	displaychanges(orgVal, temp, temp1)

}

// fmt.Println("before modifyVal", intVar, floatVar, text, list)
// 	modifyVal(intVar, floatVar, text, list)
// 	fmt.Println("after modifyVal", intVar, floatVar, text, list)
// 	list[0] = 0
// 	fmt.Println("before modifyPtr", intVar, floatVar, text, list)
// 	modifyPtr(&intVar, &floatVar, &text, &list)
// 	fmt.Println("before modifyPtr", intVar, floatVar, text, list)
