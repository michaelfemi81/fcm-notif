# fcm-notif
Firebase Cloud Messaging Library for sending push notification 
### INSTALLATION

    go get github.com/michaelfemi81/fcm-notif
### API DOCUMENTATION
```
package main
import(
	"github.com/michaelfemi81/fcm-notif"
	"fmt"
)
func main() {
	fcm:=fcm.Fcm{};
	//fcm.Init("api-key","/topics/news","Hello","Welcome to my app",[]byte(`{
	fcm.Init("key=api-key","client-id","Hello","Welcome to my app",[]byte(`{
    "message": "This is a GCM Topic Message!"
   }`));
	err,res:=fcm.SendNotif()
	if(err!=nil){
		fmt.Println(err)

	}else{
		fmt.Println(string(res));

	}
}
```