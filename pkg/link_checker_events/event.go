package link_checker_events

import (
	om "github.com/ttlzx/delinkcious/pkg/object_model"
)

type Event struct {
	Username string
	Url      string
	Status   om.LinkStatus
}
