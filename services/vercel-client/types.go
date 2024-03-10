package vercelService

import (
	aw "github.com/deanishe/awgo"
	"github.com/dustin/go-humanize"
	"time"
)

type Result struct {
	InspectUrl string `json:"InspectUrl"`
}

type GetDeploymentsResponse struct {
	Deployments []Deployment `json:"deployments"`
	Pagination  Pagination   `json:"pagination"`
}
type Creator struct {
	UID         string `json:"uid"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	GitlabLogin string `json:"gitlabLogin"`
}
type Meta struct {
	GitlabCommitAuthorName     string `json:"gitlabCommitAuthorName"`
	GitlabCommitMessage        string `json:"gitlabCommitMessage"`
	GitlabCommitRef            string `json:"gitlabCommitRef"`
	GitlabCommitSha            string `json:"gitlabCommitSha"`
	GitlabDeployment           string `json:"gitlabDeployment"`
	GitlabProjectName          string `json:"gitlabProjectName"`
	GitlabProjectNamespace     string `json:"gitlabProjectNamespace"`
	GitlabProjectPath          string `json:"gitlabProjectPath"`
	GitlabProjectRepo          string `json:"gitlabProjectRepo"`
	GitlabNamespaceKind        string `json:"gitlabNamespaceKind"`
	GitlabProjectID            string `json:"gitlabProjectId"`
	GitlabCommitRawAuthorEmail string `json:"gitlabCommitRawAuthorEmail"`
	GitlabProjectVisibility    string `json:"gitlabProjectVisibility"`
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
	BranchAlias                string `json:"branchAlias"`
}
type ProjectSettings struct {
	CommandForIgnoringBuildStep any `json:"commandForIgnoringBuildStep"`
}
type Meta0 struct {
	GitlabCommitAuthorName     string `json:"gitlabCommitAuthorName"`
	GitlabCommitMessage        string `json:"gitlabCommitMessage"`
	GitlabCommitRef            string `json:"gitlabCommitRef"`
	GitlabCommitSha            string `json:"gitlabCommitSha"`
	GitlabDeployment           string `json:"gitlabDeployment"`
	GitlabProjectName          string `json:"gitlabProjectName"`
	GitlabProjectNamespace     string `json:"gitlabProjectNamespace"`
	GitlabProjectPath          string `json:"gitlabProjectPath"`
	GitlabProjectRepo          string `json:"gitlabProjectRepo"`
	GitlabNamespaceKind        string `json:"gitlabNamespaceKind"`
	GitlabProjectID            string `json:"gitlabProjectId"`
	GitlabCommitRawAuthorEmail string `json:"gitlabCommitRawAuthorEmail"`
	GitlabProjectVisibility    string `json:"gitlabProjectVisibility"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabMrID                 string `json:"gitlabMrId"`
	BranchAlias                string `json:"branchAlias"`
}
type Meta1 struct {
	GitlabCommitAuthorName     string `json:"gitlabCommitAuthorName"`
	GitlabCommitMessage        string `json:"gitlabCommitMessage"`
	GitlabCommitRef            string `json:"gitlabCommitRef"`
	GitlabCommitSha            string `json:"gitlabCommitSha"`
	GitlabDeployment           string `json:"gitlabDeployment"`
	GitlabProjectName          string `json:"gitlabProjectName"`
	GitlabProjectNamespace     string `json:"gitlabProjectNamespace"`
	GitlabProjectPath          string `json:"gitlabProjectPath"`
	GitlabProjectRepo          string `json:"gitlabProjectRepo"`
	GitlabNamespaceKind        string `json:"gitlabNamespaceKind"`
	GitlabProjectID            string `json:"gitlabProjectId"`
	GitlabCommitRawAuthorEmail string `json:"gitlabCommitRawAuthorEmail"`
	GitlabProjectVisibility    string `json:"gitlabProjectVisibility"`
	GitlabMrID                 string `json:"gitlabMrId"`
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
	BranchAlias                string `json:"branchAlias"`
}
type Meta2 struct {
	GitlabCommitAuthorName     string `json:"gitlabCommitAuthorName"`
	GitlabCommitMessage        string `json:"gitlabCommitMessage"`
	GitlabCommitRef            string `json:"gitlabCommitRef"`
	GitlabCommitSha            string `json:"gitlabCommitSha"`
	GitlabDeployment           string `json:"gitlabDeployment"`
	GitlabProjectName          string `json:"gitlabProjectName"`
	GitlabProjectNamespace     string `json:"gitlabProjectNamespace"`
	GitlabProjectPath          string `json:"gitlabProjectPath"`
	GitlabProjectRepo          string `json:"gitlabProjectRepo"`
	GitlabNamespaceKind        string `json:"gitlabNamespaceKind"`
	GitlabProjectID            string `json:"gitlabProjectId"`
	GitlabCommitRawAuthorEmail string `json:"gitlabCommitRawAuthorEmail"`
	GitlabProjectVisibility    string `json:"gitlabProjectVisibility"`
	GitlabMrID                 string `json:"gitlabMrId"`
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
	BranchAlias                string `json:"branchAlias"`
}
type Deployment struct {
	UID                 string          `json:"uid"`
	Name                string          `json:"name"`
	URL                 string          `json:"url"`
	Created             int64           `json:"created"`
	Source              string          `json:"source"`
	State               string          `json:"state"`
	ReadyState          string          `json:"readyState"`
	ReadySubstate       string          `json:"readySubstate,omitempty"`
	Type                string          `json:"type"`
	Creator             Creator         `json:"creator"`
	InspectorURL        string          `json:"inspectorUrl"`
	Meta                Meta            `json:"meta"`
	Target              string          `json:"target"`
	AliasError          any             `json:"aliasError,omitempty"`
	AliasAssigned       int64           `json:"aliasAssigned"`
	IsRollbackCandidate bool            `json:"isRollbackCandidate"`
	CreatedAt           int64           `json:"createdAt"`
	BuildingAt          int64           `json:"buildingAt"`
	Ready               int64           `json:"ready"`
	ProjectSettings     ProjectSettings `json:"projectSettings"`
}
type Pagination struct {
	Count int   `json:"count"`
	Next  int64 `json:"next"`
	Prev  int64 `json:"prev"`
}

