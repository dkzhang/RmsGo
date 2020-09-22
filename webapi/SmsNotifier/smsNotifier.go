package SmsNotifier

import (
	"fmt"
	"github.com/dkzhang/RmsGo/myUtils/logMap"
	"github.com/dkzhang/RmsGo/myUtils/shortMessageService"
	"github.com/dkzhang/RmsGo/webapi/dataInfra/userDM"
	"github.com/dkzhang/RmsGo/webapi/model/user"
	"github.com/sirupsen/logrus"
	"time"
)

type Notifier struct {
	theUserDM     userDM.UserDM
	TheSmsService shortMessageService.SmsService
	TheLogMap     logMap.LogMap
}

func (nt Notifier) Notify(ids []int, msg []string) (err error) {
	for _, userID := range ids {
		userInfo, err := nt.theUserDM.QueryUserByID(userID)
		if err != nil {
			return fmt.Errorf("theUserDM.QueryUserByID query userID <%d> error: %v", userID, err)
		}

		msg[0] = userInfo.ChineseName

		for retryTimes := 0; retryTimes < 3; retryTimes++ {
			resp, err := nt.TheSmsService.SendSMS(shortMessageService.MessageContent{
				PhoneNumberSet:   []string{shortMessageService.ChineseMobile(userInfo.Mobile)},
				TemplateParamSet: msg,
				TemplateType:     shortMessageService.TemplatePwd,
			})

			nt.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
				"UserID": userInfo.UserID,
				"Mobile": userInfo.Mobile,
				"msg":    msg,
				"resp":   resp,
				"err":    err,
			}).Info("NotifyProjectChief success.")

			if err == nil {
				break
			}

			time.Sleep(time.Second * 10)
		}
	}
	return nil
}

func (nt Notifier) NotifyProjectChief(userID int, msg string) (err error) {
	userInfo, err := nt.theUserDM.QueryUserByID(userID)
	if err != nil {
		return fmt.Errorf("theUserDM.QueryUserByID query userID <%d> error: %v", userID, err)
	}

	resp, err := nt.TheSmsService.SendSMS(shortMessageService.MessageContent{
		PhoneNumberSet:   []string{shortMessageService.ChineseMobile(userInfo.Mobile)},
		TemplateParamSet: []string{userInfo.ChineseName, msg},
		TemplateType:     shortMessageService.TemplatePwd,
	})
	if err != nil {
		return fmt.Errorf("TheSmsService.SendSMS error: %v", err)
	}

	nt.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserID": userInfo.UserID,
		"Mobile": userInfo.Mobile,
		"msg":    msg,
		"resp":   resp,
	}).Info("NotifyProjectChief success.")

	return nil
}

func (nt Notifier) NotifyApprover(dc string, msg string) (err error) {
	users := nt.theUserDM.QueryUserByFilter(func(userInfo user.UserInfo) bool {
		return userInfo.DepartmentCode == dc && userInfo.Role == user.RoleApprover
	})

	if len(users) != 1 {
		return fmt.Errorf("theUserDM query department <%s> approver info error, expected 1 but got %d", dc, len(users))
	}

	approverInfo := users[0]

	resp, err := nt.TheSmsService.SendSMS(shortMessageService.MessageContent{
		PhoneNumberSet:   []string{shortMessageService.ChineseMobile(approverInfo.Mobile)},
		TemplateParamSet: []string{approverInfo.ChineseName, msg},
		TemplateType:     shortMessageService.TemplatePwd,
	})
	if err != nil {
		return fmt.Errorf("TheSmsService.SendSMS error: %v", err)
	}

	nt.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
		"UserID": approverInfo.UserID,
		"Mobile": approverInfo.Mobile,
		"msg":    msg,
		"resp":   resp,
	}).Info("NotifyApprover success.")

	return nil
}

func (nt Notifier) NotifyController(msg string) (err error) {
	users := nt.theUserDM.QueryUserByFilter(func(userInfo user.UserInfo) bool {
		return userInfo.Role == user.RoleController
	})

	if len(users) < 1 {
		return fmt.Errorf("theUserDM query controllers info error, expected at least 1")
	}

	for _, controllerInfo := range users {
		resp, err := nt.TheSmsService.SendSMS(shortMessageService.MessageContent{
			PhoneNumberSet:   []string{shortMessageService.ChineseMobile(controllerInfo.Mobile)},
			TemplateParamSet: []string{controllerInfo.ChineseName, msg},
			TemplateType:     shortMessageService.TemplatePwd,
		})
		if err != nil {
			return fmt.Errorf("TheSmsService.SendSMS error: %v", err)
		}

		nt.TheLogMap.Log(logMap.NORMAL).WithFields(logrus.Fields{
			"UserID": controllerInfo.UserID,
			"Mobile": controllerInfo.Mobile,
			"msg":    msg,
			"resp":   resp,
		}).Info("NotifyController success.")
	}

	return nil
}
