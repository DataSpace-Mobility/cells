// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package rest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tree "github.com/pydio/cells/common/proto/tree"
import object "github.com/pydio/cells/common/proto/object"
import ctl "github.com/pydio/cells/common/proto/ctl"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Configuration message. Data is an Json representation of any value
type Configuration struct {
	// Full slash-separated path to the config key
	FullPath string `protobuf:"bytes,1,opt,name=FullPath" json:"FullPath,omitempty"`
	// JSON-encoded data to store
	Data string `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
}

func (m *Configuration) Reset()                    { *m = Configuration{} }
func (m *Configuration) String() string            { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()               {}
func (*Configuration) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Configuration) GetFullPath() string {
	if m != nil {
		return m.FullPath
	}
	return ""
}

func (m *Configuration) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ListDataSourceRequest struct {
}

func (m *ListDataSourceRequest) Reset()                    { *m = ListDataSourceRequest{} }
func (m *ListDataSourceRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDataSourceRequest) ProtoMessage()               {}
func (*ListDataSourceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

// Collection of datasources
type DataSourceCollection struct {
	DataSources []*object.DataSource `protobuf:"bytes,1,rep,name=DataSources" json:"DataSources,omitempty"`
	Total       int32                `protobuf:"varint,2,opt,name=Total" json:"Total,omitempty"`
}

func (m *DataSourceCollection) Reset()                    { *m = DataSourceCollection{} }
func (m *DataSourceCollection) String() string            { return proto.CompactTextString(m) }
func (*DataSourceCollection) ProtoMessage()               {}
func (*DataSourceCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *DataSourceCollection) GetDataSources() []*object.DataSource {
	if m != nil {
		return m.DataSources
	}
	return nil
}

func (m *DataSourceCollection) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type DeleteDataSourceResponse struct {
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
}

func (m *DeleteDataSourceResponse) Reset()                    { *m = DeleteDataSourceResponse{} }
func (m *DeleteDataSourceResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDataSourceResponse) ProtoMessage()               {}
func (*DeleteDataSourceResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *DeleteDataSourceResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListPeersAddressesRequest struct {
}

func (m *ListPeersAddressesRequest) Reset()                    { *m = ListPeersAddressesRequest{} }
func (m *ListPeersAddressesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPeersAddressesRequest) ProtoMessage()               {}
func (*ListPeersAddressesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type ListPeersAddressesResponse struct {
	// List of peer addresses
	PeerAddresses []string `protobuf:"bytes,1,rep,name=PeerAddresses" json:"PeerAddresses,omitempty"`
}

func (m *ListPeersAddressesResponse) Reset()                    { *m = ListPeersAddressesResponse{} }
func (m *ListPeersAddressesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPeersAddressesResponse) ProtoMessage()               {}
func (*ListPeersAddressesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ListPeersAddressesResponse) GetPeerAddresses() []string {
	if m != nil {
		return m.PeerAddresses
	}
	return nil
}

type ListPeerFoldersRequest struct {
	// Restrict listing to a given peer
	PeerAddress string `protobuf:"bytes,1,opt,name=PeerAddress" json:"PeerAddress,omitempty"`
	// Path to the parent folder for listing
	Path string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
}

func (m *ListPeerFoldersRequest) Reset()                    { *m = ListPeerFoldersRequest{} }
func (m *ListPeerFoldersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPeerFoldersRequest) ProtoMessage()               {}
func (*ListPeerFoldersRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *ListPeerFoldersRequest) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *ListPeerFoldersRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type CreatePeerFolderRequest struct {
	// Restrict listing to a given peer
	PeerAddress string `protobuf:"bytes,1,opt,name=PeerAddress" json:"PeerAddress,omitempty"`
	// Path to the folder to be created
	Path string `protobuf:"bytes,2,opt,name=Path" json:"Path,omitempty"`
}

func (m *CreatePeerFolderRequest) Reset()                    { *m = CreatePeerFolderRequest{} }
func (m *CreatePeerFolderRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePeerFolderRequest) ProtoMessage()               {}
func (*CreatePeerFolderRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *CreatePeerFolderRequest) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *CreatePeerFolderRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type CreatePeerFolderResponse struct {
	Success bool       `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	Node    *tree.Node `protobuf:"bytes,2,opt,name=Node" json:"Node,omitempty"`
}

func (m *CreatePeerFolderResponse) Reset()                    { *m = CreatePeerFolderResponse{} }
func (m *CreatePeerFolderResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePeerFolderResponse) ProtoMessage()               {}
func (*CreatePeerFolderResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *CreatePeerFolderResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreatePeerFolderResponse) GetNode() *tree.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

