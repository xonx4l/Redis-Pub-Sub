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
func (b *Broker) Publish(ctx context.Context, in *pb.PublishRequest) (*pb.PublishResponse, error) {
   b.mu.RLock()

   brokerSubscribers := make([]streamKey, 0)
   for subscriberId, stream := range b.subscribers[in.GetTopic()] {
     key := streamKey{topic: in.GetTopic(), subscriberId: subscriberId}
     b.topicSubscriberStreamMutexes[key].Lock()

     err := stream.Send(&pb.Message{Topic: in.GetTopic(), Message: in.GetMessage()})
     b.topicSubscriberStreamMutexes[key].Unlock()

	   
