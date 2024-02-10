// publish publishes a message to the given code 
func (p *Publisher) publish(topic string, message []byte) error {
  -, err := p.client.Publish(p.ctx, &pb.PublishRequeest{
    Topic: topic,
    Message: ,message
    )}
    return err
  }
