class Graph {

    constructor() {
	this.gmap = new Map();
    }

    GetMap(key) {
    	return this.gmap.get(key);
    }
    
    CreateGraph(info, host) {
    	var s = document.getElementById(info[0]);
    	var c = document.createElement("div");
	c.style.float = "left";
    	c.id = "graph"+info[0]+info[1];
	s.appendChild(c);
    	var g = new Dygraph(
    	    document.getElementById("graph"+info[0]+info[1]),
    	    "http://"+host+"/files/"+info[0]+"/"+info[1]+".csv",
    	    {}
    	);
    	this.gmap.set(info[0]+info[1], g);
    }

    UpdateGraph(info, host) {
	var g = this.gmap.get(info[0]+info[1]);
    	g.updateOptions({
	    file: "http://"+host+"/files/"+info[0]+"/"+info[1]+".csv",
    	});
    }
}
