package criteria

// Type represents a search criteria
type Type struct {
	ID               int
	Name             string
	AllOfTheseWords  []string
	ThisExactPhrase  string
	AnyOfTheseWords  []string
	NoneOfTheseWords []string
	TheseHashtags    []string
	Language         string
	Since            string
	Until            string
}
