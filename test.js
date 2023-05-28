function halo(){
    console.log("halo")
}
function kopi(){
    halo()
}

function print(callback){
    callback()
}

print(kopi)
