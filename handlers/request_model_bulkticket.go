package handlers
// The following Type defines the Input JSON Body for Creating a Bulk Ticket Request
type BulkTicket_Request struct{
	UserLogin	string	`json:"UserLogin"`
	Password	string	`json:"Password"`
	Action		string	`json:"Action"`
	TicketIDs	[]string	`json:"TicketIDs"`
	MergeTo		string	`json:"MergeTo"`
	Subaction	string	`json:"Subaction"`
	Priority	string	`json:"Priority"`
	TypeID		int	`json:"TypeID"`
	StateID		int	`json:"StateID"`
	QueueID		int	`json:"QueueID"`
	Owner		string	`json:"Owner"`
	ResponsibleID	string	`json:"ResponsibleID"`
	Responsible	string	`json:"Responsible"`
	PriorityID	string	`json:"PriorityID"`
	Queue		string	`json:"Queue"`
	Subject		string	`json:"Subject"`
	Body		string	`json:"Body"`
	ArticleTypeID	string	`json:"ArticleTypeID"`
	ArticleType	string	`json:"ArticleType"`
	State		string	`json:"State"`
	MergeToSelection	string	`json:"MergeToSelection"`
	LinkTogether	string	`json:"LinkTogether"`
	EmailSubject	string	`json:"EmailSubject"`
	EmailBody	string	`json:"EmailBody"`
	EmailTimeUnits	string	`json:"EmailTimeUnits"`
	LinkTogetherParent	string	`json:"LinkTogetherParent"`
	Unlock		string	`json:"Unlock"`
	MergeToChecked	string	`json:"MergeToChecked"`
	MergeToOldestChecked	string	`json:"MergeToOldestChecked"`
	Year	string	`json:"Year"`
	Month	string	`json:"Month"`
	Day	string	`json:"Day"`
	Hour	string	`json:"Hour"`
	Minute	string	`json:"Minute"`
}
