/*
for(var i=0;i<arg.length;i++){
	console.log("arg"+i+":"+arg[i]);
}

var json={name:'zhangan',age:22};
var aa=gofunc(json);
var arrays=[1,2,3,4,5,6];
gofunc(arrays);
console.log("aa:"+aa)
var data=readFile("main.go");
writeFile("abc.txt","hello");
*/
var del=delFile("E://adcfg.json");
console.log(JSON.stringify(del));
/*
scanFile("E://");
function scanFile(dir){
	var result=listFile(dir);
	var jj=JSON.stringify(result);
	console.log(jj);
	if(result.success==0){
		var list=result.data;
		for(var i=0;i<list.length;i++){
			var item=list[i];
			console.log(dir+item.name+"//")
		}
	}else{
		console.log("####"+result.message);
	}
}
*/
