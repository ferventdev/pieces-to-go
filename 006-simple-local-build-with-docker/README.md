Dockerfile, which is present in this directory, shows the simplest approach for building and running a Go program via docker. Nevertheless, it can be applied for local builds especially with the small codebases because of the Go's high compilation speed. If needed, additional build steps, such as autoformat, linter and tests running can easily be added to this Dockerfile.


Use the following command to build an app's docker image and run a container from it in the terminal:

	docker build . -t go-app && docker run --rm -p 8080:8080 go-app

Pressing Ctrl+C will stop the container and automatically remove it due to the `--rm` option.


To run in a detached (background) mode use:

	docker build . -t go-app && docker run --rm -d -p 8080:8080 go-app


When runnning, please note that available http-server endpoints are only `/ping` and `/info`.
