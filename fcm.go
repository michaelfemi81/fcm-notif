package fcm
import (
	"errors"
	"net/http"
	"strings"
        "fmt"
	"io/ioutil"
	"encoding/json"
)
type Fcm struct{
	auth string
	to string
	title string
	body string
	data []byte

}
type Response1 struct{
	Mid string `json:"message_id"`
}

type Res struct{
	MessageId string `json:"message_id"`
}
type Response2 struct{
MultiCast string `json:"multicast_id"`
Success int `json:"success"`
Failure int `json:"failure"`
CanonicalIds int `json:"canonical_ids"`
Results []Res `json:"results"`

}
func(f *Fcm)Init (auth string ,to string ,title string ,body string,data []byte)(){
	f.auth=auth;
	f.to=to
	f.title=title
	f.body=body
        f.data=data

	return
}
func (f Fcm)SendNotif()(err error,resi []byte){
	if(f.auth==""){
		err=errors.New("App not Initialized properly");

	}else{

		url := "https://gcm-http.googleapis.com/gcm/send"

		payload := strings.NewReader("{\n  \"to\":"+f.to+",\n  \"data\": "+string(f.data)+",\n   \"notification\": {\n        \"title\":"+f.title+",\n        \"body\": "+f.body+"\n    }\n}")

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("authorization", "key="+f.auth)
		req.Header.Add("content-type", "application/json")


		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, err2 := ioutil.ReadAll(res.Body)
		if err2 != nil {
			panic(err.Error())
		}
		if(strings.HasPrefix((strings.ToLower(f.to)),"/topics/")){
                           fmt.Println(string(body))
			r, err3 := getRes1([]byte(body))
			if err3 != nil {
				panic(err.Error())
			}
			ret, err3 := json.Marshal(r)
			if err3 != nil {
				panic(err.Error())
			} else{
				resi=[]byte( ret);}
		}else{

			r, err3 := getRes2([]byte(body))
			if err3 != nil {
				panic(err.Error())
			}
			if(r.Success>1&&r.Failure<=0){
				ret, err3 := json.Marshal(r)
				if err3 != nil {
					panic(err.Error())
				} else{resi=[]byte( ret);}
			}else{
   resi=[]byte(`{error:"Something went Wrong"}`);
}
		}





}
return
}
func getRes1(body []byte) (*Response1, error) {
	var s = new(Response1)
	err := json.Unmarshal(body, &s)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getRes2(body []byte) (*Response2, error) {
	var s = new(Response2)
	err := json.Unmarshal(body, &s)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return s, err
}