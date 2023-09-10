# Punch

Punch card automation for my company

## Usage

```bash
punch [in|out] --useername={username} --password={password} --company={company}
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
