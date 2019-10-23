const fs = require('fs')
const path = require('path')
const DIRPATH = path.join(__dirname,'analyze')
const os = require('os')
// packages的每个文件夹 是相当于每一个包
const packages = fs.readdirSync(DIRPATH)

// 换行这事以后再说..
// let platform;
// (()=>{
//     if(os.platform() == 'darwin') {
//         platform = 'mac'
//     } else if (os.platform() == 'win32') {
//         platform = 'windows'
//     } else {
//         platform = 'lunix'
//     }
// })()
// const linebreak  = {
//     windows : '\r\n',
//     mac : '\r',
//     lunix : '\n'
// }

// 分析这个包内的struct依赖关系

const structRep = /^type (.*) struct {$/

packages.forEach(package=>{
    const files = fs.readdirSync(path.join(DIRPATH,package))
    files.forEach(name=>{
        const content = fs.readFileSync(path.join(DIRPATH,package,name),'utf8')
        const lines = content.split('\r\n')
        console.log(lines)
        const structs = lines.filter((line,i)=>{
            structRep.test(line)
        })
        .map(line=>line.match(structRep)[1])
        // console.log(structs)
    })
})



// function getStructContent(lines,begin) {
    
// }