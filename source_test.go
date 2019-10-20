package agenda

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMemorySource(t *testing.T) {
	source := Memory(2)
	jobs := []*Job{
		&Job{
			UUID: uuid.New().String(),
		},
	}

	for _, job := range jobs {
		source.Save(job)
	}

	if len(source.Fetch()) != len(jobs) {
		t.Errorf("Expect source.Fetch() to be %d, received: %d",
			len(jobs), len(source.Fetch()))
	}

	jobs[0].CompletedAt = time.Now().Format(time.UnixDate)

	if len(source.Fetch()) != len(jobs)-1 {
		t.Errorf("Expect source.Fetch() to be %d, received: %d",
			len(jobs)-1, len(source.Fetch()))
	}

}
