package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/micro/go-log"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	log.Logf("example call 1")
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Logf("example call decode err!")
		http.Error(w, err.Error(), 500)
		return
	}

	name,ok := request["name"]
	if ok {

	}
	log.Logf("example call 2 : %s",name)

	time.Sleep(time.Millisecond *5)
	//time.Sleep(time.Second * 15)

	log.Logf("example call 3 : %s",request)

	//client.NewClient(
	//	client.Selector(cache.NewSelector(cache.TTL(time.Second *30))),
	//	client.Transport(tcp.NewTransport()),
	//)
	//
	//// call the backend service
	//userClient := user.NewUserService("go.micro.srv.auth", client.DefaultClient)
	//rsp, err := userClient.GetUserLogin(context.TODO(), &user.ReqLogin{
	//	Nickname: "Hobo",
	//	Pwd:      "pwd",
	//})
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	// we want to augment the response
	response := map[string]interface{}{
		"user": "test",
		"ref":  time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Logf("example call encode err!")
		http.Error(w, err.Error(), 500)
		return
	}
}
