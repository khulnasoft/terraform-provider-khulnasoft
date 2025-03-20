package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"time"
)

type RuntimePolicy struct {
	AllowedExecutables         AllowedExecutables       `json:"allowed_executables"`
	AllowedRegistries          AllowedRegistries        `json:"allowed_registries"`
	ApplicationScopes          []string                 `json:"application_scopes"`
	AuditBruteForceLogin       bool                     `json:"audit_brute_force_login"`
	AuditOnFailure             bool                     `json:"audit_on_failure,omitempty"`
	Auditing                   Auditing                 `json:"auditing"`
	Author                     string                   `json:"author"`
	BlacklistedOsUsers         BlacklistedOsUsers       `json:"blacklisted_os_users,omitempty"`
	BlockDisallowedImages      bool                     `json:"block_disallowed_images,omitempty"`
	BlockFailed                bool                     `json:"block_failed,omitempty"`
	BlockFilelessExec          bool                     `json:"block_fileless_exec"`
	BlockNonCompliantWorkloads bool                     `json:"block_non_compliant_workloads"`
	BlockNonK8sContainers      bool                     `json:"block_non_k8s_containers"`
	BlockNwUnlinkCont          bool                     `json:"block_nw_unlink_cont,omitempty"`
	BypassScope                BypassScope              `json:"bypass_scope"`
	ContainerExec              ContainerExec            `json:"container_exec"`
	Created                    string                   `json:"created,omitempty"`
	Cve                        string                   `json:"cve"`
	DefaultSecurityProfile     string                   `json:"default_security_profile"`
	Description                string                   `json:"description"`
	Digest                     string                   `json:"digest"`
	Domain                     string                   `json:"domain,omitempty"`
	DomainName                 string                   `json:"domain_name,omitempty"`
	DriftPrevention            DriftPrevention          `json:"drift_prevention"`
	EnableCryptoMiningDns      bool                     `json:"enable_crypto_mining_dns,omitempty"`
	EnableForkGuard            bool                     `json:"enable_fork_guard"`
	EnableIPReputation         bool                     `json:"enable_ip_reputation"`
	EnablePortScanProtection   bool                     `json:"enable_port_scan_protection"`
	Enabled                    bool                     `json:"enabled"`
	Enforce                    bool                     `json:"enforce"`
	EnforceAfterDays           int                      `json:"enforce_after_days"`
	EnforceSchedulerAddedOn    int                      `json:"enforce_scheduler_added_on,omitempty"`
	ExecutableBlacklist        ExecutableBlacklist      `json:"executable_blacklist"`
	FailCicd                   bool                     `json:"fail_cicd,omitempty"`
	FailedKubernetesChecks     FailedKubernetesChecks   `json:"failed_kubernetes_checks"`
	FileBlock                  FileBlock                `json:"file_block"`
	FileIntegrityMonitoring    FileIntegrityMonitoring  `json:"file_integrity_monitoring"`
	ForkGuardProcessLimit      int                      `json:"fork_guard_process_limit"`
	HeuristicRefID             int                      `json:"heuristic_ref_id,omitempty"`
	ImageID                    int                      `json:"image_id,omitempty"`
	ImageName                  string                   `json:"image_name"`
	IsAuditChecked             bool                     `json:"is_audit_checked"`
	IsAutoGenerated            bool                     `json:"is_auto_generated"`
	Lastupdate                 int                      `json:"lastupdate,omitempty"`
	LimitContainerPrivileges   LimitContainerPrivileges `json:"limit_container_privileges"`
	LinuxCapabilities          LinuxCapabilities        `json:"linux_capabilities"`
	MalwareScanOptions         MalwareScanOptions       `json:"malware_scan_options"`
	Name                       string                   `json:"name"`
	NoNewPrivileges            bool                     `json:"no_new_privileges"`
	OnlyRegisteredImages       bool                     `json:"only_registered_images,omitempty"`
	PackageBlock               PackageBlock             `json:"package_block"`
	Permission                 string                   `json:"permission,omitempty"`
	PortBlock                  PortBlock                `json:"port_block"`
	//PreventOverrideDefaultConfig PreventOverrideDefaultConfig `json:"prevent_override_default_config,omitempty"`
	ReadonlyFiles             ReadonlyFiles             `json:"readonly_files"`
	ReadonlyRegistry          ReadonlyRegistry          `json:"readonly_registry"`
	Registry                  string                    `json:"registry"`
	RegistryAccessMonitoring  RegistryAccessMonitoring  `json:"registry_access_monitoring"`
	RepoID                    int                       `json:"repo_id,omitempty"`
	RepoName                  string                    `json:"repo_name"`
	ResourceName              string                    `json:"resource_name"`
	ResourceType              string                    `json:"resource_type"`
	RestrictedVolumes         RestrictedVolumes         `json:"restricted_volumes"`
	ReverseShell              ReverseShell              `json:"reverse_shell"`
	RuntimeType               string                    `json:"runtime_type"`
	Scope                     Scope                     `json:"scope"`
	SystemIntegrityProtection SystemIntegrityProtection `json:"system_integrity_protection"`
	Tripwire                  Tripwire                  `json:"tripwire"`
	Type                      string                    `json:"type"`
	Updated                   time.Time                 `json:"updated"`
	Version                   string                    `json:"version"`
	VpatchVersion             string                    `json:"vpatch_version"`
	VulnID                    int                       `json:"vuln_id,omitempty"`
	WhitelistedOsUsers        WhitelistedOsUsers        `json:"whitelisted_os_users"`
	//JSON
	//EnableCryptoMiningDNS bool `json:"enable_crypto_mining_dns"`
	BlockContainerExec       bool     `json:"block_container_exec,omitempty"`
	IsOOTBPolicy             bool     `json:"is_ootb_policy,omitempty"`
	RuntimeMode              int      `json:"runtime_mode,omitempty"`
	ExcludeApplicationScopes []string `json:"exclude_application_scopes,omitempty"`
}

