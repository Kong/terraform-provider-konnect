const SwaggerParser = require('@apidevtools/swagger-parser')
const fs = require('fs/promises')
const path = require('path')
const dataFolder = path.resolve(__dirname, 'testdata')
const schemaFolder = path.resolve(__dirname, '..', 'definitions')
const yaml = require('yaml')
const Ajv = require('ajv')
const addFormats = require('ajv-formats')

const oasStub = {
    openapi: "3.0.3",
    info: {},
    servers: [],
    paths: {}
}

async function main() {
    let errorCount = false
    const testDatas = await fs.readdir(dataFolder)
    for (const testCasesFile of testDatas) {
        const schemaName = path.parse(testCasesFile).name
        if (schemaName.startsWith('.')) {
          continue
        }
        const testCases = JSON.parse(await fs.readFile(path.resolve(dataFolder, testCasesFile), 'utf-8'))
        const rawDefinition = await fs.readFile(path.resolve(schemaFolder, `${schemaName}.yaml`), 'utf-8')

        // OAS headers are required to have swaggerparser properly dereference the schemas
        const stubbedOAS = {
            ...oasStub,
            ...yaml.parse(rawDefinition)
        }

        const parsedDefinition = await SwaggerParser.dereference(stubbedOAS, { validate: false })

        // instanciation and adding all the schemas in the AJV environment to handle the `$ref`
        let mainValidate = addFormats(new Ajv({ allErrors: true })).addVocabulary(['example', 'x-validation-message'])
        for (const def in parsedDefinition.components.schemas) {
            mainValidate = mainValidate.addSchema(parsedDefinition.components.schemas[def], def)
        }

        console.log("testing definition:", `${schemaName}.yaml`)

        // Loop over all the test cases, the json files in testdata folder
        for (const typeDefinition in testCases) {
            console.log(`schema: ${typeDefinition}`)
            if (!parsedDefinition.components.schemas[typeDefinition]) {
                throw "definition not found"
            }

            let idx = 0
            // Positive assertions
            if (testCases[typeDefinition].ok) {

                console.log(`ok:`)
                for (const ok of testCases[typeDefinition].ok) {
                    const validate = mainValidate.compile(parsedDefinition.components.schemas[typeDefinition])


                    const validateResult = validate(ok)
                    if (!validateResult) {
                        console.error(`  ⚔︎ ERROR: OK assertions item#${idx}`, validate.errors)
                        errorCount++
                    } else {
                        console.log(`  √ - item#${idx}`)
                    }
                    idx++
                }
            }

            idx = 0
            // Negative assertions
            if (testCases[typeDefinition].nok) {
                console.log(`nok:`)
                for (const nok of testCases[typeDefinition].nok) {
                    const validate = mainValidate.compile(parsedDefinition.components.schemas[typeDefinition])

                    const validateResult = validate(nok)
                    if (validateResult) {
                        console.error(`  ⚔︎ ERROR: nok item#${idx} should be invalid`)
                        errorCount++
                    } else {
                        console.log(`  √ - item#${idx}`)
                    }
                    idx++
                }
            }
            console.log('')
        }
    }
    if (errorCount) {
        console.error(`TESTS FAILED: ${errorCount} error${errorCount > 1 ? 's': ''} found`)
        process.exit(1)
    } else {
        console.log('all tests passed')
    }
}

if (require.main === module) {
    main().catch(e => { throw e })
}
