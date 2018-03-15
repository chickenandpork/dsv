# dsv

DataStore Validation: for a related project, do you have a datastore we can use?

The structure of this project needs ot remain fairly constant -- it's a part of a larger thing.  I hope with additions to the `gorm` directory, and cheanges to existing `proto` files, we can confirm functionality on more datastore backends such as MongoDB, Sqlite, whatever we can use with some performance.

## Building

Because this uses github.com/infobloxopen/protoc-gen-gorm which has an issue in latest Go that is unaddressed across its 50-ish forks, we need to do this a bit more manually:

  1. get the binary installed so that `which protoc-gen-gorm` works:
    1. git clone https://github.com/infobloxopen/protoc-gen-gorm.git
    2. cd protoc-gen-gorm && go install
  2. generate the things:
    1. go generate ./...
  3. run it:

    go run . --dialect sqlite3 --connect ":memory:"
    go run . --dialect mariadb --connect "scott:tiger@uberdatabase.example.com/dsvtablespace"

