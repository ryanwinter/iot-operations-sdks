// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

using System.Collections.Generic;

namespace Azure.Iot.Operations.Protocol.Models
{
    public class MqttClientPublishResult(ushort? packetIdentifier, MqttClientPublishReasonCode reasonCode, string reasonString, IReadOnlyCollection<MqttUserProperty> userProperties)
    {

        /// <summary>
        ///     Returns if the overall status of the publish is a success. This can be the reason code _Success_ or
        ///     _NoMatchingSubscribers_. _NoMatchingSubscribers_ usually indicates only that no other client is interested in the
        ///     topic but overall transmit
        ///     to the server etc. was a success.
        /// </summary>
        public bool IsSuccess => ReasonCode is MqttClientPublishReasonCode.Success or MqttClientPublishReasonCode.NoMatchingSubscribers;

        /// <summary>
        ///     Gets the packet identifier which was used for this publish.
        /// </summary>
        public ushort? PacketIdentifier { get; } = packetIdentifier;

        /// <summary>
        ///     Gets or sets the reason code.
        ///     <remarks>MQTT 5.0.0+ feature.</remarks>
        /// </summary>
        public MqttClientPublishReasonCode ReasonCode { get; } = reasonCode;

        /// <summary>
        ///     Gets or sets the reason string.
        ///     <remarks>MQTT 5.0.0+ feature.</remarks>
        /// </summary>
        public string ReasonString { get; } = reasonString;

        /// <summary>
        ///     Gets or sets the user properties.
        ///     In MQTT 5, user properties are basic UTF-8 string key-value pairs that you can append to almost every type of MQTT
        ///     packet.
        ///     As long as you don’t exceed the maximum message size, you can use an unlimited number of user properties to add
        ///     metadata to MQTT messages and pass information between publisher, broker, and subscriber.
        ///     The feature is very similar to the HTTP header concept.
        ///     <remarks>MQTT 5.0.0+ feature.</remarks>
        /// </summary>
        public IReadOnlyCollection<MqttUserProperty> UserProperties { get; } = userProperties;
    }
}