type ListStorageBucketsRequest struct {
	DataSource    *object.DataSource `protobuf:"bytes,1,opt,name=DataSource" json:"DataSource,omitempty"`
	BucketsRegexp string             `protobuf:"bytes,2,opt,name=BucketsRegexp" json:"BucketsRegexp,omitempty"`
}

func (m *ListStorageBucketsRequest) Reset()                    { *m = ListStorageBucketsRequest{} }
func (m *ListStorageBucketsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListStorageBucketsRequest) ProtoMessage()               {}
func (*ListStorageBucketsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *ListStorageBucketsRequest) GetDataSource() *object.DataSource {
	if m != nil {
		return m.DataSource
	}
	return nil
}

func (m *ListStorageBucketsRequest) GetBucketsRegexp() string {
	if m != nil {
		return m.BucketsRegexp
	}
	return ""
}

type Process struct {
	// Process ID
	ID string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	// Parent Process ID
	ParentID string `protobuf:"bytes,2,opt,name=ParentID" json:"ParentID,omitempty"`
	// Port to access the metrics api
	MetricsPort int32 `protobuf:"varint,3,opt,name=MetricsPort" json:"MetricsPort,omitempty"`
	// Id of peer node
	PeerId string `protobuf:"bytes,4,opt,name=PeerId" json:"PeerId,omitempty"`
	// Address of peer node
	PeerAddress string `protobuf:"bytes,5,opt,name=PeerAddress" json:"PeerAddress,omitempty"`
	// Parameters used to start this process
	StartTag string `protobuf:"bytes,6,opt,name=StartTag" json:"StartTag,omitempty"`
	// List of services running inside this process
	Services []string `protobuf:"bytes,7,rep,name=Services" json:"Services,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *Process) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Process) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *Process) GetMetricsPort() int32 {
	if m != nil {
		return m.MetricsPort
	}
	return 0
}

func (m *Process) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *Process) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *Process) GetStartTag() string {
	if m != nil {
		return m.StartTag
	}
	return ""
}

func (m *Process) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

type ListProcessesRequest struct {
	// Id of the peer node
	PeerId string `protobuf:"bytes,1,opt,name=PeerId" json:"PeerId,omitempty"`
	// Look for a service name
	ServiceName string `protobuf:"bytes,2,opt,name=ServiceName" json:"ServiceName,omitempty"`
}

func (m *ListProcessesRequest) Reset()                    { *m = ListProcessesRequest{} }
func (m *ListProcessesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListProcessesRequest) ProtoMessage()               {}
func (*ListProcessesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ListProcessesRequest) GetPeerId() string {
	if m != nil {
		return m.PeerId
	}
	return ""
}

func (m *ListProcessesRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type ListProcessesResponse struct {
	Processes []*Process `protobuf:"bytes,1,rep,name=Processes" json:"Processes,omitempty"`
}

func (m *ListProcessesResponse) Reset()                    { *m = ListProcessesResponse{} }
func (m *ListProcessesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListProcessesResponse) ProtoMessage()               {}
func (*ListProcessesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *ListProcessesResponse) GetProcesses() []*Process {
	if m != nil {
		return m.Processes
	}
	return nil
}

type ListVersioningPolicyRequest struct {
}

func (m *ListVersioningPolicyRequest) Reset()                    { *m = ListVersioningPolicyRequest{} }
func (m *ListVersioningPolicyRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVersioningPolicyRequest) ProtoMessage()               {}
func (*ListVersioningPolicyRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

type VersioningPolicyCollection struct {
	Policies []*tree.VersioningPolicy `protobuf:"bytes,1,rep,name=Policies" json:"Policies,omitempty"`
}

func (m *VersioningPolicyCollection) Reset()                    { *m = VersioningPolicyCollection{} }
func (m *VersioningPolicyCollection) String() string            { return proto.CompactTextString(m) }
func (*VersioningPolicyCollection) ProtoMessage()               {}
func (*VersioningPolicyCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *VersioningPolicyCollection) GetPolicies() []*tree.VersioningPolicy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type ListVirtualNodesRequest struct {
}

func (m *ListVirtualNodesRequest) Reset()                    { *m = ListVirtualNodesRequest{} }
func (m *ListVirtualNodesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVirtualNodesRequest) ProtoMessage()               {}
func (*ListVirtualNodesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

type ListServiceRequest struct {
	// Filter services by a given status (ANY, STOPPED, STOPPING, RUNNING)
	StatusFilter ctl.ServiceStatus `protobuf:"varint,1,opt,name=StatusFilter,enum=ctl.ServiceStatus" json:"StatusFilter,omitempty"`
}

func (m *ListServiceRequest) Reset()                    { *m = ListServiceRequest{} }
func (m *ListServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*ListServiceRequest) ProtoMessage()               {}
func (*ListServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{16} }

func (m *ListServiceRequest) GetStatusFilter() ctl.ServiceStatus {
	if m != nil {
		return m.StatusFilter
	}
	return ctl.ServiceStatus_ANY
}

type ServiceCollection struct {
	Services []*ctl.Service `protobuf:"bytes,1,rep,name=Services" json:"Services,omitempty"`
	Total    int32          `protobuf:"varint,2,opt,name=Total" json:"Total,omitempty"`
}

func (m *ServiceCollection) Reset()                    { *m = ServiceCollection{} }
func (m *ServiceCollection) String() string            { return proto.CompactTextString(m) }
func (*ServiceCollection) ProtoMessage()               {}
func (*ServiceCollection) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{17} }

func (m *ServiceCollection) GetServices() []*ctl.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ServiceCollection) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type ControlServiceRequest struct {
	// Name of the service to stop
	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName" json:"ServiceName,omitempty"`
	// Name of the node
	NodeName string `protobuf:"bytes,2,opt,name=NodeName" json:"NodeName,omitempty"`
	// Command to apply (START or STOP)
	Command ctl.ServiceCommand `protobuf:"varint,3,opt,name=Command,enum=ctl.ServiceCommand" json:"Command,omitempty"`
}

