package wg

import "sync"

type BiliWaitGroup struct {
	sync.WaitGroup
}

func (bwg *BiliWaitGroup) Run(f func()) {
	bwg.Add(1)
	go func() {
		f()
		bwg.Done()
	}()
}
