package app

import (
	"com/chat/service/pkg/prof"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type GoServer interface {
	Start() error
	Stop() error
	String() string
}

type Close func() error

type App struct {
	servers []GoServer
	closes  []Close
}

func New(servers []GoServer, closes []Close) *App {
	return &App{
		servers: servers,
		closes:  closes,
	}
}

func (a *App) watch(ctx context.Context) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGTRAP)
	profile := prof.NewProfile()

	for {
		select {
		case <-ctx.Done(): // service error
			_ = a.stop()
			return ctx.Err()
		case sigType := <-sig: // system notification signal 采集profile数据
			fmt.Printf("received system notification signal: %s\n", sigType.String())
			switch sigType {
			case syscall.SIGTRAP:
				profile.StartOrStop() // start or stop sampling profile
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP:
				if err := a.stop(); err != nil {
					return err
				}
				fmt.Println("stop app successfully")
				return nil
			}
		}
	}
}

func (a *App) Run() {
	// ctx will be notified whenever an error occurs in one of the goroutines.
	eg, ctx := errgroup.WithContext(context.Background())

	// start all servers
	for _, server := range a.servers {
		s := server
		eg.Go(func() error {
			fmt.Println(s.String())
			return s.Start()
		})
	}

	// watch and stop app
	eg.Go(func() error {
		return a.watch(ctx)
	})

	if err := eg.Wait(); err != nil {
		panic(err)
	}

}

// Stopping services and releasing resources
func (a *App) stop() error {
	for _, closeFn := range a.closes {
		if err := closeFn(); err != nil {
			return err
		}
	}
	return nil
}
