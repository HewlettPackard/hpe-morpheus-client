/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.6
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the SaveClusterDatastoreRequestCluster type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveClusterDatastoreRequestCluster{}

// SaveClusterDatastoreRequestCluster struct for SaveClusterDatastoreRequestCluster
type SaveClusterDatastoreRequestCluster struct {
	Name *string `json:"name,omitempty"`
	DatastoreType *GetAlerts200ResponseAllOfChecksInnerAccount `json:"datastoreType,omitempty"`
	StorageServer *GetAlerts200ResponseAllOfChecksInnerAccount `json:"storageServer,omitempty"`
	Visibility *string `json:"visibility,omitempty"`
	Active *bool `json:"active,omitempty"`
	DefaultStore *bool `json:"defaultStore,omitempty"`
	Config *SaveClusterDatastoreRequestClusterConfig `json:"config,omitempty"`
	Tenants []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner `json:"tenants,omitempty"`
	ResourcePermissions *SaveClusterDatastoreRequestClusterResourcePermissions `json:"resourcePermissions,omitempty"`
	Datastores []map[string]interface{} `json:"datastores,omitempty"`
}

// NewSaveClusterDatastoreRequestCluster instantiates a new SaveClusterDatastoreRequestCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveClusterDatastoreRequestCluster() *SaveClusterDatastoreRequestCluster {
	this := SaveClusterDatastoreRequestCluster{}
	return &this
}

// NewSaveClusterDatastoreRequestClusterWithDefaults instantiates a new SaveClusterDatastoreRequestCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveClusterDatastoreRequestClusterWithDefaults() *SaveClusterDatastoreRequestCluster {
	this := SaveClusterDatastoreRequestCluster{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SaveClusterDatastoreRequestCluster) SetName(v string) {
	o.Name = &v
}

// GetDatastoreType returns the DatastoreType field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetDatastoreType() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.DatastoreType) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.DatastoreType
}

// GetDatastoreTypeOk returns a tuple with the DatastoreType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetDatastoreTypeOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.DatastoreType) {
		return nil, false
	}
	return o.DatastoreType, true
}

// IsSetDatastoreType returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetDatastoreType() bool {
	if o != nil && !IsNil(o.DatastoreType) {
		return true
	}

	return false
}

// SetDatastoreType gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the DatastoreType field.
func (o *SaveClusterDatastoreRequestCluster) SetDatastoreType(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.DatastoreType = &v
}

// GetStorageServer returns the StorageServer field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetStorageServer() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.StorageServer) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetStorageServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.StorageServer) {
		return nil, false
	}
	return o.StorageServer, true
}

// IsSetStorageServer returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetStorageServer() bool {
	if o != nil && !IsNil(o.StorageServer) {
		return true
	}

	return false
}

// SetStorageServer gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the StorageServer field.
func (o *SaveClusterDatastoreRequestCluster) SetStorageServer(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.StorageServer = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *SaveClusterDatastoreRequestCluster) SetVisibility(v string) {
	o.Visibility = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SaveClusterDatastoreRequestCluster) SetActive(v bool) {
	o.Active = &v
}

// GetDefaultStore returns the DefaultStore field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetDefaultStore() bool {
	if o == nil || IsNil(o.DefaultStore) {
		var ret bool
		return ret
	}
	return *o.DefaultStore
}

// GetDefaultStoreOk returns a tuple with the DefaultStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetDefaultStoreOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultStore) {
		return nil, false
	}
	return o.DefaultStore, true
}

// IsSetDefaultStore returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetDefaultStore() bool {
	if o != nil && !IsNil(o.DefaultStore) {
		return true
	}

	return false
}

