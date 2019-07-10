
// Reload the whole messages panel
var refresh = function() {
    document.getElementById("thread").value += "Hello";
}

// Call refresh every 5 seconds
setInterval(refresh, 1000);


// var refresh = function() {
//     var xmlHttp = new XMLHttpRequest();
//     xmlHttp.open( "GET", "http://192.168.1.105:9000/refresh", false ); // false for synchronous request
//     xmlHttp.send( null );


//     var textArea = document.getElementById("thread")

//     textArea.value += "Hello There\n";
//     // textArea.value += xmlHttp.responseText;
//     // textArea.scrollTop = textArea.scrollHeight;
// }