package tutum

type ActionListResponse struct {
	Objects []Action `json:"objects"`
}

type Action struct {
	Action       string `json:"action"`
	Body         string `json:"body"`
	End_date     string `json:"end_date"`
	Ip           string `json:"ip"`
	Location     string `json:"location"`
	Logs         string `json:"logs"`
	Method       string `json:"method"`
	Object       string `json:"object"`
	Path         string `json:"path"`
	Resource_uri string `json:"resource_uri"`
	Start_date   string `json:"start_date"`
	State        string `json:"state"`
	Uuid         string `json:"uuid"`
}

type CListResponse struct {
	Objects []Container `json:"objects"`
}

type Container struct {
	Application            string    `json:"application"`
	Autodestroy            string    `json:"autodestroy"`
	Autoreplace            string    `json:"autoreplace"`
	Autorestart            string    `json:"autorestart"`
	Container_ports        []CCPInfo `json:"container_ports"`
	Container_size         string    `json:"container_size"`
	Current_num_containers int       `json:"current_num_containers"`
	Deployed_datetime      string    `json:"deployed_datetime"`
	Destroyed_datetime     string    `json:"destroyed_datetime"`
	Entrypoint             string    `json:"entrypoint"`
	Exit_code              int       `json:"exit_code"`
	Exit_code_message      string    `json:"exit_code_message"`
	Image_name             string    `json:"image_name"`
	Image_tag              string    `json:"image_tag"`
	Name                   string    `json:"name"`
	Public_dns             string    `json:"public_dns"`
	Resource_uri           string    `json:"resource_uri"`
	Run_command            string    `json:"run_command"`
	Started_datetime       string    `json:"started_datetime"`
	State                  string    `json:"state"`
	Stopped_datetime       string    `json:"stopped_datetime"`
	Unique_name            string    `json:"unique_name"`
	Uuid                   string    `json:"uuid"`
}

type CCPInfo struct {
	Container  string `json:"container"`
	Inner_port int    `json:"inner_port"`
	Outer_port int    `json:"outer_port"`
	Protocol   string `json:"protocol"`
}

type Event struct {
	Type         string   `json:"type"`
	Action       string   `json:"action"`
	Parents      []string `json:"parents"`
	Resource_uri string   `json:"resource_uri"`
	State        string   `json:"state"`
}

type Exec struct {
	Type       string `json:"type"`
	Output     string `json:"output"`
	StreamType string `json:"streamType"`
}

type Logs struct {
	Type      string `json:"type"`
	Source    string `json:"source"`
	Log       string `json:"log"`
	Timestamp int    `json:"timestamp"`
}

type NodeListResponse struct {
	Objects []Node `json:"objects"`
}

type Node struct {
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime"`
	Docker_version     string    `json:"docker_version"`
	Last_seen          string    `json:"last_seen"`
	Node_cluster       string    `json:"node_cluster"`
	Public_ip          string    `json:"public_ip"`
	Region             string    `json:"region"`
	Resource_uri       string    `json:"resource_uri"`
	State              string    `json:"state"`
	Tags               []NodeTag `json:"tags"`
	Uuid               string    `json:"uuid"`
}

type NodeTag struct {
	Name string `json:"name"`
}

type NodeClusterListResponse struct {
	Objects []NodeCluster `json:"objects"`
}

type NodeCluster struct {
	Current_num_nodes  int       `json:"current_num_nodes"`
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime"`
	Disk               int       `json:"disk"`
	Nodes              []string  `json:"nodes"`
	Region             string    `json:"region"`
	Resource_uri       string    `json:"resource_uri"`
	State              string    `json:"state"`
	Tags               []NodeTag `json:"tags"`
	Target_num_nodes   int       `json:"target_num_nodes"`
	Uuid               string    `json:"uuid"`
}

type NodeTypeListResponse struct {
	Objects []NodeType `json:"objects"`
}

type NodeType struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Provider     string   `json:"provider"`
	Regions      []string `json:"regions"`
	Resource_uri string   `json:"resource_uri"`
}

type ProviderListResponse struct {
	Objects []Provider `json:"objects"`
}

type Provider struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Regions      []string `json:"regions"`
	Resource_uri string   `json:"resource_uri"`
}

