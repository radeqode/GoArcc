## checking code linting , comments , spacing , variable name
## Dead code etc.
go get github.com/mgechev/revive
revive -formatter friendly ./...
go get fmt
go fmt ./...

