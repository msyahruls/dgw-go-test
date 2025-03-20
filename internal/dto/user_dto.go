package dto

type CreateUserRequest struct {
	NIK         string  `json:"nik" binding:"required"`
	FullName    string  `json:"full_name" binding:"required"`
	LegalName   string  `json:"legal_name" binding:"required"`
	BirthPlace  string  `json:"birth_place" binding:"required"`
	BirthDate   string  `json:"birth_date" binding:"required,datetime=2006-01-02"` // Date format
	Salary      float64 `json:"salary" binding:"required,gt=0"`
	PhotoIDCard string  `json:"photo_id_card" binding:""`
	PhotoSelfie string  `json:"photo_selfie" binding:""`
}
