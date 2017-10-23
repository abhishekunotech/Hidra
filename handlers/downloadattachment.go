package handlers

import (
	"github.com/Unotechsoftware/Hydra/utils"
	"github.com/Unotechsoftware/Hydra/lerna"
	"net/http"
)

func callGetDownloadAttachment(username string, password string, fileid string, articleid string) []uint8{

	ConfObj := lerna.ReturnConfigObject()
	felicitybaseurl := ConfObj.Sub("components.otrs").GetString("url")
	felicityapiuri := ConfObj.Sub("components.otrs.apis.GetDownloadAttachment").GetString("uri")
	sessionIDString := callSessionDetails(username, password)
	url := felicitybaseurl + felicityapiuri + "?UserLogin=" + username + "&Password=" + password + "&FileID=" + fileid + "&ArticleID=" + articleid + "&SessionID=" + sessionIDString 
	
	return utils.MakeHTTPGetCall(url)

}

// This function is a handler that creates a GET API to download attachment based on Article ID and File ID. 
//
// **Business Logic**: Function takes as an input GET Parameters UserLogin, Password, File ID and Article ID and generate the response.
//
// Returns data as shown in examples
func (h *Handler) GetDownloadAttachment(w http.ResponseWriter, r *http.Request) {
	mapHttp := utils.RequestAbstractGet(r)
	var username string
	var password string
	var fileid string
	var articleid string

	for key, value := range mapHttp {
		if key == "UserLogin" {
			for _, valueStrg := range value {
				username = valueStrg
			}
		}
		if key == "FileID" { 
                        for _, valueStrg := range value {
                                fileid = valueStrg
                        }
                }
		if key == "ArticleID" { 
                        for _, valueStrg := range value {
                                articleid = valueStrg
                        }
                }

		if key == "Password" {
			for _, valueStrg := range value {
				password = valueStrg
			}
		}
	}
	utils.ResponseAbstract(callGetDownloadAttachment(username, password, fileid, articleid),w)
}
