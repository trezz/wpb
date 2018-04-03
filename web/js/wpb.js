
/*
 * Requests the server to load the "add person" page
 */
$("button#addNewPerson")
    .click(function(event) { window.location.href = "person-add.html"; })

/*
 * Post a new Person JSON serialized object to the server
 */
$("button#addPersonButton").click(function(event) {
  // Use a form but prevent default behavior which is a http POST and
  // reloads the page... The POST is made in js below.
  event.preventDefault()

  // Get data to transmit to the server
  var desc = $("textarea#personDesc").val()
  var name = $("input#personName").val()

  // Post the data using JSON using AJAX!
  $.post("addPerson", {"PersonName" : name, "PersonDescription" : desc},
         function(data) { console.log("post succeed"); })
})

/*
 * Create a new 'location' entry in the "Add Person" table.
 */
$(".addLoc").click(function(event) {
  // get the table
  var table = document.getElementById("personAddFormTable")
  // create the latitude input object
  var latitude = document.createElement("input")
  latitude.class = "latitude"
  latitude.type = "text"

  // create the longitude input object
  var longitude = document.createElement("input")
  longitude.class = "longitude"
  longitude.type = "text"

  // create the date input object
  var date = document.createElement("input")
  date.class = "date"
  date.type = "text"
  $(date).datepicker()

  // create the references text area
  var refs = document.createElement("textarea")
  refs.class = "refs"
  refs.rows = 5
  refs.cols = 30

  // create the "Cancel" button
  var cancelButton = document.createElement("button")
  cancelButton.id = "cancelButton"
  cancelButton.innerText = "Cancel"
  cancelButton.onclick =
      function(event) {
    // get the table and the clicked item
    var table = document.getElementById("personAddFormTable")
    var clicked = $(event.target)
    var clickedRow = clicked.parent().parent()
    // remove the 2 previous rows and the button row clicked
    clickedRow.prev().prev().prev().prev().remove()
    clickedRow.prev().prev().prev().remove()
    clickedRow.prev().prev().remove()
    clickedRow.prev().remove()
    clickedRow.remove()
  }

  // Insert elements in the table

  var latitudeRow = table.insertRow(-1)
  latitudeRow.insertCell(0).innerHTML = "Latitude"
  latitudeRow.insertCell(1).append(latitude)

  var longitudeRow = table.insertRow(-1)
  longitudeRow.insertCell(0).innerHTML = "Longitude"
  longitudeRow.insertCell(1).append(longitude)

  var dateRow = table.insertRow(-1)
  dateRow.insertCell(0).innerHTML = "Date"
  dateRow.insertCell(1).append(date)

  var refsRow = table.insertRow(-1)
  refsRow.insertCell(0).innerHTML = "References"
  refsRow.insertCell(1).append(refs)

  var buttonsRow = table.insertRow(-1)
  buttonsRow.insertCell(0).append("")
  buttonsRow.insertCell(1).append(cancelButton)
})