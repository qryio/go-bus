// Copyright 2022 Thiago Souza <tcostasouza@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testHandler = "test"
const emptyHandler = ""

func TestNode_AddAndLookupStatic1(t *testing.T) {
	found := emptyHandler
	p := []string{ "test" }
	n := NewRoot()
	_ = n.Add(p, testHandler)
	n.Accept(p, func(h interface{}) { found = h.(string) })
	assert.Equal(t, testHandler, found)
}

func TestNode_AddAndLookupStatic2(t *testing.T) {
	found1 := emptyHandler
	found2 := emptyHandler
	p1 := []string{ "test1" }
	p2 := []string{ "test1", "test2" }
	n := NewRoot()
	_ = n.Add(p1, testHandler + "1")
	_ = n.Add(p2, testHandler + "2")
	n.Accept(p1, func(h interface{}) { found1 = h.(string) })
	n.Accept(p2, func(h interface{}) { found2 = h.(string) })
	assert.Equal(t, testHandler + "1", found1)
	assert.Equal(t, testHandler + "2", found2)
}

func TestNode_AddRemoveAndLookupStatic1(t *testing.T) {
	found := emptyHandler
	p := []string{ "test" }
	n := NewRoot()
	n.Add(p, testHandler).Remove()
	n.Accept(p, func(h interface{}) { found = h.(string) })
	assert.Equal(t, emptyHandler, found)
}

func TestNode_AddRemoveAndLookupStatic2(t *testing.T) {
	found1 := emptyHandler
	found2 := emptyHandler
	p1 := []string{ "test1" }
	p2 := []string{ "test1", "test2" }
	n := NewRoot()
	n.Add(p1, testHandler + "1").Remove()
	n.Add(p2, testHandler + "2")
	n.Accept(p1, func(h interface{}) { found1 = h.(string) })
	n.Accept(p2, func(h interface{}) { found2 = h.(string) })
	assert.Equal(t, emptyHandler, found1)
	assert.Equal(t, testHandler + "2", found2)
}

func TestNode_AddAndLookupWildcard1(t *testing.T) {
	found := emptyHandler
	pW := []string{ "*" }
	pS1 := []string{ "test1" }
	pS2 := []string{ "test1", "test2" }
	n := NewRoot()
	_ = n.Add(pW, testHandler)
	n.Accept(pS1, func(h interface{}) { found = h.(string) })
	found = emptyHandler
	n.Accept(pS2, func(h interface{}) { found = h.(string) })
	assert.Equal(t, emptyHandler, found)
}

func TestNode_AddAndLookupWildcard2(t *testing.T) {
	found := emptyHandler
	pW := []string{ "test1", "*" }
	pS := []string{ "test1", "test2" }
	n := NewRoot()
	_ = n.Add(pW, testHandler)
	n.Accept(pS, func(h interface{}) { found = h.(string) })
	assert.Equal(t, testHandler, found)
}

func TestNode_AddAndLookupWildcard3(t *testing.T) {
	found := emptyHandler
	pW := []string{ "test1", "*", "test3" }
	pS1 := []string{ "test1", "test2", "test3" }
	pS2 := []string{ "test2", "test1", "test3" }
	n := NewRoot()
	_ = n.Add(pW, testHandler)
	n.Accept(pS1, func(h interface{}) { found = h.(string) })
	found = emptyHandler
	n.Accept(pS2, func(h interface{}) { found = h.(string) })
	assert.Equal(t, emptyHandler, found)
}