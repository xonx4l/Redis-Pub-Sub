type subscriberStream pb.PubSubService_SubscribeServer

type streamKey struct {
	topic      string
	subscribed unit32
}

type Broker struct {
	pb.UnimplementedPubSubServiceServer
	port                         string
	listener                     net.listener
	grpcServer                   *grpc.Server
	subscribers                  map[string]map[unit32]subscriberStream
	topicSubscriberStreamMutexes map[streamKey]*sync.Mutex
	mu                           sync.RWMutex
	ctx                          context.context
	cancel                       context.cancelfunc
}
