package main

import (
	"os"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
)


func Run(rl *fn.ResourceList) (bool, error) {

	for _, kubeObject := range rl.Items {
		
		hostGenerated := kubeObject.GetAnnotation("openshift.io/host.generated")
		
		if hostGenerated == "true" {

			kubeObject.SetNestedField("", "spec", "host")
			var out map[string]interface{}
			err := kubeObject.As(&out)
			if err != nil {
				//p.Log.Error("[route-restore] Error occured while conversion", err)
		}
	}
	rl.Results = append(rl.Results, fn.GeneralResult("[route-restore] process completed", fn.Info))
}
	return true, nil
}


func main() {
	// `AsMain` accepts a `ResourceListProcessor` interface.
	if err := fn.AsMain(fn.ResourceListProcessorFunc(Run)); err != nil {
		os.Exit(1)
	}
}
