package install

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"gopkg.in/yaml.v2"
)

func ExecuteTestCommand(args []string) ([]byte, error) {
	cmd := NewCommand()
	cmd.SetArgs(args)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	err := cmd.Execute()
	if err != nil {
		return []byte{}, err
	}
	return ioutil.ReadAll(b)
}

func ExecuteTemplateYamlGenerationCommand(args []string) (map[string]interface{}, error) {
	out, err := ExecuteTestCommand(args)
	if err != nil {
		return nil, err
	}

	var template map[string]interface{}
	if err := yaml.Unmarshal(out, &template); err != nil {
		return nil, err
	}

	return template, nil
}

func VerifyTemplateParameterPresent(template map[string]interface{}, paramName string) bool {
	params := template["parameters"].([]interface{})
	for _, p := range params {
		if name, namePresent := p.(map[interface{}]interface{})["name"]; namePresent {
			if name == paramName {
				return true
			}
		}
	}
	return false
}

func TestMultiDocYamlRendering(t *testing.T) {
	out, err := ExecuteTestCommand([]string{"render", "--format", "yaml"})
	if err != nil {
		t.Fatal(err)
	}

	dec := yaml.NewDecoder(bytes.NewReader(out))
	var manifest map[string]interface{}
	cnt := 0
	for dec.Decode(&manifest) == nil {
		cnt += 1
	}
	if cnt < 2 {
		t.Fatal("no manifests found")
	}
}

func TestTemplateYamlRendering(t *testing.T) {
	template, err := ExecuteTemplateYamlGenerationCommand([]string{"render", "--format", "yaml", "--template"})
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := template["objects"]; !ok {
		t.Fatal("objects missing in template")
	}
	objects := template["objects"].([]interface{})
	if len(objects) == 0 {
		t.Fatal("no objects found in template")
	}
	for _, param := range []string{"OPERATOR_REPLICAS", "OPERATOR_IMG", "IMAGE_TAG", "NAMESPACE"} {
		if !VerifyTemplateParameterPresent(template, param) {
			t.Fatal("expected parameter", param, "not found")
		}
	}
}

func TestTemplateYamlRenderingWithOIDC(t *testing.T) {
	template, err := ExecuteTemplateYamlGenerationCommand([]string{"--oidc-storage-provider-s3-bucket-name", "bucket", "render", "--format", "yaml", "--template"})
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := template["objects"]; !ok {
		t.Fatal("objects missing in template")
	}
	objects := template["objects"].([]interface{})
	if len(objects) == 0 {
		t.Fatal("no objects found in template")
	}
	params := []string{
		"OPERATOR_REPLICAS", "OPERATOR_IMG", "IMAGE_TAG", "NAMESPACE",
		"OIDC_S3_BUCKET", "OIDC_S3_REGION", "OIDC_S3_ACCESS_KEY_ID",
		"OIDC_S3_SECRET_ACCESS_KEY",
	}
	for _, param := range params {
		if !VerifyTemplateParameterPresent(template, param) {
			t.Fatal("expected parameter", param, "not found")
		}
	}
}

func ExecuteJsonGenerationCommand(args []string) (map[string]interface{}, error) {
	out, err := ExecuteTestCommand(args)
	if err != nil {
		return nil, err
	}

	var doc map[string]interface{}
	json.Unmarshal(out, &doc)

	return doc, nil
}

func TestJsonListRendering(t *testing.T) {
	doc, err := ExecuteJsonGenerationCommand([]string{"render", "--format", "json"})
	if err != nil {
		t.Fatal(err)
	}

	if doc["kind"] != "List" {
		t.Fatal("expected kind List")
	}
	items := doc["items"].([]interface{})
	if len(items) == 0 {
		t.Fatal("no objects in items of json List")
	}
}

func TestJsonTemplateRendering(t *testing.T) {
	doc, err := ExecuteJsonGenerationCommand([]string{"render", "--format", "json", "--template"})
	if err != nil {
		t.Fatal(err)
	}

	if doc["kind"] != "Template" {
		t.Fatal("expected kind Template")
	}
	objects := doc["objects"].([]interface{})
	if len(objects) == 0 {
		t.Fatal("no objects in template objects")
	}
}
