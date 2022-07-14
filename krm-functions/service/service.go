package main

import (
	"os"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)


func Run(rl *fn.ResourceList) (bool, error) {

	for _, kubeObject := range rl.Items {


		// serviceName, _, err := kubeObject.NestedString("metadata", "name")
		//p.Log.Infof("[service-restore] service: %s", serviceName)

		serviceType, _, err := kubeObject.NestedString("spec", "type")

		// only clear ExternalIPs for LoadBalancer services
		if serviceType == "LoadBalancer" {
			// p.Log.Infof("[service-restore] Clearing externalIPs for LoadBalancer service: %s", serviceName)
			kubeObject.RemoveNestedField("spec", "externalIPs")
		}

		var out map[string]interface{}
		err = kubeObject.As(&out)

	}
	rl.Results = append(rl.Results, fn.GeneralResult("[service-restore] process completed", fn.Info))

	return true, nil
}


func main() {
	// `AsMain` accepts a `ResourceListProcessor` interface.
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
