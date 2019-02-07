// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

import (
	"fmt"

	"github.com/chnykn/bimface/bean/common"
)

//PropertyItem PropertyItem
type PropertyItem struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	Unit      string `json:"unit,omitempty"`
	ValueType int    `json:"valueType,omitempty"`
}

//NewPropertyItem ***
func NewPropertyItem(key string, value string, unit string) *PropertyItem {
	o := &PropertyItem{
		Key:   key,
		Value: value,
		Unit:  unit,
	}
	return o
}

// ToString get the string
func (o *PropertyItem) ToString() string {
	return fmt.Sprintf("PropertyItem [Key=%s, Value=%v, Unit=%s]", o.Key, o.Value, o.Unit)
}

//--------------------------------------------------------------------

//PropertyGroup PropertyGroup
type PropertyGroup struct {
	Group string         `json:"group"`
	Items []PropertyItem `json:"items"`
}

//NewPropertyGroup ***
func NewPropertyGroup(name string) *PropertyGroup {
	o := &PropertyGroup{
		Group: name,
		Items: make([]PropertyItem, 0),
	}
	return o
}

//--------------------------------------------------------------------

//PropertyPack ***
type PropertyPack struct {
	ElementID   string             `json:"elementId"`
	Name        string             `json:"name"`
	GUID        string             `json:"guid"`
	FamilyGUID  string             `json:"familyGuid"`
	BoundingBox common.BoundingBox `json:"boundingBox"`
	Properties  []PropertyGroup    `json:"properties"`
}

//NewPropertyPack ***
func NewPropertyPack(elementID string) *PropertyPack {
	o := &PropertyPack{
		ElementID: elementID,
		//Groups: make([]PropertyGroup, 0),
	}
	return o
}

//AddPropertyGroup ***
func (o *PropertyPack) AddPropertyGroup(group PropertyGroup) {
	o.Properties = append(o.Properties, group)
}

// ToString get the string
func (o *PropertyPack) ToString() string {
	return fmt.Sprintf("Properties [ElementID=%s, Name=%s, GUID=%s, FamilyGUID=%s,BoundingBox=%v, Properties=%v]",
		o.ElementID, o.Name, o.GUID, o.FamilyGUID, o.BoundingBox, o.Properties)
}
