# Punch

Punch card automation for my company

## Usage

```bash
make # build binary
./punch [in|out] --username={username} --password={password} --company={company}
```

## Development

To check if automation works, with browser and dev tools showing

```golang
package main

import "github.com/imfulee/punch/hr_system"

func main() {
    nueip := NUEIP{...} // Initialize a NUEIP instance
    nueip.Punch(PunchStatus.In) // Place a punch status, declared in hr_system/
}
```

Then run

```bash
go run . -rod=show,devtools
```

## Roadmap

Things that I would like to further develop

### Now

- Write better docs on how to use this program
- Some reporting mechanism that it doesn't work, possibly by sending email?
- Add container building definitions (ex `docker-compose.yml`)

### Backlog

- Add CI to build container
- Write test to check if punch in works
