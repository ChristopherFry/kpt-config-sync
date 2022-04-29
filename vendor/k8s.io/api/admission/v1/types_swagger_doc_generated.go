/*
Copyright The Kubernetes Authors.

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

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO:... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_AdmissionRequest = map[string]string{
	"":                   "AdmissionRequest describes the admission.Attributes for the admission request.",
	"uid":                "UID is an identifier for the individual request/response. It allows us to distinguish instances of requests which are otherwise identical (parallel requests, requests when earlier requests did not modify etc) The UID is meant to track the round trip (request/response) between the KAS and the WebHook, not the user request. It is suitable for correlating log entries between the webhook and apiserver, for either auditing or debugging.",
	"kind":               "Kind is the fully-qualified type of object being submitted (for example, v1.Pod or autoscaling.v1.Scale)",
	"resource":           "Resource is the fully-qualified resource being requested (for example, v1.pods)",
	"subResource":        "SubResource is the subresource being requested, if any (for example, \"status\" or \"scale\")",
	"requestKind":        "RequestKind is the fully-qualified type of the original API request (for example, v1.Pod or autoscaling.v1.Scale). If this is specified and differs from the value in \"kind\", an equivalent match and conversion was performed.\n\nFor example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of `apiGroups:[\"apps\"], apiVersions:[\"v1\"], resources: [\"deployments\"]` and `matchPolicy: Equivalent`, an API request to apps/v1beta1 deployments would be converted and sent to the webhook with `kind: {group:\"apps\", version:\"v1\", kind:\"Deployment\"}` (matching the rule the webhook registered for), and `requestKind: {group:\"apps\", version:\"v1beta1\", kind:\"Deployment\"}` (indicating the kind of the original API request).\n\nSee documentation for the \"matchPolicy\" field in the webhook configuration type for more details.",
	"requestResource":    "RequestResource is the fully-qualified resource of the original API request (for example, v1.pods). If this is specified and differs from the value in \"resource\", an equivalent match and conversion was performed.\n\nFor example, if deployments can be modified via apps/v1 and apps/v1beta1, and a webhook registered a rule of `apiGroups:[\"apps\"], apiVersions:[\"v1\"], resources: [\"deployments\"]` and `matchPolicy: Equivalent`, an API request to apps/v1beta1 deployments would be converted and sent to the webhook with `resource: {group:\"apps\", version:\"v1\", resource:\"deployments\"}` (matching the resource the webhook registered for), and `requestResource: {group:\"apps\", version:\"v1beta1\", resource:\"deployments\"}` (indicating the resource of the original API request).\n\nSee documentation for the \"matchPolicy\" field in the webhook configuration type.",
	"requestSubResource": "RequestSubResource is the name of the subresource of the original API request, if any (for example, \"status\" or \"scale\") If this is specified and differs from the value in \"subResource\", an equivalent match and conversion was performed. See documentation for the \"matchPolicy\" field in the webhook configuration type.",
	"name":               "Name is the name of the object as presented in the request.  On a CREATE operation, the client may omit name and rely on the server to generate the name.  If that is the case, this field will contain an empty string.",
	"namespace":          "Namespace is the namespace associated with the request (if any).",
	"operation":          "Operation is the operation being performed. This may be different than the operation requested. e.g. a patch can result in either a CREATE or UPDATE Operation.",
	"userInfo":           "UserInfo is information about the requesting user",
	"object":             "Object is the object from the incoming request.",
	"oldObject":          "OldObject is the existing object. Only populated for DELETE and UPDATE requests.",
	"dryRun":             "DryRun indicates that modifications will definitely not be persisted for this request. Defaults to false.",
	"options":            "Options is the operation option structure of the operation being performed. e.g. `meta.k8s.io/v1.DeleteOptions` or `meta.k8s.io/v1.CreateOptions`. This may be different than the options the caller provided. e.g. for a patch request the performed Operation might be a CREATE, in which case the Options will a `meta.k8s.io/v1.CreateOptions` even though the caller provided `meta.k8s.io/v1.PatchOptions`.",
}

func (AdmissionRequest) SwaggerDoc() map[string]string {
	return map_AdmissionRequest
}

var map_AdmissionResponse = map[string]string{
	"":                 "AdmissionResponse describes an admission response.",
	"uid":              "UID is an identifier for the individual request/response. This must be copied over from the corresponding AdmissionRequest.",
	"allowed":          "Allowed indicates whether or not the admission request was permitted.",
	"status":           "Result contains extra details into why an admission request was denied. This field IS NOT consulted in any way if \"Allowed\" is \"true\".",
	"patch":            "The patch body. Currently we only support \"JSONPatch\" which implements RFC 6902.",
	"patchType":        "The type of Patch. Currently we only allow \"JSONPatch\".",
	"auditAnnotations": "AuditAnnotations is an unstructured key value map set by remote admission controller (e.g. error=image-blacklisted). MutatingAdmissionWebhook and ValidatingAdmissionWebhook admission controller will prefix the keys with admission webhook name (e.g. imagepolicy.example.com/error=image-blacklisted). AuditAnnotations will be provided by the admission webhook to add additional context to the audit log for this request.",
	"warnings":         "warnings is a list of warning messages to return to the requesting API client. Warning messages describe a problem the client making the API request should correct or be aware of. Limit warnings to 120 characters if possible. Warnings over 256 characters and large numbers of warnings may be truncated.",
}

func (AdmissionResponse) SwaggerDoc() map[string]string {
	return map_AdmissionResponse
}

var map_AdmissionReview = map[string]string{
	"":         "AdmissionReview describes an admission review request/response.",
	"request":  "Request describes the attributes for the admission request.",
	"response": "Response describes the attributes for the admission response.",
}

func (AdmissionReview) SwaggerDoc() map[string]string {
	return map_AdmissionReview
}

// AUTO-GENERATED FUNCTIONS END HERE
