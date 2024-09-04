package learn_aceld

// CompareStruct 比较结构体
func CompareStruct() {
	s1 := struct {
		name string
		age  int
	}{} // 匿名结构体

	s2 := struct {
		name string
		age  int
	}{} // 匿名结构体

	if s1 == s2 {
		println("s1 == s2")
	}

	type Struct1 struct {
		name string
		age  int
	}

	type Struct2 struct {
		name string
		age  int
	}

	s3 := Struct1{}
	if s1 == s3 {
		println("s1 == s3")
	}
	//s4 := Struct2{}
	//if s3 == s4 { // Invalid operation: s3 == s4 (mismatched types Struct1 and Struct2
	//}

	//s5 := struct {
	//	age  int
	//	name string
	//}{} // 匿名结构体，字段顺序和s1相反
	//if s1 == s5 { // Invalid operation: s1 == s5 (mismatched types struct {...} and struct {...}
	//}

	//s6 := struct {
	//	name  string
	//	store map[string]int
	//}{}
	//
	//s7 := struct {
	//	name  string
	//	store map[string]int
	//}{}
	//
	//if s6 == s7{ // Invalid operation: s6 == s7 (the operator == is not defined on struct {...}
	//}

	type Struct3 struct {
		name string
		list []string
	}

	//s8 := Struct3{}
	//s9 := Struct3{}
	//if s8 == s9 {// Invalid operation: s8 == s9 (the operator == is not defined on Struct3)
	//
	//}

	// 总结：
	// - 同一作用域下，匿名结构体的成员字段名相同、类型相同（可比较类型）、字段顺序相同，可以比较；
	// - 具名结构体，结构体名称不同不可比较；
	// - 匿名结构体和具名结构体不可比较；
	// - 结构体的成员字段类型包含不可比较的类型，比如：slice、map，不可比较，但可以使用reflect.DeepEqual比较；
}