func (m *ControlServiceRequest) Reset()                    { *m = ControlServiceRequest{} }
func (m *ControlServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*ControlServiceRequest) ProtoMessage()               {}
func (*ControlServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{18} }

func (m *ControlServiceRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ControlServiceRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *ControlServiceRequest) GetCommand() ctl.ServiceCommand {
	if m != nil {
		return m.Command
	}
	return ctl.ServiceCommand_START
}

type DiscoveryRequest struct {
	// Filter result to a specific endpoint type
	EndpointType string `protobuf:"bytes,1,opt,name=EndpointType" json:"EndpointType,omitempty"`
}

func (m *DiscoveryRequest) Reset()                    { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()               {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{19} }

func (m *DiscoveryRequest) GetEndpointType() string {
	if m != nil {
		return m.EndpointType
	}
	return ""
}

type DiscoveryResponse struct {
	// Current Package Type, empty if user is not authenticated
	PackageType string `protobuf:"bytes,1,opt,name=PackageType" json:"PackageType,omitempty"`
	// Current Package Label, empty if user is not authenticated
	PackageLabel string `protobuf:"bytes,2,opt,name=PackageLabel" json:"PackageLabel,omitempty"`
	// Current Package Version, empty if user is not authenticated
	Version string `protobuf:"bytes,3,opt,name=Version" json:"Version,omitempty"`
	// Build stamp of the binary build, empty if user is not authenticated
	BuildStamp int32 `protobuf:"varint,4,opt,name=BuildStamp" json:"BuildStamp,omitempty"`
	// Revision of the current binary build, empty if user is not authenticated
	BuildRevision string `protobuf:"bytes,5,opt,name=BuildRevision" json:"BuildRevision,omitempty"`
	// List of endpoints and their corresponding URL access. Special case for grpc that just send back its port
	Endpoints map[string]string `protobuf:"bytes,6,rep,name=Endpoints" json:"Endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{20} }

func (m *DiscoveryResponse) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *DiscoveryResponse) GetPackageLabel() string {
	if m != nil {
		return m.PackageLabel
	}
	return ""
}

func (m *DiscoveryResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DiscoveryResponse) GetBuildStamp() int32 {
	if m != nil {
		return m.BuildStamp
	}
	return 0
}

func (m *DiscoveryResponse) GetBuildRevision() string {
	if m != nil {
		return m.BuildRevision
	}
	return ""
}

func (m *DiscoveryResponse) GetEndpoints() map[string]string {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ConfigFormRequest struct {
	// Retrieve a configuration form for a given service
	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName" json:"ServiceName,omitempty"`
}

func (m *ConfigFormRequest) Reset()                    { *m = ConfigFormRequest{} }
func (m *ConfigFormRequest) String() string            { return proto.CompactTextString(m) }
func (*ConfigFormRequest) ProtoMessage()               {}
func (*ConfigFormRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{21} }

func (m *ConfigFormRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

type OpenApiResponse struct {
}

func (m *OpenApiResponse) Reset()                    { *m = OpenApiResponse{} }
func (m *OpenApiResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenApiResponse) ProtoMessage()               {}
func (*OpenApiResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{22} }

type ActionDescription struct {
	// Unique name of the action
	Name string `protobuf:"bytes,1,opt,name=Name" json:"Name,omitempty"`
	// Mdi reference name for displaying icon
	Icon string `protobuf:"bytes,2,opt,name=Icon" json:"Icon,omitempty"`
	// Human-readable label
	Label string `protobuf:"bytes,3,opt,name=Label" json:"Label,omitempty"`
	// Long description and help text
	Description string `protobuf:"bytes,4,opt,name=Description" json:"Description,omitempty"`
	// Template for building a short summary of the action configuration
	SummaryTemplate string `protobuf:"bytes,5,opt,name=SummaryTemplate" json:"SummaryTemplate,omitempty"`
	// Whether this action has a form or not
	HasForm bool `protobuf:"varint,6,opt,name=HasForm" json:"HasForm,omitempty"`
	// JS module name to be used instead of loading standard form
	FormModule string `protobuf:"bytes,11,opt,name=FormModule" json:"FormModule,omitempty"`
	// JSON props used to init the FormModule (optional)
	FormModuleProps string `protobuf:"bytes,12,opt,name=FormModuleProps" json:"FormModuleProps,omitempty"`
	// User-defined category to organize actions list
	Category string `protobuf:"bytes,7,opt,name=Category" json:"Category,omitempty"`
	// User-defined hexa or rgb color
	Tint string `protobuf:"bytes,8,opt,name=Tint" json:"Tint,omitempty"`
	// Additional description for expected inputs
	InputDescription string `protobuf:"bytes,9,opt,name=InputDescription" json:"InputDescription,omitempty"`
	// Additional description describing the action output
	OutputDescription string `protobuf:"bytes,10,opt,name=OutputDescription" json:"OutputDescription,omitempty"`
	// If action is declared internal, it is hidden to avoid polluting the list.
	IsInternal bool `protobuf:"varint,13,opt,name=IsInternal" json:"IsInternal,omitempty"`
}

func (m *ActionDescription) Reset()                    { *m = ActionDescription{} }
func (m *ActionDescription) String() string            { return proto.CompactTextString(m) }
func (*ActionDescription) ProtoMessage()               {}
func (*ActionDescription) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{23} }

func (m *ActionDescription) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ActionDescription) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *ActionDescription) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *ActionDescription) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ActionDescription) GetSummaryTemplate() string {
	if m != nil {
		return m.SummaryTemplate
	}
	return ""
}

