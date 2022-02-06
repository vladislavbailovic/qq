package main

import (
	"log"
	"os"
	"sync"

	g "github.com/AllenDang/giu"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type State struct {
	opts       []string
	currentOpt int
	wnd        *g.MasterWindow
	mu         sync.RWMutex
}

func (s *State) hasWindowOpen() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.wnd != nil
}

func main() {
	wg := sync.WaitGroup{}

	state := State{
		opts: []string{
			"",
			"this is my list",
			"of selected strings",
			"whatever the fuck goes here",
			"is what goes here",
			"yay",
		},
	}
	for {
		log.Println("next")
		wg.Add(1)

		go func() {
			s := hook.Start()

			go func() {
				isCtrl := false
				for {
					i := <-s

					if i.Rawcode == KeyCtrl {
						isCtrl = i.Kind == hook.KeyDown
					}

					if state.hasWindowOpen() {
						if i.Rawcode == KeyEnter || i.Rawcode == KeyTab {
							hook.End()
							state.wnd.Close()
							state.wnd = nil
							log.Println("done")
							break
						}
					}

					if i.Kind == hook.KeyDown {
						if i.Rawcode == KeyUp {
							if state.currentOpt > 0 {
								state.currentOpt -= 1
							}
							log.Println("up", state.currentOpt)
						}
						if i.Rawcode == KeyDown {
							if state.currentOpt < len(state.opts)-1 {
								state.currentOpt += 1
							}
							log.Println("down", state.currentOpt)
						}
					}

					if i.Kind == hook.KeyDown && i.Rawcode < 255 {
						key := string(int32(i.Rawcode))
						if isCtrl {
							if key == "c" || key == "q" {
								os.Exit(0)
							} else if !state.hasWindowOpen() && "a" == key {
								go renderWindow(&state)
								log.Printf("Opening")
							}
						}
						// log.Printf("evt: %v\n", string(i.Rawcode))
					}
					// log.Printf("evt: %v\n", i)
				}
				// time.Sleep(250 * time.Millisecond)
				wg.Done()
			}()
			// state.currentOpt = 0
			// robotgo.TypeStr(state.opts[state.currentOpt])
			defer robotgo.TypeStr(state.opts[state.currentOpt])
			// time.Sleep(250 * time.Millisecond)
		}()

		// go func() {
		// 	flags := g.MasterWindowFlagsFloating | g.MasterWindowFlagsFrameless | g.MasterWindowFlagsNotResizable
		// 	wnd = g.NewMasterWindow("Hello world", 400, 200, flags)
		// 	wnd.Run(loop)
		// }()

		wg.Wait()
	}

	defer robotgo.TypeStr(state.opts[state.currentOpt])
}
