echo update openapi-generator
./bin/utils/openapi-generator-cli.sh
rm -Rf ./namsor/client/go/
echo run openapi-generator
java -jar modules/openapi-generator-cli/target/openapi-generator-cli.jar generate --git-repo-id namsor-golang-sdk2 --git-user-id namsor --additional-properties=packageVersion=2.0.13 --additional-properties=packageName=namsorapi -i https://v2.namsor.com/NamSorAPIv2/api2/openapi.json -g go -o  namsor/client/go
cd ./namsor/client/go/
cp -R /home/namsor/codegen/openapi-generator/namsor/client/go/* /home/namsor/codegen/namsor-golang-sdk2/
cp /home/namsor/codegen/openapi-generator/run-golang.bash /home/namsor/codegen/namsor-golang-sdk2/
#~/dropbox_uploader.sh upload namsor-sdk2-1.0.0.jar namsor-sdk2-1.0.0.jar
#~/dropbox_uploader.sh upload namsor-sdk2-1.0.0-javadoc.jar namsor-sdk2-1.0.0-javadoc.jar
#~/dropbox_uploader.sh upload namsor-sdk2-1.0.0-sources.jar namsor-sdk2-1.0.0-sources.jar
#~/dropbox_uploader.sh upload namsor-sdk2-1.0.0-tests.jar namsor-sdk2-1.0.0-tests.jar