type RegionListResponse struct {
	Objects []Region `json:"objects"`
}

type Region struct {
	Available    bool     `json:"available"`
	Label        string   `json:"label"`
	Name         string   `json:"name"`
	Node_types   []string `json:"node_types"`
	Provider     string   `json:"provider"`
	Resource_uri string   `json:"resource_uri"`
}

type SListResponse struct {
	Objects []Service `json: "objects"`
}

type Service struct {
	Autodestroy            string       `json:"autodestroy`
	Autoredeploy           bool         `json:"autoredeploy`
	Autorestart            string       `json:"autorestart`
	Containers             []string     `json:"containers`
	Container_ports        []SCPInfo    `json:"container_ports`
	Container_size         string       `json:"container_size`
	Current_num_containers int          `json:"current_num_containers`
	Deployed_datetime      string       `json:"deployed_datetime`
	Destroyed_datetime     string       `json:"destroyed_datetime`
	Entrypoint             string       `json:"entrypoint`
	Exit_code              int          `json:"exit_code`
	Exit_code_message      string       `json:"exit_code_message`
	Image_name             string       `json:"image_name"`
	Image_tag              string       `json:"image_tag`
	Linked_to_service      []LinkToInfo `json:"linked_to_service`
	Name                   string       `json:"name"`
	Public_dns             string       `json:"public_dns"`
	Resource_uri           string       `json:"resource_uri"`
	Run_command            string       `json:"run_command"`
	Started_datetime       string       `json:"started_datetime"`
	State                  string       `json:"state"`
	Stack                  string       `json:"stack"`
	Stopped_datetime       string       `json:"stopped_datetime"`
	Target_num_containers  int          `json:"target_num_containers"`
	Unique_name            string       `json:"unique_name"`
	Uuid                   string       `json:"uuid"`
}

type SCPInfo struct {
	Container  string `json:"container"`
	Inner_port int    `json:"inner_port"`
	Outer_port int    `json:"outer_port"`
	Protocol   string `json:"protocol"`
}

//Basic information from linked services
type LinkToInfo struct {
	From_service string `json:"from_service"`
	Name         string `json:"name"`
	To_service   string `json:"to_service"`
}

type StackListResponse struct {
	Objects []StackShort `json:"objects"`
}

type StackShort struct {
	Deployed_datetime  string   `json:"deployed_datetime"`
	Destroyed_datetime string   `json:"destroyed_datetime"`
	Name               string   `json:"name"`
	Resource_uri       string   `json:"resource_uri"`
	Service            []string `json:"services"`
	State              string   `json:"state"`
	Synchronized       bool     `json:"synchronized"`
	Uuid               string   `json:"uuid"`
}

type Stack struct {
	Deployed_datetime  string    `json:"deployed_datetime"`
	Destroyed_datetime string    `json:"destroyed_datetime`
	Name               string    `json:"name"`
	Resource_uri       string    `json:"resource_uri`
	Service            []Service `json:"services"`
	State              string    `json:"state"`
	Synchronized       bool      `json:"synchronized"`
	Uuid               string    `json:"uuid"`
}

type Token struct {
	Token string `json:"token"`
}

type TriggerListResponse struct {
	Objects []Trigger `json:"objects"`
}

type Trigger struct {
	Url          string `json:"url"`
	Name         string `json:"name"`
	Operation    string `json:"operation"`
	Resource_uri string `json:"resource_uri"`
}

type VolumeListResponse struct {
	Objects []Volume `json:"objects"`
}

type Volume struct {
	Containers   []string `json:"containers"`
	Node         string   `json:"node"`
	Resource_uri string   `json:"resource_uri"`
	State        string   `json:"state"`
	Uuid         string   `json:"uuid"`
	Volume_group string   `json:"volume_group"`
}

type VolumeGroupListResponse struct {
	Objects []VolumeGroup `json:"objects"`
}

type VolumeGroup struct {
	Name         string   `json:"name"`
	Resource_uri string   `json:"resource_uri"`
	Services     []string `json:"services"`
	State        string   `json:"state"`
	Uuid         string   `json:"uuid"`
	Volume       []string `json:"volume"`
}
