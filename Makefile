TEMPDIR := $(shell mktemp -d)

.swagger-codegen-cli.jar:
	wget http://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/2.2.1/swagger-codegen-cli-2.2.1.jar -O .swagger-codegen-cli.jar
	java -jar .swagger-codegen-cli.jar help

swagger.json: .swagger-codegen-cli.jar
	java -jar .swagger-codegen-cli.jar generate -l swagger -i https://api.newrelic.com/v2/definitions.json

generate: .swagger-codegen-cli.jar swagger.json
	java -jar .swagger-codegen-cli.jar generate \
		-i swagger.json \
		-l go \
		-c generate-config.json \
		-o $(TEMPDIR)

	mkdir -p api
	cp $(TEMPDIR)/*.go api/
	rm -rf $(TEMPDIR)
.PHONY: generate