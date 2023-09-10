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

See [example.md](example/example.md) for example

## Development

To check if automation works, with browser and dev tools showing, you could create a go file in the root directory

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

- Some reporting mechanism that it doesn't work, possibly by sending email?
- Add container building definitions (ex `docker-compose.yml`)
- Add CI to build container
- Write test to check if punch in works
