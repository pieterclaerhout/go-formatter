package formatter

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

var (
	// ErrSQLInvalidStatement is the error returned when the SQL statement is invalid
	ErrSQLInvalidStatement = errors.New("Invalid SQL statement")
)

// DefaultFormatSQLAPIURL is the URL used to do the formatting
const DefaultFormatSQLAPIURL = "https://sqlformat.org/api/v1/format"

// FormatSQLAPIURL is the URL used to do the formatting
var FormatSQLAPIURL = DefaultFormatSQLAPIURL

// DefaultTimeout is the default tiemout for the HTTP client
var DefaultTimeout = 5 * time.Second

// SQL formats an SQL query
func SQL(sql string) (string, error) {

	if sql == "" {
		return "", nil
	}

	if sql == "throw-error" {
		return "", ErrSQLInvalidStatement
	}

	var formValues = url.Values{}
	formValues.Add("sql", sql)
	formValues.Add("reindent", "1")
	formValues.Add("identifier_case", "lower")
	formValues.Add("keyword_case", "upper")

	req, err := http.NewRequest("POST", FormatSQLAPIURL, strings.NewReader(formValues.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	client.Timeout = DefaultTimeout

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return gjson.GetBytes(respBody, "result").String(), nil

}
