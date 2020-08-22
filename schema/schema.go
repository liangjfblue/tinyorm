/**
 *
 * @author liangjf
 * @create on 2020/8/21
 * @version 1.0
 */
package schema

import (
	"go/ast"
	"reflect"
	"tinyorm/dialect"
)

type Field struct {
	Name string
	Type string
	Tag  string
}

type Schema struct {
	Model      interface{} //被映射对象
	Name       string      //表名
	Fields     []*Field    //字段列表
	FieldNames []string    //字段名列表
	fieldM     map[string]*Field
}

func (s *Schema) GetField(name string) *Field {
	return s.fieldM[name]
}

//RecordValues &User{Name: "Sam", Age: 25} ==> ("Sam",25)
func (s *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	fieldValues := make([]interface{}, 0)
	for _, field := range s.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

//Parse 解析go对象, 构建sql结构
func Parse(obj interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(obj)).Type()

	s := &Schema{
		Model:  obj,
		Name:   modelType.Name(),
		fieldM: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		f := modelType.Field(i)

		if !f.Anonymous && ast.IsExported(f.Name) {
			field := &Field{
				Name: f.Name,
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(f.Type))),
			}

			if v, ok := f.Tag.Lookup("tinyorm"); ok {
				field.Tag = v
			}

			s.Fields = append(s.Fields, field)
			s.FieldNames = append(s.FieldNames, f.Name)
			s.fieldM[f.Name] = field
		}

	}

	return s
}
