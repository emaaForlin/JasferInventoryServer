package views
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)


type View struct {
	l 		*log.Logger
	file 	string
	data	interface{}
}


func NewView(l *log.Logger, file string, data interface{}) View {
	return View{l, file, data}
}

func (v *View) Render(c *gin.Context) {
	c.HTML(http.StatusOK, v.file, gin.H{
		"data": v.data,
	})

}