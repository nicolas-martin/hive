package repo

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nicolas-martin/hive/scheduler/config"
	"github.com/nicolas-martin/hive/scheduler/model"
)

type Repo struct {
	schedule []*model.Schedule
}

func NewRepo(cfg *config.Config) *Repo {
	schedule := make([]*model.Schedule, 0)
	return &Repo{schedule: schedule}
}

func (r *Repo) GetSchedule(id uuid.UUID) (*model.Schedule, error) {
	for _, v := range r.schedule {
		if v.ScheduleID == id {
			return v, nil
		}
	}

	return nil, fmt.Errorf("cannot find schedule with ID %s", id)
}

func (r *Repo) AddSchedule() (*uuid.UUID, error) {
	// id := uuid.New()
	return nil, nil

}

// import (
//   "gorm.io/driver/postgres"
//   "gorm.io/gorm"
// )

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
