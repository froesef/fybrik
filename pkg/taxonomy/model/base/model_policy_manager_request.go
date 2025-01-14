/*
 * Policy Manager Service
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package base

import (
	"encoding/json"
)

// PolicyManagerRequest struct for PolicyManagerRequest
type PolicyManagerRequest struct {
	Context  *map[string]interface{}    `json:"context,omitempty"`
	Action   PolicyManagerRequestAction `json:"action"`
	Resource Resource                   `json:"resource"`
}

// NewPolicyManagerRequest instantiates a new PolicyManagerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyManagerRequest(action PolicyManagerRequestAction, resource Resource) *PolicyManagerRequest {
	this := PolicyManagerRequest{}
	this.Action = action
	this.Resource = resource
	return &this
}

// NewPolicyManagerRequestWithDefaults instantiates a new PolicyManagerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyManagerRequestWithDefaults() *PolicyManagerRequest {
	this := PolicyManagerRequest{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *PolicyManagerRequest) GetContext() map[string]interface{} {
	if o == nil || o.Context == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyManagerRequest) GetContextOk() (*map[string]interface{}, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *PolicyManagerRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]interface{} and assigns it to the Context field.
func (o *PolicyManagerRequest) SetContext(v map[string]interface{}) {
	o.Context = &v
}

// GetAction returns the Action field value
func (o *PolicyManagerRequest) GetAction() PolicyManagerRequestAction {
	if o == nil {
		var ret PolicyManagerRequestAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PolicyManagerRequest) GetActionOk() (*PolicyManagerRequestAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PolicyManagerRequest) SetAction(v PolicyManagerRequestAction) {
	o.Action = v
}

// GetResource returns the Resource field value
func (o *PolicyManagerRequest) GetResource() Resource {
	if o == nil {
		var ret Resource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *PolicyManagerRequest) GetResourceOk() (*Resource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *PolicyManagerRequest) SetResource(v Resource) {
	o.Resource = v
}

func (o PolicyManagerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullablePolicyManagerRequest struct {
	value *PolicyManagerRequest
	isSet bool
}

func (v NullablePolicyManagerRequest) Get() *PolicyManagerRequest {
	return v.value
}

func (v *NullablePolicyManagerRequest) Set(val *PolicyManagerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyManagerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyManagerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyManagerRequest(val *PolicyManagerRequest) *NullablePolicyManagerRequest {
	return &NullablePolicyManagerRequest{value: val, isSet: true}
}

func (v NullablePolicyManagerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyManagerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
