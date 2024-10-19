package types

type User struct {
	UserName string
	Name     string
	Email    string

	BookedClasses  []string
	OfferedClasses []string
}

// TODO: Add constructor for User struct
