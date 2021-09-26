package accessor

import "fmt"

type Query interface {
	// determine where user applied query fields or not
	Active() bool
	Get() string
}

type FieldQuery struct {
	Field string
}

type OptionsQuery struct {
	Opts string
}

type TimeseriesQuery struct {
	Start    string
	End      string
	Interval string
}

func (t TimeseriesQuery) Active() bool {
	return t.Start != "" && t.End != "" && t.Interval != ""
}

// Check if user applied filters
func (f FieldQuery) Active() bool {
	return f.Field != ""
}

// Check if user applied filters
func (o OptionsQuery) Active() bool {
	return o.Opts != ""
}

func (f FieldQuery) Get() string {
	return f.Field
}

func (o OptionsQuery) Get() string {
	return o.Opts
}

func (t TimeseriesQuery) Get() string {
	return fmt.Sprintf("?%v&%v&%v", t.Start, t.End, t.Interval)
}
