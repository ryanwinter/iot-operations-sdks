﻿// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

namespace Azure.Iot.Operations.Protocol.Models
{
    public enum MqttAuthenticateReasonCode
    {
        Success = 0,
        ContinueAuthentication = 24,
        ReAuthenticate = 25
    }
}
