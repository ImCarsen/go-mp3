// Copyright 2017 Hajime Hoshi
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

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/imcarsen/go-mp3"
)

func run() error {
	f, err := os.Open("classic.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c := CreateContext(d.SampleRate())

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	fmt.Printf("Length: %d[bytes]\n", d.Length())
	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func CreateContext(sampleRate int) *oto.Context {
	op := &oto.NewContextOptions{
		SampleRate:   sampleRate,              // Match the given sample rate
		ChannelCount: 2,                       // Stereo
		Format:       oto.FormatSignedInt16LE, // Format of the source. go-mp3's format is signed 16bit integers.
	}

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}

	<-readyChan
	return otoCtx
}
