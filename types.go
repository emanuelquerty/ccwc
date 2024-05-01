package ccwc

// A Flag represents a flag that can hold flag information
type Flag struct {
	Name  string
	Value string
	Usage string
}

// FlagValue implements flag.Value interface 
// as defined in the standard library flag package
type FlagValue struct {
	filename string
}

func (f *FlagValue) String() string {
	return f.filename
}

func (f *FlagValue) Set(filename string) error {
	f.filename = filename
	return nil
}