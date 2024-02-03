func (c*Consumer) Subscribe(topic string ) error {
	 
	_, ok := c.subscriptions.Load(topic)
	if ok {
		return nil
	}
}

streamCtx, streamCancel := context.WithCancel(c.ctx)

stream, err := c.client,Subscribe(streamCtx, &pb.SubscribeRequest{subscribedId: c.ID, Topic: topic })
