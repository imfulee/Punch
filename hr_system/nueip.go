package hr_system

import (
	"errors"

	"github.com/go-rod/rod"
)

const NUEIP_URL string = "https://portal.nueip.com/login"

type NUEIP struct {
	Company  string
	Username string
	Password string
}

func (nueip NUEIP) Punch(status PunchStatus) (errs []error) {
	if !status.IsValid() {
		errs = append(errs, errors.New("wrong punch status"))
		return
	}

	// go to the login page
	browser := rod.New().MustConnect()
	defer func() {
		browserCloseErr := browser.Close()
		if browserCloseErr != nil {
			errs = append(errs, browserCloseErr)
		}
	}()

	page := browser.MustPage(NUEIP_URL)
	page.MustWindowMaximize()
	page.MustWaitStable()

	// input all the fields
	page.MustElement("input[name='inputCompany']").MustInput(nueip.Company)
	page.MustElement("input[name='inputID']").MustInput(nueip.Username)
	page.MustElement("input[name='inputPassword']").MustInput(nueip.Password)
	page.MustElement("button.login-button").MustClick()

	// redirect
	punchButton := ""
	if status == PunchIn {
		punchButton = "div.por-punch-clock__content--button > div.button-row.el-row > div:nth-child(1) > button.punch-btn"
	} else {
		punchButton = "div.por-punch-clock__content--button > div.button-row.el-row > div:nth-child(2) > button.punch-btn"
	}
	page.MustElement(punchButton)

	// deny geolocation permission
	page.MustEval(`() => {
        // Override the navigator.geolocation object
        Object.defineProperty(navigator, 'geolocation', {
            value: {
                getCurrentPosition: (successCallback, errorCallback) => {
                    errorCallback(new Error('Geolocation is disabled'));
                },
                watchPosition: () => {},
                clearWatch: () => {},
            },
            configurable: true,
        });
    }`)

	// punch
	page.MustElement(punchButton).MustClick()

	// wait for button to be punched
	punchedButton := punchButton + ".punched"
	page.MustElement(punchedButton)

	return errs
}
