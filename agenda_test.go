package agenda

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestAgenda(t *testing.T) {
	memory := Memory(5)
	agenda := New(memory, 5)
	agenda.Register("event", func(job *Job) error {
		log.Println("Event", job.UUID)
		return nil
	})

	agenda.Now("event", map[string]interface{}{
		"message": "Hello test",
	})

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)

	defer cancel()

	go agenda.Start()

	select {
	case job := <-agenda.update:
		if job == nil {
			t.Fatal("Fail to execute job")
		}
		if job.CompletedAt == "" {
			t.Fatal("Fail to complete task")
		}
	case <-ctx.Done():
		t.Fatal(ctx.Err())
	}

}
