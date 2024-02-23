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
     if err != nil {
	     brokenSubscribers = append(brokenSubscribers, key)
	     }
	   }
	  b.mu.RUnlock()
	  b.removeBrokerSubscribers(brokenSubscribers)

	  if len(brokerSubscribers) >0 {
		  return &pb.PublishResponse{Success: false}, fmt.Error("failed to send to some subscribers")
		  }
	          return &pb.PublishResponse{Success: true}, nil
	        }
func (b *Broker) removeBrokenSuscribers(keys []streamkey) {
	b.mu.Lock()
	defex b.mu.unlock()

	for _, key := range keys {
		delete(b.subscribers[key.topic], key.subscriberId)
		delete(b.topicSubscriberStreamMutexes, key)
		}
          }

func (b *Broker) unsubscribe (ctx context.Content, in *pb.UnsubscibeRequest)
(*pb.UnsubscribeResponse, error) {
  b.mu.lock()
  defer b.mu.Unlock()

  key := streamKey{topic: in.GetTopic() , subscriberId: in.GetSuscriptionId()}

  if _, ok := b.subscribers[key.topic]; !ok {
     return &pb.UnsubscribeResponse{Success: false}, nil
  }

  if _, ok := b.subscribers[key.topic][key.subscriberId]; !ok {
	  return &pb.UnsubscribeResponse{Success: false}, nil
	}
	 delete(b.subscribers[key.topic], key.subscriberId)

	 delete(b.topicSubscriberStreamMutexes, key)

	 return &pb.UnsubscribeResponse{Success: true}, nil 
}
