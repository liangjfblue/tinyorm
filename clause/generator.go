/**
 *
 * @author liangjf
 * @create on 2020/8/22
 * @version 1.0
 */
package clause

import (
	"fmt"
	"strings"
)

type generator func(values ...interface{}) (string, []interface{})

var (
	_generatorM = make(map[Type]generator)
)

func init() {
	_generatorM[INSERT] = _insert
	_generatorM[VALUES] = _values
	_generatorM[SELECT] = _select
	_generatorM[LIMIT] = _limit
	_generatorM[WHERE] = _where
	_generatorM[ORDERBY] = _orderBy
	_generatorM[UPDATE] = _update
	_generatorM[DELETE] = _delete
	_generatorM[COUNT] = _count

}

/**
SELECT col1, col2, ...
    FROM table_name
    WHERE [ conditions ]
    GROUP BY col1
    HAVING [ conditions ]
*/

//genBindVars get ?,?,?,? by field nums
func genBindVars(num int) string {
	res := make([]string, 0)
	for i := 0; i < num; i++ {
		res = append(res, "?")
	}
	return strings.Join(res, ", ")
}

//_insert insert sql
func _insert(values ...interface{}) (string, []interface{}) {
	// INSERT INTO $tableName ($fields)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("INSERT INTO %s (%v)", tableName, fields), []interface{}{}
}

//_values values sql
func _values(values ...interface{}) (string, []interface{}) {
	//VALUES ($v1), ($v2), ...
	var (
		bindVar string
		sql     strings.Builder
		vars    []interface{}
	)

	sql.WriteString("VALUES ")
	for k, value := range values {
		/**
		传入:
		{"Tom", 18},{"Sam", 25}
		[[Tom 18] [Sam 25]]
		处理后:
		[Tom 18 Sam 25]
		*/
		v := value.([]interface{})
		if bindVar == "" {
			bindVar = genBindVars(len(v))
		}
		sql.WriteString(fmt.Sprintf("(%v)", bindVar))

		if k+1 != len(values) {
			sql.WriteString(", ")
		}

		vars = append(vars, v...)
	}

	return sql.String(), vars
}

//_select select sql
func _select(values ...interface{}) (string, []interface{}) {
	//SELECT $fields FROM $tableName
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ",")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []interface{}{}
}

//_limit limit sql
func _limit(values ...interface{}) (string, []interface{}) {
	// LIMIT $num
	return "LIMIT ?", values
}

//_where where sql
func _where(values ...interface{}) (string, []interface{}) {
	// WHERE $desc
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

//_orderBy orderBy sql
func _orderBy(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []interface{}{}
}

//_update update sql, kv kv kv  map[string]interface{}
func _update(values ...interface{}) (string, []interface{}) {
	tableName := values[0]
	m := values[1].(map[string]interface{})
	var keys []string
	var vars []interface{}
	for k, v := range m {
		keys = append(keys, k+" = ?")
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), vars
}

//_delete delete sql
func _delete(values ...interface{}) (string, []interface{}) {
	return fmt.Sprintf("DELETE FROM %s", values[0].(string)), []interface{}{}
}

//_count count sql
func _count(values ...interface{}) (string, []interface{}) {
	return _select(values[0], []string{"count(*)"})
}
