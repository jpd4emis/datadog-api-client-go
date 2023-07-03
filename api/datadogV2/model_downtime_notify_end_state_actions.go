// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/goccy/go-json"
)

// DowntimeNotifyEndStateActions Action that will trigger a monitor notification if the downtime is in the `notify_end_types` state.
type DowntimeNotifyEndStateActions string

// List of DowntimeNotifyEndStateActions.
const (
	DOWNTIMENOTIFYENDSTATEACTIONS_CANCELED DowntimeNotifyEndStateActions = "canceled"
	DOWNTIMENOTIFYENDSTATEACTIONS_EXPIRED  DowntimeNotifyEndStateActions = "expired"
)

var allowedDowntimeNotifyEndStateActionsEnumValues = []DowntimeNotifyEndStateActions{
	DOWNTIMENOTIFYENDSTATEACTIONS_CANCELED,
	DOWNTIMENOTIFYENDSTATEACTIONS_EXPIRED,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DowntimeNotifyEndStateActions) GetAllowedValues() []DowntimeNotifyEndStateActions {
	return allowedDowntimeNotifyEndStateActionsEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DowntimeNotifyEndStateActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DowntimeNotifyEndStateActions(value)
	return nil
}

// NewDowntimeNotifyEndStateActionsFromValue returns a pointer to a valid DowntimeNotifyEndStateActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDowntimeNotifyEndStateActionsFromValue(v string) (*DowntimeNotifyEndStateActions, error) {
	ev := DowntimeNotifyEndStateActions(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DowntimeNotifyEndStateActions: valid values are %v", v, allowedDowntimeNotifyEndStateActionsEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DowntimeNotifyEndStateActions) IsValid() bool {
	for _, existing := range allowedDowntimeNotifyEndStateActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DowntimeNotifyEndStateActions value.
func (v DowntimeNotifyEndStateActions) Ptr() *DowntimeNotifyEndStateActions {
	return &v
}

// NullableDowntimeNotifyEndStateActions handles when a null is used for DowntimeNotifyEndStateActions.
type NullableDowntimeNotifyEndStateActions struct {
	value *DowntimeNotifyEndStateActions
	isSet bool
}

// Get returns the associated value.
func (v NullableDowntimeNotifyEndStateActions) Get() *DowntimeNotifyEndStateActions {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDowntimeNotifyEndStateActions) Set(val *DowntimeNotifyEndStateActions) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDowntimeNotifyEndStateActions) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag.
func (v *NullableDowntimeNotifyEndStateActions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDowntimeNotifyEndStateActions initializes the struct as if Set has been called.
func NewNullableDowntimeNotifyEndStateActions(val *DowntimeNotifyEndStateActions) *NullableDowntimeNotifyEndStateActions {
	return &NullableDowntimeNotifyEndStateActions{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDowntimeNotifyEndStateActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDowntimeNotifyEndStateActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
