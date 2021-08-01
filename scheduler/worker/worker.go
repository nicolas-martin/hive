package worker

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Worker struct {
	c *cron.Cron
}

func NewCron() *Worker {
	c := cron.New()
	c.Start()
	return &Worker{c: c}
}

// Cron adds function to be run at specified times
func (w *Worker) Cron(spec string, f func()) (cron.EntryID, error) {
	id, err := w.c.AddFunc(spec, f)

	return id, err
}

func (w *Worker) Remove(id cron.EntryID) {
	w.c.Remove(id)
}

func (w *Worker) Entry(id cron.EntryID) (*cron.Entry, error) {
	r := w.c.Entry(id)

	if (cron.Entry{} == r) {
		return nil, fmt.Errorf("Cron entry with id %d not found", id)

	}
	return &r, nil
}

// Parse returns a cron schedule but doesn't start it.
func Parse(spec string) (cron.Schedule, error) {
	parser := cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	return parser.Parse(spec)
}

// NewCronSpec creates a new instance of cron following a spec and a function to invoke.
// This should only be used in cases where a new cron needs to be created before
// the initial message has beeen received and the timezone has been set.
func (w *Worker) NewCronSpec(spec string, f func()) {
	w.c.AddFunc(spec, f)
}