// SetDefaultStore gets a reference to the given bool and assigns it to the DefaultStore field.
func (o *SaveClusterDatastoreRequestCluster) SetDefaultStore(v bool) {
	o.DefaultStore = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetConfig() SaveClusterDatastoreRequestClusterConfig {
	if o == nil || IsNil(o.Config) {
		var ret SaveClusterDatastoreRequestClusterConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetConfigOk() (*SaveClusterDatastoreRequestClusterConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// IsSetConfig returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given SaveClusterDatastoreRequestClusterConfig and assigns it to the Config field.
func (o *SaveClusterDatastoreRequestCluster) SetConfig(v SaveClusterDatastoreRequestClusterConfig) {
	o.Config = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetTenants() []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner {
	if o == nil || IsNil(o.Tenants) {
		var ret []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetTenantsOk() ([]ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// IsSetTenants returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner and assigns it to the Tenants field.
func (o *SaveClusterDatastoreRequestCluster) SetTenants(v []ListCloudDatastores200ResponseAllOfDatastoresInnerTenantsInner) {
	o.Tenants = v
}

// GetResourcePermissions returns the ResourcePermissions field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetResourcePermissions() SaveClusterDatastoreRequestClusterResourcePermissions {
	if o == nil || IsNil(o.ResourcePermissions) {
		var ret SaveClusterDatastoreRequestClusterResourcePermissions
		return ret
	}
	return *o.ResourcePermissions
}

// GetResourcePermissionsOk returns a tuple with the ResourcePermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetResourcePermissionsOk() (*SaveClusterDatastoreRequestClusterResourcePermissions, bool) {
	if o == nil || IsNil(o.ResourcePermissions) {
		return nil, false
	}
	return o.ResourcePermissions, true
}

// IsSetResourcePermissions returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetResourcePermissions() bool {
	if o != nil && !IsNil(o.ResourcePermissions) {
		return true
	}

	return false
}

// SetResourcePermissions gets a reference to the given SaveClusterDatastoreRequestClusterResourcePermissions and assigns it to the ResourcePermissions field.
func (o *SaveClusterDatastoreRequestCluster) SetResourcePermissions(v SaveClusterDatastoreRequestClusterResourcePermissions) {
	o.ResourcePermissions = &v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *SaveClusterDatastoreRequestCluster) GetDatastores() []map[string]interface{} {
	if o == nil || IsNil(o.Datastores) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveClusterDatastoreRequestCluster) GetDatastoresOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Datastores) {
		return nil, false
	}
	return o.Datastores, true
}

// IsSetDatastores returns a boolean if a field has been set.
func (o *SaveClusterDatastoreRequestCluster) IsSetDatastores() bool {
	if o != nil && !IsNil(o.Datastores) {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []map[string]interface{} and assigns it to the Datastores field.
func (o *SaveClusterDatastoreRequestCluster) SetDatastores(v []map[string]interface{}) {
	o.Datastores = v
}

func (o SaveClusterDatastoreRequestCluster) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveClusterDatastoreRequestCluster) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.DatastoreType) {
		toSerialize["datastoreType"] = o.DatastoreType
	}
	if !IsNil(o.StorageServer) {
		toSerialize["storageServer"] = o.StorageServer
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.DefaultStore) {
		toSerialize["defaultStore"] = o.DefaultStore
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.ResourcePermissions) {
		toSerialize["resourcePermissions"] = o.ResourcePermissions
	}
	if !IsNil(o.Datastores) {
		toSerialize["datastores"] = o.Datastores
	}
	return toSerialize, nil
}

type NullableSaveClusterDatastoreRequestCluster struct {
	value *SaveClusterDatastoreRequestCluster
	isSet bool
}

func (v NullableSaveClusterDatastoreRequestCluster) Get() *SaveClusterDatastoreRequestCluster {
	return v.value
}

func (v *NullableSaveClusterDatastoreRequestCluster) Set(val *SaveClusterDatastoreRequestCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveClusterDatastoreRequestCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveClusterDatastoreRequestCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveClusterDatastoreRequestCluster(val *SaveClusterDatastoreRequestCluster) *NullableSaveClusterDatastoreRequestCluster {
	return &NullableSaveClusterDatastoreRequestCluster{value: val, isSet: true}
}

func (v NullableSaveClusterDatastoreRequestCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveClusterDatastoreRequestCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


