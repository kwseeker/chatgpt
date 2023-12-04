package response

type Entity struct {
	Code    int
	Status  string
	Message string
	Data    map[string]interface{}
}