type AllowedExecutables struct {
	AllowExecutables     []string `json:"allow_executables,omitempty"`
	AllowRootExecutables []string `json:"allow_root_executables,omitempty"`
	Enabled              bool     `json:"enabled"`
	SeparateExecutables  bool     `json:"separate_executables,omitempty"`
}

type AllowedRegistries struct {
	AllowedRegistries []string `json:"allowed_registries"`
	Enabled           bool     `json:"enabled"`
}

type ExecutableBlacklist struct {
	Enabled     bool     `json:"enabled"`
	Executables []string `json:"executables"`
}

type FailedKubernetesChecks struct {
	Enabled      bool     `json:"enabled"`
	FailedChecks []string `json:"failed_checks"`
}
type DriftPrevention struct {
	Enabled               bool     `json:"enabled"`
	ExecLockdown          bool     `json:"exec_lockdown"`
	ImageLockdown         bool     `json:"image_lockdown"`
	PreventPrivileged     bool     `json:"prevent_privileged,omitempty"`
	ExecLockdownWhiteList []string `json:"exec_lockdown_white_list"`
}

type RestrictedVolumes struct {
	Enabled bool     `json:"enabled"`
	Volumes []string `json:"volumes"`
}

type BypassScope struct {
	Enabled bool  `json:"enabled"`
	Scope   Scope `json:"scope"`
}

type LimitContainerPrivileges struct {
	Enabled               bool `json:"enabled"`
	Privileged            bool `json:"privileged,omitempty"`
	Netmode               bool `json:"netmode,omitempty"`
	Pidmode               bool `json:"pidmode,omitempty"`
	Utsmode               bool `json:"utsmode,omitempty"`
	Usermode              bool `json:"usermode,omitempty"`
	Ipcmode               bool `json:"ipcmode,omitempty"`
	PreventRootUser       bool `json:"prevent_root_user,omitempty"`
	PreventLowPortBinding bool `json:"prevent_low_port_binding,omitempty"`
	BlockAddCapabilities  bool `json:"block_add_capabilities,omitempty"`
	UseHostUser           bool `json:"use_host_user,omitempty"`
}

