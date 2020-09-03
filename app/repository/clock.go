package repository

import "time"

type Clock interface {
	Now() time.Time
}
