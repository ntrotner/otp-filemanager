package id_manager

import (
	"errors"
	"log"
	"time"

	"github.com/madflojo/tasks"
)

// CheckAndDeleteExpiredIdentities finds identities and requests deletion
// if time delta is not within retention time anymore
func CheckAndDeleteExpiredIdentities(expirationTime *int64) error {
	minuteDelta := time.Duration(-(*expirationTime)) * time.Minute
	timeThreshhold := time.Now().Add(minuteDelta)
	var hasError error = nil

	for id, user := range ExistingIDs {
		if (*expirationTime != 0) && (*user.IssuedDate).Before(timeThreshhold) {
			err := DeleteIdentity(&id)
			log.Println("Delete Identity:", id)

			if err != nil {
				hasError = errors.New("Deleting an Identity caused issues")
			}
		}
	}

	return hasError
}

// OrchestrateExpirationCheck schedules the deletion of expired identities
func OrchestrateExpirationCheck(expirationTime *int64) {
	scheduler := tasks.New()

	scheduler.Add(&tasks.Task{
		Interval: time.Duration(10 * time.Second),
		TaskFunc: func() error {
			return CheckAndDeleteExpiredIdentities(expirationTime)
		},
		ErrFunc: func(err error) {
			log.Println("Error for Expiration Check:", err)
		},
	})
}
