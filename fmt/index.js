const editor = require("./editor")
const shell = require("shelljs")
const path = require("path")

const fileDir = "temp_src"
const pkgName = "pkga"
const files = ["age.go","name.go"]


shell.cd(fileDir)
shell.mkdir(pkgName)
files.forEach(filename=>{
    shell.cp("-r",filename,pkgName)
    shell.rm("-f",filename)
})

