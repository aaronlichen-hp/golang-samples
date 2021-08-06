// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START functions_helloworld_pubsub_v2]

// Package helloworld provides a set of Cloud Functions samples.
package helloworld

import (
	"log"
)

// A CloudEvent containing the Pub/Sub message.
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type CloudEventMessage struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a CloudEvent message and extracts the Pub/Sub message.
func HelloPubSub(cem CloudEventMessage) error {
	name := string(cem.Message.Data) // Automatically decoded from base64.
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}

// [END functions_helloworld_pubsub_v2]
