package cqrskit

import (
	"context"
	"time"
)

//*******************************************************************************
// Event Type and Functions
//*******************************************************************************

// Event embodies the data stored correlating with events
// generated for a giving type.
type Event struct {
	Version       int64       `json:"version" bson:"version" sql:"version"`
	Created       time.Time   `json:"created" bson:"created" sql:"created"`
	EventType     string      `json:"event_type" bson:"event_type" sql:"event_type"`
	EventData     interface{} `json:"event_data" bson:"event_data" sql:"event_data"`
	AggregateID   string      `json:"aggregate_id" bson:"aggregate_id" sql:"aggregate_id"`
	CorrelationID string      `json:"correlation_id" bson:"correlation_id" sql:"correlation_id"`
}

//*******************************************************************************
// Write Repository Interface and Implementation
//*******************************************************************************

// WriteRepo embodies a repository which houses the store
// of events for giving type .
type WriteRepo interface {
	SaveEvents(context.Context, ...Event) error
}

// WriteRepository defines the interface which provides
// a single method to retrieve a WriteRepository which
// stores all events for a particular  identified by it's correlationID.
type WriteRepository interface {
	New(correlationID string) (WriteRepo, error)
	Get(correlationID string) (WriteRepo, error)
}

//*******************************************************************************
// Read Repository Interface and Implementation
//*******************************************************************************

// ReadRepo embodies a repository which reads the store
// of events for giving type , returning an Applier
// to apply said events to target.
type ReadRepo interface {
	ReadAll(context.Context) ([]Event, error)
	ReadFromTime(ctx context.Context, last time.Time) ([]Event, error)
	ReadFromVersion(ctx context.Context, version int64) ([]Event, error)
	ReadFromLastCount(ctx context.Context, count int64) ([]Event, error)
}

// ReadRepository defines the interface which provides
// a single method to retrieve a ReadRepos to read
// events that occur for a giving type through
// it's correlationID which identifies that records events.
type ReadRepository interface {
	Get(correlationID string) (ReadRepo, error)
}