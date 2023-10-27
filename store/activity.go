package store

import (
	"context"

	storepb "github.com/usememos/memos/proto/gen/store"
)

type ActivityType string

const (
	ActivityTypeMemoComment ActivityType = "MEMO_COMMENT"
)

func (t ActivityType) String() string {
	return string(t)
}

type ActivityLevel string

const (
	ActivityLevelInfo ActivityLevel = "INFO"
)

func (l ActivityLevel) String() string {
	return string(l)
}

type Activity struct {
	ID int32

	// Standard fields
	CreatorID int32
	CreatedTs int64

	// Domain specific fields
	Type    ActivityType
	Level   ActivityLevel
	Payload *storepb.ActivityPayload
}

type FindActivity struct {
	ID *int32
}

func (s *Store) CreateActivity(ctx context.Context, create *Activity) (*Activity, error) {
	return s.driver.CreateActivity(ctx, create)
}

func (s *Store) ListActivities(ctx context.Context, find *FindActivity) ([]*Activity, error) {
	return s.driver.ListActivities(ctx, find)
}
