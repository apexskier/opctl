package docker

import (
	"sync"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/opctl/opctl/sdks/go/model"
)

type mockPublisher struct {
	events []model.Event
	mu     sync.Mutex
	wg     sync.WaitGroup
}

func NewMockPublisher() *mockPublisher {
	return &mockPublisher{
		events: make([]model.Event, 0),
	}
}

func (p *mockPublisher) Publish(e model.Event) {
	defer p.wg.Done()
	p.mu.Lock()
	p.events = append(p.events, e)
	p.mu.Unlock()
}

func TestStdOutWriteCloser(t *testing.T) {
	g := NewGomegaWithT(t)

	/* arrange */
	eventPublisher := make(chan model.Event, 2)
	containerId := "containerID"

	objectUnderTest := NewStdOutWriteCloser(
		eventPublisher,
		&model.ContainerCall{ContainerID: containerId},
	)

	/* act */
	_, err := objectUnderTest.Write([]byte("testing 1\ntesting 2"))
	if err != nil {
		panic(err)
	}
	err = objectUnderTest.Close()
	g.Expect(err).To(BeNil())

	/* assert */
	g.Expect((<-eventPublisher).ContainerStdOutWrittenTo).To(Equal(&model.ContainerStdOutWrittenTo{ContainerID: containerId, Data: []byte("testing 1\n")}))
	g.Expect((<-eventPublisher).ContainerStdOutWrittenTo).To(Equal(&model.ContainerStdOutWrittenTo{ContainerID: containerId, Data: []byte("testing 2")}))
}
