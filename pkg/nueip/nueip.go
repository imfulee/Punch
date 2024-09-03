package nueip

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
	var rtnErr error = nil

	if nueip.Company == "" {
		rtnErr = errors.Join(rtnErr, errors.New("nueip company is empty string"))
	}
	if nueip.Password == "" {
		rtnErr = errors.Join(rtnErr, errors.New("nueip password is empty string"))
	}
	if nueip.Username == "" {
		rtnErr = errors.Join(rtnErr, errors.New("nueip username is empty string"))
	}

	return rtnErr
}

func (nueip NUEIP) login(page *rod.Page, waitForElement string) error {
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

	if _, err = page.Timeout(DefaultTimeout).Element(waitForElement); err != nil {
		return errors.Join(errors.New("unable to redirect in 10 seconds"), err)
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

func (nueip NUEIP) punchClock(page *rod.Page, punchButtonSelector string, status PunchStatus) error {
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
	punchedButtonSelector := punchButtonSelector + ".is-punched"
	if _, err := page.Timeout(DefaultTimeout).Element(punchedButtonSelector); err != nil {
		return errors.Join(errors.New("unable to redirect in 10 seconds"), err)
	}

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
	// using a custom launcher instead of rod.New because
	// https://stackoverflow.com/questions/70254649/rod-running-in-docker-alpine-get-error-chrome-linux-chrome-no-such-file-or-dir
	path, foundPath := launcher.LookPath()
	if !foundPath {
		return errors.New("cannot find launcher path")
	}
	// also chromium needs to launch with --no-sandbox when using with root
	launcher, err := launcher.New().Bin(path).Set("--no-sandbox").Launch()
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

	// redirect and wait for punch button to show
	var punchButtonSelector string
	if status == PunchIn {
		punchButtonSelector = "div.por-punch-clock__button-group > div.el-row > div:nth-child(1) > button"
	} else {
		punchButtonSelector = "div.por-punch-clock__button-group > div.el-row > div:nth-child(2) > button"
	}

	// login to NUEiP
	if err := nueip.login(page, punchButtonSelector); err != nil {
		return fmt.Errorf("unable to login: %w", err)
	}

	// disable geolocation permission
	if err := nueip.disableGeolocationPerm(page); err != nil {
		return fmt.Errorf("unable to disable location permissions: %w", err)
	}

	// punch the clock in/out button
	if err := nueip.punchClock(page, punchButtonSelector, status); err != nil {
		return fmt.Errorf("unable to punch button: %w", err)
	}

	return err
}
