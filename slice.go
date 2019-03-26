// Copyright 2013 com authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com

import (
	"math/rand"
	"strings"
	"time"
)

// AppendStr appends string to slice with no duplicates.
func AppendStr(strs []string, str string) []string {
	for _, s := range strs {
		if s == str {
			return strs
		}
	}
	return append(strs, str)
}

// CompareSliceStr compares two 'string' type slices.
// It returns true if elements and order are both the same.
func CompareSliceStr(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// CompareSliceStrU compares two 'string' type slices.
// It returns true if elements are the same, and ignores the order.
func CompareSliceStrU(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[i] == s2[j] {
				s2 = append(s2[:j], s2[j+1:]...)
				break
			}
		}
	}
	if len(s2) > 0 {
		return false
	}
	return true
}

// IsSliceContainsStr returns true if the string exists in given slice.
func IsSliceContainsStr(sl []string, str string) bool {
	str = strings.ToLower(str)
	for _, s := range sl {
		if strings.ToLower(s) == str {
			return true
		}
	}
	return false
}

// IsSliceContainsInt64 returns true if the int64 exists in given slice.
func IsSliceContainsInt64(sl []int64, i int64) bool {
	for _, s := range sl {
		if s == i {
			return true
		}
	}
	return false
}

// ==============================
type reducetype func(interface{}) interface{}
type filtertype func(interface{}) bool

func InSlice(v string, sl []string) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InSliceIface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InStringSlice(v string, sl []string) bool {
	return InSlice(v, sl)
}

func InInterfaceSlice(v interface{}, sl []interface{}) bool {
	return InSliceIface(v, sl)
}

func InIntSlice(v int, sl []int) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InInt32Slice(v int32, sl []int32) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InInt16Slice(v int16, sl []int16) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InInt64Slice(v int64, sl []int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InUintSlice(v uint, sl []uint) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InUint32Slice(v uint32, sl []uint32) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InUint16Slice(v uint16, sl []uint16) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func InUint64Slice(v uint64, sl []uint64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func SliceRandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t0 := time.Now()
	rand.Seed(int64(t0.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

func SliceReduce(slice []interface{}, a reducetype) (dslice []interface{}) {
	for _, v := range slice {
		dslice = append(dslice, a(v))
	}
	return
}

func SliceRand(a []interface{}) (b interface{}) {
	randnum := rand.Intn(len(a))
	b = a[randnum]
	return
}

func SliceSum(intslice []int64) (sum int64) {
	for _, v := range intslice {
		sum += v
	}
	return
}

func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	for _, v := range slice {
		if a(v) {
			ftslice = append(ftslice, v)
		}
	}
	return
}

func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	for _, v := range slice1 {
		if !InSliceIface(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func SliceChunk(slice []interface{}, size int) (chunkslice [][]interface{}) {
	if size >= len(slice) {
		chunkslice = append(chunkslice, slice)
		return
	}
	end := size
	for i := 0; i <= (len(slice) - size); i += size {
		chunkslice = append(chunkslice, slice[i:end])
		end += size
	}
	return
}

func SliceRange(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !InSliceIface(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}

func SliceShuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}

func SliceInsert(slice, insertion []interface{}, index int) []interface{} {
	result := make([]interface{}, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

//SliceRemove SliceRomove(a,4,5) //a[4]
func SliceRemove(slice []interface{}, start int, args ...int) []interface{} {
	var end int
	if len(args) == 0 {
		end = start + 1
	} else {
		end = args[0]
	}
	if end > len(slice)-1 {
		return slice[:start]
	}
	return append(slice[:start], slice[end:]...)
}

// SliceRemoveCallback : 根据条件删除
// a=[]int{1,2,3,4,5,6}
// SliceRemoveCallback(len(a), func(i int) func(bool)error{
//	if a[i]!=4 {
//	 	return nil
// 	}
// 	return func(inside bool)error{
//		if inside {
//			a=append(a[0:i],a[i+1:]...)
// 		}else{
//			a=a[0:i]
//		}
//		return nil
// 	}
//})
func SliceRemoveCallback(length int, callback func(int) func(bool) error) error {
	for i, j := 0, length-1; i <= j; i++ {
		if removeFunc := callback(i); removeFunc != nil {
			var err error
			if i+1 <= j {
				err = removeFunc(true)
			} else {
				err = removeFunc(false)
			}
			if err != nil {
				return err
			}
			i--
			j--
		}
	}
	return nil
}
