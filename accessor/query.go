package accessor

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

// Check if user applied filters
func (f FieldQuery) Active() bool {
	if f.Field == "" {
		return false
	}
	return true
}

// Check if user applied filters
func (o OptionsQuery) Active() bool {
	if o.Opts == "" {
		return false
	}
	return true
}

func (f FieldQuery) Get() string {
	return f.Field
}

func (o OptionsQuery) Get() string {
	return o.Opts
}
