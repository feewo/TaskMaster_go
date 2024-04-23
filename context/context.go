package context

// переименовать
import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}
type Error struct {
	Message string
}

func (c *Context) Error(status int, error string) {
	c.Response.WriteHeader(status)

	em := Error{Message: error}
	marsh, _ := json.Marshal(em)
	c.Response.Write([]byte(marsh))
}
func (c *Context) Print(data interface{}) {
	c.Response.Header().Set("Content-Type", "application/json")

	marsh, _ := json.Marshal(data)
	c.Response.Write([]byte(marsh))
}

// decoder
// func ToStruct[t any](ctx Context) (dt T, err error) {
// 	decoder := json.NewDecoder(ctx.Request.Body)
// 	err = decoder.Decode(&dt)
// 	return dt, err
// }
