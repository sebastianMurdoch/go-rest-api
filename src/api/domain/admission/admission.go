package admission

type Admission struct {
	ID      int    `json: "ID"`
	UserId  string `json: "UserId"`
	BrandId string `json: "BrandId"`
}

type AdmissionService interface {
	GetAll() []Admission
}

type AdmissionRepository interface {
	GetAll() []Admission
}
