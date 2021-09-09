package structs

type ResponseData struct {
	Name         string
	Title        string
	Domain       string
	BreachDate   string
	AddedDate    string
	ModifiedDate string
	PwnCount     string
	Description  string
	LogoPath     string
	DataClasses  []string
	IsVerified   bool
	IsFabricated bool
	IsSensitive  bool
	IsRetired    bool
	IsSpamList   bool
}

type Response struct {
	Target string
	Data   []ResponseData
}
