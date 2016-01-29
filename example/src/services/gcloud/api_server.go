package gcloud

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/jmcvetta/randutil"

	"github.com/sr/operator"
	"golang.org/x/net/context"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/container/v1"
	"google.golang.org/api/logging/v1beta3"
)

const (
	clusterAdminUsername = "admin"
	loggingService       = "logging.googleapis.com"
	machineType          = "n1-standard-1"
	userAccountScope     = "https://www.googleapis.com/auth/cloud.useraccounts"
	userInfoEmailScope   = "https://www.googleapis.com/auth/userinfo.email"
)

var oauthScopes = []string{
	compute.CloudPlatformScope,
	compute.ComputeScope,
	compute.DevstorageReadWriteScope,
	logging.LoggingAdminScope,
	userAccountScope,
	userInfoEmailScope,
}

type apiServer struct {
	client           *http.Client
	computeService   *compute.Service
	containerService *container.Service
}

func newAPIServer(
	client *http.Client,
	computeService *compute.Service,
	containerService *container.Service,
) *apiServer {
	return &apiServer{
		client,
		computeService,
		containerService,
	}
}

func (s *apiServer) CreateContainerCluster(
	ctx context.Context,
	request *CreateContainerClusterRequest,
) (*CreateContainerClusterResponse, error) {
	nodeCount, err := strconv.ParseInt(request.NodeCount, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid node count value: %v", request.NodeCount)
	}
	password, err := randutil.String(10, randutil.Alphanumeric)
	if err != nil {
		return nil, errors.New("failed to generated password")
	}
	operation, err := s.containerService.Projects.Zones.Clusters.Create(
		request.ProjectId,
		request.Zone,
		&container.CreateClusterRequest{
			Cluster: &container.Cluster{
				Name:             request.Name,
				InitialNodeCount: nodeCount,
				LoggingService:   loggingService,
				MasterAuth: &container.MasterAuth{
					Username: clusterAdminUsername,
					Password: password,
				},
				NodeConfig: &container.NodeConfig{
					MachineType: machineType,
					OauthScopes: oauthScopes,
				},
			},
		},
	).Do()
	if err != nil {
		return nil, err
	}
	return &CreateContainerClusterResponse{
		Output: &operator.Output{PlainText: operation.SelfLink},
	}, nil
}

func (s *apiServer) ListInstances(
	ctx context.Context,
	request *ListInstancesRequest,
) (*ListInstancesResponse, error) {
	if request.ProjectId == "" {
		return nil, operator.NewArgumentRequiredError("ProjectId")
	}

	response, err := s.computeService.Instances.AggregatedList(request.ProjectId).Do()
	if err != nil {
		return nil, err
	}

	var instances []*Instance
	for _, item := range response.Items {
		for _, instance := range item.Instances {
			zoneParts := strings.Split(instance.Zone, "/")
			instances = append(instances, &Instance{
				Id:     string(instance.Id),
				Name:   instance.Name,
				Status: instance.Status,
				Zone:   zoneParts[len(zoneParts)-1],
			})
		}
	}

	output := bytes.NewBufferString("")
	w := new(tabwriter.Writer)
	w.Init(output, 0, 8, 0, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "NAME", "STATUS", "ZONE")
	for _, instance := range instances {
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\n",
			instance.Name,
			instance.Status,
			instance.Zone,
		)
	}
	w.Flush()

	return &ListInstancesResponse{
		Objects: instances,
		Output:  &operator.Output{PlainText: output.String()},
	}, nil
}
