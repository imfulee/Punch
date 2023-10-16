# Punch

Punch card automation for my company

## Usage

```bash
# build binary
make

# punch command
./punch \
    In|Out \
    --username=USERNAME \
    --password=PASSWORD \
    --company=COMPANY
```

## Development

To check if automation works, with browser and dev tools showing, you could create a go file in the root directory

```golang
package main

import hrSystem "github.com/imfulee/punch/pkg/nueip"

func main() {
    nueip := hrSystem.NUEIP{...} // Initialize a NUEIP instance
    nueip.Punch(hrSystem.PunchStatus.In) // Place a punch status, declared in hr_system/
}
```

Then run

```bash
go run . -rod=show,devtools
```

## Container

To build as a container, replace podman with docker if that's what you use.

```bash
make podman
```

To run the container, you might want to setup up a cron job

```text
0 9 * * 1-5 podman run --rm <punch-image-name> /app/punch --username=<username> --password=<password> --company=<company>
0 17 * * 1-5 podman run --rm <punch-image-name> /app/punch --username=<username> --password=<password> --company=<company>
```

and manually set the cron time of your work, image name, username etc.

## Wishlist

Things that I would like to further develop

- Add check is user is on old website, and switch to the new one
- Some reporting mechanism that it doesn't work, possibly by sending email?
- Add CI to build container
- Write test to check if punch in works
- Make sure `rod` uses the right orientation and screen size
