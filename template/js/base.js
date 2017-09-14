function del_cookie(name)
{
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
}

$("form[data-type=formAction]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var action = $(target).attr("action");
    $.post(action, $(target).serialize(), function(ret){
   // $.get(action, $(target).serialize(), function(ret){
       if(ret.Ret == "0") {
            alert(ret.Reason);
        }else{
	location.href = $(target).attr("form-redirect");
		}
	
    },"json")
})
$(document).ready(function(){
	$("#dbNameCon").click(function(){
	  var dbNameCon = $("#dbNameCon").val();
	  $(".form-control").attr("disabled",false);
      if(dbNameCon == "mysql"){
		$("#ZkRoot").attr("disabled",true);
	    $("#JsonText").attr("disabled",true);
	  }else if(dbNameCon == "oracle"){
	    $("#DbPort").attr("disabled",true);
	    $("#JsonText").attr("disabled",true);
		$("#ZkRoot").attr("disabled",true);
	  }else if(dbNameCon == "sqlite3"){
		$(".form-control:not(#dbNameCon)").attr("disabled",true);
		$("#DbName").attr("disabled",false);
	  }else if(dbNameCon == "hbase"){
		$(".form-control:not(#dbNameCon)").attr("disabled",true);
		$("#DbIP").attr("disabled",false);
	    $("#ZkRoot").attr("disabled",true);
	    $("#DbPort").attr("disabled",false);
	    $("#JsonText").attr("disabled",false);
	  }else if(dbNameCon == "redis"){
		$(".form-control:not(#dbNameCon)").attr("disabled",true);
		$("#DbIP").attr("disabled",false);
	    $("#DbPort").attr("disabled",false);
	  }else{
		$(".form-control").attr("disabled",false);
	  }
	
	})
})
