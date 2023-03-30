// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"
	"fmt"
)

// RelationshipToTeamLinkData Relationship between a link and a team
type RelationshipToTeamLinkData struct {
	// The team link's identifier
	Id string `json:"id"`
	// Team link type
	Type TeamLinkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewRelationshipToTeamLinkData instantiates a new RelationshipToTeamLinkData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRelationshipToTeamLinkData(id string, typeVar TeamLinkType) *RelationshipToTeamLinkData {
	this := RelationshipToTeamLinkData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewRelationshipToTeamLinkDataWithDefaults instantiates a new RelationshipToTeamLinkData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRelationshipToTeamLinkDataWithDefaults() *RelationshipToTeamLinkData {
	this := RelationshipToTeamLinkData{}
	var typeVar TeamLinkType = TEAMLINKTYPE_TEAM_LINKS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *RelationshipToTeamLinkData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RelationshipToTeamLinkData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *RelationshipToTeamLinkData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *RelationshipToTeamLinkData) GetType() TeamLinkType {
	if o == nil {
		var ret TeamLinkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelationshipToTeamLinkData) GetTypeOk() (*TeamLinkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *RelationshipToTeamLinkData) SetType(v TeamLinkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RelationshipToTeamLinkData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RelationshipToTeamLinkData) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	required := struct {
		Id   *string       `json:"id"`
		Type *TeamLinkType `json:"type"`
	}{}
	all := struct {
		Id   string       `json:"id"`
		Type TeamLinkType `json:"type"`
	}{}
	err = json.Unmarshal(bytes, &required)
	if err != nil {
		return err
	}
	if required.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if required.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	err = json.Unmarshal(bytes, &all)
	if err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	if v := all.Type; !v.IsValid() {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	o.Id = all.Id
	o.Type = all.Type
	return nil
}
