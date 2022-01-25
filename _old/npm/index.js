var binwrap = require('binwrap')
var path = require('path')

var packageInfo = require(path.join(__dirname, '..', 'package.json'))
var version = packageInfo.version
var root = 'https://github.com/ory/release-canary/releases/download/v' + version

module.exports = binwrap({
  dirname: __dirname,
  binaries: ['release-canary'],
  urls: {
    'linux-x64': root + '/release-canary_' + version + '_linux_amd64.tar.gz',
    'darwin-x64': root + '/release-canary_' + version + '_darwin_amd64.tar.gz',
    'darwin-arm64': root + '/release-canary_' + version + '_darwin_arm64.tar.gz'
  }
})