type PreventOverrideDefaultConfig struct {
	Enabled         bool `json:"enabled,omitempty"`
	EnforceSelinux  bool `json:"enforce_selinux,omitempty"`
	EnforceSeccomp  bool `json:"enforce_seccomp,omitempty"`
	EnforceApparmor bool `json:"enforce_apparmor,omitempty"`
}

type Auditing struct {
	AuditAllNetwork            bool `json:"audit_all_network"`
	AuditAllProcesses          bool `json:"audit_all_processes"`
	AuditFailedLogin           bool `json:"audit_failed_login"`
	AuditOsUserActivity        bool `json:"audit_os_user_activity"`
	AuditProcessCmdline        bool `json:"audit_process_cmdline"`
	AuditSuccessLogin          bool `json:"audit_success_login"`
	AuditUserAccountManagement bool `json:"audit_user_account_management"`
	Enabled                    bool `json:"enabled"`
}

type BlacklistedOsUsers struct {
	Enabled        bool     `json:"enabled"`
	UserBlackList  []string `json:"user_black_list"`
	GroupBlackList []string `json:"group_black_list"`
}

type WhitelistedOsUsers struct {
	Enabled        bool     `json:"enabled"`
	UserWhiteList  []string `json:"user_white_list"`
	GroupWhiteList []string `json:"group_white_list"`
}

type FileBlock struct {
	Enabled                        bool     `json:"enabled"`
	FilenameBlockList              []string `json:"filename_block_list"`
	ExceptionalBlockFiles          []string `json:"exceptional_block_files"`
	BlockFilesUsers                []string `json:"block_files_users"`
	BlockFilesProcesses            []string `json:"block_files_processes"`
	ExceptionalBlockFilesUsers     []string `json:"exceptional_block_files_users"`
	ExceptionalBlockFilesProcesses []string `json:"exceptional_block_files_processes"`
}

type PackageBlock struct {
	Enabled                           bool     `json:"enabled"`
	PackagesBlackList                 []string `json:"packages_black_list"`
	ExceptionalBlockPackagesFiles     []string `json:"exceptional_block_packages_files"`
	BlockPackagesUsers                []string `json:"block_packages_users"`
	BlockPackagesProcesses            []string `json:"block_packages_processes"`
	ExceptionalBlockPackagesUsers     []string `json:"exceptional_block_packages_users"`
	ExceptionalBlockPackagesProcesses []string `json:"exceptional_block_packages_processes"`
}

type LinuxCapabilities struct {
	Enabled                 bool     `json:"enabled"`
	RemoveLinuxCapabilities []string `json:"remove_linux_capabilities"`
}

type MalwareScanOptions struct {
	Action             string   `json:"action"`
	Enabled            bool     `json:"enabled"`
	ExcludeDirectories []string `json:"exclude_directories"`
	ExcludeProcesses   []string `json:"exclude_processes"`
	IncludeDirectories []string `json:"include_directories"`
}

type PortBlock struct {
	Enabled            bool     `json:"enabled"`
	BlockInboundPorts  []string `json:"block_inbound_ports"`
	BlockOutboundPorts []string `json:"block_outbound_ports"`
}

type Tripwire struct {
	Enabled       bool     `json:"enabled"`
	UserID        string   `json:"user_id"`
	UserPassword  string   `json:"user_password"`
	ApplyOn       []string `json:"apply_on"`
	ServerlessApp string   `json:"serverless_app"`
}

