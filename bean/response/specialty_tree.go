// Copyright 2019 chnykn@gmail.com All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package response

//SpecialtyFloor ***
type SpecialtyFloor struct {
	Floor      string
	Categories []Category
}

//NewSpecialtyFloor ***
func NewSpecialtyFloor(floor string) *SpecialtyFloor {
	o := &SpecialtyFloor{
		Floor:      floor,
		Categories: make([]Category, 0),
	}
	return o
}

//---------------------------------------------------------------------

//Specialty ***
type Specialty struct {
	Specialty string
	Floors    []SpecialtyFloor
}

//NewSpecialty ***
func NewSpecialty(specialty string) *Specialty {
	o := &Specialty{
		Specialty: specialty,
		Floors:    make([]SpecialtyFloor, 0),
	}
	return o
}

//---------------------------------------------------------------------

//SpecialtyTree ***
type SpecialtyTree struct {
	TreeType int64
	Tree     []Specialty
}

//NewSpecialtyTree ***
func NewSpecialtyTree(treeType int64) *SpecialtyTree {
	o := &SpecialtyTree{
		TreeType: treeType,
		Tree:     make([]Specialty, 0),
	}
	return o
}
