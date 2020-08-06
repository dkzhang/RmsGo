package SmsNotifier

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/shortMessageService"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDM"
)

type Notifier struct {
	theUserDM     userDM.UserDM
	TheSmsService shortMessageService.SmsService
}

func (nt Notifier) NotifyProjectChief(userID int, msg string) (err error) {

	return fmt.Errorf("UNDO not yet accomplied")
}

func (nt Notifier) NotifyApprover(dc int, msg string) (err error) {
	return fmt.Errorf("UNDO not yet accomplied")
}

func (nt Notifier) NotifyController(msg string) (err error) {
	return fmt.Errorf("UNDO not yet accomplied")
}
