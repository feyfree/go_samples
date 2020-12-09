package prototype

type JobManager struct {
	jobs map[string]Cloneable
}

func NewJobManager() *JobManager {
	return &JobManager{
		jobs: make(map[string]Cloneable),
	}
}

func (p *JobManager) Get(name string) Cloneable {
	return p.jobs[name]
}

func (p *JobManager) Set(name string, prototype Cloneable) {
	p.jobs[name] = prototype
}
