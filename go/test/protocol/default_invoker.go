// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
package protocol

type DefaultInvoker struct {
	CommandName         *string           `toml:"command-name"`
	Serializer          DefaultSerializer `toml:"serializer"`
	RequestTopic        *string           `toml:"request-topic"`
	TopicNamespace      *string           `toml:"topic-namespace"`
	ResponseTopicPrefix *string           `toml:"response-topic-prefix"`
	ResponseTopicSuffix *string           `toml:"response-topic-suffix"`
}

func (invoker *DefaultInvoker) GetCommandName() *string {
	if invoker.CommandName == nil {
		return nil
	}

	commandName := *invoker.CommandName
	return &commandName
}

func (invoker *DefaultInvoker) GetSerializer() TestCaseSerializer {
	return invoker.Serializer.GetSerializer()
}

func (invoker *DefaultInvoker) GetRequestTopic() *string {
	if invoker.RequestTopic == nil {
		return nil
	}

	requestTopic := *invoker.RequestTopic
	return &requestTopic
}

func (invoker *DefaultInvoker) GetTopicNamespace() *string {
	if invoker.TopicNamespace == nil {
		return nil
	}

	topicNamespace := *invoker.TopicNamespace
	return &topicNamespace
}

func (invoker *DefaultInvoker) GetResponseTopicPrefix() *string {
	if invoker.ResponseTopicPrefix == nil {
		return nil
	}

	responseTopicPrefix := *invoker.ResponseTopicPrefix
	return &responseTopicPrefix
}

func (invoker *DefaultInvoker) GetResponseTopicSuffix() *string {
	if invoker.ResponseTopicSuffix == nil {
		return nil
	}

	responseTopicSuffix := *invoker.ResponseTopicSuffix
	return &responseTopicSuffix
}
