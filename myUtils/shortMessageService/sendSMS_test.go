package shortMessageService

import "testing"

func TestSendSMS(t *testing.T) {
	filepath := `D:\test\sms.yaml`
	LoadSmsConfig(filepath)

	content := MessageContent{
		PhoneNumberSet:   []string{ChineseMobile("15383026353")},
		TemplateParamSet: []string{"张", "123456", "10"},
	}

	result, err := SendSMS(content)
	if err != nil {
		t.Errorf("SendSMS error: %v", err)
	} else {
		t.Logf("SendSMS success: %s", result)
	}
}
