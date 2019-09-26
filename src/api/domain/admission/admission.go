package admission

type Admission struct {
	ID      int    `json:"ID" db:"id"`
	UserId  string `json:"UserId" db:"user_id"`
	BrandId string `json:"BrandId" db:"brand_id"`
}

type AdmissionService interface {
	GetAll() []Admission
}

type AdmissionRepository interface {
	GetAll() []Admission
}
