Use the following command to build an app's docker image and run a container from it in the terminal:

	docker build . -t go-app && docker run --rm -p 8080:8080 go-app

Then Ctrl+C would stop the container.

Or, to run in a detached (background) mode use:

	docker build . -t go-app && docker run --rm -it -p 8080:8080 go-app

Available http-server endpoints are only /ping and /info.