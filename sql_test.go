package formatter_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pieterclaerhout/go-formatter"
)

func Test_SQL(t *testing.T) {

	type test struct {
		input         string
		shouldFormat  bool
		expectedError error
	}

	var tests = []test{
		{"", false, nil},
		{"throw-error", true, formatter.ErrSQLInvalidStatement},
		{"foo bar", false, nil},
		{"SELECT * From tbl_apps", true, nil},
		{"SELECT * From `tbl_apps`", true, nil},
		{"select left(id, 3), coalesce(a.name, \"\") as name, * from `user_items` join apps on a.id = ui.app_id where user_id=? order by created_at limit 3 offset 10", true, nil},
		{"SELECT * From table", true, nil},
		{"SELECT * From table where id in (select id from other_table)", true, nil},
		{"UPDATE tbl_test set uuid =? where id=?", true, nil},
		{"UPDATE tbl_test set uuid =?,name=? where id=? and app_id=?", true, nil},
		{"INSERT INTO tbl_test (field1, field2, field3) values (?,?,?)", true, nil},
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

func Test_ExchangeRates_InvalidURL(t *testing.T) {

	formatter.FormatSQLAPIURL = "ht&@-tp://:aa"
	defer resetSQLFormatURL()

	actual, err := formatter.SQL("SELECT * From tbl_apps")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_ExchangeRates_Timeout(t *testing.T) {

	formatter.DefaultTimeout = 250 * time.Millisecond

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(500 * time.Millisecond)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("hello"))
		}),
	)
	defer s.Close()

	formatter.FormatSQLAPIURL = s.URL
	defer resetSQLFormatURL()

	actual, err := formatter.SQL("SELECT * From tbl_apps")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_ExchangeRates_ReadBodyError(t *testing.T) {

	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Length", "1")
		}),
	)
	defer s.Close()

	formatter.FormatSQLAPIURL = s.URL
	defer resetSQLFormatURL()

	actual, err := formatter.SQL("SELECT * From tbl_apps")

	assert.Error(t, err)
	assert.Empty(t, actual)

}

func resetSQLFormatURL() {
	formatter.FormatSQLAPIURL = formatter.DefaultFormatSQLAPIURL
}
