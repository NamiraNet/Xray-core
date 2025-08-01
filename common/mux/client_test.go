package mux_test

import (
	"context"
	"testing"
	"time"

	"github.com/NamiraNet/xray-core/common"
	"github.com/NamiraNet/xray-core/common/errors"
	"github.com/NamiraNet/xray-core/common/mux"
	"github.com/NamiraNet/xray-core/common/net"
	"github.com/NamiraNet/xray-core/common/session"
	"github.com/NamiraNet/xray-core/testing/mocks"
	"github.com/NamiraNet/xray-core/transport"
	"github.com/NamiraNet/xray-core/transport/pipe"
	"github.com/golang/mock/gomock"
)

func TestIncrementalPickerFailure(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockWorkerFactory := mocks.NewMuxClientWorkerFactory(mockCtl)
	mockWorkerFactory.EXPECT().Create().Return(nil, errors.New("test"))

	picker := mux.IncrementalWorkerPicker{
		Factory: mockWorkerFactory,
	}

	_, err := picker.PickAvailable()
	if err == nil {
		t.Error("expected error, but nil")
	}
}

func TestClientWorkerEOF(t *testing.T) {
	reader, writer := pipe.New(pipe.WithoutSizeLimit())
	common.Must(writer.Close())

	worker, err := mux.NewClientWorker(transport.Link{Reader: reader, Writer: writer}, mux.ClientStrategy{})
	common.Must(err)

	time.Sleep(time.Millisecond * 500)

	f := worker.Dispatch(context.Background(), nil)
	if f {
		t.Error("expected failed dispatching, but actually not")
	}
}

func TestClientWorkerClose(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	r1, w1 := pipe.New(pipe.WithoutSizeLimit())
	worker1, err := mux.NewClientWorker(transport.Link{
		Reader: r1,
		Writer: w1,
	}, mux.ClientStrategy{
		MaxConcurrency: 4,
		MaxConnection:  4,
	})
	common.Must(err)

	r2, w2 := pipe.New(pipe.WithoutSizeLimit())
	worker2, err := mux.NewClientWorker(transport.Link{
		Reader: r2,
		Writer: w2,
	}, mux.ClientStrategy{
		MaxConcurrency: 4,
		MaxConnection:  4,
	})
	common.Must(err)

	factory := mocks.NewMuxClientWorkerFactory(mockCtl)
	gomock.InOrder(
		factory.EXPECT().Create().Return(worker1, nil),
		factory.EXPECT().Create().Return(worker2, nil),
	)

	picker := &mux.IncrementalWorkerPicker{
		Factory: factory,
	}
	manager := &mux.ClientManager{
		Picker: picker,
	}

	tr1, tw1 := pipe.New(pipe.WithoutSizeLimit())
	ctx1 := session.ContextWithOutbounds(context.Background(), []*session.Outbound{{
		Target: net.TCPDestination(net.DomainAddress("www.example.com"), 80),
	}})
	common.Must(manager.Dispatch(ctx1, &transport.Link{
		Reader: tr1,
		Writer: tw1,
	}))
	defer tw1.Close()

	common.Must(w1.Close())

	time.Sleep(time.Millisecond * 500)
	if !worker1.Closed() {
		t.Error("worker1 is not finished")
	}

	tr2, tw2 := pipe.New(pipe.WithoutSizeLimit())
	ctx2 := session.ContextWithOutbounds(context.Background(), []*session.Outbound{{
		Target: net.TCPDestination(net.DomainAddress("www.example.com"), 80),
	}})
	common.Must(manager.Dispatch(ctx2, &transport.Link{
		Reader: tr2,
		Writer: tw2,
	}))
	defer tw2.Close()

	common.Must(w2.Close())
}