type FileIntegrityMonitoring struct {
	Enabled                            bool     `json:"enabled"`
	MonitoredFiles                     []string `json:"monitored_files"`
	ExceptionalMonitoredFiles          []string `json:"exceptional_monitored_files"`
	MonitoredFilesProcesses            []string `json:"monitored_files_processes"`
	ExceptionalMonitoredFilesProcesses []string `json:"exceptional_monitored_files_processes"`
	MonitoredFilesUsers                []string `json:"monitored_files_users"`
	ExceptionalMonitoredFilesUsers     []string `json:"exceptional_monitored_files_users"`
	MonitoredFilesCreate               bool     `json:"monitored_files_create"`
	MonitoredFilesRead                 bool     `json:"monitored_files_read"`
	MonitoredFilesModify               bool     `json:"monitored_files_modify"`
	MonitoredFilesDelete               bool     `json:"monitored_files_delete"`
	MonitoredFilesAttributes           bool     `json:"monitored_files_attributes"`
}

type RegistryAccessMonitoring struct {
	Enabled                               bool     `json:"enabled"`
	ExceptionalMonitoredRegistryPaths     []string `json:"exceptional_monitored_registry_paths"`
	ExceptionalMonitoredRegistryProcesses []string `json:"exceptional_monitored_registry_processes"`
	ExceptionalMonitoredRegistryUsers     []string `json:"exceptional_monitored_registry_users"`
	MonitoredRegistryAttributes           bool     `json:"monitored_registry_attributes"`
	MonitoredRegistryCreate               bool     `json:"monitored_registry_create"`
	MonitoredRegistryDelete               bool     `json:"monitored_registry_delete"`
	MonitoredRegistryModify               bool     `json:"monitored_registry_modify"`
	MonitoredRegistryPaths                []string `json:"monitored_registry_paths"`
	MonitoredRegistryProcesses            []string `json:"monitored_registry_processes"`
	MonitoredRegistryRead                 bool     `json:"monitored_registry_read"`
	MonitoredRegistryUsers                []string `json:"monitored_registry_users"`
}

type SystemIntegrityProtection struct {
	AuditSystemtimeChange     bool `json:"audit_systemtime_change"`
	Enabled                   bool `json:"enabled"`
	MonitorAuditLogIntegrity  bool `json:"monitor_audit_log_integrity"`
	WindowsServicesMonitoring bool `json:"windows_services_monitoring"`
}

type ReadonlyFiles struct {
	Enabled                           bool     `json:"enabled"`
	ReadonlyFiles                     []string `json:"readonly_files"`
	ExceptionalReadonlyFiles          []string `json:"exceptional_readonly_files"`
	ReadonlyFilesProcesses            []string `json:"readonly_files_processes"`
	ExceptionalReadonlyFilesProcesses []string `json:"exceptional_readonly_files_processes"`
	ReadonlyFilesUsers                []string `json:"readonly_files_users"`
	ExceptionalReadonlyFilesUsers     []string `json:"exceptional_readonly_files_users"`
}

type ReadonlyRegistry struct {
	Enabled                              bool     `json:"enabled"`
	ExceptionalReadonlyRegistryPaths     []string `json:"exceptional_readonly_registry_paths"`
	ExceptionalReadonlyRegistryProcesses []string `json:"exceptional_readonly_registry_processes"`
	ExceptionalReadonlyRegistryUsers     []string `json:"exceptional_readonly_registry_users"`
	ReadonlyRegistryPaths                []string `json:"readonly_registry_paths"`
	ReadonlyRegistryProcesses            []string `json:"readonly_registry_processes"`
	ReadonlyRegistryUsers                []string `json:"readonly_registry_users"`
}

type ContainerExec struct {
	Enabled                    bool     `json:"enabled"`
	BlockContainerExec         bool     `json:"block_container_exec"`
	ContainerExecProcWhiteList []string `json:"container_exec_proc_white_list"`
}

type ReverseShell struct {
	Enabled                   bool     `json:"enabled"`
	BlockReverseShell         bool     `json:"block_reverse_shell"`
	ReverseShellProcWhiteList []string `json:"reverse_shell_proc_white_list"`
	ReverseShellIpWhiteList   []string `json:"reverse_shell_ip_white_list"`
}

// JSON

