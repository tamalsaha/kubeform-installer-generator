package main

import (
	"fmt"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

func main() {
	filename := "/home/tamal/go/src/kubeform.dev/installer/crds/kubeform-provider-linode/object-crds.yaml"

	var obj map[string]interface{}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		panic(err)
	}

	removeDescription(obj)

	data, err = yaml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func removeDescription(m map[string]interface{}) {
	for k, v := range m {
		if k == "description" {
			if _, ok := v.(string); ok {
				delete(m, k)
			}
		}
		if inner, ok := v.(map[string]interface{}); ok {
			removeDescription(inner)
		} else if arr, ok := v.([]interface{}); ok {
			for i := range arr {
				if inner, ok := arr[i].(map[string]interface{}); ok {
					removeDescription(inner)
				}
			}
		}
	}
}
