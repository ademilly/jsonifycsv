# jsonifycsv

Converts a CSV file to a JSON string dumped into stdout.

## Install

Using go toolchain:
```
    $ go get github.com/ademilly/jsonifycsv
```

## Usage

Inline arguments:
- filename: path to CSV file
- sep: separator character used in CSV file

```
    $ jsonifycsv -h
    => display inline help
    $ jsonifycsv -filename path_to_csv -sep 'some_character'
    => converts file at 'path_to_csv' with sep 'some_character' into a json string
    $ jsonifycsv -filename path_to_csv -sep 'some_character' > path_to_json
    => converts file at 'path_to_csv' with sep 'some_character' into a json string and pipe it to 'path_to_json'
```

