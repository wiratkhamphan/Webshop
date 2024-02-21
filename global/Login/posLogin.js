document.getElementById("login-form-element").addEventListener("submit", function(event) {
    event.preventDefault(); // ไม่ให้ฟอร์ม submit โดยปกติ

    var formData = new FormData(this); // เก็บข้อมูลจากฟอร์ม
    var payload = {
        username: formData.get("username"),
        password: formData.get("password")
    }; // สร้างข้อมูลที่จะส่งไปยังเซิร์ฟเวอร์

    // ส่งข้อมูลผ่าน AJAX
    fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
    })
    .then(response => {
        if (response.ok) {
            // alert("Login successful!");
            $('#successful').modal('show');
            
            setTimeout(function() {
                window.location.href = "/index.html";
            }, 900);
            
        } else {
            $('#failed').modal('show');
            // alert("Login failed. Please check your credentials.");
        }
    })
    .catch(error => {
        console.error("Error:", error);
        $('#exampleModal').modal('show'); // เรียกใช้ modal

    });
});