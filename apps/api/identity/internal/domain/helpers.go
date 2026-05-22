package domain

import (
	"os"
	"strings"
)

func (u *User) Sanitize() {
	u.Password = ""
	if u.ProfilePictureUrl != nil && *u.ProfilePictureUrl != "" && !strings.HasPrefix(*u.ProfilePictureUrl, "http") {
		uri := os.Getenv("AZURE_STORAGE_URL")
		fullUrl := uri + "/" + *u.ProfilePictureUrl
		u.ProfilePictureUrl = &fullUrl
	}

	if u.StudentOf != nil && u.StudentOf.Trainer != nil {
		u.StudentOf.Trainer.Sanitize()
	}

	for i := range u.TrainerOf {
		if u.TrainerOf[i].Student != nil {
			u.TrainerOf[i].Student.Sanitize()
		}
	}
}
