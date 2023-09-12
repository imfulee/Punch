package hr_system

import (
	"errors"
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

const NUEIP_URL string = "https://portal.nueip.com/login"

type NUEIP struct {
	Company  string
	Username string
	Password string
}

func (nueip NUEIP) mustValid() error {
	if nueip.Company == "" || nueip.Password == "" || nueip.Username == "" {
		return errors.New("nueip fields are not passed in")
	}

	return nil
}

func (nueip NUEIP) login(page *rod.Page) error {
	// input company
	companyField, err := page.Element("input[name='inputCompany']")
	if err != nil {
		return fmt.Errorf("unable to select company name input box: %w", err)
	}
	if err := companyField.Input(nueip.Company); err != nil {
		return fmt.Errorf("unable to input company name: %w", err)
	}

	// input username
	usernameField, err := page.Element("input[name='inputID']")
	if err != nil {
		return fmt.Errorf("unable to select username input box: %w", err)
	}
	if err := usernameField.Input(nueip.Username); err != nil {
		return fmt.Errorf("unable to input username: %w", err)
	}

	// input password
	passwordField, err := page.Element("input[name='inputPassword']")
	if err != nil {
		return fmt.Errorf("unable to select password input box: %w", err)
	}
	if err := passwordField.Input(nueip.Password); err != nil {
		return fmt.Errorf("unable to input password: %w", err)
	}

	loginButton, err := page.Element("button.login-button")
	if err != nil {
		return fmt.Errorf("unable to select login button: %w", err)
	}
	if err := loginButton.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return fmt.Errorf("unable to click login button: %w", err)
	}

	return nil
}

func (nueip NUEIP) disableGeolocationPerm(page *rod.Page) (err error) {
	// deny geolocation permission
	_, err = page.Eval(`() => {
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

	return err
}

func (nueip NUEIP) punch(page *rod.Page, punchButtonSelector string, status PunchStatus) error {
	if !status.IsValid() {
		return errors.New("invalid punch status")
	}

	if punchButtonSelector == "" {
		return errors.New("punch button selector cannot be empty string")
	}

	// click the punch in/out button
	punchButton, err := page.Element(punchButtonSelector)
	if err != nil {
		return fmt.Errorf("unable to find punch %v button: %w", status, err)
	}
	if err := punchButton.Click(proto.InputMouseButtonLeft, 1); err != nil {
		return fmt.Errorf("unable to click punch %v button: %w", status, err)
	}

	// wait for button to be punched
	punchedButtonSelector := punchButtonSelector + ".punched"
	page.MustElement(punchedButtonSelector)

	return nil
}

func (nueip NUEIP) Punch(status PunchStatus) error {
	var err error

	if !status.IsValid() {
		return errors.New("wrong punch status")
	}

	if err := nueip.mustValid(); err != nil {
		return errors.New("NUEIP not yet configured")
	}

	// open browser
	path, foundPath := launcher.LookPath()
	if !foundPath {
		return errors.New("cannot find launcher path")
	}
	launcher, err := launcher.New().Bin(path).Launch()
	if err != nil {
		return errors.New("cannot launch launcher")
	}
	browser := rod.New().ControlURL(launcher)
	if err := browser.Connect(); err != nil {
		return fmt.Errorf("unable to connect to browser: %w", err)
	}
	defer func() {
		browserErr := browser.Close()
		if browserErr != nil {
			err = fmt.Errorf("cannot close browser: %w", err)
		}
	}()

	// redirect to NUEiP login page
	page, err := browser.Page(proto.TargetCreateTarget{
		URL: NUEIP_URL,
	})
	if err != nil {
		return fmt.Errorf("unable to go to page: %w", err)
	}

	// login to NUEiP
	if err := nueip.login(page); err != nil {
		return fmt.Errorf("unable to login: %w", err)
	}

	// redirect and wait for punch button to show
	var punchButtonSelector string
	if status == PunchIn {
		punchButtonSelector = "div.por-punch-clock__content--button > div.button-row.el-row > div:nth-child(1) > button.punch-btn"
	} else {
		punchButtonSelector = "div.por-punch-clock__content--button > div.button-row.el-row > div:nth-child(2) > button.punch-btn"
	}

	if _, err := page.Element(punchButtonSelector); err != nil {
		return err
	}

	// disable geolocation permission
	if err := nueip.disableGeolocationPerm(page); err != nil {
		return fmt.Errorf("unable to disable location permissions: %w", err)
	}

	// punch the clock in/out button
	if err := nueip.punch(page, punchButtonSelector, status); err != nil {
		return fmt.Errorf("unable to punch button: %w", err)
	}

	return err
}
