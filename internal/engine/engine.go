package engine

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

type Engine struct {
	c    *cron.Cron
	l    *log.Logger
	tick Job
}

type Job interface {
	Expression() string
	Execute() func()
}

func NewEngine(l *log.Logger) (*Engine, error) {
	e := &Engine{l: l}

	if err := e.Start(); err != nil {
		return &Engine{}, fmt.Errorf("starting engine: %w", err)
	}

	return e, nil
}

// Register a new job in the engine.
func (e *Engine) Register(job Job) (int, error) {
	id, err := e.c.AddFunc(job.Expression(), job.Execute())
	if err != nil {
		return 0, fmt.Errorf("unable to register job: %w", err)
	}

	e.l.Printf("registered requester with ID %d\n", id)

	return int(id), err
}

func (e *Engine) Start() error {
	c := cron.New(cron.WithSeconds())
	defer c.Start()
	e.c = c

	_, err := c.AddFunc("0 * * * * *", func() {
		e.l.Println("tick")
	})
	if err != nil {
		return fmt.Errorf("error adding pulse check: %w", err)
	}

	return nil
}
