package Observer

type Subscriber interface {
	Notify(interface{}) error
	Close()
}

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out,
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {

	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}

	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

type publisher struct {
	subscribers []Subscriber

	// nilai dari channel in mesti di-declare ketika akan melakukan AddSubscriber
	addSubCh chan Subscriber

	// nilai dari channel in mesti di-declare ketika akan melakukan RemoveSubscriber
	removeSubCh chan Subscriber

	// nilai dari channel in mesti di-declare ketika akan melakukan PublishMessage
	in chan interface{}

	stop chan struct{}
}

func (p *publisher) AddSubscriber() {
	return p.addSubCh
}

func (p *publisher) RemoveSubscriberCh() {
	return p.removeSubCh
}

func (p *publisher) PublishMessage() {
	return p.in
}

func (p *publisher) Stop() {
	close(p.stop)
}

func (p *publisher) start() {

	// infinite loop untuk waiting request

	for {
		select {
		case msg := <-p.in:

			// sub = NewWriterSubscriber(0, nil)

			for _, ch := range p.subscribers {
				sub.Notify(msg)
			}
		case sub := <-p.addSubCh:
			p.subscribers = append(p.subscribers, sub)
		case sub := <-p.removeSubCh:
			for i, candidate := range p.subscribers {
				if candidate == sub {
					p.subscribers = append(p.subscribers[:i],
						p.subscribers[i+1:]...)
					candidate.Close()
					break
				}
			}
		case <-p.stop:
			for _, sub := range p.subscribers {
				sub.Close()
			}
			close(p.addSubCh)
			close(p.in)
			close(p.removeSubCh)
			return
		}
	}
}
