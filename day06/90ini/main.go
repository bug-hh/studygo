package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ini 文件解析器
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(x interface{}) {
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Ptr {
		fmt.Errorf("传入的参数不是指针类型：%v", t.Kind())
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		fmt.Errorf("传入的参数不是结构体指针类型：%v", t.Kind())
		return
	}

	fileObj, err := os.Open("./conf.ini")
	defer fileObj.Close()
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	reader := bufio.NewReader(fileObj)

	var structName string

	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		// 去掉空行
		if len(s) == 0 {
			continue
		}
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}
		if strings.HasPrefix(s, "[") {
			if s[0] != '[' || s[len(s)-1] != ']' {
				fmt.Errorf("%s 语法错误", s)
				return
			}
			// [mysql] 去掉 [], 取 mysql
			sectionName := strings.TrimSpace(s[1 : len(s)-1])
			if len(sectionName) == 0 {
				fmt.Errorf("%s 语法错误", s)
				return
			}

			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到 %s 对应的结构体 %s\n", sectionName, structName)
				}
			}

		} else {
			// 将每行内容用 = 分割
			arr := strings.Split(s, "=")
			name := arr[0]
			value := arr[1]

			v := reflect.ValueOf(x)
			structValue := v.Elem().FieldByName(structName)
			sType := structValue.Type()

			if sType.Kind() != reflect.Struct {
				fmt.Errorf("x 中 %s 字段应该是一个结构体", structName)
				return
			}
			var fieldType reflect.StructField
			var fieldName string
			for i := 0; i < sType.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == name {
					fieldType = field
					fieldName = field.Name
					break
				}
			}
			if len(fieldName) == 0 {
				continue
			}
			// 子结构体
			structObj := structValue.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				structObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueInt, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					fmt.Errorf("解析数字失败")
					return
				}
				structObj.SetInt(valueInt)
			}

		}

	}
}

func main() {
	// 通过反射，将 conf.ini 文件中的值，赋值到 MysqlConfig 结构体对应的字段中
	var conf Config
	loadIni(&conf)
	fmt.Println(conf)
}
