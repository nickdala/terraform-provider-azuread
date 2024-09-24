package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHealthIssueStatus string

const (
	SecurityHealthIssueStatus_Closed     SecurityHealthIssueStatus = "closed"
	SecurityHealthIssueStatus_Open       SecurityHealthIssueStatus = "open"
	SecurityHealthIssueStatus_Suppressed SecurityHealthIssueStatus = "suppressed"
)

func PossibleValuesForSecurityHealthIssueStatus() []string {
	return []string{
		string(SecurityHealthIssueStatus_Closed),
		string(SecurityHealthIssueStatus_Open),
		string(SecurityHealthIssueStatus_Suppressed),
	}
}

func (s *SecurityHealthIssueStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHealthIssueStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHealthIssueStatus(input string) (*SecurityHealthIssueStatus, error) {
	vals := map[string]SecurityHealthIssueStatus{
		"closed":     SecurityHealthIssueStatus_Closed,
		"open":       SecurityHealthIssueStatus_Open,
		"suppressed": SecurityHealthIssueStatus_Suppressed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHealthIssueStatus(input)
	return &out, nil
}