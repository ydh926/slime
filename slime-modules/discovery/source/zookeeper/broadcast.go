package zookeeper

type BroadcastService struct {
	chListeners map[chan struct{}]bool
}

func (b *BroadcastService) Register(ch chan struct{}) {
	if !b.chListeners[ch] {
		b.chListeners[ch] = true
	}
}

func (b *BroadcastService) BroadCast(message struct{}) {
	for l := range b.chListeners {
		func() {
			defer func() {
				if r := recover(); r != nil {
					// todo log
				}
			}()
			l <- message
		}()
	}
}

func (b *BroadcastService) DeRegister(ch chan struct{}) {
	close(ch)
	delete(b.chListeners, ch)
}
