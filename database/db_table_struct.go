package models

type Config struct {
	AutoRunScore string
}

type Logs struct {
	TNo              int
	SubmitLog        string
	UploadLog        string
	SetupLog         string
	WaitLog          string
	CompleteLog      string
	InfraLog         string
	ScoreLog         string
	InferenceLogPath string
	ScoreLogPath     string
}

type ResourceBoard struct {
	No             int
	NodeName       string
	GpuId          string
	IpAddress      string
	GpuInformation string
	Available      string
	Status         string
	TNo            int
	Id             string
	ImageTag       string
}

type Scoring struct {
	TNo                int
	Id              string
	ScoreEndDate    string
	Score           float64
	AnswerSheetPath string
}

type SessionLog struct {
	No         int
	Id         string
	IpAddress  string
	AccessDate string
}

type Submit struct {
	TNo                int
	TrackName          string
	Id                 string
	UseCreditFlag      string
	SubmitStartDate    string
	SubmitEndDate      string
	UploadStartDate    string
	UploadEndDate      string
	SetupStartDate     string
	SetupEndDate       string
	WaitStartDate      string
	WaitEndDate        string
	InferenceStartDate string
	InferenceEndDate   string
	InferenceStopDate  string
	InferenceErrorDate string
	ScoreStartDate     string
	ScoreEndDate       string
	ImageTag           string
	NodeName           string
	SubmitPath         string
	DockerfilePath     string
	EnvInfo            map[string]interface{}
	AnswerSheetPath    string
	Status             string
	Retries            int
	Progress           int
}

type Users struct {
	Uuid       int
	TrackName  string
	Id         string
	Password   string
	TeamName   string
	UserName   string
	GroupInfo  string
	GroupName  string
	Email      string
	CellPhone  string
	AdminFlag  string
	Credit     int
	Activation string
}