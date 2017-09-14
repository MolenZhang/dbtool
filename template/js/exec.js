$("#logout").click(function(event){
    event.preventDefault();
    del_cookie("admin_id");

	var target = event.target;
	$.post("/logout", $(target).serialize(), function(ret) {
		if(ret.Ret == "0") {
			alert(ret.Reason);
		}else{
			window.location.href = "/login/index";
		}
	}, "json")	
	
	//		window.location.href = "/login/index";
	
})

function del_cookie(name)
{
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
}

$("form[data-type=formAction]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var action = $(target).attr("action");
	$("#sqlTbody").empty();
	$("#sqlThead").empty();
	$.get(action, $(target).serialize(), function(ret){

       if(ret.Ret == "0") {
            alert(ret.Reason);
        }else{
		var tbodyTr= "";
		var theadTr= "";
		for(var i in ret){
			theadTr ='<tr>';
			tbodyTr +='<tr>';
			for( var j in ret[i]){
				theadTr +='<td>'+j+'</td>';
			
				tbodyTr +='<td>'+ret[i][j]+'</td>';
			}
			
			
			theadTr += '</tr>';
			tbodyTr += '</tr>';
		//	alert(ret[i]);
		
		/*	tbodyTr += '<tr>';
			for(var j in ret[i])
		'<td>'+ret[i].admin_id+'</td>'+
		'<td>'+ret[i].admin_name+'</td>'+
			'<td>'+ret[i].admin_password+'</td>'+
		'</tr>';*/
		}
		$("#sqlThead").append(theadTr);
		$("#sqlTbody").append(tbodyTr);
		}
    //    if(ret.Ret == "0") {
    //        alert(ret.Reason);
    //    } else {
	  //location.href = $(target).attr("form-redirect");
	//	}
    },"json")
})

