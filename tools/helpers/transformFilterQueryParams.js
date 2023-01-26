const FILTER_QUERY_NAME='filter'

const filterSchemaTransformerMap = {
    '#/components/schemas/StringFieldFilter': transformStringFieldFilter,
    '#/components/schemas/StringFieldEqualsFilter': transformStringFieldEqualsFilter,
    '#/components/schemas/BooleanFieldFilter': transformBooleanFieldFilter,
};

/**
 * In our standard spec we use a nested object representation for filter query parameters.
 * Since OAS does not officially support nested object params, we can use this function
 * to transform them into unnested objects which suits things like rendering (i.e. API docs) better.
 */
function transformFilterQueryParams(spec) {
  if (!isObject(spec)) {
    throw new Error(`Unexpected Open API spec object: ${spec}`);
  }

  for (const path of Object.keys(spec.paths ?? {})) {
    for (const method of Object.keys(spec.paths[path])) {
      let currParamIdx = 0;
      while (currParamIdx < spec.paths[path][method].parameters?.length) {
        const param = spec.paths[path][method].parameters[currParamIdx];
        if (param.in === 'query' && param.name === FILTER_QUERY_NAME && param.schema?.type === 'object') {
          const transformedParams = transformFilterQueryParam(param);
          spec.paths[path][method].parameters.splice(currParamIdx, 1, ...transformedParams);
          currParamIdx += transformedParams.length
        } else {
          currParamIdx++;
        }
      }
    }
  }

  return spec
}

function transformFilterQueryParam(param) {
  if (!isObject(param)) {
    throw new Error(`Unexpected OAS Parameter object: ${param}`);
  }

  const transformedProperties = [];

  for (const [name, value] of Object.entries(param.schema.properties)) {
    if (isFilterSchemaRef(value)) {
      transformedProperties.push(...filterSchemaTransformerMap[value.$ref](name));
    } else {
      transformedProperties.push({ [name]: value});
    }
  }

  return transformedProperties;
}

function transformStringFieldFilter(filterName) {
  const properties = [];
  properties.push(...transformStringFieldEqualsFilter(filterName));
  properties.push({
    schema: { type: 'string' },
    in: 'query',
    name: `${FILTER_QUERY_NAME}[${filterName}][contains]`,
    description: `Filter results by ${filterName} that contains the given value.`
  })

  return properties;
}

function transformStringFieldEqualsFilter(filterName) {
  const properties = [];
  properties.push({
    schema: { type: 'string' },
    in: 'query',
    name: `${FILTER_QUERY_NAME}[${filterName}]`,
    description: `Filter results by ${filterName} that equals the given value.`
  });

  return properties;
}

function transformBooleanFieldFilter(filterName) {
    const properties = [];
    properties.push({
        schema: { type: 'boolean' },
        in: 'query',
        name: `${FILTER_QUERY_NAME}[${filterName}]`,
        description: `Filter results by ${filterName} being true or false.`
    });

    return properties;
}

function isFilterSchemaRef(val) {
  return isObject(val) && 
    typeof val.$ref === 'string' &&
    Object.keys(filterSchemaTransformerMap).includes(val.$ref);
}

function isObject(val) {
  return typeof val === 'object' && val !== null && !Array.isArray(val);
}

module.exports = {
    transformFilterQueryParams,
}
