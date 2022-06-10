package email

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	sender, err := NewSender("smtp.gmail.com", "operator@dcrn.xyz", "",
		"operator@dcrn.xyz", true, nil, false)
	if err != nil {
		t.Errorf("err:%s\n", err)
		return
	}

	err = sender.sendMail("wei.zeng01@antpool.com", "start email", "regestis email")
	if err != nil {
		t.Errorf("err:%s\n", err)
		return
	}
}
