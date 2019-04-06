const fs = require('fs')
const execSync = require('child_process').execSync
const deps = require('./package.json').dependencies
const devDeps = require('./package.json').devDependencies

let shouldInstall = false

for (const depFolder in Object.assign({}, deps, devDeps)) {
  if (!fs.existsSync(`./node_modules/${depFolder}`)) {
    console.log(`Dependency "${depFolder}" is NOT installed - installing now...`)
    shouldInstall = true
  }
}

if (shouldInstall) {
  const command = 'npm install'
  console.log('Installing...')
  console.log(command)
  execSync(command)
  console.log('All dependencies are just installed.')
  process.exit(0)
}
console.log('All dependencies are already installed.')
