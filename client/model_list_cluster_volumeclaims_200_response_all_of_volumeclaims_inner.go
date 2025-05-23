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

// checks if the ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner{}

// ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner struct for ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner
type ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ControllerId *int64 `json:"controllerId,omitempty"`
	ControllerMountPoint *string `json:"controllerMountPoint,omitempty"`
	Resizeable *bool `json:"resizeable,omitempty"`
	RootVolume *bool `json:"rootVolume,omitempty"`
	UnitNumber *string `json:"unitNumber,omitempty"`
	DeviceName *string `json:"deviceName,omitempty"`
	DeviceDisplayName *string `json:"deviceDisplayName,omitempty"`
	Type *ListBackupSettings200ResponseBackupSettingsDefaultSchedule `json:"type,omitempty"`
	TypeId *int64 `json:"typeId,omitempty"`
	Category *string `json:"category,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMessage *string `json:"statusMessage,omitempty"`
	ConfigurableIOPS *bool `json:"configurableIOPS,omitempty"`
	MaxStorage *int64 `json:"maxStorage,omitempty"`
	DisplayOrder *int64 `json:"displayOrder,omitempty"`
	MaxIOPS *string `json:"maxIOPS,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
	Active *bool `json:"active,omitempty"`
	Zone *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"zone,omitempty"`
	ZoneId *int64 `json:"zoneId,omitempty"`
	Datastore *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"datastore,omitempty"`
	DatastoreId *int64 `json:"datastoreId,omitempty"`
	StorageGroup *string `json:"storageGroup,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	StorageServer *string `json:"storageServer,omitempty"`
	Source *string `json:"source,omitempty"`
	Owner *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"owner,omitempty"`
}

// NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner instantiates a new ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner {
	this := ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner{}
	return &this
}

// NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerWithDefaults instantiates a new ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListClusterVolumeclaims200ResponseAllOfVolumeclaimsInnerWithDefaults() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner {
	this := ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDescription(v string) {
	o.Description = &v
}

// GetControllerId returns the ControllerId field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerId() int64 {
	if o == nil || IsNil(o.ControllerId) {
		var ret int64
		return ret
	}
	return *o.ControllerId
}

// GetControllerIdOk returns a tuple with the ControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ControllerId) {
		return nil, false
	}
	return o.ControllerId, true
}

// IsSetControllerId returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetControllerId() bool {
	if o != nil && !IsNil(o.ControllerId) {
		return true
	}

	return false
}

// SetControllerId gets a reference to the given int64 and assigns it to the ControllerId field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerId(v int64) {
	o.ControllerId = &v
}

// GetControllerMountPoint returns the ControllerMountPoint field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerMountPoint() string {
	if o == nil || IsNil(o.ControllerMountPoint) {
		var ret string
		return ret
	}
	return *o.ControllerMountPoint
}

// GetControllerMountPointOk returns a tuple with the ControllerMountPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetControllerMountPointOk() (*string, bool) {
	if o == nil || IsNil(o.ControllerMountPoint) {
		return nil, false
	}
	return o.ControllerMountPoint, true
}

// IsSetControllerMountPoint returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetControllerMountPoint() bool {
	if o != nil && !IsNil(o.ControllerMountPoint) {
		return true
	}

	return false
}

// SetControllerMountPoint gets a reference to the given string and assigns it to the ControllerMountPoint field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetControllerMountPoint(v string) {
	o.ControllerMountPoint = &v
}

// GetResizeable returns the Resizeable field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetResizeable() bool {
	if o == nil || IsNil(o.Resizeable) {
		var ret bool
		return ret
	}
	return *o.Resizeable
}

// GetResizeableOk returns a tuple with the Resizeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetResizeableOk() (*bool, bool) {
	if o == nil || IsNil(o.Resizeable) {
		return nil, false
	}
	return o.Resizeable, true
}

// IsSetResizeable returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetResizeable() bool {
	if o != nil && !IsNil(o.Resizeable) {
		return true
	}

	return false
}

// SetResizeable gets a reference to the given bool and assigns it to the Resizeable field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetResizeable(v bool) {
	o.Resizeable = &v
}

// GetRootVolume returns the RootVolume field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRootVolume() bool {
	if o == nil || IsNil(o.RootVolume) {
		var ret bool
		return ret
	}
	return *o.RootVolume
}

// GetRootVolumeOk returns a tuple with the RootVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetRootVolumeOk() (*bool, bool) {
	if o == nil || IsNil(o.RootVolume) {
		return nil, false
	}
	return o.RootVolume, true
}

// IsSetRootVolume returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetRootVolume() bool {
	if o != nil && !IsNil(o.RootVolume) {
		return true
	}

	return false
}

