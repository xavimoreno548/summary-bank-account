package file_manager

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/xavimoreno548/summary-bank-account/internal/model"
)

func TestGetRecordsFromFile(t *testing.T) {

	// Success File
	data := `Id, Date, Transaction
1, 2/12/21, +20.5
2, 2/15/21, -10.2
3, 3/18/21, -8.4
4, 4/3/21, -6
5, 5/5/21, +20`

	tmpfile, err := os.CreateTemp("", "testTrx.*.csv")
	if err != nil {
		t.Fatal("create temp file failed:", err)
	}
	defer os.Remove(tmpfile.Name())

	_, err = tmpfile.Write([]byte(data))
	if err != nil {
		t.Fatal("write data in temp file failed:", err)
	}
	tmpfile.Close()

	testTransactions := []model.Transaction{
		{
			ID:          1,
			Date:        time.Date(2021, 2, 12, 0, 0, 0, 0, time.UTC),
			Transaction: 20.5,
		},
		{
			ID:          2,
			Date:        time.Date(2021, 2, 15, 0, 0, 0, 0, time.UTC),
			Transaction: -10.2,
		},
		{
			ID:          3,
			Date:        time.Date(2021, 3, 18, 0, 0, 0, 0, time.UTC),
			Transaction: -8.4,
		},
		{
			ID:          4,
			Date:        time.Date(2021, 4, 3, 0, 0, 0, 0, time.UTC),
			Transaction: -6,
		},
		{
			ID:          5,
			Date:        time.Date(2021, 5, 5, 0, 0, 0, 0, time.UTC),
			Transaction: 20,
		},
	}

	pathToError, _ := filepath.Abs("---")

	testCases := []struct {
		name     string
		fileName string
		trxsWant []model.Transaction
		errWant  error
	}{
		{
			name:     "happy_path",
			fileName: tmpfile.Name(),
			trxsWant: testTransactions,
		},
		{
			name:     "file_path_error",
			fileName: "---",
			trxsWant: nil,
			errWant:  fmt.Errorf("open %v: %s", pathToError, "no such file or directory"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			fm := NewFileManager()
			trxs, err := fm.GetRecordsFromFile(tt.fileName)
			if tt.errWant != nil {
				assert.Equal(t, tt.errWant.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				if !reflect.DeepEqual(tt.trxsWant, trxs) {
					t.Error("the transactions are not equal")
				}
			}
		})
	}
}

func TestParseID(t *testing.T) {

	testCases := []struct {
		name      string
		value     string
		valueWant uint64
		errWant   error
	}{
		{
			name:      "happy_path",
			value:     "1",
			valueWant: uint64(1),
		},
		{
			name:      "try_to_parse_letter",
			value:     "A",
			valueWant: 0,
			errWant:   fmt.Errorf("strconv.ParseUint: parsing \"A\": invalid syntax"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			v, err := parseID(tt.value)
			if tt.errWant != nil {
				assert.Equal(t, tt.errWant.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.valueWant, v)
			}
		})
	}
}

func TestParseDate(t *testing.T) {

	testCases := []struct {
		name      string
		value     string
		valueWant time.Time
		errWant   error
	}{
		{
			name:      "happy_path",
			value:     "2/12/21",
			valueWant: time.Date(2021, 2, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			name:      "parse_failed",
			value:     "A",
			valueWant: time.Time{},
			errWant:   fmt.Errorf("parsing time \"A\" as \"1/2/06\": cannot parse \"A\" as \"1\""),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			v, err := parseDate(tt.value)
			if tt.errWant != nil {
				assert.Equal(t, tt.errWant.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.valueWant, v)
			}
		})
	}
}

func TestParseTransaction(t *testing.T) {

	testCases := []struct {
		name      string
		value     string
		valueWant float64
		errWant   error
	}{
		{
			name:      "happy_path",
			value:     "+20.6",
			valueWant: 20.6,
		},
		{
			name:      "parse_failed",
			value:     "A",
			valueWant: 0,
			errWant:   fmt.Errorf("strconv.ParseFloat: parsing \"A\": invalid syntax"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			v, err := parseTransaction(tt.value)
			if tt.errWant != nil {
				assert.Equal(t, tt.errWant.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.valueWant, v)
			}
		})
	}
}
