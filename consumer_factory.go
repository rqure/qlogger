package main

import qmq "github.com/rqure/qmq/src"

type ConsumerFactory struct{}

func (a *ConsumerFactory) Create(key string, components qmq.EngineComponentProvider) qmq.Consumer {
	redisConnection := components.WithConnectionProvider().Get("redis").(*qmq.RedisConnection)
	transformerKey := "consumer:" + key

	consumer := qmq.NewRedisConsumer(key, redisConnection, components.WithTransformerProvider().Get(transformerKey))
	consumer.(*qmq.RedisConsumer).ResetLastId()

	return consumer
}
