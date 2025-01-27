package kubetest

import (
	"fmt"

	"github.com/kiali/kiali/kubernetes"
	osapps_v1 "github.com/openshift/api/apps/v1"
	osproject_v1 "github.com/openshift/api/project/v1"
	osroutes_v1 "github.com/openshift/api/route/v1"
	"github.com/stretchr/testify/mock"
	apps_v1 "k8s.io/api/apps/v1"
	auth_v1 "k8s.io/api/authorization/v1"
	batch_v1 "k8s.io/api/batch/v1"
	batch_apps_v1 "k8s.io/api/batch/v1beta1"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
)

//// Mock for the K8SClientFactory

type K8SClientFactoryMock struct {
	mock.Mock
	k8s kubernetes.IstioClientInterface
}

// Constructor
func NewK8SClientFactoryMock(k8s kubernetes.IstioClientInterface) *K8SClientFactoryMock {
	k8sClientFactory := new(K8SClientFactoryMock)
	k8sClientFactory.k8s = k8s
	return k8sClientFactory
}

// Business Methods
func (o *K8SClientFactoryMock) GetClient(token string) (kubernetes.IstioClientInterface, error) {
	return o.k8s, nil
}

/////

type K8SClientMock struct {
	mock.Mock
}

// Constructor

func NewK8SClientMock() *K8SClientMock {
	k8s := new(K8SClientMock)
	k8s.On("IsOpenShift").Return(true)
	return k8s
}

// Business methods

// MockEmptyWorkloads setup the current mock to return empty workloads for every type of workloads (deployment, dc, rs, jobs, etc.)
func (o *K8SClientMock) MockEmptyWorkloads(namespace interface{}) {
	o.On("GetDeployments", namespace).Return([]apps_v1.Deployment{}, nil)
	o.On("GetReplicaSets", namespace).Return([]apps_v1.ReplicaSet{}, nil)
	o.On("GetReplicationControllers", namespace).Return([]core_v1.ReplicationController{}, nil)
	o.On("GetDeploymentConfigs", namespace).Return([]osapps_v1.DeploymentConfig{}, nil)
	o.On("GetStatefulSets", namespace).Return([]apps_v1.StatefulSet{}, nil)
	o.On("GetJobs", namespace).Return([]batch_v1.Job{}, nil)
	o.On("GetCronJobs", namespace).Return([]batch_apps_v1.CronJob{}, nil)
}

// MockEmptyWorkload setup the current mock to return an empty workload for every type of workloads (deployment, dc, rs, jobs, etc.)
func (o *K8SClientMock) MockEmptyWorkload(namespace interface{}, workload interface{}) {
	notfound := fmt.Errorf("not found")
	o.On("GetDeployment", namespace, workload).Return(&apps_v1.Deployment{}, notfound)
	o.On("GetStatefulSet", namespace, workload).Return(&apps_v1.StatefulSet{}, notfound)
	o.On("GetDeploymentConfig", namespace, workload).Return(&osapps_v1.DeploymentConfig{}, notfound)
	o.On("GetReplicaSets", namespace).Return([]apps_v1.ReplicaSet{}, nil)
	o.On("GetReplicationControllers", namespace).Return([]core_v1.ReplicationController{}, nil)
	o.On("GetJobs", namespace).Return([]batch_v1.Job{}, nil)
	o.On("GetCronJobs", namespace).Return([]batch_apps_v1.CronJob{}, nil)
}

func (o *K8SClientMock) CreateIstioObject(api, namespace, resourceType, json string) (kubernetes.IstioObject, error) {
	args := o.Called(api, namespace, resourceType, json)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) DeleteIstioObject(api, namespace, objectType, objectName string) error {
	args := o.Called(api, namespace, objectType, objectName)
	return args.Error(0)
}

func (o *K8SClientMock) GetAdapter(namespace, adapterType, adapterName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, adapterType, adapterName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetAdapters(namespace, labelSelector string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetCronJobs(namespace string) ([]batch_apps_v1.CronJob, error) {
	args := o.Called(namespace)
	return args.Get(0).([]batch_apps_v1.CronJob), args.Error(1)
}

func (o *K8SClientMock) GetDeployment(namespace string, deploymentName string) (*apps_v1.Deployment, error) {
	args := o.Called(namespace, deploymentName)
	return args.Get(0).(*apps_v1.Deployment), args.Error(1)
}

func (o *K8SClientMock) GetDeployments(namespace string) ([]apps_v1.Deployment, error) {
	args := o.Called(namespace)
	return args.Get(0).([]apps_v1.Deployment), args.Error(1)
}

// It returns an error on any problem.
func (o *K8SClientMock) GetDeploymentsByLabel(namespace string, labelSelector string) ([]apps_v1.Deployment, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]apps_v1.Deployment), args.Error(1)
}

