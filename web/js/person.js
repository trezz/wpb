
// Post a new Person JSON serialized object to the server
$("button#AddPersonButton").click(function(event) {
    // Use a form but prevent default behavior which is a http POST and
    // reloads the page... The POST is made in js below.
    event.preventDefault()

    // Get data to transmit to the server
    var desc = $("textarea#PersonDescription").val()
    var name = $("input#PersonName").val()

    // Post the data using JSON using AJAX!
    $.post("/addPerson", {"PersonName" : name, "PersonDescription" : desc},
           function(data) { console.log("post succeed"); })
})