// SetRootVolume gets a reference to the given bool and assigns it to the RootVolume field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetRootVolume(v bool) {
	o.RootVolume = &v
}

// GetUnitNumber returns the UnitNumber field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUnitNumber() string {
	if o == nil || IsNil(o.UnitNumber) {
		var ret string
		return ret
	}
	return *o.UnitNumber
}

// GetUnitNumberOk returns a tuple with the UnitNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUnitNumberOk() (*string, bool) {
	if o == nil || IsNil(o.UnitNumber) {
		return nil, false
	}
	return o.UnitNumber, true
}

// IsSetUnitNumber returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetUnitNumber() bool {
	if o != nil && !IsNil(o.UnitNumber) {
		return true
	}

	return false
}

// SetUnitNumber gets a reference to the given string and assigns it to the UnitNumber field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUnitNumber(v string) {
	o.UnitNumber = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// IsSetDeviceName returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceDisplayName returns the DeviceDisplayName field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceDisplayName() string {
	if o == nil || IsNil(o.DeviceDisplayName) {
		var ret string
		return ret
	}
	return *o.DeviceDisplayName
}

// GetDeviceDisplayNameOk returns a tuple with the DeviceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDeviceDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceDisplayName) {
		return nil, false
	}
	return o.DeviceDisplayName, true
}

// IsSetDeviceDisplayName returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDeviceDisplayName() bool {
	if o != nil && !IsNil(o.DeviceDisplayName) {
		return true
	}

	return false
}

// SetDeviceDisplayName gets a reference to the given string and assigns it to the DeviceDisplayName field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDeviceDisplayName(v string) {
	o.DeviceDisplayName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule {
	if o == nil || IsNil(o.Type) {
		var ret ListBackupSettings200ResponseBackupSettingsDefaultSchedule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ListBackupSettings200ResponseBackupSettingsDefaultSchedule and assigns it to the Type field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule) {
	o.Type = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeId() int64 {
	if o == nil || IsNil(o.TypeId) {
		var ret int64
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// IsSetTypeId returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given int64 and assigns it to the TypeId field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetTypeId(v int64) {
	o.TypeId = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// IsSetCategory returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetCategory(v string) {
	o.Category = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// IsSetStatusMessage returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetConfigurableIOPS returns the ConfigurableIOPS field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetConfigurableIOPS() bool {
	if o == nil || IsNil(o.ConfigurableIOPS) {
		var ret bool
		return ret
	}
	return *o.ConfigurableIOPS
}

// GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetConfigurableIOPSOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfigurableIOPS) {
		return nil, false
	}
	return o.ConfigurableIOPS, true
}

// IsSetConfigurableIOPS returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetConfigurableIOPS() bool {
	if o != nil && !IsNil(o.ConfigurableIOPS) {
		return true
	}

	return false
}

// SetConfigurableIOPS gets a reference to the given bool and assigns it to the ConfigurableIOPS field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetConfigurableIOPS(v bool) {
	o.ConfigurableIOPS = &v
}

// GetMaxStorage returns the MaxStorage field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxStorage() int64 {
	if o == nil || IsNil(o.MaxStorage) {
		var ret int64
		return ret
	}
	return *o.MaxStorage
}

// GetMaxStorageOk returns a tuple with the MaxStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxStorageOk() (*int64, bool) {
	if o == nil || IsNil(o.MaxStorage) {
		return nil, false
	}
	return o.MaxStorage, true
}

// IsSetMaxStorage returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetMaxStorage() bool {
	if o != nil && !IsNil(o.MaxStorage) {
		return true
	}

	return false
}

// SetMaxStorage gets a reference to the given int64 and assigns it to the MaxStorage field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetMaxStorage(v int64) {
	o.MaxStorage = &v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDisplayOrder() int64 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int64
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDisplayOrderOk() (*int64, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// IsSetDisplayOrder returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int64 and assigns it to the DisplayOrder field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDisplayOrder(v int64) {
	o.DisplayOrder = &v
}

// GetMaxIOPS returns the MaxIOPS field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxIOPS() string {
	if o == nil || IsNil(o.MaxIOPS) {
		var ret string
		return ret
	}
	return *o.MaxIOPS
}

// GetMaxIOPSOk returns a tuple with the MaxIOPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetMaxIOPSOk() (*string, bool) {
	if o == nil || IsNil(o.MaxIOPS) {
		return nil, false
	}
	return o.MaxIOPS, true
}

// IsSetMaxIOPS returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetMaxIOPS() bool {
	if o != nil && !IsNil(o.MaxIOPS) {
		return true
	}

	return false
}