// CreateRuntimePolicy creates an Khulnasoft RuntimePolicy
func (cli *Client) CreateRuntimePolicy(runtimePolicy *RuntimePolicy) error {
	payload, err := json.Marshal(runtimePolicy)
	if err != nil {
		return err
	}

	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v2/runtime_policies")
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	// Create a copy of the payload with sensitive information removed
	var sanitizedPayload map[string]interface{}
	json.Unmarshal(payload, &sanitizedPayload)
	if tripwire, ok := sanitizedPayload["tripwire"].(map[string]interface{}); ok {
		tripwire["user_password"] = "REDACTED"
	}
	sanitatedPayloadBytes, _ := json.Marshal(sanitizedPayload)
	log.Println(string(sanitatedPayloadBytes))
	resp, body, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Post(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed creating runtime policy.")
	}
	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		var errorResponse ErrorResponse
		err = json.Unmarshal([]byte(body), &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v", body)
			return fmt.Errorf("failed creating runtime policy with name %v. Status: %v, Response: %v", runtimePolicy.Name, resp.StatusCode, body)
		}
		return fmt.Errorf("failed creating runtime policy. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}

	return nil
}

// GetRuntimePolicy gets an Khulnasoft runtime policy by name
func (cli *Client) GetRuntimePolicy(name string) (*RuntimePolicy, error) {
	var err error
	var response RuntimePolicy
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v2/runtime_policies/%v", name)
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return nil, err
	}
	events, body, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Get(cli.url + apiPath).End()
	if errs != nil {
		return nil, errors.Wrap(getMergedError(errs), "failed getting runtime policy with name "+name)
	}
	if events.StatusCode == 200 {
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			log.Printf("Error unmarshaling response body")
			return nil, errors.Wrap(err, fmt.Sprintf("couldn't unmarshal get runtime policy response. Body: %v", body))
		}
	} else {
		var errorReponse ErrorResponse
		err = json.Unmarshal([]byte(body), &errorReponse)
		if err != nil {
			log.Println("failed to unmarshal error response")
			return nil, fmt.Errorf("failed getting runtime policy with name %v. Status: %v, Response: %v", name, events.StatusCode, body)
		}

		return nil, fmt.Errorf("failed getting runtime policy with name %v. Status: %v, error message: %v", name, events.StatusCode, errorReponse.Message)
	}

	return &response, nil
}

// UpdateRuntimePolicy updates an existing runtime policy policy
func (cli *Client) UpdateRuntimePolicy(runtimePolicy *RuntimePolicy) error {
	payload, err := json.Marshal(runtimePolicy)
	if err != nil {
		return err
	}
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v2/runtime_policies/%s", runtimePolicy.Name)
	err = cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, _, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Put(cli.url + apiPath).Send(string(payload)).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed modifying runtime policy")
	}
	if resp.StatusCode != 201 && resp.StatusCode != 204 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response Body")
			return err
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(body, &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v. error: %v", string(body), err)
			return err
		}
		return fmt.Errorf("failed modifying runtime policy. status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	return nil
}

// DeleteRuntimePolicy removes a Khulnasoft runtime policy
func (cli *Client) DeleteRuntimePolicy(name string) error {
	request := cli.gorequest
	apiPath := fmt.Sprintf("/api/v2/runtime_policies/%s", name)
	err := cli.limiter.Wait(context.Background())
	if err != nil {
		return err
	}
	resp, body, errs := request.Clone().Set("Authorization", "Bearer "+cli.token).Delete(cli.url + apiPath).End()
	if errs != nil {
		return errors.Wrap(getMergedError(errs), "failed deleting runtime policy")
	}
	if resp.StatusCode != 204 {
		var errorResponse ErrorResponse
		err := json.Unmarshal([]byte(body), &errorResponse)
		if err != nil {
			log.Printf("Failed to Unmarshal response Body to ErrorResponse. Body: %v.", body)
			return err
		}
		return fmt.Errorf("failed deleting runtime policy, status: %v. error message: %v", resp.Status, errorResponse.Message)
	}
	return nil
}
