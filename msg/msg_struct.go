package agcmsg

type TestMsg struct {
	Msg_Type	string	`json:"msg_type"`
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Email		string	`json:"email"`
}

type RecvMessage struct {
	Status			string	`json:"status" binding:"required"`
}

type EnvInfo struct {
	OS_Info		string	`json:"os_version"`
	Cuda_Info	string	`json:"cuda_version"`
	Python_Info	string	`json:"python_version"`
	User_Setup	string	`json:"user_setup"`
	STT_Info	string	`json:"stt_version"`
}

// [FrontWeb to Communicator] Message Type 1.1
type SubmitStart struct {
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	EnvInfo			EnvInfo	`json:"env_info"`
	SubmitStartDate	string 	`json:"submit_start_date"`
	SubmitPath		string	`json:"submit_path"`
	Status			string	`json:"status"`
}

// [FrontWeb to Communicator] Message Type 1.2
type SubmitStop struct {
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	SubmitEndDate	string 	`json:"submit_end_date"`
	SubmitPath		string	`json:"submit_path"`
	Status			string	`json:"status"`
	Log				string	`json:"log"`
}

// [FrontWeb to Communicator] Message Type 1.3
type SubmitComplete struct {
	Tno				int		`json:"t_no"`
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	SubmitStartDate	string 	`json:"submit_start_date"`
	SubmitEndDate	string 	`json:"submit_end_date"`
	SubmitPath		string	`json:"submit_path"`
	Status			string	`json:"status"`
}

// [FrontWeb to Communicator] Message Type 1.4
type UserAccessInfo struct {
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	IpAddress		string 	`json:"ip_address"`
	AccessDate		string 	`json:"access_date"`
	Status			string	`json:"status"`
}

// [FrontWeb to Communicator] Message Type 1.5
type SubmitError struct {
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	SubmitStartDate	string 	`json:"submit_start_date"`
	SubmitEndDate	string 	`json:"submit_end_date"`
	SubmitPath		string	`json:"submit_path"`
	Status			string	`json:"status"`
	Log				string	`json:"log"`
}

// [FrontWeb to Communicator] Message Type 1.6
type InferenceWaitStart struct {
	MsgType			string	`json:"msg_type"`
	Tno				int		`json:"t_no"`
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	ImageTag		string	`json:"image_tag"`
	Status			string	`json:"status"`
}

// [FrontWeb to Communicator] Message Type 1.7
type ProcessingStop struct {
	Tno				int		`json:"t_no"`
	Id				string	`json:"id"`
	TrackName		string	`json:"track_name"`
	Status			string	`json:"status"`
	Log				string	`json:"log"`
}

//////////////////////////////////////////////////////////
// [FrontWeb to Communicator] New AGC2022 test env message
//////////////////////////////////////////////////////////

// type AnswerSheetTest struct {
// 	No				int		`json:"no"`
// 	Value			string	`json:"value"`
// }

type AnswerSheetTest struct {
	Answer			[]string	`json:"answer"`
}


type TestMessage struct {
	TeamName		string			`json:"team"`
	AnswerSheet		AnswerSheetTest	`json:"answer_sheet"`
	Status			string			`json:"status"`
}

//////////////////////////////////////////////////////////

//////////////////////////////////////////////////////////
// strapi test message structure
//////////////////////////////////////////////////////////
type Strapi_AGC2022 struct {
	RecvMsg			string		`json:"recvmsg"`
	Uid				int			`json:"uid"`
	Date			string		`json:"date"`
}