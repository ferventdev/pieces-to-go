Dockerfile, which is present in this directory, makes use of a multi-stage docker build. 
It may (and probably should) be applied for CI/CD pipelines instead of a simple build (shown previously), because with a multi-stage approach the resultant image has significantly smaller size (as it mainly holds just the app binary). The first stage builds the executable artifact, which is then inserted into the lightweight distroless image.


Use the following command to build an app's docker image and run a container from it in the terminal:

	docker build . -t go-app && docker run --rm -p 8080:8080 go-app

Pressing Ctrl+C will stop the container and automatically remove it due to the `--rm` option.


To run in a detached (background) mode use:

	docker build . -t go-app && docker run --rm -d -p 8080:8080 go-app


When runnning, please note that available http-server endpoints are only `/ping` and `/info`.
