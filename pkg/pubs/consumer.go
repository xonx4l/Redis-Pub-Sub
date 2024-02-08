func (c*Consumer) Subscribe(topic string ) error {
	 
	_, ok := c.subscriptions.Load(topic)
	if ok {
		return nil
	}
}

streamCtx, streamCancel := context.WithCancel(c.ctx)

stream, err := c.client,Subscribe(streamCtx, &pb.SubscribeRequest{subscribedId: c.ID, Topic: topic })

if err != nil {
	streamCancel()
	return err
}

c.subscription.store(topic, streamCancel)

go c.recieve(stream, streamCtx)

return nil 

}

func (c*Consumer) recieve (stream pb.PubSubService_subscribeClient, ctx content.Content){
	 for {
		select {
		case <-ctx.Done():
			stream.CloseSend()
			return
		default:
			msg,err := stream.Recv
			if err != nil {
				if err == io.EOF{
					return
				}
				return
			}

			c.message <- msg
		}
	 }
}
