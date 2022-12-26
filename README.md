# OTP-Filemanager

A file server that allows to access files with a OTP-based authentication.
This server offers a simple interface, but also allows it do integrate into your application through
the endpoints, as it's stateless and relies on an up-to-date OTP.

## Run a Build
1. Check branch `release/x.x`
2. Set variables in `.env` (see below)
3. Run `./otp-filemanager-x.x`


## Run for Development
1. Run `./scripts/setup.sh` 
  * Set `HTTPPORT` 
  * Set `IDSEED` to a random numeric sequence
  * Set `ISSUER` to your service name with URL compatible characters 
  * Set `PERIOD` to seconds a OTP is valid
  * Set `MAXFILESIZEMB` to set the max size for an uploaded file
2. `go run main.go`


## How to Use

1. [Create Account](/docs/createAccount.md)
2. [Authentication](/docs/authentication.md)


## Inner Workings
The authentication relies on a randomly generated username and OTP (currently only TOTP).
The password has to be up-to-date, as no other authentication method is used for the user actions.

Files and user related information are saved directly to the file system.
### Architecture

![Overview](./docs/architecture/actual.png)


## TODO
- Extensive load and security tests
- Separate (better) UI
- Extend to other OTP methods other than TOTP
- Implement alternative methods of saving users and files
- Encrypt users and files
