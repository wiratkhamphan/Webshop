// Your inline JavaScript code for the form submission
document.getElementById("register-form-element").addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the default form submission

    var formData = new FormData(this); // Collect form data
    var payload = {
        username: formData.get("username"),
        password: formData.get("password"),
        email: formData.get("email")
    }; // Prepare payload for the server

    // Send data to server using fetch API
    fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (response.ok) {
            $('#successful').modal('show'); // Show success modal
        } else {
            throw new Error('Server responded with an error!');
        }
    })
    .catch(error => {
        console.error("Error:", error);
        $('#exampleModal').modal('show'); // Show error modal
    });
});