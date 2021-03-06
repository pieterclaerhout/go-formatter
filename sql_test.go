package formatter_test

import (
	"testing"

	"github.com/pieterclaerhout/go-formatter"
	"github.com/stretchr/testify/assert"
)

func TestSQL(t *testing.T) {

	type test struct {
		input         string
		shouldFormat  bool
		expectedError error
	}

	var tests = []test{
		{"", false, nil},
		{"throw-error", true, formatter.ErrSQLInvalidStatement},
		{"foo bar", false, nil},
		{"select left(id, 3), coalesce(a.name, \"\") as name, * from `user_items` join apps on a.id = ui.app_id where user_id=? order by created_at limit 3 offset 10", true, nil},
		{"UPDATE tbl_test set uuid =? where id=?", true, nil},
		{"UPDATE tbl_test set uuid =?,name=? where id=? and app_id=?", true, nil},
		{"INSERT INTO tbl_test (field1, field2, field3) values (?,?,?)", true, nil},
		{"select * from tbl_apps", true, nil},
		{"SELECT * from `tbl_users`", true, nil},
		{"select * from table where id in (select id from other_table)", true, nil},
		{"SELECT DATE_FORMAT(b.t_create, '%Y-%c-%d') dateID, b.title memo     FROM (SELECT id FROM orc_scheme_detail d WHERE d.business=208     AND d.type IN (29,30,31,321,33,34,3542,361,327,3, nil8,39,40,41,42,431,4422,415,4546,47,48,'a',    29,30,31,321,33,34,3542,361,327,38,39,40,41,42,431,4422,415,4546,47,48,'a')     AND d.title IS NOT NULL AND t_create >=     DATE_FORMAT((DATE_SUB(NOW(),INTERVAL 1 DAY)),'%Y-%c-%d') AND t_create     < DATE_FORMAT(NOW(), '%Y-%c-%d') ORDER BY d.id LIMIT 2,10) a,     orc_scheme_detail b WHERE a.id = b.id", true, nil},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {

			actual, err := formatter.SQL(tc.input)

			if tc.shouldFormat {
				assert.NotEqual(t, tc.input, actual)
			} else {
				assert.Equal(t, tc.input, actual)
			}

			assert.Equal(t, tc.expectedError, err)

		})
	}

}