func (o *K8SClientMock) GetRoute(namespace, name string) (*osroutes_v1.Route, error) {
	args := o.Called(namespace, name)
	return args.Get(0).(*osroutes_v1.Route), args.Error(1)
}

func (o *K8SClientMock) GetSidecar(namespace string, sidecar string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetSidecars(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetDeploymentConfig(namespace string, deploymentName string) (*osapps_v1.DeploymentConfig, error) {
	args := o.Called(namespace, deploymentName)
	return args.Get(0).(*osapps_v1.DeploymentConfig), args.Error(1)
}

func (o *K8SClientMock) GetDeploymentConfigs(namespace string) ([]osapps_v1.DeploymentConfig, error) {
	args := o.Called(namespace)
	return args.Get(0).([]osapps_v1.DeploymentConfig), args.Error(1)
}

func (o *K8SClientMock) GetDestinationRules(namespace string, serviceName string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetDestinationRule(namespace string, destinationrule string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, destinationrule)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetEndpoints(namespace string, serviceName string) (*core_v1.Endpoints, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*core_v1.Endpoints), args.Error(1)
}

func (o *K8SClientMock) GetGateways(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetGateway(namespace string, gateway string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, gateway)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetIstioDetails(namespace string, serviceName string) (*kubernetes.IstioDetails, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*kubernetes.IstioDetails), args.Error(1)
}

func (o *K8SClientMock) GetIstioRule(namespace string, istiorule string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, istiorule)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetIstioRules(namespace string, labelSelector string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetJobs(namespace string) ([]batch_v1.Job, error) {
	args := o.Called(namespace)
	return args.Get(0).([]batch_v1.Job), args.Error(1)
}

func (o *K8SClientMock) GetNamespace(namespace string) (*core_v1.Namespace, error) {
	args := o.Called(namespace)
	return args.Get(0).(*core_v1.Namespace), args.Error(1)
}

func (o *K8SClientMock) GetNamespaces(labelSelector string) ([]core_v1.Namespace, error) {
	args := o.Called()
	return args.Get(0).([]core_v1.Namespace), args.Error(1)
}

func (o *K8SClientMock) GetPods(namespace, labelSelector string) ([]core_v1.Pod, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]core_v1.Pod), args.Error(1)
}

func (o *K8SClientMock) GetPod(namespace, name string) (*core_v1.Pod, error) {
	args := o.Called(namespace, name)
	return args.Get(0).(*core_v1.Pod), args.Error(1)
}

func (o *K8SClientMock) GetPodLogs(namespace, name string, opts *core_v1.PodLogOptions) (*kubernetes.PodLogs, error) {
	args := o.Called(namespace, name, opts)
	return args.Get(0).(*kubernetes.PodLogs), args.Error(1)
}

func (o *K8SClientMock) GetProject(project string) (*osproject_v1.Project, error) {
	args := o.Called(project)
	return args.Get(0).(*osproject_v1.Project), args.Error(1)
}

func (o *K8SClientMock) GetProjects(labelSelector string) ([]osproject_v1.Project, error) {
	args := o.Called()
	return args.Get(0).([]osproject_v1.Project), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpec(namespace string, quotaSpecName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, quotaSpecName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecs(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecBinding(namespace string, quotaSpecBindingName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, quotaSpecBindingName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetQuotaSpecBindings(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetReplicationControllers(namespace string) ([]core_v1.ReplicationController, error) {
	args := o.Called(namespace)
	return args.Get(0).([]core_v1.ReplicationController), args.Error(1)
}

func (o *K8SClientMock) GetReplicaSets(namespace string) ([]apps_v1.ReplicaSet, error) {
	args := o.Called(namespace)
	return args.Get(0).([]apps_v1.ReplicaSet), args.Error(1)
}

func (o *K8SClientMock) GetSelfSubjectAccessReview(namespace, api, resourceType string, verbs []string) ([]*auth_v1.SelfSubjectAccessReview, error) {
	args := o.Called(namespace, api, resourceType, verbs)
	return args.Get(0).([]*auth_v1.SelfSubjectAccessReview), args.Error(1)
}

func (o *K8SClientMock) GetService(namespace string, serviceName string) (*core_v1.Service, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).(*core_v1.Service), args.Error(1)
}

func (o *K8SClientMock) GetServices(namespace string, selectorLabels map[string]string) ([]core_v1.Service, error) {
	args := o.Called(namespace, selectorLabels)
	return args.Get(0).([]core_v1.Service), args.Error(1)
}

func (o *K8SClientMock) GetServiceEntries(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceEntry(namespace string, serviceEntryName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceEntryName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetStatefulSet(namespace string, statefulsetName string) (*apps_v1.StatefulSet, error) {
	args := o.Called(namespace, statefulsetName)
	return args.Get(0).(*apps_v1.StatefulSet), args.Error(1)
}

func (o *K8SClientMock) GetStatefulSets(namespace string) ([]apps_v1.StatefulSet, error) {
	args := o.Called(namespace)
	return args.Get(0).([]apps_v1.StatefulSet), args.Error(1)
}

func (o *K8SClientMock) GetTemplate(namespace, templateType, templateName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, templateType, templateName)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetTemplates(namespace, labelSelector string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, labelSelector)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetVirtualServices(namespace string, serviceName string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace, serviceName)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetVirtualService(namespace string, virtualservice string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace, virtualservice)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetPolicies(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetPolicy(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetMeshPolicies() ([]kubernetes.IstioObject, error) {
	args := o.Called()
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetMeshPolicy(policyName string) (kubernetes.IstioObject, error) {
	args := o.Called()
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceMeshPolicies(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceMeshPolicy(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetClusterRbacConfigs() ([]kubernetes.IstioObject, error) {
	args := o.Called()
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetClusterRbacConfig(policyName string) (kubernetes.IstioObject, error) {
	args := o.Called()
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceMeshRbacConfigs(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceMeshRbacConfig(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetRbacConfigs(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetRbacConfig(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceRoles(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceRole(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceRoleBindings(namespace string) ([]kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).([]kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetServiceRoleBinding(namespace string, policyName string) (kubernetes.IstioObject, error) {
	args := o.Called(namespace)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) GetAuthorizationDetails(namespace string) (*kubernetes.RBACDetails, error) {
	args := o.Called(namespace)
	return args.Get(0).(*kubernetes.RBACDetails), args.Error(1)
}

func (o *K8SClientMock) IsOpenShift() bool {
	args := o.Called()
	return args.Get(0).(bool)
}

func (o *K8SClientMock) IsMaistraApi() bool {
	args := o.Called()
	return args.Get(0).(bool)
}

func (o *K8SClientMock) GetServerVersion() (*version.Info, error) {
	args := o.Called()
	return args.Get(0).(*version.Info), args.Error(1)
}

func (o *K8SClientMock) Stop() {
}

func (o *K8SClientMock) UpdateIstioObject(api, namespace, resourceType, name, jsonPatch string) (kubernetes.IstioObject, error) {
	args := o.Called(api, namespace, resourceType, name, jsonPatch)
	return args.Get(0).(kubernetes.IstioObject), args.Error(1)
}

func (o *K8SClientMock) MockService(namespace, name string) {
	s := fakeService(namespace, name)
	o.On("GetService", namespace, name).Return(&s, nil)
}

func (o *K8SClientMock) MockServices(namespace string, names []string) {
	services := []core_v1.Service{}
	for _, name := range names {
		services = append(services, fakeService(namespace, name))
	}
	o.On("GetServices", namespace, mock.AnythingOfType("map[string]string")).Return(services, nil)
	o.On("GetDeployments", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return([]apps_v1.Deployment{}, nil)
}

func fakeService(namespace, name string) core_v1.Service {
	return core_v1.Service{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: core_v1.ServiceSpec{
			ClusterIP: "fromservice",
			Type:      "ClusterIP",
			Selector:  map[string]string{"app": name},
			Ports: []core_v1.ServicePort{
				{
					Name:     "http",
					Protocol: "TCP",
					Port:     3001,
				},
				{
					Name:     "http",
					Protocol: "TCP",
					Port:     3000,
				},
			},
		},
	}
}

func FakePodListWithoutSidecar() []core_v1.Pod {
	return []core_v1.Pod{
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "reviews-v1",
				Labels: map[string]string{"app": "reviews", "version": "v1"}}},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "reviews-v2",
				Labels: map[string]string{"app": "reviews", "version": "v2"}}},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:   "httpbin-v1",
				Labels: map[string]string{"app": "httpbin", "version": "v1"}}},
	}
}

func FakePodList() []core_v1.Pod {
	return []core_v1.Pod{
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:        "reviews-v1",
				Labels:      map[string]string{"app": "reviews", "version": "v1"},
				Annotations: FakeIstioAnnotations(),
			},
		},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:        "reviews-v2",
				Labels:      map[string]string{"app": "reviews", "version": "v2"},
				Annotations: FakeIstioAnnotations(),
			},
		},
		{
			ObjectMeta: meta_v1.ObjectMeta{
				Name:        "httpbin-v1",
				Labels:      map[string]string{"app": "httpbin", "version": "v1"},
				Annotations: FakeIstioAnnotations(),
			},
		},
	}
}

func FakeIstioAnnotations() map[string]string {
	return map[string]string{"sidecar.istio.io/status": "{\"version\":\"\",\"initContainers\":[\"istio-init\",\"enable-core-dump\"],\"containers\":[\"istio-proxy\"],\"volumes\":[\"istio-envoy\",\"istio-certs\"]}"}
}

func FakeNamespace(name string) *core_v1.Namespace {
	return &core_v1.Namespace{
		ObjectMeta: meta_v1.ObjectMeta{
			Name: name,
		},
	}
}