type Deployments []Deployment

func (d Deployments) SendAsAlfredFilter() {
	wf := aw.New()

	alfredItems := make([]*aw.Item, 0, len(d))

	statusEmoji := map[string]string{
		"READY":        "🟢",
		"ERROR":        "🔴",
		"BUILDING":     "🔵",
		"QUEUED":       "🟡",
		"INITIALIZING": "🟡",
		"DEPLOYING":    "🟡",
		"UPLOADING":    "🟡",
		"CANCELED":     "⚪️",
	}

	for _, deployment := range d {
		localCreatedAt := time.Unix(deployment.Created/1000, 0).Local()

		deploymentStatus := statusEmoji[deployment.ReadyState]
		alfredItems = append(alfredItems,
			wf.NewItem(deploymentStatus+" "+deployment.Meta.GitlabCommitMessage).
				Subtitle("Created at: "+humanize.Time(localCreatedAt)+", By: "+deployment.Meta.GitlabCommitAuthorName).
				Arg(deployment.InspectorURL).
				Var("url", deployment.URL).
				Var("inspectorUrl", deployment.InspectorURL).
				Quicklook(deployment.URL).
				Valid(true))
	}

	wf.SendFeedback()
	return
}

type GetProjectsResponse struct {
	Projects   []Project  `json:"projects"`
	Pagination Pagination `json:"pagination"`
}
type SpeedInsights struct {
	ID      string `json:"id"`
	HasData bool   `json:"hasData"`
}
type Crons struct {
	EnabledAt    int64  `json:"enabledAt"`
	DisabledAt   any    `json:"disabledAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	DeploymentID string `json:"deploymentId"`
	Definitions  []any  `json:"definitions"`
}
type Env struct {
	Key             string   `json:"key"`
	Target          []string `json:"target"`
	ConfigurationID any      `json:"configurationId"`
	CreatedAt       int64    `json:"createdAt"`
	UpdatedAt       int64    `json:"updatedAt"`
	CreatedBy       string   `json:"createdBy"`
	UpdatedBy       string   `json:"updatedBy"`
	ID              string   `json:"id"`
	Type            string   `json:"type"`
	Value           string   `json:"value"`
	GitBranch       string   `json:"gitBranch,omitempty"`
}
type WebAnalytics struct {
	ID string `json:"id"`
}
type Link struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	Sourceless               bool   `json:"sourceless"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}

