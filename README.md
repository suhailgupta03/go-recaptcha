# go-recaptcha

Go utility to verify google captcha requests. Extensible towards other captcha providers
# Install
```shell
go get -u github.com/suhailgupta03/go-recaptcha@latest
```

# Example usage to directly use recaptcha
```go
import (
	"fmt"
	"github.com/suhailgupta03/go-recaptcha/models"
	"github.com/suhailgupta03/go-recaptcha/recaptcha"
)

func main() {
	params := recaptcha.New("secretKey", "secretResponseToken", "")
	resp, err := params.Verify()
	if err != nil {
		fmt.Println(err)
	} else {
		recaptchaResp := resp.(*models.RecaptchaResponse)
		fmt.Println(recaptchaResp.Success)
		fmt.Println(recaptchaResp.ErrorCodes)
	}
}
```