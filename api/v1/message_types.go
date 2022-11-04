/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
  ここで定義したstructに基づきcontroller-genがCRDを作成する
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MessageSpec defines the desired state of Message
type MessageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Message. Edit message_types.go to remove/update

	Word string `json:"word,omitempty"`

	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:default=1

	Number *int32 `json:"number,omitempty"` // 数値はポインタで指定することで値が指定されなかった場合はnullにできる
}

// MessageStatus defines the observed state of Message
type MessageStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Word   string `json:"word,omitempty"`
	Number int32  `json:"number,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:resource:shortName="mg"
// +kubebuilder:printcolumn:JSONPath=".status.word",name=Word,type=string
// +kubebuilder:printcolumn:JSONPath=".status.number",name=Number,type=integer

// Message is the Schema for the messages API
// MessageリソースのSchema定義に該当
type Message struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MessageSpec   `json:"spec,omitempty"`
	Status MessageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MessageList contains a list of Message
type MessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Message `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Message{}, &MessageList{})
}
