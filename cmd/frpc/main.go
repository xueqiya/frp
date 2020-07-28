// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package frpclib

import (
	"math/rand"
	"os"
	"time"

	_ "github.com/fatedier/frp/assets/frpc/statik"
	"github.com/fatedier/frp/cmd/frpc/sub"

	"github.com/fatedier/golib/crypto"
)

type JniMember struct {
	callback Callback
}

var jniMember JniMember

//func main() {
//	crypto.DefaultSalt = "frp"
//	rand.Seed(time.Now().UnixNano())
//	jniMember.callback = new(JniMember)
//	sub.Execute(&jniMember)
//}
//
//func (ac *JniMember) CallByGo() {
//	log.Println("````````````")
//}

type Callback interface {
	CallByGo()
}

func (ac *JniMember) SendCallByGo() {
	ac.callback.CallByGo()
}

func Run(cfgFilePath string, c Callback) {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())

	jniMember.callback = c
	sub.RunClient(cfgFilePath, &jniMember)
}

func Stop() {
	os.Exit(1)
}
