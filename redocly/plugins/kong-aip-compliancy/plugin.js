function AddAddionalProp(options) {
    return {
        Schema(schema) {
            if (schema.type != "object") { return }
            if (!schema.hasOwnProperty('additionalProperties')) {
                schema.additionalProperties = false
            }
        }
    }
}

module.exports = {
    id: 'kong-aip-compliancy',
    preprocessors: {
        oas3: {
            'add-additional-prop': AddAddionalProp
        }
    }
}
