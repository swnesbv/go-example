package sqlcsv

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"time"
)


func WriteFile(csvFileName string, rows *sql.Rows) error {
	return NewConverter(rows).WriteFile(csvFileName)
}


func WriteString(rows *sql.Rows) (string, error) {
	return NewConverter(rows).WriteString()
}


func Write(writer io.Writer, rows *sql.Rows) error {
	return NewConverter(rows).Write(writer)
}

type CsvRowPostProcessorFunc func(row []string, columnTypes []*sql.ColumnType) (outputRow bool, processedRow []string)


type Converter struct {
	Headers      []string
	WriteHeaders bool
	TimeFormat   string
	Delimiter    rune
	rows         *sql.Rows
	rowPostProcessor CsvRowPostProcessorFunc
}

func (c *Converter) SetRowPostProcessor(processor CsvRowPostProcessorFunc) {
	c.rowPostProcessor = processor
}

func (c Converter) String() string {
	s,err := c.WriteString()
	if err != nil {
		return ""
	}
	return s
}

func (c Converter) WriteString() (string, error) {
	buffer := bytes.Buffer{}
	err := c.Write(&buffer)
	return buffer.String(), err
}

func (c Converter) WriteFile(csvFileName string) error {
	f,err := os.Create(csvFileName)
	if err != nil {
		return err
	}

	err = c.Write(f)
	if err != nil {
		_ = f.Close() // close, but only return/handle the write error
		return err
	}

	return f.Close()
}


func (c Converter) Write(writer io.Writer) error {
	rows := c.rows
	csvWriter := csv.NewWriter(writer)
	if c.Delimiter != '\x00' {
		csvWriter.Comma = c.Delimiter
	}

	columns,err := rows.ColumnTypes()
	if err != nil {
		return err
	}

	if c.WriteHeaders {

		var headers []string
		if len(c.Headers) > 0 {
			headers = c.Headers
		} else {
			columnNames := make([]string, len(columns))
			for i, col := range columns {
				columnNames[i] = col.Name()
			}
			headers = columnNames
		}
		err = csvWriter.Write(headers)
		if err != nil {
			return err
		}
	}

	count := len(columns)
	values := make([]interface{}, count)
	valPointers := make([]interface{}, count)

	for rows.Next() {
		row := make([]string, count)

		for i := range columns {
			valPointers[i] = &values[i]
		}

		if err = rows.Scan(valPointers...); err != nil {
			return err
		}

		for i, column := range columns {
			if b, isSliceOfBytes := values[i].([]byte); isSliceOfBytes {
				switch column.DatabaseTypeName() {
				case "UNIQUEIDENTIFIER":
					var v uuid.UUID
					if v, err = uuid.FromBytes(b); err != nil {
						return err
					}
					row[i] = v.String()
				default:
					row[i] = string(b)
				}
			} else {
				var value interface{}
				value = values[i]
				if c.TimeFormat != "" {
					if timeValue, ok := value.(time.Time); ok {
						value = timeValue.Format(c.TimeFormat)
					}
				}
				if value == nil {
					row[i] = ""
				} else {
					row[i] = fmt.Sprintf("%v", value)
				}
			}
		}

		writeRow := true
		if c.rowPostProcessor != nil {
			writeRow, row = c.rowPostProcessor(row, columns)
		}
		if writeRow {
			err = csvWriter.Write(row)
			if err != nil {
				return err
			}
		}
	}
	err = rows.Err()

	csvWriter.Flush()

	return err
}

func NewConverter(rows *sql.Rows) *Converter {
	return &Converter{
		rows:         rows,
		WriteHeaders: true,
		Delimiter:    ',',
	}
}
