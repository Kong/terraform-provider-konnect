# common definitions tests

Can be run using `node index.js` . Test assertions live in `testdata` folder, each testcase file should match the filename
of its common definition in the `defitions` folder. Example `common/test/testdata/errors.json` will map the schemas from
`common/definitions/errors.yaml` . Assertion file is a map with the schema name as key and each key can contains positive
or negative payload assertions like so:

```json
{
    "BadRequestError": {
        "ok": [
            {
                "foo": "bar" // This payload needs to be validated
            }
        ],
        "nok": [
            {
                "bar": "foo" // this payload needs to be invalidated
            }
        ]
    },
}
```
