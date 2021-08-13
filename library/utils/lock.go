package utils

type MutexLock struct {
	mutex chan struct{}
}

func (l *MutexLock) Lock() {
	if l.mutex == nil {
		l.mutex = make(chan struct{}, 1)
	}
	l.mutex <- struct{}{}
}

func (l *MutexLock) Unlock() {
	if l.mutex == nil {
		return
	}
	<-l.mutex
}
