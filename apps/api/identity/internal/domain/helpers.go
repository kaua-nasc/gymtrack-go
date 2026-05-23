package domain

import (
	"strings"

	"github.com/kaua-nasc/gymtrack-go/libs/storage"
)

func (u *User) Sanitize() {
	u.Password = ""
	if u.ProfilePictureUrl != nil && *u.ProfilePictureUrl != "" && !strings.HasPrefix(*u.ProfilePictureUrl, "http") {
		u.ProfilePictureUrl = storage.GetBlobURL(*u.ProfilePictureUrl)
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