type LatestDeployments struct {
	Alias                  []string `json:"alias"`
	AliasAssigned          int64    `json:"aliasAssigned"`
	AliasError             any      `json:"aliasError"`
	AutomaticAliases       []string `json:"automaticAliases"`
	Builds                 []any    `json:"builds"`
	CreatedAt              int64    `json:"createdAt"`
	CreatedIn              string   `json:"createdIn"`
	Creator                Creator  `json:"creator"`
	DeploymentHostname     string   `json:"deploymentHostname"`
	Forced                 bool     `json:"forced"`
	ID                     string   `json:"id"`
	Meta                   Meta     `json:"meta"`
	Name                   string   `json:"name"`
	Plan                   string   `json:"plan"`
	Private                bool     `json:"private"`
	ReadyState             string   `json:"readyState"`
	ReadySubstate          string   `json:"readySubstate"`
	Target                 string   `json:"target"`
	TeamID                 string   `json:"teamId"`
	Type                   string   `json:"type"`
	URL                    string   `json:"url"`
	UserID                 string   `json:"userId"`
	WithCache              bool     `json:"withCache"`
	BuildingAt             int64    `json:"buildingAt"`
	ReadyAt                int64    `json:"readyAt"`
	PreviewCommentsEnabled bool     `json:"previewCommentsEnabled"`
}
type Production struct {
	Alias                  []string `json:"alias"`
	AliasAssigned          int64    `json:"aliasAssigned"`
	AliasError             any      `json:"aliasError"`
	AutomaticAliases       []string `json:"automaticAliases"`
	Builds                 []any    `json:"builds"`
	CreatedAt              int64    `json:"createdAt"`
	CreatedIn              string   `json:"createdIn"`
	Creator                Creator  `json:"creator"`
	DeploymentHostname     string   `json:"deploymentHostname"`
	Forced                 bool     `json:"forced"`
	ID                     string   `json:"id"`
	Meta                   Meta     `json:"meta"`
	Name                   string   `json:"name"`
	Plan                   string   `json:"plan"`
	Private                bool     `json:"private"`
	ReadyState             string   `json:"readyState"`
	ReadySubstate          string   `json:"readySubstate"`
	Target                 string   `json:"target"`
	TeamID                 string   `json:"teamId"`
	Type                   string   `json:"type"`
	URL                    string   `json:"url"`
	UserID                 string   `json:"userId"`
	WithCache              bool     `json:"withCache"`
	BuildingAt             int64    `json:"buildingAt"`
	ReadyAt                int64    `json:"readyAt"`
	PreviewCommentsEnabled bool     `json:"previewCommentsEnabled"`
}
type Targets struct {
	Production Production `json:"production"`
}
type Link0 struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}
type Link1 struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}
type SsoProtection struct {
	DeploymentType string `json:"deploymentType"`
}
type GitComments struct {
	OnCommit      bool `json:"onCommit"`
	OnPullRequest bool `json:"onPullRequest"`
}
type Analytics struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	CanceledAt any    `json:"canceledAt"`
	PaidAt     int64  `json:"paidAt"`
}
type SpeedInsights0 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	PaidAt     int64  `json:"paidAt"`
	HasData    bool   `json:"hasData"`
}
type WebAnalytics0 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int64  `json:"disabledAt"`
	CanceledAt int64  `json:"canceledAt"`
}
type Analytics0 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	PaidAt     int64  `json:"paidAt"`
}
type SpeedInsights1 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	PaidAt     int64  `json:"paidAt"`
	HasData    bool   `json:"hasData"`
}
type Link2 struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}
type Analytics1 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	PaidAt     int64  `json:"paidAt"`
}
type SpeedInsights2 struct {
	ID         string `json:"id"`
	EnabledAt  int64  `json:"enabledAt"`
	DisabledAt int    `json:"disabledAt"`
	PaidAt     int64  `json:"paidAt"`
	HasData    bool   `json:"hasData"`
}
type Link3 struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}
type Link4 struct {
	Type                     string `json:"type"`
	ProjectID                string `json:"projectId"`
	ProjectName              string `json:"projectName"`
	ProjectNameWithNamespace string `json:"projectNameWithNamespace"`
	ProjectNamespace         string `json:"projectNamespace"`
	ProjectURL               string `json:"projectUrl"`
	GitCredentialID          string `json:"gitCredentialId"`
	ProductionBranch         string `json:"productionBranch"`
	CreatedAt                int64  `json:"createdAt"`
	UpdatedAt                int64  `json:"updatedAt"`
	DeployHooks              []any  `json:"deployHooks"`
}
type Project struct {
	AccountID                        string              `json:"accountId"`
	SpeedInsights                    SpeedInsights       `json:"speedInsights,omitempty"`
	AutoExposeSystemEnvs             bool                `json:"autoExposeSystemEnvs"`
	BuildCommand                     string              `json:"buildCommand"`
	CommandForIgnoringBuildStep      string              `json:"commandForIgnoringBuildStep,omitempty"`
	CreatedAt                        int64               `json:"createdAt"`
	Crons                            Crons               `json:"crons,omitempty"`
	DevCommand                       any                 `json:"devCommand"`
	DirectoryListing                 bool                `json:"directoryListing"`
	Env                              []Env               `json:"env"`
	Framework                        string              `json:"framework"`
	GitForkProtection                bool                `json:"gitForkProtection,omitempty"`
	ID                               string              `json:"id"`
	InstallCommand                   string              `json:"installCommand"`
	LastRollbackTarget               any                 `json:"lastRollbackTarget"`
	LastAliasRequest                 any                 `json:"lastAliasRequest"`
	Name                             string              `json:"name"`
	NodeVersion                      string              `json:"nodeVersion"`
	OutputDirectory                  any                 `json:"outputDirectory"`
	PublicSource                     any                 `json:"publicSource"`
	RootDirectory                    any                 `json:"rootDirectory"`
	ServerlessFunctionRegion         string              `json:"serverlessFunctionRegion"`
	SourceFilesOutsideRootDirectory  bool                `json:"sourceFilesOutsideRootDirectory"`
	UpdatedAt                        int64               `json:"updatedAt"`
	Live                             bool                `json:"live,omitempty"`
	WebAnalytics                     WebAnalytics        `json:"webAnalytics,omitempty"`
	Link                             Link                `json:"link,omitempty"`
	LatestDeployments                []LatestDeployments `json:"latestDeployments"`
	Targets                          Targets             `json:"targets"`
	Link0                            Link0               `json:"link,omitempty"`
	Link1                            Link1               `json:"link,omitempty"`
	TransferStartedAt                int64               `json:"transferStartedAt,omitempty"`
	TransferCompletedAt              int64               `json:"transferCompletedAt,omitempty"`
	TransferredFromAccountID         string              `json:"transferredFromAccountId,omitempty"`
	AutoAssignCustomDomains          bool                `json:"autoAssignCustomDomains,omitempty"`
	AutoAssignCustomDomainsUpdatedBy string              `json:"autoAssignCustomDomainsUpdatedBy,omitempty"`
	GitLFS                           bool                `json:"gitLFS,omitempty"`
	SsoProtection                    SsoProtection       `json:"ssoProtection,omitempty"`
	GitComments                      GitComments         `json:"gitComments,omitempty"`
	Analytics                        Analytics           `json:"analytics,omitempty"`
	SpeedInsights0                   SpeedInsights0      `json:"speedInsights,omitempty"`
	WebAnalytics0                    WebAnalytics0       `json:"webAnalytics,omitempty"`
	Analytics0                       Analytics0          `json:"analytics,omitempty"`
	SpeedInsights1                   SpeedInsights1      `json:"speedInsights,omitempty"`
	Link2                            Link2               `json:"link,omitempty"`
	Analytics1                       Analytics1          `json:"analytics,omitempty"`
	SpeedInsights2                   SpeedInsights2      `json:"speedInsights,omitempty"`
	PasswordProtection               any                 `json:"passwordProtection,omitempty"`
	Link3                            Link3               `json:"link,omitempty"`
	Link4                            Link4               `json:"link,omitempty"`
}

type Projects []Project

func (p Projects) SendAsAlfredFilter() {
	wf := aw.New()

	alfredItems := make([]*aw.Item, 0, len(p))

	for _, project := range p {
		localUpdatedAt := time.Unix(project.UpdatedAt/1000, 0).Local()

		alfredItems = append(alfredItems, wf.NewItem(project.Name).
			Subtitle("Updated: "+humanize.Time(localUpdatedAt)).
			Arg().
			Var("project_id", project.ID).
			Var("project_name", project.Name).
			Var("project_url", project.Link.ProjectURL).
			Quicklook(project.Targets.Production.URL).
			Valid(true))
	}

	wf.SendFeedback()
	return
}
