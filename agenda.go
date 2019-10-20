package agenda

import (
	"time"

	"github.com/google/uuid"
)

// Processor of events
type Processor = func(job *Job) error

// Agenda worker
type Agenda struct {
	source    Source
	delay     time.Duration
	procesors map[string]Processor
	update    chan *Job
}

// Start to process jobs
func (agenda *Agenda) Start() {
	for {
		for _, job := range agenda.source.Fetch() {
			processor, found := agenda.procesors[job.Name]
			if !found {
				continue
			}
			if err := processor(job); err != nil {
				job.FailedAt = time.Now().Format(time.UnixDate)
				job.FailedReason = err.Error()
			} else {
				job.CompletedAt = time.Now().Format(time.UnixDate)
			}
			agenda.update <- job
			agenda.source.Save(job)
		}
		time.Sleep(agenda.delay)
	}
}

// Register a new processor
func (agenda *Agenda) Register(name string, processor Processor) {
	agenda.procesors[name] = processor
}

// Now add a job for being executed in the moment
func (agenda *Agenda) Now(name string, data interface{}) error {
	return agenda.source.Save(&Job{
		UUID:    uuid.New().String(),
		Name:    name,
		Data:    data,
		NextRun: time.Now().Format(time.UnixDate),
	})
}

// New returns a new agenda instance
func New(source Source, delay time.Duration) *Agenda {
	return &Agenda{
		source:    source,
		delay:     delay,
		procesors: make(map[string]Processor),
		update:    make(chan *Job),
	}
}
