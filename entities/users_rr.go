package entities

type UserDetails struct {
	UID           *string `json:"uid,omitempty" example:"0MhHcnVNBMeCIygoBHDDt0SvT053"`
	Email         string  `json:"email" validate:"required,email" example:"test@test.com"`
	EmailVerified bool    `json:"emailVerified" validate:"required" example:"false"`
	PhoneNumber   string  `json:"phoneNumber" validate:"required,e164" example:"+911234567890"`
	Password      string  `json:"password" validate:"required,min=6" example:"testpassword"`
	DisplayName   string  `json:"name" validate:"required,min=3,max=24" example:"TEST"`
	PhotoURL      *string `json:"photoURL,omitempty" validate:"url" example:"https://example.com/photo.jpg"`
	Disabled      bool    `json:"disabled" validate:"required" example:"false"`
}
