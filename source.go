package agenda

// Source of events
type Source interface {
	Save(job *Job) error
	Fetch() []*Job
	Remove(uuid string) error
}

type memory struct {
	jobs  []*Job
	limit int
}

func (r *memory) filter(jobs []*Job) []*Job {
	var result []*Job
	for _, job := range jobs {
		if job.CompletedAt == "" {
			result = append(result, job)
		}
	}
	return result
}

func (r *memory) Fetch() []*Job {
	jobs := r.filter(r.jobs)
	if r.limit <= len(jobs) {
		return jobs[:r.limit]
	}
	return jobs
}

func (r *memory) Save(job *Job) error {
	for index, jb := range r.jobs {
		if jb.UUID == job.UUID {
			r.jobs[index] = job
			return nil
		}
	}
	r.jobs = append(r.jobs, job)
	return nil
}

func (r *memory) Remove(uuid string) error {
	var jobs []*Job
	for _, job := range r.jobs {
		if job.UUID == uuid {
			continue
		}
		jobs = append(jobs, job)
	}
	r.jobs = jobs
	return nil
}

// Memory returns memory source
func Memory(limit int) Source {
	return &memory{
		limit: limit,
	}
}
