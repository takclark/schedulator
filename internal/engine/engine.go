package engine

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

type Engine struct {
	c *cron.Cron
	l *log.Logger
}

type JobExecutor interface {
	Expression() string
	Execute() func()
}

func NewEngine(l *log.Logger) *Engine {
	return &Engine{l: l}
}

func (e *Engine) Start() error {
	c := cron.New(cron.WithSeconds())
	defer c.Start()

	_, err := c.AddFunc("0 * * * * *", func() {
		e.l.Println("tick")
	})
	if err != nil {
		return fmt.Errorf("error adding pulse check: %w", err)
	}

	return nil
}
