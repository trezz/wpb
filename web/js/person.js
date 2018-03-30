
// Post a new Person JSON serialized object to the server
$("button#AddPersonButton").click(function() {

    // Get data to transmit to the server
    var desc = $("textarea#PersonDescription").val()
    var name = $("input#PersonName").val()

    // Post the data using JSON using AJAX!
    $.post("/addPerson", {"PersonName" : name, "PersonDescription" : desc},
           function(data) { console.log("post succeed"); })
})