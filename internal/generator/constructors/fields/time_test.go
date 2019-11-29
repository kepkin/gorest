package fields

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kepkin/gorest/internal/generator/translator"
)

func TestMakeDateFieldConstructor(t *testing.T) {
	t.Run("No datetime field", func(t *testing.T) {
		_, err := MakeDateFieldConstructor(translator.Field{
			Type: translator.DateTimeField * 2,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("No date field", func(t *testing.T) {
		_, err := MakeDateFieldConstructor(translator.Field{
			Type: translator.DateField * 2,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("Time field example", func(t *testing.T) {
		s, err := MakeDateFieldConstructor(translator.Field{
			Name:      "FromDate",
			Parameter: "fromDate",
			GoType:    "time.Time",
			Type:      translator.DateField,
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.FromDate, err = time.Parse("2006-01-02", fromDateStr)
if err != nil {
	errors = append(errors, NewFieldError(InQuery, "fromDate", "can't parse as RFC3339 date", err))
}`, s)
	})
}

func TestMakeDateTimeFieldConstructor(t *testing.T) {
	t.Run("No datetime field", func(t *testing.T) {
		_, err := MakeDateTimeFieldConstructor(translator.Field{
			Type: translator.DateTimeField * 2,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("No date field", func(t *testing.T) {
		_, err := MakeDateTimeFieldConstructor(translator.Field{
			Type: translator.DateField * 2,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("Time field example", func(t *testing.T) {
		s, err := MakeDateTimeFieldConstructor(translator.Field{
			Name:      "FromDate",
			Parameter: "fromDate",
			GoType:    "time.Time",
			Type:      translator.DateTimeField,
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `result.FromDate, err = time.Parse(time.RFC3339, fromDateStr)
if err != nil {
	errors = append(errors, NewFieldError(InQuery, "fromDate", "can't parse as RFC3339 time", err))
}`, s)
	})
}

func TestMakeUnixTimeFieldConstructor(t *testing.T) {
	t.Run("No unix time field", func(t *testing.T) {
		_, err := MakeUnixTimeFieldConstructor(translator.Field{
			Type: translator.UnixTimeField * 2,
		}, "InCookie")
		assert.Error(t, err)
	})

	t.Run("Unix time field example", func(t *testing.T) {
		s, err := MakeUnixTimeFieldConstructor(translator.Field{
			Name:      "FromTs",
			Parameter: "fromTs",
			GoType:    "time.Time",
			Type:      translator.UnixTimeField,
		}, "InQuery")
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, `fromTsSec, err := strconv.ParseInt(fromTsStr, 10, 64)
if err != nil {
	errors = append(errors, NewFieldError(InQuery, "fromTs", "can't parse as 64 bit integer", err))
} else {
	result.FromTs = time.Unix(fromTsSec, 0)
}`, s)
	})
}
