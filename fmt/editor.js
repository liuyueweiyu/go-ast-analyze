const cp = require("child_process")
const promisefy = require("util").promisify
const fs = require("fs")
const path = require("path")

const exec = promisefy(cp.exec)
const stat = promisefy(fs.stat)
const readdir = promisefy(fs.readdir)

// 注:因为go中的变量不能带特殊字符,所以在提换的时候,var的后面的定义的变量不能带 '.'
async function changeName(from, to , filePath){
    try{
        for(let i = 0, len = from.length; i < len ; i++) {
            const cmd = `gofmt -r="${from[i]} -> ${to[i]}" -w ${filePath}`
            fin = await exec(cmd)
            console.log(`${filePath}: 解析成功`)
        }
    }
    catch(e){
        console.log(`${filePath}: 解析成功,原因 `)
        console.log(e)
    }
}


function AddPackagePrefix_File(pkgName,args,filePath){
    let to = args.map(v=>{
        return pkgName + '.' + v
    })
    changeName(args,to,filePath)
}

async function  AddPackagePrefix_Dir(pkgName,args,dirPath){
    try{
        let files = await readdir(dirPath)
        files.forEach(async filename=>{
            let fileDir = path.join(dirPath,filename)
            let fileStat = await stat(fileDir)
            if(fileStat.isDirectory()) {
                AddPackagePrefix_Dir(pkgName,args,fileDir)
            }
            else {
                AddPackagePrefix_File(pkgName,args,fileDir)
            }
        })
    } catch(e){
        console.log(e)
    }
}


async function EditorGoImport(pkgName,dirPath) {
    
}


async function EditorGoPackageName(pkgName,dirPath) {

}
// module.exports = {
//     AddPackagePrefix_Dir,
//     AddPackagePrefix_File
// }
