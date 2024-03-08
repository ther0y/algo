package vercelclient

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
	GitlabMrID                 string `json:"gitlabMrId"`
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
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
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
	GitlabCommitAuthorLogin    string `json:"gitlabCommitAuthorLogin"`
	GitlabCommitAuthorAvatar   string `json:"gitlabCommitAuthorAvatar"`
	BranchAlias                string `json:"branchAlias"`
}
type Meta3 struct {
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
type Meta4 struct {
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
type Meta5 struct {
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
type Meta6 struct {
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
type Meta7 struct {
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
	Action                     string `json:"action"`
	OriginalDeploymentID       string `json:"originalDeploymentId"`
}
type Meta8 struct {
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
type Meta9 struct {
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
type Meta10 struct {
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
type Meta11 struct {
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
type Meta12 struct {
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

type Deployment struct {
	UID                 string          `json:"uid"`
	Name                string          `json:"name"`
	URL                 string          `json:"url"`
	Created             int64           `json:"created"`
	Source              string          `json:"source,omitempty"`
	State               string          `json:"state"`
	ReadyState          string          `json:"readyState"`
	ReadySubstate       string          `json:"readySubstate,omitempty"`
	Type                string          `json:"type"`
	Creator             Creator         `json:"creator"`
	InspectorURL        string          `json:"inspectorUrl"`
	Meta                Meta            `json:"meta,omitempty"`
	Target              any             `json:"target"`
	AliasError          any             `json:"aliasError,omitempty"`
	AliasAssigned       int64           `json:"aliasAssigned"`
	IsRollbackCandidate bool            `json:"isRollbackCandidate"`
	CreatedAt           int64           `json:"createdAt"`
	BuildingAt          int64           `json:"buildingAt"`
	Ready               int64           `json:"ready"`
	ProjectSettings     ProjectSettings `json:"projectSettings"`
	Meta0               Meta0           `json:"meta,omitempty"`
	Meta1               Meta1           `json:"meta,omitempty"`
	Meta2               Meta2           `json:"meta,omitempty"`
	Meta3               Meta3           `json:"meta,omitempty"`
	Meta4               Meta4           `json:"meta,omitempty"`
	Meta5               Meta5           `json:"meta,omitempty"`
	Meta6               Meta6           `json:"meta,omitempty"`
	Meta7               Meta7           `json:"meta,omitempty"`
	Meta8               Meta8           `json:"meta,omitempty"`
	Meta9               Meta9           `json:"meta,omitempty"`
	Meta10              Meta10          `json:"meta,omitempty"`
	Meta11              Meta11          `json:"meta,omitempty"`
	Meta12              Meta12          `json:"meta,omitempty"`
}

type Pagination struct {
	Count int   `json:"count"`
	Next  int64 `json:"next"`
	Prev  int64 `json:"prev"`
}