func (m *ActionDescription) GetHasForm() bool {
	if m != nil {
		return m.HasForm
	}
	return false
}

func (m *ActionDescription) GetFormModule() string {
	if m != nil {
		return m.FormModule
	}
	return ""
}

func (m *ActionDescription) GetFormModuleProps() string {
	if m != nil {
		return m.FormModuleProps
	}
	return ""
}

func (m *ActionDescription) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ActionDescription) GetTint() string {
	if m != nil {
		return m.Tint
	}
	return ""
}

func (m *ActionDescription) GetInputDescription() string {
	if m != nil {
		return m.InputDescription
	}
	return ""
}

func (m *ActionDescription) GetOutputDescription() string {
	if m != nil {
		return m.OutputDescription
	}
	return ""
}

func (m *ActionDescription) GetIsInternal() bool {
	if m != nil {
		return m.IsInternal
	}
	return false
}

type SchedulerActionsRequest struct {
}

func (m *SchedulerActionsRequest) Reset()                    { *m = SchedulerActionsRequest{} }
func (m *SchedulerActionsRequest) String() string            { return proto.CompactTextString(m) }
func (*SchedulerActionsRequest) ProtoMessage()               {}
func (*SchedulerActionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{24} }

type SchedulerActionsResponse struct {
	// List of all registered actions
	Actions map[string]*ActionDescription `protobuf:"bytes,1,rep,name=Actions" json:"Actions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SchedulerActionsResponse) Reset()                    { *m = SchedulerActionsResponse{} }
func (m *SchedulerActionsResponse) String() string            { return proto.CompactTextString(m) }
func (*SchedulerActionsResponse) ProtoMessage()               {}
func (*SchedulerActionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{25} }

func (m *SchedulerActionsResponse) GetActions() map[string]*ActionDescription {
	if m != nil {
		return m.Actions
	}
	return nil
}

type SchedulerActionFormRequest struct {
	// Name of the action to load
	ActionName string `protobuf:"bytes,1,opt,name=ActionName" json:"ActionName,omitempty"`
}

func (m *SchedulerActionFormRequest) Reset()                    { *m = SchedulerActionFormRequest{} }
func (m *SchedulerActionFormRequest) String() string            { return proto.CompactTextString(m) }
func (*SchedulerActionFormRequest) ProtoMessage()               {}
func (*SchedulerActionFormRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{26} }

func (m *SchedulerActionFormRequest) GetActionName() string {
	if m != nil {
		return m.ActionName
	}
	return ""
}

type SchedulerActionFormResponse struct {
}

func (m *SchedulerActionFormResponse) Reset()                    { *m = SchedulerActionFormResponse{} }
func (m *SchedulerActionFormResponse) String() string            { return proto.CompactTextString(m) }
func (*SchedulerActionFormResponse) ProtoMessage()               {}
func (*SchedulerActionFormResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{27} }

func init() {
	proto.RegisterType((*Configuration)(nil), "rest.Configuration")
	proto.RegisterType((*ListDataSourceRequest)(nil), "rest.ListDataSourceRequest")
	proto.RegisterType((*DataSourceCollection)(nil), "rest.DataSourceCollection")
	proto.RegisterType((*DeleteDataSourceResponse)(nil), "rest.DeleteDataSourceResponse")
	proto.RegisterType((*ListPeersAddressesRequest)(nil), "rest.ListPeersAddressesRequest")
	proto.RegisterType((*ListPeersAddressesResponse)(nil), "rest.ListPeersAddressesResponse")
	proto.RegisterType((*ListPeerFoldersRequest)(nil), "rest.ListPeerFoldersRequest")
	proto.RegisterType((*CreatePeerFolderRequest)(nil), "rest.CreatePeerFolderRequest")
	proto.RegisterType((*CreatePeerFolderResponse)(nil), "rest.CreatePeerFolderResponse")
	proto.RegisterType((*ListStorageBucketsRequest)(nil), "rest.ListStorageBucketsRequest")
	proto.RegisterType((*Process)(nil), "rest.Process")
	proto.RegisterType((*ListProcessesRequest)(nil), "rest.ListProcessesRequest")
	proto.RegisterType((*ListProcessesResponse)(nil), "rest.ListProcessesResponse")
	proto.RegisterType((*ListVersioningPolicyRequest)(nil), "rest.ListVersioningPolicyRequest")
	proto.RegisterType((*VersioningPolicyCollection)(nil), "rest.VersioningPolicyCollection")
	proto.RegisterType((*ListVirtualNodesRequest)(nil), "rest.ListVirtualNodesRequest")
	proto.RegisterType((*ListServiceRequest)(nil), "rest.ListServiceRequest")
	proto.RegisterType((*ServiceCollection)(nil), "rest.ServiceCollection")
	proto.RegisterType((*ControlServiceRequest)(nil), "rest.ControlServiceRequest")
	proto.RegisterType((*DiscoveryRequest)(nil), "rest.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "rest.DiscoveryResponse")
	proto.RegisterType((*ConfigFormRequest)(nil), "rest.ConfigFormRequest")
	proto.RegisterType((*OpenApiResponse)(nil), "rest.OpenApiResponse")
	proto.RegisterType((*ActionDescription)(nil), "rest.ActionDescription")
	proto.RegisterType((*SchedulerActionsRequest)(nil), "rest.SchedulerActionsRequest")
	proto.RegisterType((*SchedulerActionsResponse)(nil), "rest.SchedulerActionsResponse")
	proto.RegisterType((*SchedulerActionFormRequest)(nil), "rest.SchedulerActionFormRequest")
	proto.RegisterType((*SchedulerActionFormResponse)(nil), "rest.SchedulerActionFormResponse")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x5b, 0x6f, 0x1b, 0x45,
	0x14, 0x96, 0x73, 0xb3, 0x7d, 0xec, 0xa4, 0xf5, 0x92, 0x26, 0x5b, 0x57, 0x44, 0xd1, 0x0a, 0xa1,
	0x88, 0x52, 0x47, 0xb8, 0xa5, 0x20, 0x54, 0x09, 0x25, 0x76, 0x22, 0x2c, 0xa5, 0x89, 0xb5, 0xb6,
	0x78, 0x9f, 0xec, 0x1e, 0x9c, 0x25, 0xb3, 0x3b, 0xcb, 0xcc, 0x6c, 0x84, 0xdf, 0xf9, 0x37, 0xbc,
	0xf3, 0xc6, 0x0f, 0xe0, 0x5f, 0xa1, 0xb9, 0xec, 0x7a, 0x6c, 0xa7, 0x22, 0x12, 0x0f, 0x89, 0xe7,
	0x7c, 0xe7, 0x32, 0xe7, 0xf2, 0xcd, 0xb1, 0xa1, 0x1d, 0xb1, 0xec, 0x97, 0x64, 0xd6, 0xcb, 0x39,
	0x93, 0xcc, 0xdb, 0xe2, 0x28, 0x64, 0xf7, 0xed, 0x2c, 0x91, 0x77, 0xc5, 0x6d, 0x2f, 0x62, 0xe9,
	0x69, 0x3e, 0x8f, 0x13, 0x76, 0x1a, 0x21, 0xa5, 0xe2, 0x34, 0x62, 0x69, 0xca, 0xb2, 0x53, 0x6d,
	0x7a, 0x2a, 0x39, 0xa2, 0xfe, 0x67, 0x5c, 0xbb, 0xdf, 0x3d, 0xc5, 0x89, 0xdd, 0xfe, 0x8a, 0x91,
	0xb4, 0x1f, 0xd6, 0xf1, 0x9b, 0xa7, 0x38, 0x46, 0x92, 0xaa, 0x3f, 0xe3, 0x12, 0xfc, 0x08, 0xbb,
	0x03, 0x9d, 0x76, 0xc1, 0x89, 0x4c, 0x58, 0xe6, 0x75, 0xa1, 0x71, 0x59, 0x50, 0x3a, 0x26, 0xf2,
	0xce, 0xaf, 0x1d, 0xd7, 0x4e, 0x9a, 0x61, 0x25, 0x7b, 0x1e, 0x6c, 0x0d, 0x89, 0x24, 0xfe, 0x86,
	0xc6, 0xf5, 0x39, 0x38, 0x84, 0x17, 0x57, 0x89, 0x90, 0xea, 0x3c, 0x61, 0x05, 0x8f, 0x30, 0xc4,
	0xdf, 0x0a, 0x14, 0x32, 0xb8, 0x85, 0xfd, 0x05, 0x38, 0x60, 0x94, 0x62, 0xa4, 0x2f, 0x78, 0x07,
	0xad, 0x05, 0x2e, 0xfc, 0xda, 0xf1, 0xe6, 0x49, 0xab, 0xef, 0xf5, 0x6c, 0x21, 0x4e, 0x1c, 0xd7,
	0xcc, 0xdb, 0x87, 0xed, 0x29, 0x93, 0x84, 0xea, 0xbb, 0xb7, 0x43, 0x23, 0x04, 0xef, 0xc0, 0x1f,
	0x22, 0x45, 0x89, 0xee, 0xf5, 0x22, 0x67, 0x99, 0x40, 0xcf, 0x87, 0xfa, 0xa4, 0x88, 0x22, 0x14,
	0x42, 0xd7, 0xd1, 0x08, 0x4b, 0x31, 0x78, 0x05, 0x2f, 0x55, 0xca, 0x63, 0x44, 0x2e, 0xce, 0xe2,
	0x98, 0xa3, 0x10, 0x28, 0xca, 0xb4, 0xcf, 0xa1, 0xfb, 0x98, 0xd2, 0x06, 0xfd, 0x02, 0x76, 0x95,
	0xa6, 0x52, 0xe8, 0xf4, 0x9b, 0xe1, 0x32, 0x18, 0x5c, 0xc3, 0x41, 0x19, 0xe3, 0x92, 0xd1, 0x18,
	0x79, 0x19, 0xdd, 0x3b, 0x86, 0x96, 0x63, 0x6a, 0x1b, 0xec, 0x42, 0xaa, 0xc7, 0xba, 0xf7, 0xb6,
	0xc7, 0xea, 0x1c, 0xdc, 0xc0, 0xe1, 0x80, 0x23, 0x91, 0xb8, 0x88, 0xf8, 0xff, 0x02, 0x4e, 0xc1,
	0x5f, 0x0f, 0xf8, 0x5f, 0x7d, 0xf3, 0x8e, 0x60, 0xeb, 0x9a, 0xc5, 0xa8, 0x23, 0xb5, 0xfa, 0xd0,
	0xd3, 0x94, 0x55, 0x48, 0xa8, 0xf1, 0xa0, 0x30, 0x7d, 0x9d, 0x48, 0xc6, 0xc9, 0x0c, 0xcf, 0x8b,
	0xe8, 0x1e, 0x65, 0x55, 0x79, 0x1f, 0x60, 0x31, 0x24, 0x1d, 0xf9, 0xf1, 0xa9, 0x3b, 0x56, 0xaa,
	0xdb, 0x55, 0x94, 0x19, 0xfe, 0x9e, 0xdb, 0x1a, 0x96, 0xc1, 0xe0, 0x9f, 0x1a, 0xd4, 0xc7, 0x9c,
	0xe9, 0x14, 0xf7, 0x60, 0x63, 0x34, 0xb4, 0x5d, 0xd8, 0x18, 0x0d, 0x15, 0x9b, 0xc7, 0x84, 0x63,
	0x26, 0x47, 0x43, 0xeb, 0x5c, 0xc9, 0xaa, 0x75, 0x1f, 0x51, 0xf2, 0x24, 0x12, 0x63, 0xc6, 0xa5,
	0xbf, 0xa9, 0x89, 0xe5, 0x42, 0xde, 0x01, 0xec, 0xa8, 0x06, 0x8d, 0x62, 0x7f, 0x4b, 0xfb, 0x5a,
	0x69, 0xb5, 0xe9, 0xdb, 0xeb, 0x4d, 0xef, 0x42, 0x63, 0x22, 0x09, 0x97, 0x53, 0x32, 0xf3, 0x77,
	0xcc, 0xbd, 0xa5, 0xac, 0x75, 0xc8, 0x1f, 0x12, 0xc5, 0xfe, 0xba, 0xa6, 0x4f, 0x25, 0x07, 0x63,
	0xd8, 0xd7, 0xcc, 0x31, 0xe5, 0x54, 0xac, 0x74, 0x32, 0xa9, 0xad, 0x66, 0x62, 0x7d, 0xaf, 0x49,
	0x8a, 0xb6, 0x44, 0x17, 0x0a, 0x86, 0xe6, 0x7d, 0x3a, 0x11, 0xed, 0x9c, 0x5f, 0x43, 0xb3, 0x02,
	0xed, 0x2b, 0xdc, 0xed, 0xa9, 0xa5, 0xd5, 0xb3, 0x70, 0xb8, 0xd0, 0x07, 0x9f, 0xc3, 0x2b, 0x15,
	0xe5, 0x67, 0xe4, 0x22, 0x61, 0x59, 0x92, 0xcd, 0xc6, 0x8c, 0x26, 0xd1, 0xbc, 0x7c, 0x34, 0x63,
	0xe8, 0xae, 0xaa, 0x9c, 0x17, 0xdf, 0x87, 0x86, 0xc6, 0x92, 0xea, 0xa2, 0x03, 0xc3, 0x9d, 0xb5,
	0x70, 0x95, 0x5d, 0xf0, 0x12, 0x0e, 0xf5, 0x85, 0x09, 0x97, 0x05, 0xa1, 0x8a, 0x5e, 0xd5, 0x0b,
	0xbd, 0x02, 0x4f, 0xd3, 0xcc, 0x14, 0x59, 0x76, 0xe8, 0x3d, 0xb4, 0x27, 0x92, 0xc8, 0x42, 0x5c,
	0x26, 0x54, 0x22, 0xd7, 0x7d, 0xda, 0xeb, 0x7b, 0x3d, 0xb5, 0xea, 0xac, 0xa9, 0xd1, 0x87, 0x4b,
	0x76, 0xc1, 0x04, 0x3a, 0x56, 0xed, 0x64, 0x7c, 0xe2, 0x8c, 0xc8, 0x64, 0xdc, 0x76, 0x03, 0x2d,
	0x06, 0xf6, 0x89, 0xbd, 0xf4, 0x47, 0x0d, 0x5e, 0x0c, 0x58, 0x26, 0x39, 0xa3, 0x2b, 0x69, 0xae,
	0x0c, 0xac, 0xb6, 0x36, 0x30, 0x45, 0x0f, 0x55, 0xae, 0x33, 0xcf, 0x4a, 0xf6, 0xde, 0x40, 0x7d,
	0xc0, 0xd2, 0x94, 0x64, 0xb1, 0xa6, 0xeb, 0x5e, 0xff, 0x33, 0x37, 0x2d, 0xab, 0x0a, 0x4b, 0x9b,
	0xe0, 0x3d, 0x3c, 0x1f, 0x26, 0x22, 0x62, 0x0f, 0xc8, 0xcb, 0x51, 0x79, 0x01, 0xb4, 0x2f, 0xb2,
	0x38, 0x67, 0x49, 0x26, 0xa7, 0xf3, 0xbc, 0xcc, 0x60, 0x09, 0x0b, 0xfe, 0xde, 0x80, 0x8e, 0xe3,
	0x68, 0x09, 0xa3, 0x58, 0x4f, 0xa2, 0x7b, 0x32, 0x43, 0xc7, 0xd1, 0x85, 0x54, 0x6c, 0x2b, 0x5e,
	0x91, 0x5b, 0xa4, 0x36, 0xfd, 0x25, 0x4c, 0xad, 0x17, 0x3b, 0x76, 0x5d, 0x42, 0x33, 0x2c, 0x45,
	0xef, 0x08, 0xe0, 0xbc, 0x48, 0x68, 0x3c, 0x91, 0x24, 0xcd, 0xf5, 0x8b, 0xdb, 0x0e, 0x1d, 0xc4,
	0x6c, 0x83, 0x84, 0xc6, 0x21, 0x3e, 0x24, 0xda, 0x7f, 0xbb, 0xdc, 0x06, 0x0e, 0xe8, 0x0d, 0xa1,
	0x59, 0xd6, 0x22, 0xfc, 0x1d, 0x3d, 0xbb, 0x2f, 0x0d, 0xad, 0xd7, 0x2a, 0xea, 0x55, 0x86, 0x17,
	0x99, 0xe4, 0xf3, 0x70, 0xe1, 0xd8, 0xfd, 0x00, 0x7b, 0xcb, 0x4a, 0xef, 0x39, 0x6c, 0xde, 0xe3,
	0xdc, 0x56, 0xad, 0x8e, 0x6a, 0xf4, 0x0f, 0x84, 0x16, 0xe5, 0x94, 0x8c, 0xf0, 0xc3, 0xc6, 0xf7,
	0xb5, 0xe0, 0x5b, 0xe8, 0x98, 0x2f, 0xd5, 0x4b, 0xc6, 0xd3, 0x27, 0x4f, 0x3e, 0xe8, 0xc0, 0xb3,
	0x9b, 0x1c, 0xb3, 0xb3, 0x3c, 0x29, 0x33, 0x0c, 0xfe, 0xdc, 0x84, 0xce, 0x99, 0xe6, 0xe4, 0x10,
	0x45, 0xc4, 0x93, 0x5c, 0xd3, 0xd3, 0x83, 0x2d, 0x27, 0x86, 0x3e, 0x2b, 0x6c, 0x14, 0xb1, 0xac,
	0x5c, 0xf3, 0xea, 0xac, 0x32, 0x34, 0x83, 0x30, 0x9d, 0x36, 0x82, 0x4a, 0xc4, 0x09, 0x66, 0x57,
	0x9b, 0x0b, 0x79, 0x27, 0xf0, 0x6c, 0x52, 0xa4, 0x29, 0xe1, 0xf3, 0x29, 0xa6, 0x39, 0x25, 0x12,
	0x6d, 0xaf, 0x57, 0x61, 0x35, 0xcd, 0x9f, 0x88, 0x50, 0x65, 0xea, 0x35, 0xd7, 0x08, 0x4b, 0x51,
	0x4d, 0x53, 0x7d, 0x7e, 0x64, 0x71, 0x41, 0xd1, 0x6f, 0x69, 0x77, 0x07, 0x51, 0x77, 0x2c, 0xa4,
	0x31, 0x67, 0xb9, 0xf0, 0xdb, 0xe6, 0x8e, 0x15, 0x58, 0x3d, 0x88, 0x01, 0x91, 0x38, 0x63, 0x7c,
	0xee, 0xd7, 0xcd, 0x83, 0x28, 0x65, 0x55, 0xf5, 0x34, 0xc9, 0xa4, 0xdf, 0x30, 0x55, 0xab, 0xb3,
	0xf7, 0x15, 0x3c, 0x1f, 0x65, 0x79, 0x21, 0xdd, 0x22, 0x9b, 0x5a, 0xbf, 0x86, 0x7b, 0x5f, 0x43,
	0xe7, 0xa6, 0x90, 0x2b, 0xc6, 0xa0, 0x8d, 0xd7, 0x15, 0xaa, 0xa6, 0x91, 0x18, 0x65, 0x12, 0x79,
	0x46, 0xa8, 0xbf, 0xab, 0x0b, 0x76, 0x10, 0xb5, 0xb4, 0x26, 0xd1, 0x1d, 0xaa, 0xd4, 0xb9, 0x99,
	0x5a, 0xb5, 0xb4, 0xfe, 0xaa, 0x81, 0xbf, 0xae, 0xb3, 0x2f, 0xeb, 0x02, 0xea, 0x16, 0xb2, 0xdb,
	0xe6, 0xb5, 0x61, 0xec, 0xa7, 0x1c, 0x7a, 0x56, 0x36, 0xb4, 0x2d, 0x7d, 0xbb, 0x13, 0x68, 0xbb,
	0x8a, 0x47, 0x28, 0xfb, 0xc6, 0xa5, 0x6c, 0xab, 0x7f, 0x68, 0xae, 0x59, 0x23, 0x98, 0xcb, 0xe5,
	0x0f, 0xd0, 0x5d, 0x49, 0xc3, 0x25, 0xf5, 0x11, 0x80, 0x01, 0x1d, 0x3e, 0x3a, 0x88, 0xfa, 0xde,
	0x78, 0xd4, 0xdb, 0xd4, 0x71, 0xbb, 0xa3, 0x7f, 0x84, 0xbe, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff,
	0xf4, 0xf5, 0xc2, 0xac, 0x3b, 0x0b, 0x00, 0x00,
}
