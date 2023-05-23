package jsongen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

const ObjectFormat = `package %s;

import org.json.JSONObject;

public class %s {
%s
%s
}`

const CreateMethodFormat = `	public static %s create(JSONObject root) {
		%s
%s
		return object;
	}`

var interfaceData interface{}

type Class struct {
	Pkg  string
	Name string
	Kind reflect.Kind
	Vars []Variable
}

func (c *Class) genClass() {
	// 包名

	// 注释

	// 类名

	//变量
	vStr := ""
	funStr := ""
	for i := 0; i < len(c.Vars); i++ {
		v := c.Vars[i]
		vStr += v.gen()
		funStr += v.genFun()
	}
	newStr := fmt.Sprintf("%s object = new %s();", c.Name, c.Name)
	create := fmt.Sprintf(CreateMethodFormat, c.Name, newStr, funStr)
	result := fmt.Sprintf(ObjectFormat, c.Pkg, c.Name, vStr, create)
	WriteFile(result)
}

// 变量
type Variable struct {
	Kind reflect.Kind
	Key  string
}

func (v *Variable) gen() string {
	kind := "unknow"
	if v.Kind == reflect.String {
		kind = "String"
	}

	return fmt.Sprintf("\tpublic %s %s;\n", kind, v.Key)
}

func (v *Variable) genFun() string {
	kind := "unknow"
	if v.Kind == reflect.String {
		kind = "String"
	} else if v.Kind == reflect.Bool {
		kind = "Boolean"
	} else if v.Kind == reflect.Map {
		kind = "JSONObject"
	} else if v.Kind == reflect.Slice {
		kind = "JSONArray"
	}

	return fmt.Sprintf("\t\tobject.%s = root.opt%s(\"%s\");\n", v.Key, kind, v.Key)
}

func WriteFile(wireteString string) {
	/*****************************  第二种方式: 使用 ioutil.WriteFile 写入文件 ***********************************************/
	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile("./Json.java", d1, 0666) //写入文件(字节数组)
	if err2 != nil {
		panic(err2)
	}
}
func genClazz() {

	// map
	userJson := "{\"username\":\"system\",\"password\":\"123456\"}"

	// slice
	//userJson := "[\"username\",\"system\",\"password\",\"123456\"]"

	err := json.Unmarshal([]byte(userJson), &interfaceData)
	if err != nil {
		fmt.Println(err.Error())
	}

	typeOfA := reflect.TypeOf(interfaceData)
	if typeOfA.Kind() == reflect.Map {
		parseMap(nil, interfaceData)
	}
	fmt.Println(interfaceData)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	var v Variable
	v.Kind = reflect.String
	v.Key = "id"

	var c Class
	c.Pkg = "com.qb.monetization.demo"
	c.Name = "Json"
	c.Vars = append(c.Vars, v)
	v.Key = "name"
	c.Vars = append(c.Vars, v)
	c.genClass()
}

func parseMap(vars []Variable, interfaceData interface{}) []Variable {
	for key, value := range interfaceData.(map[string]interface{}) {
		fmt.Println("==>>", key, value)
		typeOfValue := reflect.TypeOf(value)
		var v Variable
		v.Kind = typeOfValue.Kind()
		v.Key = key
		vars = append(vars, v)
	}
	return vars
}

// json数据 输入包名 主类名
func genRootClass(data, packageName, rootClassName string) {
	err := json.Unmarshal([]byte(data), &interfaceData)
	if err != nil {
		panic(err)
	}

	var rootClass Class
	rootClass.Pkg = packageName
	rootClass.Name = rootClassName

	typeOfA := reflect.TypeOf(interfaceData)
	if typeOfA.Kind() == reflect.Map {
		rootClass.Kind = reflect.Map
		rootClass.Vars = parseMap(rootClass.Vars, interfaceData)
	} else if typeOfA.Kind() == reflect.Slice {
		rootClass.Kind = reflect.Slice
	} else {
		panic("当前json数据不是JSONObject或不是JSONArray")
	}
	rootClass.genClass()
	//rootClass.Kind =
}
