/* The following declarations are to get the user input loaded into local 
    storage. This will allow the input to be used in another page. This is done
    using the HTMLFormControlsCollection API.
*/

// Get the form we need.
var form = document.forms['shop_day_form'];
// <input> on index.html
var data = form.elements.shop_days;

// Print the value of data in the console
console.log(data.value);
// Register the 'submit' event on the <form>
form.addEventListener('submit', saveData);

// Callback called on submit on index.html
function saveData(e) {
    // Prevent the <form> from sending data to server and resetting itself. This also prevented the submit button from loading
    // the next page.
    //e.preventDefault();

    // Get the data from <input> on index.html 
    var str = data.value;

    // Save data to localStorage under the key of "data"
    localStorage.setItem('data', JSON.stringify(str));
}
