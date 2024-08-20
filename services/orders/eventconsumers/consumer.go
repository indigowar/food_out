package eventconsumers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// fatalError an error that is used to signal to abort the consumer
var fatalError = errors.New("fatal error")

// event - this type is used for consumer template
type event interface {
	protoreflect.ProtoMessage
}

// validator validates incoming data (and may modify it)
type validator[E event, D any] func(E) (D, error)

// unpacker unpacks the value from [[kafka.Message] into incoming data
type unpacker[E event] func(kafka.Message) (E, error)

// Executor is used to execute the command on consumer
type executioner[D any] func(context.Context, D) error

type Consumer interface {
	Run(ctx context.Context) error
}

type consumer[E event, D any] struct {
	id string

	unpacker    unpacker[E]
	validator   validator[E, D]
	executioner executioner[D]

	logger *slog.Logger
	reader *kafka.Reader

	waitGroup sync.WaitGroup
}

func (c *consumer[E, D]) Run(ctx context.Context) error {
	requests := make(chan E)
	validatedData := make(chan D)
	errChan := make(chan error)

	runCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go c.consume(runCtx, requests, errChan)
	go c.validate(runCtx, requests, validatedData, errChan)
	go c.execute(runCtx, validatedData, errChan)

	go func() {
		c.waitGroup.Done()
		close(errChan)
	}()

	for err := range errChan {
		if errors.Is(err, fatalError) {
			c.logger.Error("Fatal Consumer error", "consumer", c.id, "err", err)
			return err
		}
		c.logger.Warn("Consumer error", "consumer", c.id, "err", err)
	}

	return nil
}

func (c *consumer[E, D]) consume(ctx context.Context, requests chan<- E, errors chan<- error) {
	defer c.waitGroup.Done()
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			errors <- fmt.Errorf("consume: read: %w", err)
			continue
		}

		event, err := c.unpacker(message)
		if err != nil {
			errors <- fmt.Errorf("consume: unpack: %w", err)
			continue
		}

		requests <- event
	}
}

func (c *consumer[Event, Data]) validate(
	ctx context.Context,
	requests <-chan Event,
	results chan<- Data,
	errors chan<- error,
) {
	defer c.waitGroup.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-requests:
			data, err := c.validator(request)
			if err != nil {
				errors <- fmt.Errorf("validation: %w", err)
				continue
			}

			results <- data
		}
	}
}

func (c *consumer[Event, Data]) execute(
	ctx context.Context,
	requests <-chan Data,
	errors chan<- error,
) {
	defer c.waitGroup.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-requests:
			err := c.executioner(ctx, request)
			if err != nil {
				errors <- fmt.Errorf("executioner: %w", err)
			}
		}
	}
}

func newConsumer[E event, D any](
	id string,
	unpacker unpacker[E],
	validator validator[E, D],
	executioner executioner[D],
	logger *slog.Logger,
	brokers []string,
	groupID string,
	topic string,
	partition int,
) (*consumer[E, D], error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		GroupID:   groupID,
		Topic:     topic,
		Partition: partition,
	})

	return &consumer[E, D]{
		id:          id,
		unpacker:    unpacker,
		validator:   validator,
		executioner: executioner,
		logger:      logger,
		reader:      reader,
		waitGroup:   sync.WaitGroup{},
	}, nil
}
