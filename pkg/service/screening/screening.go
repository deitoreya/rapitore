package screening

import (
	"context"
	"fmt"

	"github.com/deitoreya/rapitore/pkg/client/jquants"
	"golang.org/x/xerrors"
)

type Screening struct {
}

func NewScreening() *Screening {
	return &Screening{}
}

func (s *Screening) Exec(ctx context.Context) error {
	jquantClient := jquants.NewAPIClient(jquants.NewConfiguration())
	fmt.Println(jquantClient)
	req := jquantClient.DefaultAPI.TokenAuthUserPost(ctx)
	mailPass := buildMailPass("test", "pass")
	req.MailPassword(*mailPass)
	token, resp, err := jquantClient.DefaultAPI.TokenAuthUserPostExecute(req)
	if err != nil {
		return xerrors.Errorf("Error TokenAuthUserPostExecute: %w", err)
	}
	fmt.Println("auth user post", token, resp)
	return nil
}

func buildMailPass(mail, pass string) *jquants.MailPassword {
	mailPass := jquants.NewMailPassword()
	mailPass.Mailaddress = &mail
	mailPass.Password = &pass
	return mailPass
}
