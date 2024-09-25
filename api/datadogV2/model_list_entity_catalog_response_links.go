// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListEntityCatalogResponseLinks List entity response links.
type ListEntityCatalogResponseLinks struct {
	// next link.
	Next *string `json:"next,omitempty"`
	// previous link.
	Previous *string `json:"previous,omitempty"`
	// current link.
	Self *string `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListEntityCatalogResponseLinks instantiates a new ListEntityCatalogResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListEntityCatalogResponseLinks() *ListEntityCatalogResponseLinks {
	this := ListEntityCatalogResponseLinks{}
	return &this
}

// NewListEntityCatalogResponseLinksWithDefaults instantiates a new ListEntityCatalogResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListEntityCatalogResponseLinksWithDefaults() *ListEntityCatalogResponseLinks {
	this := ListEntityCatalogResponseLinks{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListEntityCatalogResponseLinks) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponseLinks) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListEntityCatalogResponseLinks) HasNext() bool {
	return o != nil && o.Next != nil
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListEntityCatalogResponseLinks) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ListEntityCatalogResponseLinks) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponseLinks) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListEntityCatalogResponseLinks) HasPrevious() bool {
	return o != nil && o.Previous != nil
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *ListEntityCatalogResponseLinks) SetPrevious(v string) {
	o.Previous = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ListEntityCatalogResponseLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEntityCatalogResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ListEntityCatalogResponseLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ListEntityCatalogResponseLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListEntityCatalogResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListEntityCatalogResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Next     *string `json:"next,omitempty"`
		Previous *string `json:"previous,omitempty"`
		Self     *string `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next", "previous", "self"})
	} else {
		return err
	}
	o.Next = all.Next
	o.Previous = all.Previous
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}