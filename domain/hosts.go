package domain

type Json struct {
	Host Host
}

type Host []struct {
	Address             	string       		 	`json:"address"`
	AutoNumaStatus      	string              	`json:"auto_numa_status"`
	Certificate         	*Certificate       		`json:"certificate"`
	Cpu                 	*Cpu                 	`json:"cpu"`
	DevicePassThrough   	*DevicePassThrough   	`json:"device_passthrough"`
	ExternalStatus      	string               	`json:"external_status"`
	HardwareInformation 	*HardwareInformation 	`json:"hardware_information"`
	Iscsi              		*Iscsi               	`json:"iscsi"`
	KDumpStatus             string           		`json:"kdump_status"`
	Ksm                     *Ksm             		`json:"ksm"`
	LibVirtVersion          *LibVirtVersion  		`json:"libvirt_version"`
	MaxSchedulingMemory     string           		`json:"max_scheduling_memory"`
	Memory                  string           		`json:"memory"`
	NumaSupported           string           		`json:"numa_supported"`
	Os                      *Os              		`json:"os"`
	Port                    string           		`json:"port"`
	PowerManagement         *PowerManagement 		`json:"power_management"`
	Protocol                string           		`json:"protocol"`
	ReInstallationRequired  string           		`json:"reinstallation_required"`
	SeLinux                 *SeLinux                `json:"se_linux"`
	Spm                     *Spm              		`json:"spm"`
	Ssh                     *Ssh              		`json:"ssh"`
	Status                  string           		`json:"status"`
	Summary                 *Summary                `json:"summary"`
	TransparentHugePages  	*TransparentHugePages   `json:"transparent_hugepages"`
	Type               		string             		`json:"type"`
	UpdateAvailable         string                  `json:"update_available"`
	Versions                *Versions               `json:"version"`
	VGpuPlacement           string                  `json:"vgpu_placement"`
	Cluster                 *Cluster                `json:"cluster"`
	Actions                 *Actions                `json:"actions"`
	Name                    string                  `json:"name"`
	Comment                 string                  `json:"comment"`
	Href                    string                  `json:"href"`
	Id                      string                  `json:"id"`
	Link                    Link                    `json:"link"`
}

type Certificate struct {
	Organization   string 		`json:"organization"`
	Subject        string 		`json:"subject"`
}

type Cpu struct {
	Name      	   string 		`json:"name"`
	Speed          string 		`json:"speed"`
	Topology 	   *Topology 	`json:"topology"`
	Type           string 		`json:"type"`
}

type Topology struct {
	Cores   	   string 		`json:"cores"`
	Sockets        string 		`json:"sockets"`
	Threads 	   string 		`json:"threads"`
}

type DevicePassThrough struct {
	Enabled   string 		`json:"enabled"`
}

type HardwareInformation struct {
	Family              	string 					`json:"family"`
	Manufacturer        	string 					`json:"manufacturer"`
	ProductName        		string 					`json:"product_name"`
	SerialNumber       		string 					`json:"serial_number"`
	SupportedRngSources 	*SupportedRngSources 	`json:"supported_rng_sources"`
	Uuid        			string					`json:"uuid"`
	Version        			string 					`json:"version"`
}

type SupportedRngSources struct {
	SupportedRngSource   string 		`json:"supported_rng_source"`
}

type Iscsi struct {
	Initiator   string 		`json:"initiator"`
}

type Ksm struct {
	Enabled   string 		`json:"enabled"`
}

type LibVirtVersion struct {
	Build   		string 		`json:"build"`
	FullVersion     string 		`json:"full_version"`
	Major   		string 		`json:"major"`
	Minor   		string 		`json:"minor"`
	Revision   		string 		`json:"revision"`
}

type Os struct {
	CustomKernelCmdline   string   `json:"custom_kernel_cmdline"`
	ReportedKernelCmdline string   `json:"reported_kernel_cmdline"`
	Type                  string   `json:"type"`
	Version               *Version `json:"version"`
}

type Version struct {
	FullVersion   string 		`json:"full_version"`
	Major 		  string 		`json:"major"`
	Minor   	  string 		`json:"minor"`
}

type PowerManagement struct {
	AutomaticPmEnabled string `json:"automatic_pm_enabled"`
	Enabled            string `json:"enabled"`
	KDumpDetection     string `json:"kdump_detection"`
	PmProxies          string `json:"pm_proxies"`
}

type SeLinux struct {
	Mode   string 		`json:"mode"`
}

type Spm struct {
	Priority   string 		`json:"priority"`
	Status     string 		`json:"status"`
}

type Ssh struct {
	FingerPrint   string 		`json:"fingerprint"`
	Port          string 		`json:"port"`
	PublicKey     string 		`json:"public_key"`
}

type Summary struct {
	Active      string 		`json:"active"`
	Migrating   string 		`json:"migrating"`
	Total       string 		`json:"total"`
}

type TransparentHugePages struct {
	Enabled   string 		`json:"enabled"`
}

type Versions struct {
	Build         string 		`json:"build"`
	FullVersion   string 		`json:"full_version"`
	Major   	  string 		`json:"major"`
	Minor         string 		`json:"minor"`
	Revision      string 		`json:"revision"`
}
type Cluster struct {
	Href   string 		`json:"href"`
	Id     string 		`json:"id"`
}

type Actions struct {
	Links   Links 		`json:"link"`
}

type Links []struct {
	Href   string 		`json:"href"`
	Id     string 		`json:"id"`
}

type Link []struct {
	Href   string 		`json:"href"`
	Id     string 		`json:"id"`
}