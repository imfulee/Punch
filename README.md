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

1. ~~Add check is user is on old website, and switch to the new one~~ I think NUEIP updated their website to use the new layout by default now
2. Some reporting mechanism that it doesn't work, possibly by sending email?
3. Add CI to build container
4. Write test to check if punch in works
5. Make sure `rod` uses the right orientation and screen size

### Workaround

2. Just print the output to a file in a crontab like `COMMAND > file.txt`
