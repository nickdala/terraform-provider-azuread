package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallPacketQueueingMethodType string

const (
	FirewallPacketQueueingMethodType_DeviceDefault FirewallPacketQueueingMethodType = "deviceDefault"
	FirewallPacketQueueingMethodType_Disabled      FirewallPacketQueueingMethodType = "disabled"
	FirewallPacketQueueingMethodType_QueueBoth     FirewallPacketQueueingMethodType = "queueBoth"
	FirewallPacketQueueingMethodType_QueueInbound  FirewallPacketQueueingMethodType = "queueInbound"
	FirewallPacketQueueingMethodType_QueueOutbound FirewallPacketQueueingMethodType = "queueOutbound"
)

func PossibleValuesForFirewallPacketQueueingMethodType() []string {
	return []string{
		string(FirewallPacketQueueingMethodType_DeviceDefault),
		string(FirewallPacketQueueingMethodType_Disabled),
		string(FirewallPacketQueueingMethodType_QueueBoth),
		string(FirewallPacketQueueingMethodType_QueueInbound),
		string(FirewallPacketQueueingMethodType_QueueOutbound),
	}
}

func (s *FirewallPacketQueueingMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallPacketQueueingMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallPacketQueueingMethodType(input string) (*FirewallPacketQueueingMethodType, error) {
	vals := map[string]FirewallPacketQueueingMethodType{
		"devicedefault": FirewallPacketQueueingMethodType_DeviceDefault,
		"disabled":      FirewallPacketQueueingMethodType_Disabled,
		"queueboth":     FirewallPacketQueueingMethodType_QueueBoth,
		"queueinbound":  FirewallPacketQueueingMethodType_QueueInbound,
		"queueoutbound": FirewallPacketQueueingMethodType_QueueOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallPacketQueueingMethodType(input)
	return &out, nil
}