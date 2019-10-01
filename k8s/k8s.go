package k8s

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// APIWrapper is used to hold the kubernetesClient Set
// and IP address nd Node details for the node the
// pod is currently scheduled on.
type APIWrapper struct {
	Client   *kubernetes.Clientset
	IP       string
	NodeName string
}

// NewAPIWrapper will initialise the wrapper struct
// need to perform api calls to the K8S api
func NewAPIWrapper() (api *APIWrapper, err error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &APIWrapper{
		Client: clientset,
	}, nil

}

// findNodeIP will use the APIWrapper to query the
// K8S api for POD spec to identify the node IP
func (a *APIWrapper) findNodeIP() (err error) {
	hostname, err := os.Hostname()
	if err != nil {
		logrus.Error(err)
		return err
	}

	// Fetched from K8S spec using the Downward API
	nameSpace, ok := os.LookupEnv("NAMESPACE")

	logrus.Debug("POD running in namespace: " + nameSpace)
	if !ok {
		logrus.Error(err)
		return err
	}

	p, err := a.Client.CoreV1().Pods(nameSpace).Get(hostname, v1.GetOptions{})

	if err != nil {
		logrus.Error(err)
		return err
	}

	// Lets query the PodStatus to identify the IP
	a.IP = p.Status.HostIP

	logrus.Debug("Node IP returned is " + a.IP)
	return nil
}

// FindNodeName will use the APIWrapper IP to query
// K8S api for all nodes and get the node name back
func (a *APIWrapper) findNodeName() (err error) {
	nodeList, err := a.Client.CoreV1().Nodes().List(v1.ListOptions{})
	if err != nil {
		logrus.Error(err)
		return err
	}

	var nodeMap = make(map[string]string)

	// populate the wrapper to identify the Hostname //
	for _, n := range nodeList.Items {
		hostname := n.ObjectMeta.GetName()
		for _, address := range n.Status.Addresses {
			if address.Type == corev1.NodeExternalIP || address.Type == corev1.NodeInternalIP {
				nodeMap[address.Address] = hostname
			}
		}
	}

	if nodeName, ok := nodeMap[a.IP]; !ok {
		logrus.Error("Unable to find a NodeName which matches the IP found from Node Lookup")
		return errors.New("Unable to find a NodeName in findNodeName")
	} else {
		a.NodeName = nodeName
	}

	return nil

}

// NodeFinder will use wrap the calls to the Node functions to return the Nodename the
// pod is running on.
func NodeFinder() (hostname string, err error) {
	a, err := NewAPIWrapper()
	if err != nil {
		logrus.Error("Unable to create a new APIWrapper")
		return "", err
	}

	err = a.findNodeIP()
	if err != nil {
		logrus.Error("Unable to find the IP Pod is running on")
		return "", err
	}

	err = a.findNodeName()

	return a.NodeName, err
}
