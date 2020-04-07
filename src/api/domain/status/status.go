package status

type Status struct {
	Value string `json:"value"`
}

type StatusService interface {
	Get() (*Status, error)
}