// SetMaxIOPS gets a reference to the given string and assigns it to the MaxIOPS field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetMaxIOPS(v string) {
	o.MaxIOPS = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// IsSetUuid returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetUuid(v string) {
	o.Uuid = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetActive(v bool) {
	o.Active = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZone() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Zone) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// IsSetZone returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Zone field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetZone(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Zone = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneId() int64 {
	if o == nil || IsNil(o.ZoneId) {
		var ret int64
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetZoneIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// IsSetZoneId returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int64 and assigns it to the ZoneId field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetZoneId(v int64) {
	o.ZoneId = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastore() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Datastore) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// IsSetDatastore returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Datastore field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastore(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Datastore = &v
}

// GetDatastoreId returns the DatastoreId field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreId() int64 {
	if o == nil || IsNil(o.DatastoreId) {
		var ret int64
		return ret
	}
	return *o.DatastoreId
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetDatastoreIdOk() (*int64, bool) {
	if o == nil || IsNil(o.DatastoreId) {
		return nil, false
	}
	return o.DatastoreId, true
}

// IsSetDatastoreId returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetDatastoreId() bool {
	if o != nil && !IsNil(o.DatastoreId) {
		return true
	}

	return false
}

// SetDatastoreId gets a reference to the given int64 and assigns it to the DatastoreId field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetDatastoreId(v int64) {
	o.DatastoreId = &v
}

// GetStorageGroup returns the StorageGroup field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageGroup() string {
	if o == nil || IsNil(o.StorageGroup) {
		var ret string
		return ret
	}
	return *o.StorageGroup
}

// GetStorageGroupOk returns a tuple with the StorageGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageGroupOk() (*string, bool) {
	if o == nil || IsNil(o.StorageGroup) {
		return nil, false
	}
	return o.StorageGroup, true
}

// IsSetStorageGroup returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetStorageGroup() bool {
	if o != nil && !IsNil(o.StorageGroup) {
		return true
	}

	return false
}

// SetStorageGroup gets a reference to the given string and assigns it to the StorageGroup field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageGroup(v string) {
	o.StorageGroup = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// IsSetNamespace returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetNamespace(v string) {
	o.Namespace = &v
}

// GetStorageServer returns the StorageServer field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServer() string {
	if o == nil || IsNil(o.StorageServer) {
		var ret string
		return ret
	}
	return *o.StorageServer
}

// GetStorageServerOk returns a tuple with the StorageServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetStorageServerOk() (*string, bool) {
	if o == nil || IsNil(o.StorageServer) {
		return nil, false
	}
	return o.StorageServer, true
}

// IsSetStorageServer returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetStorageServer() bool {
	if o != nil && !IsNil(o.StorageServer) {
		return true
	}

	return false
}

// SetStorageServer gets a reference to the given string and assigns it to the StorageServer field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetStorageServer(v string) {
	o.StorageServer = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// IsSetSource returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetSource(v string) {
	o.Source = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Owner) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Owner field.
func (o *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Owner = &v
}

func (o ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ControllerId) {
		toSerialize["controllerId"] = o.ControllerId
	}
	if !IsNil(o.ControllerMountPoint) {
		toSerialize["controllerMountPoint"] = o.ControllerMountPoint
	}
	if !IsNil(o.Resizeable) {
		toSerialize["resizeable"] = o.Resizeable
	}
	if !IsNil(o.RootVolume) {
		toSerialize["rootVolume"] = o.RootVolume
	}
	if !IsNil(o.UnitNumber) {
		toSerialize["unitNumber"] = o.UnitNumber
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.DeviceDisplayName) {
		toSerialize["deviceDisplayName"] = o.DeviceDisplayName
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.TypeId) {
		toSerialize["typeId"] = o.TypeId
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	if !IsNil(o.ConfigurableIOPS) {
		toSerialize["configurableIOPS"] = o.ConfigurableIOPS
	}
	if !IsNil(o.MaxStorage) {
		toSerialize["maxStorage"] = o.MaxStorage
	}
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.MaxIOPS) {
		toSerialize["maxIOPS"] = o.MaxIOPS
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.ZoneId) {
		toSerialize["zoneId"] = o.ZoneId
	}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}
	if !IsNil(o.DatastoreId) {
		toSerialize["datastoreId"] = o.DatastoreId
	}
	if !IsNil(o.StorageGroup) {
		toSerialize["storageGroup"] = o.StorageGroup
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.StorageServer) {
		toSerialize["storageServer"] = o.StorageServer
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	return toSerialize, nil
}

type NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner struct {
	value *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner
	isSet bool
}

func (v NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) Get() *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner {
	return v.value
}

func (v *NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) Set(val *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner(val *ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) *NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner {
	return &NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner{value: val, isSet: true}
}

func (v NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


