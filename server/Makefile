SHELL := /bin/bash

run:
	@npm run init && npm start

clean:
	@docker ps > cmd.txt && LOL=`sed -nrz 's/.*NAMES\n(.*)   postgres.*/\1/p' cmd.txt` && docker stop "$$LOL" && docker container rm "$$LOL" && rm cmd.txt
