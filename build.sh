echo "BUILDING APPLICATION..."

rm -r -f tmp
mkdir tmp

cp -r content tmp/content
cp Dockerfile tmp/Dockerfile

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-w' -v -o ./tmp/app ./

