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

- Some reporting mechanism that it doesn't work, possibly by sending email? Currently there's a workaround print the output to a file in a crontab like `COMMAND > file.txt`. 
- Add CI to build container
- Write test to check if punch in works
- Make sure `rod` uses the right orientation and screen size

