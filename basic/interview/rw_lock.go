package interview

import "sync"

type RWLock struct {
	rLock sync.Mutex
	wLock sync.Mutex
	rCnt  int
}

func (l *RWLock) RLock() {
	l.rLock.Lock()
	defer l.rLock.Unlock()

	l.rCnt += 1
	if l.rCnt == 1 {
		l.wLock.Lock()
	}
}

func (l *RWLock) RUnlock() {
	l.rLock.Lock()
	defer l.rLock.Unlock()

	l.rCnt -= 1
	if l.rCnt == 0 {
		l.wLock.Unlock()
	}
}

func (l *RWLock) WLock() {	
	l.wLock.Lock()
}

func (l *RWLock) WUnlock() {
	l.wLock.Unlock()
}
