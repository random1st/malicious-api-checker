.PHONY: build

build:
	sam build

deploy:
	sam deploy --guided

invoke:
	sam local invoke   --env-vars=env.json
