package main

import(
	"net/http"
	"fmt"
	"encoding/json"
)

// 主函数入口
func main(){
	fmt.Print("interesting start!\n")
	http.HandleFunc("/getbook", Index)
	fmt.Print(http.ListenAndServe("0.0.0.0:1766", nil))

}
func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	fmt.Fprint(w,string(reParseJson(interesting())))
}
// struct -> json
func reParseJson(v interface{}) []byte{ 
    textbyte,err := json.Marshal(v)
    if err !=nil {
        panic(err)
    }
    return textbyte
}
