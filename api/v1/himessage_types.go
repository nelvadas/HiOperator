/*

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HiMessageSpec defines the desired state of HiMessage

type HiMessageSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Message is a Tex to be displayed by HiMessages Pods
	// +kubebuilder:validation:MaxLength:=160
	Message string `json:"message,omitempty"`

	// Docker image to be runed by the HiMessage Pods
	Image string `json:"image,omitempty"`
}

// HiMessageStatus defines the observed state of HiMessage
type HiMessageStatus struct {

	// Printed=True if the message is already printed false otherwise
	Printed bool `json:"printed"`
	// PrintedDate: Time elapsed since the message was printed
	PrintedDate string `json:"printeddate"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=messaging,path=himessages,singular=himessage,shortName=hi;him;himesg
// +kubebuilder:printcolumn:name="Image",type="string",JSONPath=".spec.image",description="Image to Run"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".spec.message", format=password,description="Message to display"
// +kubebuilder:printcolumn:name="Printed",type="boolean",JSONPath=".status.printed",description="Printed Status"
// +kubebuilder:printcolumn:name="PrintedDate",type="date",JSONPath=".status.printeddate",description="Printed Date"

// HiMessage is the Schema for the himessages API
type HiMessage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HiMessageSpec   `json:"spec,omitempty"`
	Status HiMessageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HiMessageList contains a list of HiMessage
type HiMessageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HiMessage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HiMessage{}, &HiMessageList{})
}
