const fs = require('fs').promises

async function tryMkdir(path, options) {
  try {
    await fs.mkdir(path, options)
  } catch (e) {
    if (e.code !== "EEXIST") {
      throw e;
    }
  }
}

module.exports = { tryMkdir }
