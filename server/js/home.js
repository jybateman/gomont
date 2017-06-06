// [0] server id
// [1] command name
// [2] command return value
// [3] command output
function explodeString(str) {
    var info = []
    for (i = 0, c = 0; c != 4; c++, i += +len + +off) {
	var sstr = str.substring(i);
	var len = sstr.split(":")[0];
	var off = len.length + 1;
	info.push(str.substring(+i + +off, +i + +len + +off));
    }
    return info
}

function checkStatus(cells, stat) {
    var status;
    switch(stat) {
    case "0":
	status = "success";
	break;
    case "1":
	status = "warning";
	break;
    case "2":
	status = "danger";
	break;
    default:
	status = "active";
	break;
    }
    for (var i = 0; i < cells.length; i++) {
	cells[i].className = status;
    }
}


var graph = new Graph();
var ip = location.host;
var ws = new WebSocket("ws://"+ip+"/ws");
ws.onmessage = function (event) {
    if (event.data != "") {
	var res = explodeString(event.data);
	var table = document.getElementById(res[0]);
	var cmd = table.querySelector("#"+res[1]);
	if (cmd == null) {
    	    var body = table.getElementsByTagName("tbody");
    	    var row = body[0].insertRow(0);
    	    var name = row.insertCell(0);
    	    var out = row.insertCell(1);
    	    row.id = res[1];
	    name.innerHTML = res[1];
	    out.innerHTML = res[3];
	    checkStatus(row.cells, res[2]);
	} else {
	    var cells = cmd.getElementsByTagName("td");
	    cells[0].innerHTML = res[1];
	    cells[1].innerHTML = res[3];
	    checkStatus(cells, res[2]);
	}

	if (graph.GetMap(res[0]+res[1])) {
	    graph.UpdateGraph(res, ip);
	} else {
	    graph.CreateGraph(res, ip);
	}
    }
}
