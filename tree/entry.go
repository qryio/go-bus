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

type Entry struct {
	key		uintptr
	owner   *Node
	handler interface{}
}

func (e *Entry) Remove() {
	if e.owner == nil {
		return
	}
	delete(e.owner.entries, e.key)
	e.owner.prune()
	e.key = 0
	e.owner = nil
	e.handler = nil
}

