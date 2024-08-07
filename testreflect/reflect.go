package testreflect

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

// <Value, Type>
// reflect.TypeOf 获取接口类型
// reflect.ValueOf 获取接口值

func TypeOf() {
	var num int
	typ := reflect.TypeOf(num)
	if typ.Kind() != reflect.Int {
		// 精确匹配
		println("match type")
	}

	var num2 int64
	typ2 := reflect.TypeOf(num2)
	if typ2.Kind() != reflect.Int {
		// 基本类型匹配
		println("match int64 to int type")
	}

}

type Sample interface {
	Print()
}

type SampleObj struct{}

func (s SampleObj) Print() {}

func InteraceAble() {
	if reflect.TypeOf((*Sample)(nil)).Elem().AssignableTo(reflect.TypeOf(SampleObj{})) {
		// 接口匹配判定
		println("match interface")
	}
}

// TypeConvert 类型转换匹配
func TypeConvert() {
	var num int = 1
	if reflect.TypeOf(num).ConvertibleTo(reflect.TypeOf(float64(0))) {
		println("convertible to float64")
	}
}

// ModifyFloat 修改浮点数
func ModifyFloat() {
	var f float64 = 3.1415926
	p := reflect.ValueOf(&f) // 必须传入指针
	if p.Elem().CanSet() {   // 判断是否可设置参数
		p.Elem().SetFloat(80) // 设置参数
		println(f)
	}
}

func ReflectValueToInterfaceValue() {
	type User struct {
		Id  int
		Age int
	}
	u := User{Id: 1, Age: 18}
	rf := reflect.ValueOf(u)

	var num int = 20
	rfn := reflect.ValueOf(num)

	user := rf.Interface().(User) // 通过接口值断言为具体对象

	numi := rfn.Interface()

	fmt.Printf("%v:%v %d\n", user, reflect.TypeOf(user).Kind(), numi)

	// 打印字段名和值
	uv := reflect.ValueOf(u) // 获取uv，复用uv
	uvt := uv.Type()
	for i := 0; i < uv.NumField(); i++ {
		field := uv.Field(i)
		fmt.Printf("name:%s value:%v\n", uvt.Field(i).Name, field.Interface())
	}

}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Description(hi string) string {
	return fmt.Sprintf("%s, My name is %s and age is %d.", hi, u.Name, u.Age)
}

func ModifyValueCallMethod() {
	user := &User{Id: 1, Name: "Tom", Age: 18}
	ru := reflect.ValueOf(user)

	name := ru.Elem().FieldByName("Name")
	if name.IsValid() {
		fmt.Printf("name:%v\n", name.Interface())
	}

	age := ru.Elem().FieldByName("Age")
	if age.IsValid() {
		fmt.Printf("age:%v\n", age.Interface())
	}
	if age.CanSet() && age.Kind() == reflect.Int {
		age.SetInt(20)
		fmt.Println("new age:", user.Age)
	}

	method := ru.Elem().MethodByName("Description")
	if method.IsValid() {
		params := []reflect.Value{reflect.ValueOf("Hi")}
		say := method.Call(params)
		fmt.Println("say:", say[0].Interface().(string))
	}

}

type Struct struct {
	FieldA string
	FiledB int
	FieldC float64
}

func ReflectTypeToFields(obj interface{}) ([]string, string) {
	typ := reflect.TypeOf(obj)
	var fields []string
	for i := 0; i < typ.NumField(); i++ {
		fields = append(fields, typ.Field(i).Name)
	}
	str := fmt.Sprintf("`%s`", strings.Join(fields, "`,`"))
	return fields, str
}

func GetFields() {
	list, str := ReflectTypeToFields(Struct{})
	fmt.Printf("%v\n%s\n", list, str)
}

// 使用反射设置配置
// 系统变量存在则设置到对象

type Config struct {
	Name    string `json:"server-name"`
	IP      string `json:"server-ip"`
	URL     string `json:"server-url"`
	Timeout string `json:"timeout"`
}

func ReadConfig() *Config {
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func ReflectEnv() {
	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "baidu.com")
	c := ReadConfig()
	fmt.Printf("%+v", c)
}

// MakeObjByReflectType 通过反射类型创建对象
func MakeObjByReflectType() {
	var config *Config
	typ := reflect.TypeOf(config)
	config, _ = reflect.New(typ).Interface().(*Config)
	fmt.Printf("%+v", config)
}

// SetFiledValue 修改字段值
func SetFiledValue() {
	typ := reflect.TypeOf(Config{})
	elem := reflect.New(typ).Elem()
	elem.Field(0).SetString("name")
	elem.Field(1).SetString("ip")
	elem.Field(2).SetString("url")
	elem.Field(3).SetString("timeout")
	fmt.Printf("%+v\n", elem.Interface())
}

// 使用二进制报文+反射反序列化为结构体对象

type Example struct {
	ID   uint32
	Name string
	Age  uint16
}

// BinToObj 二进制报文赋值到对象
func BinToObj() {
	// 报文
	name := "Alice"
	binaryData := make(
		[]byte, 0,
		binary.MaxVarintLen32+(binary.MaxVarintLen32+len(name))+binary.MaxVarintLen16,
		// ID + (字符串值的长度+字符串值实际占用字节) + Age
	)
	order := binary.LittleEndian
	buf := bytes.NewBuffer(binaryData)
	var err error
	// 写入ID数值
	if err = binary.Write(buf, order, uint32(123456)); err != nil {
		panic(err)
	}
	// 写入描述字符串长度的值和字符串实际占用字节
	if err = binary.Write(buf, order, uint32(len(name))); err != nil {
		panic(err)
	}
	if _, err = buf.WriteString(name); err != nil {
		panic(err)
	}
	// 写入Age数值
	if err = binary.Write(buf, order, uint16(30)); err != nil {
		panic(err)
	}

	example := Example{}
	if err = binaryToStruct(buf.Bytes(), &example); err != nil {
		panic(err)
	}
	fmt.Printf("%+v", example)
}

// binaryToStruct 将二进制数据反序列化为结构体
func binaryToStruct(data []byte, s interface{}) error {
	val := reflect.ValueOf(s).Elem()
	t := val.Type()

	reader := bytes.NewReader(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		// 根据字段类型读取二进制数据
		switch field.Type.Kind() {
		case reflect.Uint32:
			var u uint32
			err := binary.Read(reader, binary.LittleEndian, &u)
			if err != nil {
				return err
			}
			fieldVal.SetUint(uint64(u))
		case reflect.String:
			var n uint32
			err := binary.Read(reader, binary.LittleEndian, &n)
			if err != nil {
				return err
			}
			b := make([]byte, n)
			_, err = io.ReadFull(reader, b)
			if err != nil {
				return err
			}
			fieldVal.SetString(string(b))
		case reflect.Uint16:
			var u uint16
			err := binary.Read(reader, binary.LittleEndian, &u)
			if err != nil {
				return err
			}
			fieldVal.SetUint(uint64(u))
		default:
			return fmt.Errorf("unsupported type %s for field %s", field.Type, field.Name)
		}
	}

	return nil
}
